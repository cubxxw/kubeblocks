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

package parameters

import (
	"context"
	"reflect"

	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/record"
	"k8s.io/utils/pointer"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	appsv1 "github.com/apecloud/kubeblocks/apis/apps/v1"
	parametersv1alpha1 "github.com/apecloud/kubeblocks/apis/parameters/v1alpha1"
	configcore "github.com/apecloud/kubeblocks/pkg/configuration/core"
	intctrlutil "github.com/apecloud/kubeblocks/pkg/controllerutil"
	"github.com/apecloud/kubeblocks/pkg/generics"
)

// ParameterTemplateExtensionReconciler reconciles a ParameterTemplateExtension object
type ParameterTemplateExtensionReconciler struct {
	client.Client
	Scheme   *runtime.Scheme
	Recorder record.EventRecorder
}

// +kubebuilder:rbac:groups=apps.kubeblocks.io,resources=clusters,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=apps.kubeblocks.io,resources=clusters/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=apps.kubeblocks.io,resources=clusters/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the ParameterTemplateExtension object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.14.4/pkg/reconcile
func (r *ParameterTemplateExtensionReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	reqCtx := intctrlutil.RequestCtx{
		Ctx:      ctx,
		Req:      req,
		Recorder: r.Recorder,
		Log: log.FromContext(ctx).
			WithName("ParameterExtensionReconciler").
			WithValues("Namespace", req.Namespace, "ParameterExtension", req.Name),
	}

	cluster := &appsv1.Cluster{}
	if err := r.Client.Get(reqCtx.Ctx, reqCtx.Req.NamespacedName, cluster); err != nil {
		return intctrlutil.CheckedRequeueWithError(err, reqCtx.Log, "")
	}
	if cluster.IsDeleting() {
		return intctrlutil.Reconciled()
	}
	if !intctrlutil.ObjectAPIVersionSupported(cluster) {
		return intctrlutil.Reconciled()
	}
	return r.reconcile(reqCtx, cluster)
}

// SetupWithManager sets up the controller with the Manager.
func (r *ParameterTemplateExtensionReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&appsv1.Cluster{}).
		Watches(&parametersv1alpha1.ComponentParameter{}, handler.EnqueueRequestsFromMapFunc(filterComponentParameterResources)).
		Complete(r)
}

func filterComponentParameterResources(_ context.Context, object client.Object) []reconcile.Request {
	cr := object.(*parametersv1alpha1.ComponentParameter)
	return []reconcile.Request{{
		NamespacedName: client.ObjectKey{
			Name:      cr.Spec.ClusterName,
			Namespace: cr.Namespace,
		}}}
}

func (r *ParameterTemplateExtensionReconciler) reconcile(reqCtx intctrlutil.RequestCtx, runningCluster *appsv1.Cluster) (ctrl.Result, error) {
	expectedCluster, err := updateConfigsForParameterTemplate(reqCtx, r.Client, runningCluster)
	if err != nil {
		return intctrlutil.RequeueWithError(err, reqCtx.Log, "")
	}
	return r.update(reqCtx, runningCluster, expectedCluster)
}

func updateConfigsForParameterTemplate(reqCtx intctrlutil.RequestCtx, reader client.Reader, cluster *appsv1.Cluster) (*appsv1.Cluster, error) {
	updateParameterObject := func(compName string, config *appsv1.ClusterComponentConfig) error {
		cmKey := client.ObjectKey{
			Name:      configcore.GetComponentCfgName(cluster.Name, compName, pointer.StringDeref(config.Name, "")),
			Namespace: cluster.Namespace,
		}
		cm := corev1.ConfigMap{}
		if err := reader.Get(reqCtx.Ctx, cmKey, &cm); err != nil {
			return client.IgnoreNotFound(err)
		}
		config.ConfigMap = &corev1.ConfigMapVolumeSource{
			LocalObjectReference: corev1.LocalObjectReference{Name: cm.Name},
		}
		return nil
	}
	defaultExternalManagedHandle := func(compSpec *appsv1.ClusterComponentSpec, tpl *appsv1.ComponentFileTemplate) error {
		if !pointer.BoolDeref(tpl.ExternalManaged, false) {
			return nil
		}

		config := appsv1.ClusterComponentConfig{
			Name: pointer.String(tpl.Name),
		}
		if err := updateParameterObject(compSpec.Name, &config); err != nil {
			return err
		}
		if config.ConfigMap != nil {
			compSpec.Configs = append(compSpec.Configs, config)
		}
		return nil
	}

	expectedCluster := cluster.DeepCopy()
	for i := range expectedCluster.Spec.ComponentSpecs {
		compSpec := &expectedCluster.Spec.ComponentSpecs[i]
		templates, err := getSupportedParameterTemplates(reqCtx.Ctx, reader, cluster.Name, compSpec.Name, cluster.Namespace)
		if err != nil {
			return nil, err
		}
		for _, template := range templates {
			index := generics.FindFirstFunc(compSpec.Configs, func(config appsv1.ClusterComponentConfig) bool {
				return config.Name != nil && *config.Name == template.Name
			})
			if index < 0 {
				err = defaultExternalManagedHandle(compSpec, template)
			} else if needUpdateExternalManaged(compSpec.Configs[index]) {
				err = updateParameterObject(compSpec.Name, &compSpec.Configs[index])
			}
			if err != nil {
				return nil, err
			}
		}
	}
	return expectedCluster, nil
}

func (r *ParameterTemplateExtensionReconciler) update(reqCtx intctrlutil.RequestCtx, running, expected *appsv1.Cluster) (ctrl.Result, error) {
	if reflect.DeepEqual(running.Spec, expected.Spec) {
		return ctrl.Result{}, nil
	}
	if err := r.Client.Patch(reqCtx.Ctx, expected, client.MergeFrom(running)); err != nil {
		return intctrlutil.CheckedRequeueWithError(err, reqCtx.Log, "")
	}
	return intctrlutil.Reconciled()
}

func needUpdateExternalManaged(config appsv1.ClusterComponentConfig) bool {
	if !pointer.BoolDeref(config.ExternalManaged, false) {
		return false
	}
	if pointer.StringDeref(config.Name, "") == "" {
		return false
	}
	if config.ConfigMap != nil {
		return false
	}
	return true
}

func getSupportedParameterTemplates(ctx context.Context, reader client.Reader, clusterName, compName, namespace string) ([]*appsv1.ComponentFileTemplate, error) {
	parameterKey := client.ObjectKey{
		Name:      configcore.GenerateComponentConfigurationName(clusterName, compName),
		Namespace: namespace,
	}
	parameterCr := &parametersv1alpha1.ComponentParameter{}
	if err := reader.Get(ctx, parameterKey, parameterCr); err != nil {
		if !apierrors.IsNotFound(err) {
			return nil, err
		}
		return nil, nil
	}

	transformer := func(item parametersv1alpha1.ConfigTemplateItemDetail) *appsv1.ComponentFileTemplate {
		return item.ConfigSpec
	}
	parameterTemplates := generics.Map(parameterCr.Spec.ConfigItemDetails, transformer)
	return parameterTemplates, nil
}
