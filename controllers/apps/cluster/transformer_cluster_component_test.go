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
	"fmt"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/onsi/gomega/types"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/utils/pointer"
	"sigs.k8s.io/controller-runtime/pkg/client"

	appsv1 "github.com/apecloud/kubeblocks/apis/apps/v1"
	appsutil "github.com/apecloud/kubeblocks/controllers/apps/util"
	"github.com/apecloud/kubeblocks/pkg/constant"
	"github.com/apecloud/kubeblocks/pkg/controller/component"
	"github.com/apecloud/kubeblocks/pkg/controller/graph"
	"github.com/apecloud/kubeblocks/pkg/controller/model"
	testapps "github.com/apecloud/kubeblocks/pkg/testutil/apps"
)

var _ = Describe("cluster component transformer test", func() {
	const (
		clusterDefName                              = "test-clusterdef"
		clusterTopologyDefault                      = "test-topology-default"
		clusterTopologyNoOrders                     = "test-topology-no-orders"
		clusterTopologyProvisionNUpdateOOD          = "test-topology-ood"
		clusterTopologyStop                         = "test-topology-stop"
		clusterTopologyTemplate                     = "test-topology-template"
		clusterTopologyDefault4Sharding             = "test-topology-default-sharding"
		clusterTopologyNoOrders4Sharding            = "test-topology-no-orders-sharding"
		clusterTopologyProvisionNUpdateOOD4Sharding = "test-topology-ood-sharding"
		clusterTopologyStop4Sharding                = "test-topology-stop-sharding"
		clusterTopologyCompNSharding                = "test-topology-comp-sharding"
		clusterTopologyShardingNComp                = "test-topology-sharding-comp"
		clusterTopologyCompNShardingOOD             = "test-topology-ood-comp-sharding"
		clusterTopologyShardingNCompOOD             = "test-topology-ood-sharding-comp"
		compDefName                                 = "test-compdef"
		shardingDefName                             = "test-shardingdef"
		clusterName                                 = "test-cluster"
		comp1aName                                  = "comp-1a"
		comp1bName                                  = "comp-1b"
		comp2aName                                  = "comp-2a"
		comp2bName                                  = "comp-2b"
		comp3aName                                  = "comp-3a"
		sharding1aName                              = "sharding-1a"
		sharding1bName                              = "sharding-1b"
		sharding2aName                              = "sharding-2a"
		sharding2bName                              = "sharding-2b"
		sharding3aName                              = "sharding-3a"
	)

	var (
		clusterDef *appsv1.ClusterDefinition
	)

	BeforeEach(func() {
		clusterDef = testapps.NewClusterDefFactory(clusterDefName).
			AddClusterTopology(appsv1.ClusterTopology{
				Name: clusterTopologyDefault,
				Components: []appsv1.ClusterTopologyComponent{
					{
						Name:    comp1aName,
						CompDef: compDefName,
					},
					{
						Name:    comp1bName,
						CompDef: compDefName,
					},
					{
						Name:    comp2aName,
						CompDef: compDefName,
					},
					{
						Name:    comp2bName,
						CompDef: compDefName,
					},
				},
				Orders: &appsv1.ClusterTopologyOrders{
					Provision: []string{
						fmt.Sprintf("%s,%s", comp1aName, comp1bName),
						fmt.Sprintf("%s,%s", comp2aName, comp2bName),
					},
					Terminate: []string{
						fmt.Sprintf("%s,%s", comp2aName, comp2bName),
						fmt.Sprintf("%s,%s", comp1aName, comp1bName),
					},
					Update: []string{
						fmt.Sprintf("%s,%s", comp1aName, comp1bName),
						fmt.Sprintf("%s,%s", comp2aName, comp2bName),
					},
				},
			}).
			AddClusterTopology(appsv1.ClusterTopology{
				Name: clusterTopologyNoOrders,
				Components: []appsv1.ClusterTopologyComponent{
					{
						Name:    comp1aName,
						CompDef: compDefName,
					},
					{
						Name:    comp1bName,
						CompDef: compDefName,
					},
					{
						Name:    comp2aName,
						CompDef: compDefName,
					},
					{
						Name:    comp2bName,
						CompDef: compDefName,
					},
				},
			}).
			AddClusterTopology(appsv1.ClusterTopology{
				Name: clusterTopologyProvisionNUpdateOOD,
				Components: []appsv1.ClusterTopologyComponent{
					{
						Name:    comp1aName,
						CompDef: compDefName,
					},
					{
						Name:    comp1bName,
						CompDef: compDefName,
					},
					{
						Name:    comp2aName,
						CompDef: compDefName,
					},
					{
						Name:    comp2bName,
						CompDef: compDefName,
					},
				},
				Orders: &appsv1.ClusterTopologyOrders{
					Provision: []string{
						fmt.Sprintf("%s,%s", comp1aName, comp1bName),
						fmt.Sprintf("%s,%s", comp2aName, comp2bName),
					},
					Update: []string{
						fmt.Sprintf("%s,%s", comp2aName, comp2bName),
						fmt.Sprintf("%s,%s", comp1aName, comp1bName),
					},
				},
			}).
			AddClusterTopology(appsv1.ClusterTopology{
				Name: clusterTopologyStop,
				Components: []appsv1.ClusterTopologyComponent{
					{
						Name:    comp1aName,
						CompDef: compDefName,
					},
					{
						Name:    comp2aName,
						CompDef: compDefName,
					},
					{
						Name:    comp3aName,
						CompDef: compDefName,
					},
				},
				Orders: &appsv1.ClusterTopologyOrders{
					Update: []string{comp1aName, comp2aName, comp3aName},
				},
			}).
			AddClusterTopology(appsv1.ClusterTopology{
				Name: clusterTopologyTemplate,
				Components: []appsv1.ClusterTopologyComponent{
					{
						Name:    comp1aName,
						CompDef: compDefName,
					},
					{
						Name:     comp2aName,
						CompDef:  compDefName,
						Template: pointer.Bool(true),
					},
					{
						Name:    comp2bName,
						CompDef: compDefName,
					},
					{
						Name:    comp3aName,
						CompDef: compDefName,
					},
				},
				Orders: &appsv1.ClusterTopologyOrders{
					Provision: []string{comp1aName, fmt.Sprintf("%s,%s", comp2aName, comp2bName), comp3aName},
				},
			}).
			AddClusterTopology(appsv1.ClusterTopology{
				Name: clusterTopologyDefault4Sharding,
				Shardings: []appsv1.ClusterTopologySharding{
					{
						Name:        sharding1aName,
						ShardingDef: shardingDefName,
					},
					{
						Name:        sharding1bName,
						ShardingDef: shardingDefName,
					},
					{
						Name:        sharding2aName,
						ShardingDef: shardingDefName,
					},
					{
						Name:        sharding2bName,
						ShardingDef: shardingDefName,
					},
				},
				Orders: &appsv1.ClusterTopologyOrders{
					Provision: []string{
						fmt.Sprintf("%s,%s", sharding1aName, sharding1bName),
						fmt.Sprintf("%s,%s", sharding2aName, sharding2bName),
					},
					Terminate: []string{
						fmt.Sprintf("%s,%s", sharding2aName, sharding2bName),
						fmt.Sprintf("%s,%s", sharding1aName, sharding1bName),
					},
					Update: []string{
						fmt.Sprintf("%s,%s", sharding1aName, sharding1bName),
						fmt.Sprintf("%s,%s", sharding2aName, sharding2bName),
					},
				},
			}).
			AddClusterTopology(appsv1.ClusterTopology{
				Name: clusterTopologyNoOrders4Sharding,
				Shardings: []appsv1.ClusterTopologySharding{
					{
						Name:        sharding1aName,
						ShardingDef: shardingDefName,
					},
					{
						Name:        sharding1bName,
						ShardingDef: shardingDefName,
					},
					{
						Name:        sharding2aName,
						ShardingDef: shardingDefName,
					},
					{
						Name:        sharding2bName,
						ShardingDef: shardingDefName,
					},
				},
			}).
			AddClusterTopology(appsv1.ClusterTopology{
				Name: clusterTopologyProvisionNUpdateOOD4Sharding,
				Shardings: []appsv1.ClusterTopologySharding{
					{
						Name:        sharding1aName,
						ShardingDef: shardingDefName,
					},
					{
						Name:        sharding1bName,
						ShardingDef: shardingDefName,
					},
					{
						Name:        sharding2aName,
						ShardingDef: shardingDefName,
					},
					{
						Name:        sharding2bName,
						ShardingDef: shardingDefName,
					},
				},
				Orders: &appsv1.ClusterTopologyOrders{
					Provision: []string{
						fmt.Sprintf("%s,%s", sharding1aName, sharding1bName),
						fmt.Sprintf("%s,%s", sharding2aName, sharding2bName),
					},
					Update: []string{
						fmt.Sprintf("%s,%s", sharding2aName, sharding2bName),
						fmt.Sprintf("%s,%s", sharding1aName, sharding1bName),
					},
				},
			}).
			AddClusterTopology(appsv1.ClusterTopology{
				Name: clusterTopologyStop4Sharding,
				Shardings: []appsv1.ClusterTopologySharding{
					{
						Name:        sharding1aName,
						ShardingDef: shardingDefName,
					},
					{
						Name:        sharding2aName,
						ShardingDef: shardingDefName,
					},
					{
						Name:        sharding3aName,
						ShardingDef: shardingDefName,
					},
				},
				Orders: &appsv1.ClusterTopologyOrders{
					Update: []string{sharding1aName, sharding2aName, sharding3aName},
				},
			}).
			AddClusterTopology(appsv1.ClusterTopology{
				Name: clusterTopologyCompNSharding,
				Components: []appsv1.ClusterTopologyComponent{
					{
						Name:    comp1aName,
						CompDef: compDefName,
					},
				},
				Shardings: []appsv1.ClusterTopologySharding{
					{
						Name:        sharding1aName,
						ShardingDef: shardingDefName,
					},
				},
				Orders: &appsv1.ClusterTopologyOrders{
					Provision: []string{comp1aName, sharding1aName},
					Terminate: []string{sharding1aName, comp1aName},
					Update:    []string{comp1aName, sharding1aName},
				},
			}).
			AddClusterTopology(appsv1.ClusterTopology{
				Name: clusterTopologyShardingNComp,
				Components: []appsv1.ClusterTopologyComponent{
					{
						Name:    comp1aName,
						CompDef: compDefName,
					},
				},
				Shardings: []appsv1.ClusterTopologySharding{
					{
						Name:        sharding1aName,
						ShardingDef: shardingDefName,
					},
				},
				Orders: &appsv1.ClusterTopologyOrders{
					Provision: []string{sharding1aName, comp1aName},
					Terminate: []string{comp1aName, sharding1aName},
					Update:    []string{sharding1aName, comp1aName},
				},
			}).
			AddClusterTopology(appsv1.ClusterTopology{
				Name: clusterTopologyCompNShardingOOD,
				Components: []appsv1.ClusterTopologyComponent{
					{
						Name:    comp1aName,
						CompDef: compDefName,
					},
				},
				Shardings: []appsv1.ClusterTopologySharding{
					{
						Name:        sharding1aName,
						ShardingDef: shardingDefName,
					},
				},
				Orders: &appsv1.ClusterTopologyOrders{
					Provision: []string{comp1aName, sharding1aName},
					Update:    []string{sharding1aName, comp1aName},
				},
			}).
			AddClusterTopology(appsv1.ClusterTopology{
				Name: clusterTopologyShardingNCompOOD,
				Components: []appsv1.ClusterTopologyComponent{
					{
						Name:    comp1aName,
						CompDef: compDefName,
					},
				},
				Shardings: []appsv1.ClusterTopologySharding{
					{
						Name:        sharding1aName,
						ShardingDef: shardingDefName,
					},
				},
				Orders: &appsv1.ClusterTopologyOrders{
					Provision: []string{sharding1aName, comp1aName},
					Update:    []string{comp1aName, sharding1aName},
				},
			}).
			GetObject()
	})

	AfterEach(func() {})

	newDAG := func(graphCli model.GraphClient, cluster *appsv1.Cluster) *graph.DAG {
		d := graph.NewDAG()
		graphCli.Root(d, cluster, cluster, model.ActionStatusPtr())
		return d
	}

	normalizeTransformContext := func(transCtx *clusterTransformContext) {
		var (
			clusterDef = transCtx.clusterDef
			cluster    = transCtx.Cluster
			err        error
		)
		transformer := clusterNormalizationTransformer{}
		transCtx.components, transCtx.shardings, err = transformer.resolveCompsNShardingsFromTopology(clusterDef, cluster)
		Expect(err).Should(BeNil())

		transCtx.shardingComps, transCtx.shardingCompsWithTpl, err = transformer.buildShardingComps(transCtx)
		Expect(err).Should(BeNil())
	}

	newTransformerNCtx := func(topology string, processors ...func(*testapps.MockClusterFactory)) (graph.Transformer, *clusterTransformContext, *graph.DAG) {
		f := testapps.NewClusterFactory(testCtx.DefaultNamespace, clusterName, clusterDefName).
			WithRandomName().
			SetTopology(topology)
		if len(processors) > 0 {
			for _, processor := range processors {
				processor(f)
			}
		} else {
			f.SetReplicas(1)
		}
		cluster := f.GetObject()

		graphCli := model.NewGraphClient(k8sClient)
		transCtx := &clusterTransformContext{
			Context:       ctx,
			Client:        graphCli,
			EventRecorder: nil,
			Logger:        logger,
			Cluster:       cluster,
			OrigCluster:   cluster.DeepCopy(),
			clusterDef:    clusterDef,
		}
		normalizeTransformContext(transCtx)

		return &clusterComponentTransformer{}, transCtx, newDAG(graphCli, cluster)
	}

	newCompObj := func(transCtx *clusterTransformContext, compSpec *appsv1.ClusterComponentSpec, setters ...func(*appsv1.Component)) *appsv1.Component {
		comp, err := component.BuildComponent(transCtx.Cluster, compSpec, nil, nil)
		Expect(err).Should(BeNil())
		for _, setter := range setters {
			if setter != nil {
				setter(comp)
			}
		}
		return comp
	}

	mockCompObj := func(transCtx *clusterTransformContext, compName string, setters ...func(*appsv1.Component)) *appsv1.Component {
		var compSpec *appsv1.ClusterComponentSpec
		for i, spec := range transCtx.components {
			if spec.Name == compName {
				compSpec = transCtx.components[i]
				break
			}
		}
		Expect(compSpec).ShouldNot(BeNil())
		return newCompObj(transCtx, compSpec, setters...)
	}

	mockShardingCompObj := func(transCtx *clusterTransformContext, shardingName string, setters ...func(*appsv1.Component)) *appsv1.Component {
		specs := transCtx.shardingComps[shardingName]
		Expect(specs).Should(HaveLen(1))
		Expect(specs[0]).ShouldNot(BeNil())

		if setters == nil {
			setters = []func(*appsv1.Component){}
		}
		setters = append(setters, func(comp *appsv1.Component) {
			comp.Labels[constant.KBAppShardingNameLabelKey] = shardingName
		})
		return newCompObj(transCtx, specs[0], setters...)
	}

	Context("component orders", func() {
		It("w/o orders", func() {
			transformer, transCtx, dag := newTransformerNCtx(clusterTopologyNoOrders)
			err := transformer.Transform(transCtx, dag)
			Expect(err).Should(BeNil())

			// check the components
			graphCli := transCtx.Client.(model.GraphClient)
			objs := graphCli.FindAll(dag, &appsv1.Component{})
			Expect(len(objs)).Should(Equal(4))
			for _, obj := range objs {
				comp := obj.(*appsv1.Component)
				Expect(graphCli.IsAction(dag, comp, model.ActionCreatePtr())).Should(BeTrue())
			}
		})

		It("w/ orders provision - has no predecessors", func() {
			transformer, transCtx, dag := newTransformerNCtx(clusterTopologyDefault)
			err := transformer.Transform(transCtx, dag)
			Expect(err).ShouldNot(BeNil())
			Expect(err.Error()).Should(And(ContainSubstring("retry later"), ContainSubstring(comp2aName)))

			// check the first two components
			graphCli := transCtx.Client.(model.GraphClient)
			objs := graphCli.FindAll(dag, &appsv1.Component{})
			Expect(len(objs)).Should(Equal(2))
			for _, obj := range objs {
				comp := obj.(*appsv1.Component)
				Expect(component.ShortName(transCtx.Cluster.Name, comp.Name)).Should(Or(Equal(comp1aName), Equal(comp1bName)))
				Expect(graphCli.IsAction(dag, comp, model.ActionCreatePtr())).Should(BeTrue())
			}
		})

		It("w/ orders provision - has a predecessor not ready", func() {
			transformer, transCtx, dag := newTransformerNCtx(clusterTopologyDefault)

			// mock first two components status as running and creating
			reader := &appsutil.MockReader{
				Objects: []client.Object{
					mockCompObj(transCtx, comp1aName, func(comp *appsv1.Component) {
						comp.Status.Phase = appsv1.RunningComponentPhase
					}),
					mockCompObj(transCtx, comp1bName, func(comp *appsv1.Component) {
						comp.Status.Phase = appsv1.CreatingComponentPhase
					}),
				},
			}
			transCtx.Client = model.NewGraphClient(reader)

			err := transformer.Transform(transCtx, dag)
			Expect(err).ShouldNot(BeNil())
			Expect(err.Error()).Should(And(ContainSubstring("retry later"), ContainSubstring(comp2aName)))

			// should have no components to update
			graphCli := transCtx.Client.(model.GraphClient)
			objs := graphCli.FindAll(dag, &appsv1.Component{})
			Expect(len(objs)).Should(Equal(0))
		})

		It("w/ orders provision - has a predecessor in DAG", func() {
			transformer, transCtx, dag := newTransformerNCtx(clusterTopologyDefault)

			// mock one of first two components status as running
			reader := &appsutil.MockReader{
				Objects: []client.Object{
					mockCompObj(transCtx, comp1aName, func(comp *appsv1.Component) {
						comp.Status.Phase = appsv1.RunningComponentPhase
					}),
				},
			}
			transCtx.Client = model.NewGraphClient(reader)

			err := transformer.Transform(transCtx, dag)
			Expect(err).ShouldNot(BeNil())
			Expect(err.Error()).Should(And(ContainSubstring("retry later"), ContainSubstring(comp2aName)))

			// should have one component to create
			graphCli := transCtx.Client.(model.GraphClient)
			objs := graphCli.FindAll(dag, &appsv1.Component{})
			Expect(len(objs)).Should(Equal(1))
			comp := objs[0].(*appsv1.Component)
			Expect(component.ShortName(transCtx.Cluster.Name, comp.Name)).Should(Equal(comp1bName))
			Expect(graphCli.IsAction(dag, comp, model.ActionCreatePtr())).Should(BeTrue())
		})

		It("w/ orders provision - all predecessors ready", func() {
			transformer, transCtx, dag := newTransformerNCtx(clusterTopologyDefault)

			// mock first two components status as running
			reader := &appsutil.MockReader{
				Objects: []client.Object{
					mockCompObj(transCtx, comp1aName, func(comp *appsv1.Component) {
						comp.Status.Phase = appsv1.RunningComponentPhase
					}),
					mockCompObj(transCtx, comp1bName, func(comp *appsv1.Component) {
						comp.Status.Phase = appsv1.RunningComponentPhase
					}),
				},
			}
			transCtx.Client = model.NewGraphClient(reader)

			err := transformer.Transform(transCtx, dag)
			Expect(err).Should(BeNil())

			// check the last two components
			graphCli := transCtx.Client.(model.GraphClient)
			objs := graphCli.FindAll(dag, &appsv1.Component{})
			Expect(len(objs)).Should(Equal(2))
			for _, obj := range objs {
				comp := obj.(*appsv1.Component)
				Expect(component.ShortName(transCtx.Cluster.Name, comp.Name)).Should(Or(Equal(comp2aName), Equal(comp2bName)))
				Expect(graphCli.IsAction(dag, comp, model.ActionCreatePtr())).Should(BeTrue())
			}
		})

		It("w/ orders update - has no predecessors", func() {
			transformer, transCtx, dag := newTransformerNCtx(clusterTopologyDefault)

			// mock first two components
			reader := &appsutil.MockReader{
				Objects: []client.Object{
					mockCompObj(transCtx, comp1aName, func(comp *appsv1.Component) {
						comp.Spec.Replicas = 2 // to update
						comp.Status.Phase = appsv1.RunningComponentPhase
					}),
					mockCompObj(transCtx, comp1bName, func(comp *appsv1.Component) {
						comp.Status.Phase = appsv1.RunningComponentPhase
					}),
				},
			}
			transCtx.Client = model.NewGraphClient(reader)

			err := transformer.Transform(transCtx, dag)
			Expect(err).ShouldNot(BeNil())
			Expect(err.Error()).Should(And(ContainSubstring("retry later"), ContainSubstring(comp2aName)))

			// check the first component
			graphCli := transCtx.Client.(model.GraphClient)
			objs := graphCli.FindAll(dag, &appsv1.Component{})
			Expect(len(objs)).Should(Equal(1))
			comp := objs[0].(*appsv1.Component)
			Expect(component.ShortName(transCtx.Cluster.Name, comp.Name)).Should(Equal(comp1aName))
			Expect(graphCli.IsAction(dag, comp, model.ActionUpdatePtr())).Should(BeTrue())
		})

		It("w/ orders update - has a predecessor not ready", func() {
			transformer, transCtx, dag := newTransformerNCtx(clusterTopologyDefault)

			// mock components
			reader := &appsutil.MockReader{
				Objects: []client.Object{
					mockCompObj(transCtx, comp1aName, func(comp *appsv1.Component) {
						comp.Status.Phase = appsv1.CreatingComponentPhase // not ready
					}),
					mockCompObj(transCtx, comp1bName, func(comp *appsv1.Component) {
						comp.Status.Phase = appsv1.RunningComponentPhase
					}),
					mockCompObj(transCtx, comp2aName, func(comp *appsv1.Component) {
						comp.Spec.Replicas = 2 // to update
						comp.Status.Phase = appsv1.RunningComponentPhase
					}),
					mockCompObj(transCtx, comp2bName, func(comp *appsv1.Component) {
						comp.Status.Phase = appsv1.RunningComponentPhase
					}),
				},
			}
			transCtx.Client = model.NewGraphClient(reader)
			transCtx.OrigCluster.Generation += 1 // mock cluster spec update

			err := transformer.Transform(transCtx, dag)
			Expect(err).ShouldNot(BeNil())
			Expect(err.Error()).Should(And(ContainSubstring("retry later"), ContainSubstring(comp2aName)))

			// should have no components to update
			graphCli := transCtx.Client.(model.GraphClient)
			objs := graphCli.FindAll(dag, &appsv1.Component{})
			Expect(len(objs)).Should(Equal(0))
		})

		It("w/ orders update - has a predecessor in DAG", func() {
			transformer, transCtx, dag := newTransformerNCtx(clusterTopologyDefault)

			// mock components
			reader := &appsutil.MockReader{
				Objects: []client.Object{
					mockCompObj(transCtx, comp1aName, func(comp *appsv1.Component) {
						comp.Spec.Replicas = 2 // to update
						comp.Status.Phase = appsv1.RunningComponentPhase
					}),
					mockCompObj(transCtx, comp1bName, func(comp *appsv1.Component) {
						comp.Status.Phase = appsv1.RunningComponentPhase
					}),
					mockCompObj(transCtx, comp2aName, func(comp *appsv1.Component) {
						comp.Spec.Replicas = 2 // to update
						comp.Status.Phase = appsv1.RunningComponentPhase
					}),
					mockCompObj(transCtx, comp2bName, func(comp *appsv1.Component) {
						comp.Status.Phase = appsv1.RunningComponentPhase
					}),
				},
			}
			transCtx.Client = model.NewGraphClient(reader)
			transCtx.OrigCluster.Generation += 1 // mock cluster spec update

			err := transformer.Transform(transCtx, dag)
			Expect(err).ShouldNot(BeNil())
			Expect(err.Error()).Should(And(ContainSubstring("retry later"), ContainSubstring(comp2aName)))

			// should have one component to update
			graphCli := transCtx.Client.(model.GraphClient)
			objs := graphCli.FindAll(dag, &appsv1.Component{})
			Expect(len(objs)).Should(Equal(1))
			comp := objs[0].(*appsv1.Component)
			Expect(component.ShortName(transCtx.Cluster.Name, comp.Name)).Should(Equal(comp1aName))
			Expect(graphCli.IsAction(dag, comp, model.ActionUpdatePtr())).Should(BeTrue())
		})

		It("w/ orders update - all predecessors ready", func() {
			transformer, transCtx, dag := newTransformerNCtx(clusterTopologyDefault)

			// mock components
			reader := &appsutil.MockReader{
				Objects: []client.Object{
					mockCompObj(transCtx, comp1aName, func(comp *appsv1.Component) {
						comp.Status.Phase = appsv1.RunningComponentPhase
					}),
					mockCompObj(transCtx, comp1bName, func(comp *appsv1.Component) {
						comp.Status.Phase = appsv1.RunningComponentPhase
					}),
					mockCompObj(transCtx, comp2aName, func(comp *appsv1.Component) {
						comp.Spec.Replicas = 2 // to update
						comp.Status.Phase = appsv1.RunningComponentPhase
					}),
					mockCompObj(transCtx, comp2bName, func(comp *appsv1.Component) {
						comp.Status.Phase = appsv1.RunningComponentPhase
					}),
				},
			}
			transCtx.Client = model.NewGraphClient(reader)
			transCtx.OrigCluster.Generation += 1 // mock cluster spec update

			err := transformer.Transform(transCtx, dag)
			Expect(err).Should(BeNil())

			graphCli := transCtx.Client.(model.GraphClient)
			objs := graphCli.FindAll(dag, &appsv1.Component{})
			Expect(len(objs)).Should(Equal(1))
			comp := objs[0].(*appsv1.Component)
			Expect(component.ShortName(transCtx.Cluster.Name, comp.Name)).Should(Equal(comp2aName))
			Expect(graphCli.IsAction(dag, comp, model.ActionUpdatePtr())).Should(BeTrue())
		})

		It("w/ orders update - stop", func() {
			transformer, transCtx, dag := newTransformerNCtx(clusterTopologyStop)

			// mock to stop all components
			reader := &appsutil.MockReader{
				Objects: []client.Object{
					mockCompObj(transCtx, comp1aName, func(comp *appsv1.Component) {
						comp.Status.Phase = appsv1.RunningComponentPhase
					}),
					mockCompObj(transCtx, comp2aName, func(comp *appsv1.Component) {
						comp.Status.Phase = appsv1.RunningComponentPhase
					}),
					mockCompObj(transCtx, comp3aName, func(comp *appsv1.Component) {
						comp.Status.Phase = appsv1.RunningComponentPhase
					}),
				},
			}
			transCtx.Client = model.NewGraphClient(reader)
			for i := range transCtx.components {
				transCtx.components[i].Stop = pointer.Bool(true)
			}
			transCtx.OrigCluster.Generation += 1 // mock cluster spec update

			err := transformer.Transform(transCtx, dag)
			Expect(err).ShouldNot(BeNil())
			Expect(err.Error()).Should(And(ContainSubstring("retry later"), ContainSubstring(comp2aName)))

			// should have the first component to update only
			graphCli := transCtx.Client.(model.GraphClient)
			objs := graphCli.FindAll(dag, &appsv1.Component{})
			Expect(len(objs)).Should(Equal(1))
			comp := objs[0].(*appsv1.Component)
			Expect(component.ShortName(transCtx.Cluster.Name, comp.Name)).Should(Equal(comp1aName))
			Expect(graphCli.IsAction(dag, comp, model.ActionUpdatePtr())).Should(BeTrue())
			Expect(comp.Spec.Stop).ShouldNot(BeNil())
			Expect(*comp.Spec.Stop).Should(BeTrue())
		})

		It("w/ orders update - stop the second component", func() {
			transformer, transCtx, dag := newTransformerNCtx(clusterTopologyStop)

			// mock to stop all components and the first component has been stopped
			reader := &appsutil.MockReader{
				Objects: []client.Object{
					mockCompObj(transCtx, comp1aName, func(comp *appsv1.Component) {
						comp.Spec.Stop = pointer.Bool(true)
						comp.Status.Phase = appsv1.StoppedComponentPhase
					}),
					mockCompObj(transCtx, comp2aName, func(comp *appsv1.Component) {
						comp.Status.Phase = appsv1.RunningComponentPhase
					}),
					mockCompObj(transCtx, comp3aName, func(comp *appsv1.Component) {
						comp.Status.Phase = appsv1.RunningComponentPhase
					}),
				},
			}
			transCtx.Client = model.NewGraphClient(reader)
			for i := range transCtx.components {
				transCtx.components[i].Stop = pointer.Bool(true)
			}
			transCtx.OrigCluster.Generation += 1 // mock cluster spec update

			err := transformer.Transform(transCtx, dag)
			Expect(err).ShouldNot(BeNil())
			Expect(err.Error()).Should(And(ContainSubstring("retry later"), ContainSubstring(comp3aName)))

			// should have the second component to update only
			graphCli := transCtx.Client.(model.GraphClient)
			objs := graphCli.FindAll(dag, &appsv1.Component{})
			Expect(len(objs)).Should(Equal(1))
			comp := objs[0].(*appsv1.Component)
			Expect(component.ShortName(transCtx.Cluster.Name, comp.Name)).Should(Equal(comp2aName))
			Expect(graphCli.IsAction(dag, comp, model.ActionUpdatePtr())).Should(BeTrue())
			Expect(comp.Spec.Stop).ShouldNot(BeNil())
			Expect(*comp.Spec.Stop).Should(BeTrue())
		})

		It("w/ orders provision & update - OOD", func() {
			transformer, transCtx, dag := newTransformerNCtx(clusterTopologyProvisionNUpdateOOD)

			// mock first two components status as running
			reader := &appsutil.MockReader{
				Objects: []client.Object{
					mockCompObj(transCtx, comp1aName, func(comp *appsv1.Component) {
						comp.Status.Phase = appsv1.RunningComponentPhase
					}),
					mockCompObj(transCtx, comp1bName, func(comp *appsv1.Component) {
						comp.Status.Phase = appsv1.RunningComponentPhase
					}),
				},
			}
			transCtx.Client = model.NewGraphClient(reader)

			// comp2aName and comp2bName are not ready (exist) when updating comp1aName and comp1bName
			err := transformer.Transform(transCtx, dag)
			Expect(err).Should(BeNil())

			// check the last two components under provisioning
			graphCli := transCtx.Client.(model.GraphClient)
			objs := graphCli.FindAll(dag, &appsv1.Component{})
			Expect(len(objs)).Should(Equal(2))
			for _, obj := range objs {
				comp := obj.(*appsv1.Component)
				Expect(component.ShortName(transCtx.Cluster.Name, comp.Name)).Should(Or(Equal(comp2aName), Equal(comp2bName)))
				Expect(graphCli.IsAction(dag, comp, model.ActionCreatePtr())).Should(BeTrue())
			}

			// mock last two components status as running
			reader.Objects = append(reader.Objects, []client.Object{
				mockCompObj(transCtx, comp2aName, func(comp *appsv1.Component) {
					comp.Status.Phase = appsv1.RunningComponentPhase
				}),
				mockCompObj(transCtx, comp2bName, func(comp *appsv1.Component) {
					comp.Status.Phase = appsv1.RunningComponentPhase
				}),
			}...)

			// try again
			err = transformer.Transform(transCtx, newDAG(graphCli, transCtx.Cluster))
			Expect(err).Should(BeNil())
		})

		It("template component - has no components instantiated", func() {
			transformer, transCtx, dag := newTransformerNCtx(clusterTopologyTemplate)

			// check the components created, no components should be instantiated from the template automatically
			Expect(transCtx.components).Should(HaveLen(3))
			Expect(transCtx.components[0].Name).Should(Equal(comp1aName))
			Expect(transCtx.components[1].Name).Should(Equal(comp2bName))
			Expect(transCtx.components[2].Name).Should(Equal(comp3aName))

			// mock first component status as running
			reader := &appsutil.MockReader{
				Objects: []client.Object{
					mockCompObj(transCtx, comp1aName, func(comp *appsv1.Component) {
						comp.Status.Phase = appsv1.RunningComponentPhase
					}),
				},
			}
			transCtx.Client = model.NewGraphClient(reader)

			err := transformer.Transform(transCtx, dag)
			Expect(err).ShouldNot(BeNil())
			Expect(err.Error()).Should(And(ContainSubstring("retry later"), ContainSubstring(comp3aName)))

			// check other components
			graphCli := transCtx.Client.(model.GraphClient)
			objs := graphCli.FindAll(dag, &appsv1.Component{})
			Expect(len(objs)).Should(Equal(1))
			for _, obj := range objs {
				comp := obj.(*appsv1.Component)
				Expect(component.ShortName(transCtx.Cluster.Name, comp.Name)).Should(Equal(comp2bName))
				Expect(graphCli.IsAction(dag, comp, model.ActionCreatePtr())).Should(BeTrue())
			}
		})

		It("template component - has components instantiated", func() {
			transformer, transCtx, dag := newTransformerNCtx(clusterTopologyTemplate, func(f *testapps.MockClusterFactory) {
				f.AddComponent(fmt.Sprintf("%s-0", comp2aName), compDefName).
					AddComponent(fmt.Sprintf("%s-1", comp2aName), compDefName)
			})

			// check the components created
			Expect(transCtx.components).Should(HaveLen(5))
			Expect(transCtx.components[0].Name).Should(Equal(comp1aName))
			Expect(transCtx.components[1].Name).Should(HavePrefix(comp2aName))
			Expect(transCtx.components[2].Name).Should(HavePrefix(comp2aName))
			Expect(transCtx.components[3].Name).Should(Equal(comp2bName))
			Expect(transCtx.components[4].Name).Should(Equal(comp3aName))

			// mock first component status as running
			reader := &appsutil.MockReader{
				Objects: []client.Object{
					mockCompObj(transCtx, comp1aName, func(comp *appsv1.Component) {
						comp.Status.Phase = appsv1.RunningComponentPhase
					}),
				},
			}
			transCtx.Client = model.NewGraphClient(reader)

			err := transformer.Transform(transCtx, dag)
			Expect(err).ShouldNot(BeNil())
			Expect(err.Error()).Should(And(ContainSubstring("retry later"), ContainSubstring(comp3aName)))

			// check other components
			graphCli := transCtx.Client.(model.GraphClient)
			objs := graphCli.FindAll(dag, &appsv1.Component{})
			Expect(len(objs)).Should(Equal(3))
			for _, obj := range objs {
				comp := obj.(*appsv1.Component)
				Expect(component.ShortName(transCtx.Cluster.Name, comp.Name)).Should(Or(HavePrefix(comp2aName), Equal(comp2bName)))
				Expect(graphCli.IsAction(dag, comp, model.ActionCreatePtr())).Should(BeTrue())
			}
		})
	})

	Context("sharding orders", func() {
		It("w/o orders", func() {
			transformer, transCtx, dag := newTransformerNCtx(clusterTopologyNoOrders4Sharding, func(f *testapps.MockClusterFactory) {
				f.AddSharding(sharding1aName, "", "").
					AddSharding(sharding1bName, "", "").
					AddSharding(sharding2aName, "", "").
					AddSharding(sharding2bName, "", "")
			})
			err := transformer.Transform(transCtx, dag)
			Expect(err).Should(BeNil())

			// check the components
			graphCli := transCtx.Client.(model.GraphClient)
			objs := graphCli.FindAll(dag, &appsv1.Component{})
			Expect(len(objs)).Should(Equal(4))
			for _, obj := range objs {
				comp := obj.(*appsv1.Component)
				Expect(graphCli.IsAction(dag, comp, model.ActionCreatePtr())).Should(BeTrue())
			}
		})

		It("w/ orders provision - has no predecessors", func() {
			transformer, transCtx, dag := newTransformerNCtx(clusterTopologyDefault4Sharding, func(f *testapps.MockClusterFactory) {
				f.AddSharding(sharding1aName, "", "").
					AddSharding(sharding1bName, "", "").
					AddSharding(sharding2aName, "", "").
					AddSharding(sharding2bName, "", "")
			})
			err := transformer.Transform(transCtx, dag)
			Expect(err).ShouldNot(BeNil())
			Expect(err.Error()).Should(And(ContainSubstring("retry later"), ContainSubstring(sharding2aName)))

			// check the first two components
			graphCli := transCtx.Client.(model.GraphClient)
			objs := graphCli.FindAll(dag, &appsv1.Component{})
			Expect(len(objs)).Should(Equal(2))
			for _, obj := range objs {
				comp := obj.(*appsv1.Component)
				Expect(component.ShortName(transCtx.Cluster.Name, comp.Name)).Should(Or(HavePrefix(sharding1aName), HavePrefix(sharding1bName)))
				Expect(graphCli.IsAction(dag, comp, model.ActionCreatePtr())).Should(BeTrue())
			}
		})

		It("w/ orders provision - has a predecessor not ready", func() {
			transformer, transCtx, dag := newTransformerNCtx(clusterTopologyDefault4Sharding, func(f *testapps.MockClusterFactory) {
				f.AddSharding(sharding1aName, "", "").
					AddSharding(sharding1bName, "", "").
					AddSharding(sharding2aName, "", "").
					AddSharding(sharding2bName, "", "")
			})

			// mock first two components status as running and creating
			reader := &appsutil.MockReader{
				Objects: []client.Object{
					mockShardingCompObj(transCtx, sharding1aName, func(comp *appsv1.Component) {
						comp.Status.Phase = appsv1.RunningComponentPhase
					}),
					mockShardingCompObj(transCtx, sharding1bName, func(comp *appsv1.Component) {
						comp.Status.Phase = appsv1.CreatingComponentPhase
					}),
				},
			}
			transCtx.Client = model.NewGraphClient(reader)

			err := transformer.Transform(transCtx, dag)
			Expect(err).ShouldNot(BeNil())
			Expect(err.Error()).Should(And(ContainSubstring("retry later"), ContainSubstring(sharding2aName)))

			// should have no components to update
			graphCli := transCtx.Client.(model.GraphClient)
			objs := graphCli.FindAll(dag, &appsv1.Component{})
			Expect(len(objs)).Should(Equal(0))
		})

		It("w/ orders provision - has a predecessor in DAG", func() {
			transformer, transCtx, dag := newTransformerNCtx(clusterTopologyDefault4Sharding, func(f *testapps.MockClusterFactory) {
				f.AddSharding(sharding1aName, "", "").
					AddSharding(sharding1bName, "", "").
					AddSharding(sharding2aName, "", "").
					AddSharding(sharding2bName, "", "")
			})

			// mock one of first two components status as running
			reader := &appsutil.MockReader{
				Objects: []client.Object{
					mockShardingCompObj(transCtx, sharding1aName, func(comp *appsv1.Component) {
						comp.Status.Phase = appsv1.RunningComponentPhase
					}),
				},
			}
			transCtx.Client = model.NewGraphClient(reader)

			err := transformer.Transform(transCtx, dag)
			Expect(err).ShouldNot(BeNil())
			Expect(err.Error()).Should(And(ContainSubstring("retry later"), ContainSubstring(sharding2aName)))

			// should have one component to create
			graphCli := transCtx.Client.(model.GraphClient)
			objs := graphCli.FindAll(dag, &appsv1.Component{})
			Expect(len(objs)).Should(Equal(1))
			comp := objs[0].(*appsv1.Component)
			Expect(component.ShortName(transCtx.Cluster.Name, comp.Name)).Should(HavePrefix(sharding1bName))
			Expect(graphCli.IsAction(dag, comp, model.ActionCreatePtr())).Should(BeTrue())
		})

		It("w/ orders provision - all predecessors ready", func() {
			transformer, transCtx, dag := newTransformerNCtx(clusterTopologyDefault4Sharding, func(f *testapps.MockClusterFactory) {
				f.AddSharding(sharding1aName, "", "").
					AddSharding(sharding1bName, "", "").
					AddSharding(sharding2aName, "", "").
					AddSharding(sharding2bName, "", "")
			})

			// mock first two components status as running
			reader := &appsutil.MockReader{
				Objects: []client.Object{
					mockShardingCompObj(transCtx, sharding1aName, func(comp *appsv1.Component) {
						comp.Status.Phase = appsv1.RunningComponentPhase
					}),
					mockShardingCompObj(transCtx, sharding1bName, func(comp *appsv1.Component) {
						comp.Status.Phase = appsv1.RunningComponentPhase
					}),
				},
			}
			transCtx.Client = model.NewGraphClient(reader)

			err := transformer.Transform(transCtx, dag)
			Expect(err).Should(BeNil())

			// check the last two components
			graphCli := transCtx.Client.(model.GraphClient)
			objs := graphCli.FindAll(dag, &appsv1.Component{})
			Expect(len(objs)).Should(Equal(2))
			for _, obj := range objs {
				comp := obj.(*appsv1.Component)
				Expect(component.ShortName(transCtx.Cluster.Name, comp.Name)).Should(Or(HavePrefix(sharding2aName), HavePrefix(sharding2bName)))
				Expect(graphCli.IsAction(dag, comp, model.ActionCreatePtr())).Should(BeTrue())
			}
		})

		It("w/ orders update - has no predecessors", func() {
			transformer, transCtx, dag := newTransformerNCtx(clusterTopologyDefault4Sharding, func(f *testapps.MockClusterFactory) {
				f.AddSharding(sharding1aName, "", "").
					AddSharding(sharding1bName, "", "").
					AddSharding(sharding2aName, "", "").
					AddSharding(sharding2bName, "", "")
			})

			// mock first two components
			reader := &appsutil.MockReader{
				Objects: []client.Object{
					mockShardingCompObj(transCtx, sharding1aName, func(comp *appsv1.Component) {
						comp.Spec.Replicas = 2 // to update
						comp.Status.Phase = appsv1.RunningComponentPhase
					}),
					mockShardingCompObj(transCtx, sharding1bName, func(comp *appsv1.Component) {
						comp.Status.Phase = appsv1.RunningComponentPhase
					}),
				},
			}
			transCtx.Client = model.NewGraphClient(reader)

			err := transformer.Transform(transCtx, dag)
			Expect(err).ShouldNot(BeNil())
			Expect(err.Error()).Should(And(ContainSubstring("retry later"), ContainSubstring(sharding2aName)))

			// check the first component
			graphCli := transCtx.Client.(model.GraphClient)
			objs := graphCli.FindAll(dag, &appsv1.Component{})
			Expect(len(objs)).Should(Equal(1))
			comp := objs[0].(*appsv1.Component)
			Expect(component.ShortName(transCtx.Cluster.Name, comp.Name)).Should(HavePrefix(sharding1aName))
			Expect(graphCli.IsAction(dag, comp, model.ActionUpdatePtr())).Should(BeTrue())
		})

		It("w/ orders update - has a predecessor not ready", func() {
			transformer, transCtx, dag := newTransformerNCtx(clusterTopologyDefault4Sharding, func(f *testapps.MockClusterFactory) {
				f.AddSharding(sharding1aName, "", "").
					AddSharding(sharding1bName, "", "").
					AddSharding(sharding2aName, "", "").
					AddSharding(sharding2bName, "", "")
			})

			// mock components
			reader := &appsutil.MockReader{
				Objects: []client.Object{
					mockShardingCompObj(transCtx, sharding1aName, func(comp *appsv1.Component) {
						comp.Status.Phase = appsv1.CreatingComponentPhase // not ready
					}),
					mockShardingCompObj(transCtx, sharding1bName, func(comp *appsv1.Component) {
						comp.Status.Phase = appsv1.RunningComponentPhase
					}),
					mockShardingCompObj(transCtx, sharding2aName, func(comp *appsv1.Component) {
						comp.Spec.Replicas = 2 // to update
						comp.Status.Phase = appsv1.RunningComponentPhase
					}),
					mockShardingCompObj(transCtx, sharding2bName, func(comp *appsv1.Component) {
						comp.Status.Phase = appsv1.RunningComponentPhase
					}),
				},
			}
			transCtx.Client = model.NewGraphClient(reader)
			transCtx.OrigCluster.Generation += 1 // mock cluster spec update

			err := transformer.Transform(transCtx, dag)
			Expect(err).ShouldNot(BeNil())
			Expect(err.Error()).Should(And(ContainSubstring("retry later"), ContainSubstring(sharding2aName)))

			// should have no components to update
			graphCli := transCtx.Client.(model.GraphClient)
			objs := graphCli.FindAll(dag, &appsv1.Component{})
			Expect(len(objs)).Should(Equal(0))
		})

		It("w/ orders update - has a predecessor in DAG", func() {
			transformer, transCtx, dag := newTransformerNCtx(clusterTopologyDefault4Sharding, func(f *testapps.MockClusterFactory) {
				f.AddSharding(sharding1aName, "", "").
					AddSharding(sharding1bName, "", "").
					AddSharding(sharding2aName, "", "").
					AddSharding(sharding2bName, "", "")
			})

			// mock components
			reader := &appsutil.MockReader{
				Objects: []client.Object{
					mockShardingCompObj(transCtx, sharding1aName, func(comp *appsv1.Component) {
						comp.Spec.Replicas = 2 // to update
						comp.Status.Phase = appsv1.RunningComponentPhase
					}),
					mockShardingCompObj(transCtx, sharding1bName, func(comp *appsv1.Component) {
						comp.Status.Phase = appsv1.RunningComponentPhase
					}),
					mockShardingCompObj(transCtx, sharding2aName, func(comp *appsv1.Component) {
						comp.Spec.Replicas = 2 // to update
						comp.Status.Phase = appsv1.RunningComponentPhase
					}),
					mockShardingCompObj(transCtx, sharding2bName, func(comp *appsv1.Component) {
						comp.Status.Phase = appsv1.RunningComponentPhase
					}),
				},
			}
			transCtx.Client = model.NewGraphClient(reader)
			transCtx.OrigCluster.Generation += 1 // mock cluster spec update

			err := transformer.Transform(transCtx, dag)
			Expect(err).ShouldNot(BeNil())
			Expect(err.Error()).Should(And(ContainSubstring("retry later"), ContainSubstring(sharding2aName)))

			// should have one component to update
			graphCli := transCtx.Client.(model.GraphClient)
			objs := graphCli.FindAll(dag, &appsv1.Component{})
			Expect(len(objs)).Should(Equal(1))
			comp := objs[0].(*appsv1.Component)
			Expect(component.ShortName(transCtx.Cluster.Name, comp.Name)).Should(HavePrefix(sharding1aName))
			Expect(graphCli.IsAction(dag, comp, model.ActionUpdatePtr())).Should(BeTrue())
		})

		It("w/ orders update - all predecessors ready", func() {
			transformer, transCtx, dag := newTransformerNCtx(clusterTopologyDefault4Sharding, func(f *testapps.MockClusterFactory) {
				f.AddSharding(sharding1aName, "", "").
					AddSharding(sharding1bName, "", "").
					AddSharding(sharding2aName, "", "").
					AddSharding(sharding2bName, "", "")
			})

			// mock components
			reader := &appsutil.MockReader{
				Objects: []client.Object{
					mockShardingCompObj(transCtx, sharding1aName, func(comp *appsv1.Component) {
						comp.Status.Phase = appsv1.RunningComponentPhase
					}),
					mockShardingCompObj(transCtx, sharding1bName, func(comp *appsv1.Component) {
						comp.Status.Phase = appsv1.RunningComponentPhase
					}),
					mockShardingCompObj(transCtx, sharding2aName, func(comp *appsv1.Component) {
						comp.Spec.Replicas = 2 // to update
						comp.Status.Phase = appsv1.RunningComponentPhase
					}),
					mockShardingCompObj(transCtx, sharding2bName, func(comp *appsv1.Component) {
						comp.Status.Phase = appsv1.RunningComponentPhase
					}),
				},
			}
			transCtx.Client = model.NewGraphClient(reader)
			transCtx.OrigCluster.Generation += 1 // mock cluster spec update

			err := transformer.Transform(transCtx, dag)
			Expect(err).Should(BeNil())

			graphCli := transCtx.Client.(model.GraphClient)
			objs := graphCli.FindAll(dag, &appsv1.Component{})
			Expect(len(objs)).Should(Equal(1))
			comp := objs[0].(*appsv1.Component)
			Expect(component.ShortName(transCtx.Cluster.Name, comp.Name)).Should(HavePrefix(sharding2aName))
			Expect(graphCli.IsAction(dag, comp, model.ActionUpdatePtr())).Should(BeTrue())
		})

		It("w/ orders update - stop", func() {
			transformer, transCtx, dag := newTransformerNCtx(clusterTopologyStop4Sharding, func(f *testapps.MockClusterFactory) {
				f.AddSharding(sharding1aName, "", "").
					AddSharding(sharding2aName, "", "").
					AddSharding(sharding3aName, "", "")
			})

			// mock to stop all components
			reader := &appsutil.MockReader{
				Objects: []client.Object{
					mockShardingCompObj(transCtx, sharding1aName, func(comp *appsv1.Component) {
						comp.Status.Phase = appsv1.RunningComponentPhase
					}),
					mockShardingCompObj(transCtx, sharding2aName, func(comp *appsv1.Component) {
						comp.Status.Phase = appsv1.RunningComponentPhase
					}),
					mockShardingCompObj(transCtx, sharding3aName, func(comp *appsv1.Component) {
						comp.Status.Phase = appsv1.RunningComponentPhase
					}),
				},
			}
			transCtx.Client = model.NewGraphClient(reader)
			for i, sharding := range transCtx.shardings {
				transCtx.shardings[i].Template.Stop = pointer.Bool(true)
				for j := range transCtx.shardingComps[sharding.Name] {
					transCtx.shardingComps[sharding.Name][j].Stop = pointer.Bool(true)
				}
			}
			transCtx.OrigCluster.Generation += 1 // mock cluster spec update

			err := transformer.Transform(transCtx, dag)
			Expect(err).ShouldNot(BeNil())
			Expect(err.Error()).Should(And(ContainSubstring("retry later"), ContainSubstring(sharding2aName)))

			// should have the first component to update only
			graphCli := transCtx.Client.(model.GraphClient)
			objs := graphCli.FindAll(dag, &appsv1.Component{})
			Expect(len(objs)).Should(Equal(1))
			comp := objs[0].(*appsv1.Component)
			Expect(component.ShortName(transCtx.Cluster.Name, comp.Name)).Should(HavePrefix(sharding1aName))
			Expect(graphCli.IsAction(dag, comp, model.ActionUpdatePtr())).Should(BeTrue())
			Expect(comp.Spec.Stop).ShouldNot(BeNil())
			Expect(*comp.Spec.Stop).Should(BeTrue())
		})

		It("w/ orders update - stop the second component", func() {
			transformer, transCtx, dag := newTransformerNCtx(clusterTopologyStop4Sharding, func(f *testapps.MockClusterFactory) {
				f.AddSharding(sharding1aName, "", "").
					AddSharding(sharding2aName, "", "").
					AddSharding(sharding3aName, "", "")
			})

			// mock to stop all components and the first component has been stopped
			reader := &appsutil.MockReader{
				Objects: []client.Object{
					mockShardingCompObj(transCtx, sharding1aName, func(comp *appsv1.Component) {
						comp.Spec.Stop = pointer.Bool(true)
						comp.Status.Phase = appsv1.StoppedComponentPhase
					}),
					mockShardingCompObj(transCtx, sharding2aName, func(comp *appsv1.Component) {
						comp.Status.Phase = appsv1.RunningComponentPhase
					}),
					mockShardingCompObj(transCtx, sharding3aName, func(comp *appsv1.Component) {
						comp.Status.Phase = appsv1.RunningComponentPhase
					}),
				},
			}
			transCtx.Client = model.NewGraphClient(reader)
			for i, sharding := range transCtx.shardings {
				transCtx.shardings[i].Template.Stop = pointer.Bool(true)
				for j := range transCtx.shardingComps[sharding.Name] {
					transCtx.shardingComps[sharding.Name][j].Stop = pointer.Bool(true)
				}
			}
			transCtx.OrigCluster.Generation += 1 // mock cluster spec update

			err := transformer.Transform(transCtx, dag)
			Expect(err).ShouldNot(BeNil())
			Expect(err.Error()).Should(And(ContainSubstring("retry later"), ContainSubstring(sharding3aName)))

			// should have the second component to update only
			graphCli := transCtx.Client.(model.GraphClient)
			objs := graphCli.FindAll(dag, &appsv1.Component{})
			Expect(len(objs)).Should(Equal(1))
			comp := objs[0].(*appsv1.Component)
			Expect(component.ShortName(transCtx.Cluster.Name, comp.Name)).Should(HavePrefix(sharding2aName))
			Expect(graphCli.IsAction(dag, comp, model.ActionUpdatePtr())).Should(BeTrue())
			Expect(comp.Spec.Stop).ShouldNot(BeNil())
			Expect(*comp.Spec.Stop).Should(BeTrue())
		})

		It("w/ orders provision & update - OOD", func() {
			transformer, transCtx, dag := newTransformerNCtx(clusterTopologyProvisionNUpdateOOD4Sharding, func(f *testapps.MockClusterFactory) {
				f.AddSharding(sharding1aName, "", "").
					AddSharding(sharding1bName, "", "").
					AddSharding(sharding2aName, "", "").
					AddSharding(sharding2bName, "", "")
			})

			// mock first two components status as running
			reader := &appsutil.MockReader{
				Objects: []client.Object{
					mockShardingCompObj(transCtx, sharding1aName, func(comp *appsv1.Component) {
						comp.Status.Phase = appsv1.RunningComponentPhase
					}),
					mockShardingCompObj(transCtx, sharding1bName, func(comp *appsv1.Component) {
						comp.Status.Phase = appsv1.RunningComponentPhase
					}),
				},
			}
			transCtx.Client = model.NewGraphClient(reader)

			// sharding2aName and sharding2bName are not ready (exist) when updating sharding1aName and sharding1bName
			err := transformer.Transform(transCtx, dag)
			Expect(err).Should(BeNil())

			// check the last two components under provisioning
			graphCli := transCtx.Client.(model.GraphClient)
			objs := graphCli.FindAll(dag, &appsv1.Component{})
			Expect(len(objs)).Should(Equal(2))
			for _, obj := range objs {
				comp := obj.(*appsv1.Component)
				Expect(component.ShortName(transCtx.Cluster.Name, comp.Name)).Should(Or(HavePrefix(sharding2aName), HavePrefix(sharding2bName)))
				Expect(graphCli.IsAction(dag, comp, model.ActionCreatePtr())).Should(BeTrue())
			}

			// mock last two components status as running
			reader.Objects = append(reader.Objects, []client.Object{
				mockShardingCompObj(transCtx, sharding2aName, func(comp *appsv1.Component) {
					comp.Status.Phase = appsv1.RunningComponentPhase
				}),
				mockShardingCompObj(transCtx, sharding2bName, func(comp *appsv1.Component) {
					comp.Status.Phase = appsv1.RunningComponentPhase
				}),
			}...)

			// try again
			err = transformer.Transform(transCtx, newDAG(graphCli, transCtx.Cluster))
			Expect(err).Should(BeNil())
		})
	})

	Context("component and sharding orders", func() {
		It("provision", func() {
			for _, suit := range []struct {
				topology                 string
				errMatcher               types.GomegaMatcher
				firstCreatedNameMatcher  types.GomegaMatcher
				secondCreatedNameMatcher types.GomegaMatcher
				mockObjects              func(*clusterTransformContext) []client.Object
			}{
				{
					topology:                 clusterTopologyCompNSharding,
					errMatcher:               ContainSubstring(sharding1aName),
					firstCreatedNameMatcher:  Equal(comp1aName),
					secondCreatedNameMatcher: HavePrefix(sharding1aName),
					mockObjects: func(transCtx *clusterTransformContext) []client.Object {
						return []client.Object{
							mockCompObj(transCtx, comp1aName, func(comp *appsv1.Component) {
								comp.Status.Phase = appsv1.RunningComponentPhase
							}),
						}
					},
				},
				{
					topology:                 clusterTopologyShardingNComp,
					errMatcher:               ContainSubstring(comp1aName),
					firstCreatedNameMatcher:  HavePrefix(sharding1aName),
					secondCreatedNameMatcher: Equal(comp1aName),
					mockObjects: func(transCtx *clusterTransformContext) []client.Object {
						return []client.Object{
							mockShardingCompObj(transCtx, sharding1aName, func(comp *appsv1.Component) {
								comp.Status.Phase = appsv1.RunningComponentPhase
							}),
						}
					},
				},
			} {
				By(suit.topology)
				transformer, transCtx, dag := newTransformerNCtx(suit.topology, func(f *testapps.MockClusterFactory) {
					f.AddSharding(sharding1aName, "", "")
				})
				err := transformer.Transform(transCtx, dag)
				Expect(err).ShouldNot(BeNil())
				Expect(err.Error()).Should(And(ContainSubstring("retry later"), suit.errMatcher))

				// check the first component
				graphCli := transCtx.Client.(model.GraphClient)
				objs := graphCli.FindAll(dag, &appsv1.Component{})
				Expect(len(objs)).Should(Equal(1))
				for _, obj := range objs {
					comp := obj.(*appsv1.Component)
					Expect(component.ShortName(transCtx.Cluster.Name, comp.Name)).Should(suit.firstCreatedNameMatcher)
					Expect(graphCli.IsAction(dag, comp, model.ActionCreatePtr())).Should(BeTrue())
				}

				// mock first component status as running
				reader := &appsutil.MockReader{Objects: suit.mockObjects(transCtx)}
				transCtx.Client = model.NewGraphClient(reader)

				// try again and check the last component
				dag = newDAG(graphCli, transCtx.Cluster)
				err = transformer.Transform(transCtx, dag)
				Expect(err).Should(BeNil())

				graphCli = transCtx.Client.(model.GraphClient)
				objs = graphCli.FindAll(dag, &appsv1.Component{})
				Expect(len(objs)).Should(Equal(1))
				for _, obj := range objs {
					comp := obj.(*appsv1.Component)
					Expect(component.ShortName(transCtx.Cluster.Name, comp.Name)).Should(suit.secondCreatedNameMatcher)
					Expect(graphCli.IsAction(dag, comp, model.ActionCreatePtr())).Should(BeTrue())
				}
			}
		})

		It("update", func() {
			for _, suit := range []struct {
				topology           string
				errMatcher         types.GomegaMatcher
				updatedNameMatcher types.GomegaMatcher
			}{
				{clusterTopologyCompNSharding, ContainSubstring(sharding1aName), Equal(comp1aName)},
				{clusterTopologyShardingNComp, ContainSubstring(comp1aName), HavePrefix(sharding1aName)},
			} {
				By(suit.topology)
				transformer, transCtx, dag := newTransformerNCtx(suit.topology, func(f *testapps.MockClusterFactory) {
					f.AddSharding(sharding1aName, "", "")
				})

				reader := &appsutil.MockReader{
					Objects: []client.Object{
						mockCompObj(transCtx, comp1aName, func(comp *appsv1.Component) {
							comp.Spec.Replicas = 2 // to update
							comp.Status.Phase = appsv1.RunningComponentPhase
						}),
						mockShardingCompObj(transCtx, sharding1aName, func(comp *appsv1.Component) {
							comp.Spec.Replicas = 2 // to update
							comp.Status.Phase = appsv1.RunningComponentPhase
						}),
					},
				}
				transCtx.Client = model.NewGraphClient(reader)
				transCtx.OrigCluster.Generation += 1 // mock cluster spec update

				err := transformer.Transform(transCtx, dag)
				Expect(err).ShouldNot(BeNil())
				Expect(err.Error()).Should(And(ContainSubstring("retry later"), suit.errMatcher))

				// check the updated component
				graphCli := transCtx.Client.(model.GraphClient)
				objs := graphCli.FindAll(dag, &appsv1.Component{})
				Expect(len(objs)).Should(Equal(1))
				comp := objs[0].(*appsv1.Component)
				Expect(component.ShortName(transCtx.Cluster.Name, comp.Name)).Should(suit.updatedNameMatcher)
				Expect(graphCli.IsAction(dag, comp, model.ActionUpdatePtr())).Should(BeTrue())
			}
		})

		It("update - stop", func() {
			for _, suit := range []struct {
				topology           string
				errMatcher         types.GomegaMatcher
				updatedNameMatcher types.GomegaMatcher
			}{
				{clusterTopologyCompNSharding, ContainSubstring(sharding1aName), Equal(comp1aName)},
				{clusterTopologyShardingNComp, ContainSubstring(comp1aName), HavePrefix(sharding1aName)},
			} {
				By(suit.topology)
				transformer, transCtx, dag := newTransformerNCtx(suit.topology, func(f *testapps.MockClusterFactory) {
					f.AddSharding(sharding1aName, "", "").
						AddSharding(sharding2aName, "", "")
				})

				// mock to stop all components and shardings
				reader := &appsutil.MockReader{
					Objects: []client.Object{
						mockCompObj(transCtx, comp1aName, func(comp *appsv1.Component) {
							comp.Status.Phase = appsv1.RunningComponentPhase
						}),
						mockShardingCompObj(transCtx, sharding1aName, func(comp *appsv1.Component) {
							comp.Status.Phase = appsv1.RunningComponentPhase
						}),
					},
				}
				transCtx.Client = model.NewGraphClient(reader)
				for i := range transCtx.components {
					transCtx.components[i].Stop = pointer.Bool(true)
				}
				for i, sharding := range transCtx.shardings {
					transCtx.shardings[i].Template.Stop = pointer.Bool(true)
					for j := range transCtx.shardingComps[sharding.Name] {
						transCtx.shardingComps[sharding.Name][j].Stop = pointer.Bool(true)
					}
				}
				transCtx.OrigCluster.Generation += 1 // mock cluster spec update

				err := transformer.Transform(transCtx, dag)
				Expect(err).ShouldNot(BeNil())
				Expect(err.Error()).Should(And(ContainSubstring("retry later"), suit.errMatcher))

				// should have the first component to update only
				graphCli := transCtx.Client.(model.GraphClient)
				objs := graphCli.FindAll(dag, &appsv1.Component{})
				Expect(len(objs)).Should(Equal(1))
				comp := objs[0].(*appsv1.Component)
				Expect(component.ShortName(transCtx.Cluster.Name, comp.Name)).Should(suit.updatedNameMatcher)
				Expect(graphCli.IsAction(dag, comp, model.ActionUpdatePtr())).Should(BeTrue())
				Expect(comp.Spec.Stop).ShouldNot(BeNil())
				Expect(*comp.Spec.Stop).Should(BeTrue())
			}
		})

		It("provision & update OOD", func() {
			for _, suit := range []struct {
				topology           string
				createdNameMatcher types.GomegaMatcher
				firstMockObjects   func(*clusterTransformContext) []client.Object
				secondMockObjects  func(*clusterTransformContext) []client.Object
			}{
				{
					topology:           clusterTopologyCompNShardingOOD,
					createdNameMatcher: HavePrefix(sharding1aName),
					firstMockObjects: func(transCtx *clusterTransformContext) []client.Object {
						return []client.Object{
							mockCompObj(transCtx, comp1aName, func(comp *appsv1.Component) {
								comp.Status.Phase = appsv1.RunningComponentPhase
							}),
						}
					},
					secondMockObjects: func(transCtx *clusterTransformContext) []client.Object {
						return []client.Object{
							mockShardingCompObj(transCtx, sharding1aName, func(comp *appsv1.Component) {
								comp.Status.Phase = appsv1.RunningComponentPhase
							}),
						}
					},
				},
				{
					topology:           clusterTopologyShardingNCompOOD,
					createdNameMatcher: Equal(comp1aName),
					firstMockObjects: func(transCtx *clusterTransformContext) []client.Object {
						return []client.Object{
							mockShardingCompObj(transCtx, sharding1aName, func(comp *appsv1.Component) {
								comp.Status.Phase = appsv1.RunningComponentPhase
							}),
						}
					},
					secondMockObjects: func(transCtx *clusterTransformContext) []client.Object {
						return []client.Object{
							mockCompObj(transCtx, comp1aName, func(comp *appsv1.Component) {
								comp.Status.Phase = appsv1.RunningComponentPhase
							}),
						}
					},
				},
			} {
				By(suit.topology)
				transformer, transCtx, dag := newTransformerNCtx(suit.topology, func(f *testapps.MockClusterFactory) {
					f.AddSharding(sharding1aName, "", "")
				})

				// mock first component status as running
				reader := &appsutil.MockReader{Objects: suit.firstMockObjects(transCtx)}
				transCtx.Client = model.NewGraphClient(reader)

				// sharding1aName(comp1aName) is not ready (exist) when updating comp1aName(sharding1aName)
				err := transformer.Transform(transCtx, dag)
				Expect(err).Should(BeNil())

				// check another component under provisioning
				graphCli := transCtx.Client.(model.GraphClient)
				objs := graphCli.FindAll(dag, &appsv1.Component{})
				Expect(len(objs)).Should(Equal(1))
				for _, obj := range objs {
					comp := obj.(*appsv1.Component)
					Expect(component.ShortName(transCtx.Cluster.Name, comp.Name)).Should(suit.createdNameMatcher)
					Expect(graphCli.IsAction(dag, comp, model.ActionCreatePtr())).Should(BeTrue())
				}

				// mock another component status as running
				reader.Objects = append(reader.Objects, suit.secondMockObjects(transCtx)...)

				// try again
				err = transformer.Transform(transCtx, newDAG(graphCli, transCtx.Cluster))
				Expect(err).Should(BeNil())
			}
		})
	})

	Context("sharding components", func() {
		It("shard pod anti-affinity", func() {
			transformer, transCtx, dag := newTransformerNCtx(clusterTopologyNoOrders4Sharding, func(f *testapps.MockClusterFactory) {
				f.AddAnnotations(constant.ShardPodAntiAffinityAnnotationKey, strings.Join([]string{sharding1bName, sharding2aName}, ",")).
					AddSharding(sharding1aName, "", "").SetShards(2).
					AddSharding(sharding1bName, "", "").SetShards(2).
					AddSharding(sharding2aName, "", "").SetShards(2).
					AddSharding(sharding2bName, "", "").SetShards(2)
			})
			err := transformer.Transform(transCtx, dag)
			Expect(err).Should(BeNil())

			// check the components
			graphCli := transCtx.Client.(model.GraphClient)
			objs := graphCli.FindAll(dag, &appsv1.Component{})
			Expect(len(objs)).Should(Equal(8))
			for _, obj := range objs {
				comp := obj.(*appsv1.Component)
				Expect(graphCli.IsAction(dag, comp, model.ActionCreatePtr())).Should(BeTrue())

				shardingName := comp.Labels[constant.KBAppShardingNameLabelKey]
				compName := comp.Labels[constant.KBAppComponentLabelKey]
				if shardingName == sharding1aName || shardingName == sharding2bName {
					Expect(comp.Spec.SchedulingPolicy).Should(BeNil())
				} else {
					Expect(comp.Spec.SchedulingPolicy).ShouldNot(BeNil())
					Expect(comp.Spec.SchedulingPolicy.Affinity).ShouldNot(BeNil())
					Expect(comp.Spec.SchedulingPolicy.Affinity.PodAntiAffinity).ShouldNot(BeNil())
					Expect(comp.Spec.SchedulingPolicy.Affinity.PodAntiAffinity.RequiredDuringSchedulingIgnoredDuringExecution).Should(HaveLen(1))
					term := comp.Spec.SchedulingPolicy.Affinity.PodAntiAffinity.RequiredDuringSchedulingIgnoredDuringExecution[0]
					Expect(term.LabelSelector.MatchLabels).Should(HaveKeyWithValue(constant.KBAppShardingNameLabelKey, shardingName))
					Expect(term.LabelSelector.MatchLabels).Should(HaveKeyWithValue(constant.KBAppComponentLabelKey, compName))
					Expect(term.TopologyKey).Should(Equal(corev1.LabelHostname))
				}
			}
		})
	})

	Context("testing component merge functionality", func() {
		var (
			oldCompObj *appsv1.Component
			newCompObj *appsv1.Component
		)

		BeforeEach(func() {
			// Initialize a base component
			oldCompObj = &appsv1.Component{
				Spec: appsv1.ComponentSpec{
					Resources: corev1.ResourceRequirements{
						Limits: corev1.ResourceList{
							corev1.ResourceCPU:    resource.MustParse("1"),
							corev1.ResourceMemory: resource.MustParse("1Gi"),
						},
						Requests: corev1.ResourceList{
							corev1.ResourceCPU:    resource.MustParse("500m"),
							corev1.ResourceMemory: resource.MustParse("512Mi"),
						},
					},
				},
			}
			newCompObj = oldCompObj.DeepCopy()
		})

		It("should return nil when no changes are made", func() {
			result := copyAndMergeComponent(oldCompObj, newCompObj)
			Expect(result).To(BeNil())
		})

		It("should detect annotation changes", func() {
			newCompObj.Annotations = map[string]string{"key": "value"}
			result := copyAndMergeComponent(oldCompObj, newCompObj)
			Expect(result).NotTo(BeNil())
			Expect(result.Annotations).To(HaveKeyWithValue("key", "value"))
		})

		It("should detect label changes", func() {
			newCompObj.Labels = map[string]string{"app": "test"}
			result := copyAndMergeComponent(oldCompObj, newCompObj)
			Expect(result).NotTo(BeNil())
			Expect(result.Labels).To(HaveKeyWithValue("app", "test"))
		})

		It("should detect resource changes", func() {
			// Change CPU resource
			newCompObj.Spec.Resources.Limits[corev1.ResourceCPU] = resource.MustParse("2")
			result := copyAndMergeComponent(oldCompObj, newCompObj)
			Expect(result).NotTo(BeNil())
			Expect(result.Spec.Resources.Limits[corev1.ResourceCPU]).To(Equal(resource.MustParse("2")))
		})

		It("should detect VolumeClaimTemplate changes", func() {
			// Add a volume claim template
			newCompObj.Spec.VolumeClaimTemplates = []appsv1.PersistentVolumeClaimTemplate{
				{
					Name: "app-data",
					Spec: corev1.PersistentVolumeClaimSpec{
						Resources: corev1.VolumeResourceRequirements{
							Requests: corev1.ResourceList{
								corev1.ResourceStorage: resource.MustParse("1Gi"),
							},
						},
					},
				},
			}
			result := copyAndMergeComponent(oldCompObj, newCompObj)
			Expect(result).NotTo(BeNil())
			Expect(result.Spec.VolumeClaimTemplates).To(HaveLen(1))
		})

		It("should normalize CPU resources", func() {
			// 1000m is equivalent to 1
			oldCompObj.Spec.Resources.Limits[corev1.ResourceCPU] = resource.MustParse("1")
			newCompObj.Spec.Resources.Limits[corev1.ResourceCPU] = resource.MustParse("1000m")
			result := copyAndMergeComponent(oldCompObj, newCompObj)
			Expect(result).To(BeNil()) // No change after normalization
		})

		It("should normalize memory resources", func() {
			// 1024Mi is equivalent to 1Gi
			oldCompObj.Spec.Resources.Limits[corev1.ResourceMemory] = resource.MustParse("1Gi")
			newCompObj.Spec.Resources.Limits[corev1.ResourceMemory] = resource.MustParse("1024Mi")
			result := copyAndMergeComponent(oldCompObj, newCompObj)
			Expect(result).To(BeNil()) // No change after normalization

			// 1536.5Mi is equivalent to 1611137024, and 1611137026 = 1611137024 + 2
			oldCompObj.Spec.Resources.Limits[corev1.ResourceMemory] = resource.MustParse("1611137024")
			newCompObj.Spec.Resources.Limits[corev1.ResourceMemory] = resource.MustParse("1536.5Mi")
			result = copyAndMergeComponent(oldCompObj, newCompObj)
			Expect(result).To(BeNil())

			oldCompObj.Spec.Resources.Limits[corev1.ResourceMemory] = resource.MustParse("1611137026")
			newCompObj.Spec.Resources.Limits[corev1.ResourceMemory] = resource.MustParse("1536.5Mi")
			result = copyAndMergeComponent(oldCompObj, newCompObj)
			Expect(result).NotTo(BeNil())

			oldCompObj.Spec.Resources.Limits[corev1.ResourceMemory] = resource.MustParse("1.5Gi")
			newCompObj.Spec.Resources.Limits[corev1.ResourceMemory] = resource.MustParse("1.512Gi")
			result = copyAndMergeComponent(oldCompObj, newCompObj)
			Expect(result).NotTo(BeNil())
		})

		It("should handle nil resource limits", func() {
			oldCompObj.Spec.Resources.Limits = nil
			newCompObj.Spec.Resources.Limits = nil
			result := copyAndMergeComponent(oldCompObj, newCompObj)
			Expect(result).To(BeNil())
		})

		It("should handle nil resource requests", func() {
			oldCompObj.Spec.Resources.Requests = nil
			newCompObj.Spec.Resources.Requests = nil
			result := copyAndMergeComponent(oldCompObj, newCompObj)
			Expect(result).To(BeNil())
		})

		It("should detect changes when adding limits", func() {
			oldCompObj.Spec.Resources.Limits = nil
			result := copyAndMergeComponent(oldCompObj, newCompObj)
			Expect(result).NotTo(BeNil())
			Expect(result.Spec.Resources.Limits).NotTo(BeNil())
		})

		It("should detect changes in VolumeClaimTemplate storage requests", func() {
			vct := appsv1.PersistentVolumeClaimTemplate{
				Name: "app-data",
				Spec: corev1.PersistentVolumeClaimSpec{
					Resources: corev1.VolumeResourceRequirements{
						Requests: corev1.ResourceList{
							corev1.ResourceStorage: resource.MustParse("1Gi"),
						},
					},
				},
			}

			oldCompObj.Spec.VolumeClaimTemplates = []appsv1.PersistentVolumeClaimTemplate{vct}
			newCompObj.Spec.VolumeClaimTemplates = []appsv1.PersistentVolumeClaimTemplate{*vct.DeepCopy()}
			newCompObj.Spec.VolumeClaimTemplates[0].Spec.Resources.Requests[corev1.ResourceStorage] =
				resource.MustParse("2Gi")
			// Change storage request
			result := copyAndMergeComponent(oldCompObj, newCompObj)
			Expect(result).NotTo(BeNil())
			Expect(result.Spec.VolumeClaimTemplates[0].Spec.Resources.Requests[corev1.ResourceStorage]).
				To(Equal(resource.MustParse("2Gi")))
		})

		It("should normalize storage resources in VolumeClaimTemplates", func() {
			vct := appsv1.PersistentVolumeClaimTemplate{
				Name: "app-data",
				Spec: corev1.PersistentVolumeClaimSpec{
					Resources: corev1.VolumeResourceRequirements{
						Requests: corev1.ResourceList{
							corev1.ResourceStorage: resource.MustParse("1Gi"),
						},
					},
				},
			}
			oldCompObj.Spec.VolumeClaimTemplates = []appsv1.PersistentVolumeClaimTemplate{vct}
			newCompObj.Spec.VolumeClaimTemplates = []appsv1.PersistentVolumeClaimTemplate{*vct.DeepCopy()}

			// 1536Mi is equivalent to 1.5Gi
			oldCompObj.Spec.VolumeClaimTemplates[0].Spec.Resources.Requests[corev1.ResourceStorage] =
				resource.MustParse("1.5Gi")
			newCompObj.Spec.VolumeClaimTemplates[0].Spec.Resources.Requests[corev1.ResourceStorage] =
				resource.MustParse("1536Mi")
			result := copyAndMergeComponent(oldCompObj, newCompObj)
			Expect(result).To(BeNil()) // No change after normalization
		})

		It("should handle zero resource values", func() {
			oldCompObj.Spec.Resources.Limits[corev1.ResourceCPU] = resource.MustParse("0")
			newCompObj.Spec.Resources.Limits[corev1.ResourceCPU] = resource.MustParse("0m")

			result := copyAndMergeComponent(oldCompObj, newCompObj)
			Expect(result).To(BeNil()) // No change after normalization
		})

		It("should handle non-standard resource types", func() {
			customResource := "example.com/custom-resource"
			oldCompObj.Spec.Resources.Limits[corev1.ResourceName(customResource)] = resource.MustParse("5")
			newCompObj.Spec.Resources.Limits[corev1.ResourceName(customResource)] = resource.MustParse("10")
			result := copyAndMergeComponent(oldCompObj, newCompObj)
			Expect(result).NotTo(BeNil())
			Expect(result.Spec.Resources.Limits[corev1.ResourceName(customResource)]).
				To(Equal(resource.MustParse("10")))
		})

		It("should detect all changes when multiple fields change", func() {
			newCompObj.Labels = map[string]string{"app": "test"}
			newCompObj.Spec.Resources.Limits[corev1.ResourceCPU] = resource.MustParse("2")
			newCompObj.Spec.Replicas = 3

			result := copyAndMergeComponent(oldCompObj, newCompObj)
			Expect(result).NotTo(BeNil())
			Expect(result.Labels).To(HaveKeyWithValue("app", "test"))
			Expect(result.Spec.Resources.Limits[corev1.ResourceCPU]).To(Equal(resource.MustParse("2")))
			Expect(result.Spec.Replicas).To(Equal(int32(3)))
		})
	})
})
