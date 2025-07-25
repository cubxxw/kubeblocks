/*
Copyright (C) 2022-2025 ApeCloud Co., Ltd

This file is part of KubeBlocks project

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package cluster

import (
	"context"
	"fmt"
	"reflect"
	"slices"
	"strconv"
	"strings"

	"golang.org/x/exp/maps"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/utils/pointer"
	"sigs.k8s.io/controller-runtime/pkg/client"

	appsv1 "github.com/apecloud/kubeblocks/apis/apps/v1"
	"github.com/apecloud/kubeblocks/pkg/constant"
	"github.com/apecloud/kubeblocks/pkg/controller/component"
	"github.com/apecloud/kubeblocks/pkg/controller/graph"
	"github.com/apecloud/kubeblocks/pkg/controller/model"
	"github.com/apecloud/kubeblocks/pkg/controller/sharding"
	ictrlutil "github.com/apecloud/kubeblocks/pkg/controllerutil"
	"github.com/apecloud/kubeblocks/pkg/generics"
)

// clusterComponentTransformer transforms components and shardings to mapping Component objects
type clusterComponentTransformer struct{}

var _ graph.Transformer = &clusterComponentTransformer{}

func (t *clusterComponentTransformer) Transform(ctx graph.TransformContext, dag *graph.DAG) error {
	transCtx, _ := ctx.(*clusterTransformContext)
	if transCtx.OrigCluster.IsDeleting() {
		return nil
	}

	updateToDate, err := checkAllCompsUpToDate(transCtx, transCtx.Cluster)
	if err != nil {
		return err
	}

	// if the cluster is not updating and all components are up-to-date, skip the reconciliation
	if !transCtx.OrigCluster.IsUpdating() && updateToDate {
		return nil
	}

	return t.transform(transCtx, dag)
}

func (t *clusterComponentTransformer) transform(transCtx *clusterTransformContext, dag *graph.DAG) error {
	runningSet, err := t.runningSet(transCtx)
	if err != nil {
		return err
	}
	protoSet := t.protoSet(transCtx)

	createSet, deleteSet, updateSet := setDiff(runningSet, protoSet)

	if err := deleteCompNShardingInOrder(transCtx, dag, deleteSet, pointer.Bool(true)); err != nil {
		return err
	}

	var delayedErr error
	if err := t.handleUpdate(transCtx, dag, updateSet); err != nil {
		if !ictrlutil.IsDelayedRequeueError(err) {
			return err
		}
		delayedErr = err
	}

	if err := t.handleCreate(transCtx, dag, createSet); err != nil {
		return err
	}

	return delayedErr
}

func (t *clusterComponentTransformer) runningSet(transCtx *clusterTransformContext) (sets.Set[string], error) {
	return clusterRunningCompNShardingSet(transCtx.Context, transCtx.Client, transCtx.Cluster)
}

func (t *clusterComponentTransformer) protoSet(transCtx *clusterTransformContext) sets.Set[string] {
	names := sets.Set[string]{}
	for _, comp := range transCtx.components {
		names.Insert(comp.Name)
	}
	for _, sharding := range transCtx.shardings {
		names.Insert(sharding.Name)
	}
	return names
}

func (t *clusterComponentTransformer) handleCreate(transCtx *clusterTransformContext, dag *graph.DAG, createSet sets.Set[string]) error {
	handler := newCompNShardingHandler(transCtx, createOp)
	return handleCompNShardingInOrder(transCtx, dag, createSet, handler)
}

func (t *clusterComponentTransformer) handleUpdate(transCtx *clusterTransformContext, dag *graph.DAG, updateSet sets.Set[string]) error {
	handler := newCompNShardingHandler(transCtx, updateOp)
	return handleCompNShardingInOrder(transCtx, dag, updateSet, handler)
}

func deleteCompNShardingInOrder(transCtx *clusterTransformContext, dag *graph.DAG, deleteSet sets.Set[string], scaleIn *bool) error {
	handler := newCompNShardingHandler(transCtx, deleteOp)
	if h, ok := handler.(*clusterParallelHandler); ok {
		h.scaleIn = scaleIn
	}
	if h, ok := handler.(*orderedDeleteHandler); ok {
		h.scaleIn = scaleIn
	}
	return handleCompNShardingInOrder(transCtx, dag, deleteSet, handler)
}

func handleCompNShardingInOrder(transCtx *clusterTransformContext, dag *graph.DAG, nameSet sets.Set[string], handler clusterConditionalHandler) error {
	orderedNames, err := handler.ordered(sets.List(nameSet))
	if err != nil {
		return err
	}
	unmatched := ""
	for _, name := range orderedNames {
		ok, err := handler.match(transCtx, dag, name)
		if err != nil {
			return err
		}
		if !ok {
			unmatched = name
			break
		}
		if err = handler.handle(transCtx, dag, name); err != nil {
			return err
		}
	}
	if len(unmatched) > 0 {
		return ictrlutil.NewDelayedRequeueError(0, fmt.Sprintf("retry later: %s are not ready", unmatched))
	}
	return nil
}

func checkAllCompsUpToDate(transCtx *clusterTransformContext, cluster *appsv1.Cluster) (bool, error) {
	compList := &appsv1.ComponentList{}
	labels := constant.GetClusterLabels(cluster.Name)
	if err := transCtx.Client.List(transCtx.Context, compList, client.InNamespace(cluster.Namespace), client.MatchingLabels(labels)); err != nil {
		return false, err
	}
	if len(compList.Items) != transCtx.total() {
		return false, nil
	}
	for _, comp := range compList.Items {
		generation, ok := comp.Annotations[constant.KubeBlocksGenerationKey]
		if !ok {
			return false, nil
		}
		if comp.Generation != comp.Status.ObservedGeneration || generation != strconv.FormatInt(cluster.Generation, 10) {
			return false, nil
		}
	}
	return true, nil
}

// copyAndMergeComponent merges two component objects for updating:
// 1. new a component object targetCompObj by copying from oldCompObj
// 2. merge all fields can be updated from newCompObj into targetCompObj
func copyAndMergeComponent(oldCompObj, newCompObj *appsv1.Component) *appsv1.Component {
	compObjCopy := oldCompObj.DeepCopy()
	compProto := newCompObj

	normalizeQuantity := func(name corev1.ResourceName, q resource.Quantity) resource.Quantity {
		switch name {
		case corev1.ResourceCPU:
			return *resource.NewMilliQuantity(q.MilliValue(), resource.DecimalSI)
		case corev1.ResourceMemory, corev1.ResourceStorage, corev1.ResourceEphemeralStorage:
			return *resource.NewQuantity(q.Value(), resource.BinarySI)
		default:
			return q.DeepCopy()
		}
	}

	normalizeResourceList := func(resources *corev1.ResourceList) {
		if resources == nil {
			return
		}
		for k, v := range *resources {
			(*resources)[k] = normalizeQuantity(k, v)
		}
	}

	normalize := func(spec appsv1.ComponentSpec) appsv1.ComponentSpec {
		normalized := spec.DeepCopy()
		if normalized.Resources.Limits != nil {
			normalizeResourceList(&normalized.Resources.Limits)
		}
		if normalized.Resources.Requests != nil {
			normalizeResourceList(&normalized.Resources.Requests)
		}
		for i := range normalized.VolumeClaimTemplates {
			vct := &normalized.VolumeClaimTemplates[i]
			if vct.Spec.Resources.Limits != nil {
				normalizeResourceList(&vct.Spec.Resources.Limits)
			}
			if vct.Spec.Resources.Requests != nil {
				normalizeResourceList(&vct.Spec.Resources.Requests)
			}
		}
		return *normalized
	}

	mergeConfigs := func(running, expected []appsv1.ClusterComponentConfig) []appsv1.ClusterComponentConfig {
		var mergedConfigs []appsv1.ClusterComponentConfig

		mergedConfigs = append(mergedConfigs, expected...)
		for _, config := range running {
			if config.Name == nil || config.ConfigMap == nil {
				continue
			}
			matchConfig := func(c appsv1.ClusterComponentConfig) bool {
				return pointer.StringEqual(c.Name, config.Name)
			}
			index := generics.FindFirstFunc(mergedConfigs, matchConfig)
			if index < 0 {
				mergedConfigs = append(mergedConfigs, config)
				continue
			}
			if mergedConfigs[index].ConfigMap == nil {
				mergedConfigs[index].ConfigMap = config.ConfigMap
			}
		}
		return mergedConfigs
	}

	// Merge metadata
	ictrlutil.MergeMetadataMapInplace(compProto.Annotations, &compObjCopy.Annotations)
	ictrlutil.MergeMetadataMapInplace(compProto.Labels, &compObjCopy.Labels)

	// Merge all spec fields
	compObjCopy.Spec.TerminationPolicy = compProto.Spec.TerminationPolicy
	compObjCopy.Spec.CompDef = compProto.Spec.CompDef
	compObjCopy.Spec.ServiceVersion = compProto.Spec.ServiceVersion
	compObjCopy.Spec.ServiceRefs = compProto.Spec.ServiceRefs
	compObjCopy.Spec.Labels = compProto.Spec.Labels
	compObjCopy.Spec.Annotations = compProto.Spec.Annotations
	compObjCopy.Spec.Env = compProto.Spec.Env
	compObjCopy.Spec.VolumeClaimTemplates = compProto.Spec.VolumeClaimTemplates
	compObjCopy.Spec.PersistentVolumeClaimRetentionPolicy = compProto.Spec.PersistentVolumeClaimRetentionPolicy
	compObjCopy.Spec.Volumes = compProto.Spec.Volumes
	compObjCopy.Spec.Network = compProto.Spec.Network
	compObjCopy.Spec.Services = compProto.Spec.Services
	compObjCopy.Spec.SystemAccounts = compProto.Spec.SystemAccounts
	compObjCopy.Spec.Replicas = compProto.Spec.Replicas
	compObjCopy.Spec.Configs = mergeConfigs(compObjCopy.Spec.Configs, compProto.Spec.Configs)
	compObjCopy.Spec.ServiceAccountName = compProto.Spec.ServiceAccountName
	compObjCopy.Spec.ParallelPodManagementConcurrency = compProto.Spec.ParallelPodManagementConcurrency
	compObjCopy.Spec.PodUpdatePolicy = compProto.Spec.PodUpdatePolicy
	compObjCopy.Spec.InstanceUpdateStrategy = compProto.Spec.InstanceUpdateStrategy
	compObjCopy.Spec.SchedulingPolicy = compProto.Spec.SchedulingPolicy
	compObjCopy.Spec.TLSConfig = compProto.Spec.TLSConfig
	compObjCopy.Spec.Instances = compProto.Spec.Instances
	compObjCopy.Spec.FlatInstanceOrdinal = compProto.Spec.FlatInstanceOrdinal
	compObjCopy.Spec.OfflineInstances = compProto.Spec.OfflineInstances
	compObjCopy.Spec.RuntimeClassName = compProto.Spec.RuntimeClassName
	compObjCopy.Spec.DisableExporter = compProto.Spec.DisableExporter
	compObjCopy.Spec.Stop = compProto.Spec.Stop
	compObjCopy.Spec.Sidecars = compProto.Spec.Sidecars
	compObjCopy.Spec.Resources = compProto.Spec.Resources

	metadataChanged := !reflect.DeepEqual(oldCompObj.Annotations, compObjCopy.Annotations) ||
		!reflect.DeepEqual(oldCompObj.Labels, compObjCopy.Labels)
	specChanged := !reflect.DeepEqual(normalize(oldCompObj.Spec), normalize(compObjCopy.Spec))

	// If nothing changed after normalization, return nil
	if !metadataChanged && !specChanged {
		return nil
	}

	return compObjCopy
}

const (
	createOp int = 0
	deleteOp int = 1
	updateOp int = 2
)

func newCompNShardingHandler(transCtx *clusterTransformContext, op int) clusterConditionalHandler {
	topology, orders := definedOrders(transCtx, op)
	if len(orders) == 0 {
		return newParallelHandler(op)
	}
	return newOrderedHandler(topology, orders, op)
}

func definedOrders(transCtx *clusterTransformContext, op int) (appsv1.ClusterTopology, []string) {
	var (
		cluster    = transCtx.Cluster
		clusterDef = transCtx.clusterDef
	)
	if len(cluster.Spec.Topology) != 0 && clusterDef != nil {
		for _, topology := range clusterDef.Spec.Topologies {
			if topology.Name == cluster.Spec.Topology {
				if topology.Orders != nil {
					switch op {
					case createOp:
						return topology, topology.Orders.Provision
					case deleteOp:
						return topology, topology.Orders.Terminate
					case updateOp:
						return topology, topology.Orders.Update
					default:
						panic("runtime error: unknown operation: " + strconv.Itoa(op))
					}
				}
			}
		}
	}
	return appsv1.ClusterTopology{}, nil
}

func newParallelHandler(op int) clusterConditionalHandler {
	if op == createOp || op == deleteOp || op == updateOp {
		return &clusterParallelHandler{
			clusterCompNShardingHandler: clusterCompNShardingHandler{op: op},
		}
	}
	panic("runtime error: unknown operation: " + strconv.Itoa(op))
}

func newOrderedHandler(topology appsv1.ClusterTopology, orders []string, op int) clusterConditionalHandler {
	switch op {
	case createOp, updateOp:
		return &orderedCreateNUpdateHandler{
			clusterOrderedOrder: clusterOrderedOrder{
				topology: topology,
				orders:   orders,
			},
			phasePrecondition: phasePrecondition{
				topology:       topology,
				orders:         orders,
				ignoreNotExist: op == updateOp,
			},
			clusterCompNShardingHandler: clusterCompNShardingHandler{op: op},
		}
	case deleteOp:
		return &orderedDeleteHandler{
			clusterOrderedOrder: clusterOrderedOrder{
				topology: topology,
				orders:   orders,
			},
			notExistPrecondition: notExistPrecondition{
				topology: topology,
				orders:   orders,
			},
			clusterCompNShardingHandler: clusterCompNShardingHandler{op: op},
		}
	default:
		panic("runtime error: unknown operation: " + strconv.Itoa(op))
	}
}

type clusterConditionalHandler interface {
	ordered([]string) ([]string, error)
	match(transCtx *clusterTransformContext, dag *graph.DAG, name string) (bool, error)
	handle(transCtx *clusterTransformContext, dag *graph.DAG, name string) error
}

type clusterParallelOrder struct{}

func (o *clusterParallelOrder) ordered(names []string) ([]string, error) {
	return names, nil
}

type clusterOrderedOrder struct {
	topology appsv1.ClusterTopology
	orders   []string
}

func (o *clusterOrderedOrder) ordered(names []string) ([]string, error) {
	result := make([]string, 0)
	for _, order := range o.orders {
		entities := strings.Split(order, ",")
		for _, name := range names {
			if slices.ContainsFunc(entities, func(e string) bool {
				return clusterTopologyEntityMatched(o.topology, e, name)
			}) {
				result = append(result, name)
			}
		}
	}
	if len(result) != len(names) {
		return nil, fmt.Errorf("cannot find order for components and shardings: %s", strings.Join(names, ","))
	}
	return result, nil
}

type dummyPrecondition struct{}

func (c *dummyPrecondition) match(*clusterTransformContext, *graph.DAG, string) (bool, error) {
	return true, nil
}

type notExistPrecondition struct {
	topology appsv1.ClusterTopology
	orders   []string
}

func (c *notExistPrecondition) match(transCtx *clusterTransformContext, dag *graph.DAG, name string) (bool, error) {
	predecessors, err := predecessors(c.topology, c.orders, name)
	if err != nil {
		return false, err
	}
	for _, predecessor := range predecessors {
		exist, err := c.predecessorExist(transCtx, dag, predecessor)
		if err != nil {
			return false, err
		}
		if exist {
			return false, nil
		}
	}
	return true, nil
}

func (c *notExistPrecondition) predecessorExist(transCtx *clusterTransformContext, dag *graph.DAG, name string) (bool, error) {
	isSharding, err := transCtx.sharding(name)
	if err != nil {
		return false, err
	}
	if isSharding {
		return c.shardingExist(transCtx, dag, name)
	}
	return c.compExist(transCtx, dag, name)
}

func (c *notExistPrecondition) compExist(transCtx *clusterTransformContext, dag *graph.DAG, name string) (bool, error) {
	var (
		compKey = types.NamespacedName{
			Namespace: transCtx.Cluster.Namespace,
			Name:      component.FullName(transCtx.Cluster.Name, name),
		}
	)
	get := func() (bool, error) {
		comp := &appsv1.Component{}
		err := transCtx.Client.Get(transCtx.Context, compKey, comp)
		if err != nil && !apierrors.IsNotFound(err) {
			return false, err
		}
		return err == nil, nil
	}
	dagCreate := func() bool {
		graphCli, _ := transCtx.Client.(model.GraphClient)
		comp := &appsv1.Component{
			ObjectMeta: metav1.ObjectMeta{
				Namespace: compKey.Namespace,
				Name:      compKey.Name,
			},
		}
		return graphCli.IsAction(dag, comp, model.ActionCreatePtr())
	}

	exist, err := get()
	if err != nil {
		return false, err
	}
	if exist {
		return true, nil
	}
	if dagCreate() {
		return true, nil
	}
	return false, nil
}

func (c *notExistPrecondition) shardingExist(transCtx *clusterTransformContext, dag *graph.DAG, name string) (bool, error) {
	list := func() (bool, error) {
		comps, err := sharding.ListShardingComponents(transCtx.Context, transCtx.Client, transCtx.Cluster, name)
		if err != nil {
			return false, err
		}
		return len(comps) > 0, nil
	}
	dagCreate := func() bool {
		graphCli, _ := transCtx.Client.(model.GraphClient)
		for _, obj := range graphCli.FindAll(dag, &appsv1.Component{}) {
			if shardingCompWithName(obj.(*appsv1.Component), name) &&
				graphCli.IsAction(dag, obj, model.ActionCreatePtr()) {
				return true
			}
		}
		return false
	}

	exist, err := list()
	if err != nil {
		return false, err
	}
	if exist {
		return true, nil
	}
	if dagCreate() {
		return true, nil
	}
	return false, nil
}

type phasePrecondition struct {
	topology       appsv1.ClusterTopology
	orders         []string
	ignoreNotExist bool
}

func (c *phasePrecondition) match(transCtx *clusterTransformContext, dag *graph.DAG, name string) (bool, error) {
	predecessors, err := predecessors(c.topology, c.orders, name)
	if err != nil {
		return false, err
	}
	for _, predecessor := range predecessors {
		match, err := c.predecessorMatch(transCtx, dag, predecessor)
		if err != nil {
			return false, err
		}
		if !match {
			return false, nil
		}
	}
	return true, nil
}

func (c *phasePrecondition) predecessorMatch(transCtx *clusterTransformContext, dag *graph.DAG, name string) (bool, error) {
	isSharding, err := transCtx.sharding(name)
	if err != nil {
		return false, err
	}
	if isSharding {
		return c.shardingMatch(transCtx, dag, name)
	}
	return c.compMatch(transCtx, dag, name)
}

func (c *phasePrecondition) compMatch(transCtx *clusterTransformContext, dag *graph.DAG, name string) (bool, error) {
	var (
		compKey = types.NamespacedName{
			Namespace: transCtx.Cluster.Namespace,
			Name:      component.FullName(transCtx.Cluster.Name, name),
		}
	)
	dagGet := func() bool {
		graphCli, _ := transCtx.Client.(model.GraphClient)
		for _, obj := range graphCli.FindAll(dag, &appsv1.Component{}) {
			if client.ObjectKeyFromObject(obj) == compKey {
				return true // TODO: should check the action?
			}
		}
		return false
	}

	comp := &appsv1.Component{}
	if err := transCtx.Client.Get(transCtx.Context, compKey, comp); err != nil {
		return c.ignoreNotExist, client.IgnoreNotFound(err)
	}
	if !c.expected(comp) {
		transCtx.Logger.Info("waiting for predecessor component in expected phase",
			"component", comp.Name, "predecessor", name)
		return false, nil
	}
	// create or update in DAG?
	if dagGet() {
		transCtx.Logger.Info("waiting for predecessor component in DAG",
			"component", comp.Name, "predecessor", name)
		return false, nil
	}
	return true, nil
}

func (c *phasePrecondition) shardingMatch(transCtx *clusterTransformContext, dag *graph.DAG, name string) (bool, error) {
	dagList := func() bool {
		graphCli, _ := transCtx.Client.(model.GraphClient)
		for _, obj := range graphCli.FindAll(dag, &appsv1.Component{}) {
			if shardingCompWithName(obj.(*appsv1.Component), name) {
				return true // TODO: should check the action?
			}
		}
		return false
	}

	protoComps, ok := transCtx.shardingComps[name]
	if !ok {
		return false, fmt.Errorf("cluster sharding %s not found", name)
	}

	comps, err := sharding.ListShardingComponents(transCtx.Context, transCtx.Client, transCtx.Cluster, name)
	if err != nil {
		return false, err
	}
	if len(comps) != len(protoComps) {
		if len(comps) == 0 {
			return c.ignoreNotExist, nil
		}
		return false, nil
	}
	for _, comp := range comps {
		if !c.expected(&comp) {
			transCtx.Logger.Info("waiting for predecessor sharding in expected phase",
				"shard", comp.Name, "predecessor sharding", name)
			return false, nil
		}
	}
	// create or update in DAG?
	if dagList() {
		transCtx.Logger.Info("waiting for predecessor sharding in DAG",
			"shards", comps, "predecessor sharding", name)
		return false, nil
	}
	return true, nil
}

func (c *phasePrecondition) expected(comp *appsv1.Component) bool {
	if comp.Generation == comp.Status.ObservedGeneration {
		expect := appsv1.RunningComponentPhase
		if comp.Spec.Stop != nil && *comp.Spec.Stop {
			expect = appsv1.StoppedComponentPhase
		}
		return comp.Status.Phase == expect
	}
	return false
}

type clusterCompNShardingHandler struct {
	op      int
	scaleIn *bool
}

func (h *clusterCompNShardingHandler) handle(transCtx *clusterTransformContext, dag *graph.DAG, name string) error {
	isSharding, err := transCtx.sharding(name)
	if err != nil {
		return err
	}
	if isSharding {
		handler := &clusterShardingHandler{scaleIn: h.scaleIn}
		switch h.op {
		case createOp:
			return handler.create(transCtx, dag, name)
		case deleteOp:
			return handler.delete(transCtx, dag, name)
		default:
			return handler.update(transCtx, dag, name)
		}
	} else {
		handler := &clusterComponentHandler{}
		switch h.op {
		case createOp:
			return handler.create(transCtx, dag, name)
		case deleteOp:
			return handler.delete(transCtx, dag, name)
		default:
			return handler.update(transCtx, dag, name)
		}
	}
}

func predecessors(topology appsv1.ClusterTopology, orders []string, name string) ([]string, error) {
	var previous []string
	for _, order := range orders {
		entities := strings.Split(order, ",")
		if slices.ContainsFunc(entities, func(e string) bool {
			return clusterTopologyEntityMatched(topology, e, name)
		}) {
			return previous, nil
		}
		previous = entities
	}
	return nil, fmt.Errorf("cannot find predecessor for component or sharding %s", name)
}

func clusterTopologyEntityMatched(topology appsv1.ClusterTopology, entityName, name string) bool {
	for _, sharding := range topology.Shardings {
		if sharding.Name == entityName {
			return entityName == name // full match for sharding
		}
	}
	for _, comp := range topology.Components {
		if comp.Name == entityName {
			return clusterTopologyCompMatched(comp, name)
		}
	}
	return false // TODO: runtime error
}

type clusterParallelHandler struct {
	clusterParallelOrder
	dummyPrecondition
	clusterCompNShardingHandler
}

type orderedCreateNUpdateHandler struct {
	clusterOrderedOrder
	phasePrecondition
	clusterCompNShardingHandler
}

type orderedDeleteHandler struct {
	clusterOrderedOrder
	notExistPrecondition
	clusterCompNShardingHandler
}

func setDiff(s1, s2 sets.Set[string]) (sets.Set[string], sets.Set[string], sets.Set[string]) {
	return s2.Difference(s1), s1.Difference(s2), s1.Intersection(s2)
}

func mapDiff[T interface{}](m1, m2 map[string]T) (sets.Set[string], sets.Set[string], sets.Set[string]) {
	s1, s2 := sets.KeySet(m1), sets.KeySet(m2)
	return setDiff(s1, s2)
}

type clusterComponentHandler struct{}

func (h *clusterComponentHandler) create(transCtx *clusterTransformContext, dag *graph.DAG, name string) error {
	proto, err := h.protoComp(transCtx, name, nil)
	if err != nil {
		return err
	}
	graphCli, _ := transCtx.Client.(model.GraphClient)
	graphCli.Create(dag, proto)

	// initClusterCompNShardingStatus(transCtx, name)

	return nil
}

func (h *clusterComponentHandler) delete(transCtx *clusterTransformContext, dag *graph.DAG, name string) error {
	comp, err := h.runningComp(transCtx, name)
	if err != nil && !apierrors.IsNotFound(err) {
		return err
	}
	if apierrors.IsNotFound(err) || model.IsObjectDeleting(comp) {
		return nil
	}

	transCtx.Logger.Info(fmt.Sprintf("deleting component %s", comp.Name))
	graphCli, _ := transCtx.Client.(model.GraphClient)
	graphCli.Delete(dag, comp)

	return nil
}

func (h *clusterComponentHandler) update(transCtx *clusterTransformContext, dag *graph.DAG, name string) error {
	running, err1 := h.runningComp(transCtx, name)
	if err1 != nil {
		return err1
	}
	proto, err2 := h.protoComp(transCtx, name, running)
	if err2 != nil {
		return err2
	}

	if obj := copyAndMergeComponent(running, proto); obj != nil {
		graphCli, _ := transCtx.Client.(model.GraphClient)
		graphCli.Update(dag, running, obj)
	}
	return nil
}

func (h *clusterComponentHandler) runningComp(transCtx *clusterTransformContext, name string) (*appsv1.Component, error) {
	compKey := types.NamespacedName{
		Namespace: transCtx.Cluster.Namespace,
		Name:      component.FullName(transCtx.Cluster.Name, name),
	}
	comp := &appsv1.Component{}
	if err := transCtx.Client.Get(transCtx.Context, compKey, comp); err != nil {
		return nil, err
	}
	return comp, nil
}

func (h *clusterComponentHandler) protoComp(transCtx *clusterTransformContext, name string, running *appsv1.Component) (*appsv1.Component, error) {
	for _, comp := range transCtx.components {
		if comp.Name == name {
			return buildComponentWrapper(transCtx, comp, nil, nil, running)
		}
	}
	return nil, fmt.Errorf("cluster component %s not found", name)
}

type clusterShardingHandler struct {
	scaleIn *bool
}

func (h *clusterShardingHandler) create(transCtx *clusterTransformContext, dag *graph.DAG, name string) error {
	protoComps, err := h.protoComps(transCtx, name, nil)
	if err != nil {
		return err
	}
	graphCli, _ := transCtx.Client.(model.GraphClient)
	for i := range protoComps {
		graphCli.Create(dag, protoComps[i])
	}

	// initClusterCompNShardingStatus(transCtx, name)

	// TODO:
	//  1. sharding post-provision
	//  2. provision strategy

	return nil
}

// delete handles the sharding component deletion when cluster is Deleting
func (h *clusterShardingHandler) delete(transCtx *clusterTransformContext, dag *graph.DAG, name string) error {
	runningComps, err := sharding.ListShardingComponents(transCtx.Context, transCtx.Client, transCtx.Cluster, name)
	if err != nil {
		return err
	}

	// TODO: sharding pre-terminate

	graphCli, _ := transCtx.Client.(model.GraphClient)
	for i := range runningComps {
		h.deleteComp(transCtx, graphCli, dag, &runningComps[i], nil)
	}

	return nil
}

func (h *clusterShardingHandler) deleteComp(transCtx *clusterTransformContext,
	graphCli model.GraphClient, dag *graph.DAG, comp *appsv1.Component, scaleIn *bool) {
	if !model.IsObjectDeleting(comp) {
		transCtx.Logger.Info(fmt.Sprintf("deleting sharding component %s", comp.Name))

		vertex := graphCli.Do(dag, nil, comp, model.ActionDeletePtr(), nil)
		if scaleIn != nil && *scaleIn {
			compCopy := comp.DeepCopy()
			if comp.Annotations == nil {
				comp.Annotations = make(map[string]string)
			}
			comp.Annotations[constant.ComponentScaleInAnnotationKey] = "true"
			graphCli.Do(dag, compCopy, comp, model.ActionUpdatePtr(), vertex)
		}
	}
}

func (h *clusterShardingHandler) update(transCtx *clusterTransformContext, dag *graph.DAG, name string) error {
	runningComps, err1 := sharding.ListShardingComponents(transCtx.Context, transCtx.Client, transCtx.Cluster, name)
	if err1 != nil {
		return err1
	}

	runningCompsMap := make(map[string]*appsv1.Component)
	for i, comp := range runningComps {
		runningCompsMap[comp.Name] = &runningComps[i]
	}

	var running *appsv1.Component
	if len(runningComps) > 0 {
		running = &runningComps[0]
	}
	protoComps, err2 := h.protoComps(transCtx, name, running)
	if err2 != nil {
		return err2
	}
	protoCompsMap := make(map[string]*appsv1.Component)
	for i, comp := range protoComps {
		protoCompsMap[comp.Name] = protoComps[i]
	}

	toCreate, toDelete, toUpdate := mapDiff(runningCompsMap, protoCompsMap)

	// TODO: update strategy
	h.deleteComps(transCtx, dag, runningCompsMap, toDelete)
	h.updateComps(transCtx, dag, runningCompsMap, protoCompsMap, toUpdate)
	h.createComps(transCtx, dag, protoCompsMap, toCreate)

	return nil
}

func (h *clusterShardingHandler) createComps(transCtx *clusterTransformContext, dag *graph.DAG,
	protoComps map[string]*appsv1.Component, createSet sets.Set[string]) {
	graphCli, _ := transCtx.Client.(model.GraphClient)
	for name := range createSet {
		graphCli.Create(dag, protoComps[name])
		// TODO: shard post-provision
	}
}

// deleteComps deletes the subcomponents of the sharding when the shards count is updated.
func (h *clusterShardingHandler) deleteComps(transCtx *clusterTransformContext, dag *graph.DAG,
	runningComps map[string]*appsv1.Component, deleteSet sets.Set[string]) {
	graphCli, _ := transCtx.Client.(model.GraphClient)
	for name := range deleteSet {
		h.deleteComp(transCtx, graphCli, dag, runningComps[name], pointer.Bool(true))
	}
}

func (h *clusterShardingHandler) updateComps(transCtx *clusterTransformContext, dag *graph.DAG,
	runningComps map[string]*appsv1.Component, protoComps map[string]*appsv1.Component, updateSet sets.Set[string]) {
	graphCli, _ := transCtx.Client.(model.GraphClient)
	for name := range updateSet {
		running, proto := runningComps[name], protoComps[name]
		if obj := copyAndMergeComponent(running, proto); obj != nil {
			graphCli.Update(dag, running, obj)
		}
	}
}

func (h *clusterShardingHandler) protoComps(transCtx *clusterTransformContext, name string, running *appsv1.Component) ([]*appsv1.Component, error) {
	for _, sharding := range transCtx.shardings {
		if sharding.Name == name {
			return h.buildComps(transCtx, sharding, running)
		}
	}
	return nil, fmt.Errorf("cluster sharding %s not found", name)
}

func (h *clusterShardingHandler) buildComps(transCtx *clusterTransformContext,
	sharding *appsv1.ClusterSharding, running *appsv1.Component) ([]*appsv1.Component, error) {
	objs := make([]*appsv1.Component, 0)
	shardingComps := transCtx.shardingCompsWithTpl[sharding.Name]
	for tplName, tplComps := range shardingComps {
		for i := range tplComps {
			spec := shardingComps[tplName][i]
			labels := h.buildLabels(sharding, tplName)
			annotations := h.buildAnnotations(transCtx, sharding.Name, spec.Name)
			obj, err := buildComponentWrapper(transCtx, spec, labels, annotations, running)
			if err != nil {
				return nil, err
			}
			h.buildShardPodAntiAffinity(transCtx, sharding.Name, spec.Name, obj)
			objs = append(objs, obj)
		}
	}
	return objs, nil
}

func (h *clusterShardingHandler) buildLabels(sharding *appsv1.ClusterSharding, shardTplName string) map[string]string {
	labels := map[string]string{
		constant.KBAppShardingNameLabelKey: sharding.Name,
	}
	if len(sharding.ShardingDef) > 0 {
		labels[constant.ShardingDefLabelKey] = sharding.ShardingDef
	}
	if len(shardTplName) > 0 {
		labels[constant.KBAppShardTemplateLabelKey] = shardTplName
	}
	return labels
}

func (h *clusterShardingHandler) buildAnnotations(transCtx *clusterTransformContext, shardingName, compName string) map[string]string {
	var annotations map[string]string
	if compAnnotations := transCtx.annotations[compName]; len(compAnnotations) > 0 {
		annotations = maps.Clone(compAnnotations)
	}

	// convert the sharding hostNetwork annotation to the component annotation
	if hnKey, ok := transCtx.Cluster.Annotations[constant.HostNetworkAnnotationKey]; ok {
		hns := strings.Split(hnKey, ",")
		if slices.Index(hns, shardingName) >= 0 {
			if annotations == nil {
				annotations = make(map[string]string)
			}
			annotations[constant.HostNetworkAnnotationKey] = compName
		}
	}
	return annotations
}

func (h *clusterShardingHandler) buildShardPodAntiAffinity(transCtx *clusterTransformContext,
	shardingName, compName string, comp *appsv1.Component) {
	var (
		cluster = transCtx.Cluster
	)
	key := cluster.Annotations[constant.ShardPodAntiAffinityAnnotationKey]
	if !slices.Contains(strings.Split(key, ","), shardingName) {
		return
	}

	shardPodAntiAffinity := corev1.PodAffinityTerm{
		LabelSelector: &metav1.LabelSelector{
			MatchLabels: constant.GetCompLabels(cluster.Name, compName,
				map[string]string{constant.KBAppShardingNameLabelKey: shardingName}),
		},
		TopologyKey: corev1.LabelHostname,
	}

	if comp.Spec.SchedulingPolicy == nil {
		comp.Spec.SchedulingPolicy = &appsv1.SchedulingPolicy{}
	}
	if comp.Spec.SchedulingPolicy.Affinity == nil {
		comp.Spec.SchedulingPolicy.Affinity = &corev1.Affinity{}
	}
	if comp.Spec.SchedulingPolicy.Affinity.PodAntiAffinity == nil {
		comp.Spec.SchedulingPolicy.Affinity.PodAntiAffinity = &corev1.PodAntiAffinity{}
	}
	if comp.Spec.SchedulingPolicy.Affinity.PodAntiAffinity.RequiredDuringSchedulingIgnoredDuringExecution == nil {
		comp.Spec.SchedulingPolicy.Affinity.PodAntiAffinity.RequiredDuringSchedulingIgnoredDuringExecution = []corev1.PodAffinityTerm{}
	}
	comp.Spec.SchedulingPolicy.Affinity.PodAntiAffinity.RequiredDuringSchedulingIgnoredDuringExecution = append(
		comp.Spec.SchedulingPolicy.Affinity.PodAntiAffinity.RequiredDuringSchedulingIgnoredDuringExecution, shardPodAntiAffinity)
}

func clusterRunningCompNShardingSet(ctx context.Context, cli client.Reader, cluster *appsv1.Cluster) (sets.Set[string], error) {
	compList := &appsv1.ComponentList{}
	ml := client.MatchingLabels{constant.AppInstanceLabelKey: cluster.Name}
	if err := cli.List(ctx, compList, client.InNamespace(cluster.Namespace), ml); err != nil {
		return nil, err
	}

	names := sets.Set[string]{}
	for _, comp := range compList.Items {
		if shardingName := shardingCompNName(&comp); len(shardingName) > 0 {
			names.Insert(shardingName)
		} else {
			name, err := component.ShortName(cluster.Name, comp.Name)
			if err != nil {
				return nil, err
			}
			names.Insert(name)
		}
	}
	return names, nil
}

func shardingCompWithName(comp *appsv1.Component, shardingName string) bool {
	if comp == nil || comp.Labels == nil {
		return false
	}
	name, ok := comp.Labels[constant.KBAppShardingNameLabelKey]
	return ok && name == shardingName
}

func shardingCompNName(comp *appsv1.Component) string {
	if comp != nil && comp.Labels != nil {
		name, ok := comp.Labels[constant.KBAppShardingNameLabelKey]
		if ok {
			return name
		}
	}
	return ""
}

func buildComponentWrapper(transCtx *clusterTransformContext,
	spec *appsv1.ClusterComponentSpec, labels, annotations map[string]string, running *appsv1.Component) (*appsv1.Component, error) {
	// cluster.spec.components[*] has no sidecars defined, so we need to build sidecars for it here
	comp, err := component.BuildComponent(transCtx.Cluster, spec, labels, annotations)
	if err != nil {
		return nil, err
	}
	if err = buildComponentSidecars(transCtx, comp, running); err != nil {
		return nil, err
	}
	return comp, nil
}

func buildComponentSidecars(transCtx *clusterTransformContext, proto, running *appsv1.Component) error {
	// component definitions used by all components and shardings of the cluster
	compDefs := func() sets.Set[string] {
		defs := sets.New[string]()
		for _, spec := range transCtx.components {
			defs.Insert(spec.ComponentDef)
		}
		for _, spec := range transCtx.shardings {
			defs.Insert(spec.Template.ComponentDef)
		}
		return defs
	}()

	sidecars, err := hostedSidecarsOfCompDef(transCtx.Context, transCtx.Client, compDefs, proto.Spec.CompDef)
	if err != nil {
		return err
	}
	if len(sidecars) > 0 {
		for name, sidecar := range sidecars {
			if err = buildComponentSidecar(proto, running, name, sidecar); err != nil {
				return err
			}
		}
		// keep the sidecars ordered
		slices.SortFunc(proto.Spec.Sidecars, func(a, b appsv1.Sidecar) int {
			return strings.Compare(a.Name, b.Name)
		})
	}
	return nil
}

func hostedSidecarsOfCompDef(ctx context.Context, cli client.Reader, compDefs sets.Set[string], compDef string) (map[string][]any, error) {
	sidecarList := &appsv1.SidecarDefinitionList{}
	if err := cli.List(ctx, sidecarList); err != nil {
		return nil, err
	}

	match := func(sidecarDef *appsv1.SidecarDefinition) (any, error) {
		owners := sets.New(strings.Split(sidecarDef.Status.Owners, ",")...)
		selectors := sets.New(strings.Split(sidecarDef.Status.Selectors, ",")...)

		owned := compDefs.Intersection(owners)
		if len(owned) == 0 {
			return nil, nil
		}
		selected := compDefs.Intersection(selectors)
		if len(selected) == 0 {
			return nil, fmt.Errorf("no comp-def selected by sidecar definition: %s", sidecarDef.Name)
		}
		// double check
		if selected.Intersection(owned).Len() != 0 {
			return nil, fmt.Errorf("owner and selectors should not be overlapped: %s",
				strings.Join(sets.List(selected.Intersection(owned)), ","))
		}
		if !selected.Has(compDef) {
			return nil, nil // it's not me
		}
		ownerList := sets.List(owned)
		slices.SortFunc(ownerList, func(a, b string) int {
			return strings.Compare(a, b) * -1
		})
		// tuple<sidecarDef, owners>
		return []any{sidecarDef, ownerList}, nil
	}

	// sidecarName -> []tuple<sidecarDef, owners>
	result := make(map[string][]any)
	for i, sidecarDef := range sidecarList.Items {
		matched, err := match(&sidecarList.Items[i])
		if err != nil {
			return nil, err
		}
		if matched != nil {
			sidecars, ok := result[sidecarDef.Spec.Name]
			if !ok {
				result[sidecarDef.Spec.Name] = []any{matched}
			} else {
				result[sidecarDef.Spec.Name] = append(sidecars, matched)
			}
		}
	}

	for name := range result {
		sidecars := result[name]
		// ordered by sidecarDef.Name from latest to oldest
		slices.SortFunc(sidecars, func(a1, a2 any) int {
			sidecarDef1 := a1.([]any)[0].(*appsv1.SidecarDefinition)
			sidecarDef2 := a2.([]any)[0].(*appsv1.SidecarDefinition)
			return strings.Compare(sidecarDef1.Name, sidecarDef2.Name) * -1
		})
		result[name] = sidecars
	}
	return result, nil
}

func buildComponentSidecar(proto, running *appsv1.Component, sidecarName string, ctx []any) error {
	exist := func() int {
		if running == nil {
			return -1
		}
		return slices.IndexFunc(running.Spec.Sidecars, func(s appsv1.Sidecar) bool {
			return s.Name == sidecarName
		})
	}()

	checkedAppend := func(sidecar appsv1.Sidecar, sidecarDef *appsv1.SidecarDefinition) error {
		if sidecarDef.Generation != sidecarDef.Status.ObservedGeneration {
			return fmt.Errorf("the SidecarDefinition is not up to date: %s", sidecarDef.Name)
		}
		if sidecarDef.Status.Phase != appsv1.AvailablePhase {
			return fmt.Errorf("the SidecarDefinition is unavailable: %s", sidecarDef.Name)
		}
		if proto.Spec.Sidecars == nil {
			proto.Spec.Sidecars = make([]appsv1.Sidecar, 0)
		}
		proto.Spec.Sidecars = append(proto.Spec.Sidecars, sidecar)
		if proto.Annotations == nil {
			proto.Annotations = make(map[string]string)
		}
		proto.Annotations[constant.SidecarDefLabelKey] = sidecar.SidecarDef
		return nil
	}

	if exist >= 0 {
		sidecar := running.Spec.Sidecars[exist]
		for _, a := range ctx {
			sidecarDef, owners := a.([]any)[0].(*appsv1.SidecarDefinition), a.([]any)[1].([]string)
			if sidecar.SidecarDef == sidecarDef.Name && slices.Contains(owners, sidecar.Owner) {
				// has the fully matched owner comp-def and sidecar def, use it directly
				return checkedAppend(sidecar, sidecarDef)
			}
		}
	}

	// otherwise, use the latest one, new created or upgraded
	sidecarDef := ctx[0].([]any)[0].(*appsv1.SidecarDefinition)
	sidecar := appsv1.Sidecar{
		Name:       sidecarDef.Spec.Name,
		Owner:      ctx[0].([]any)[1].([]string)[0],
		SidecarDef: sidecarDef.Name,
	}
	return checkedAppend(sidecar, sidecarDef)
}
