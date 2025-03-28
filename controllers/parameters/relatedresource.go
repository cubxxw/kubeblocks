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
	"fmt"

	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	appsv1 "github.com/apecloud/kubeblocks/apis/apps/v1"
	"github.com/apecloud/kubeblocks/pkg/constant"
	configctrl "github.com/apecloud/kubeblocks/pkg/controller/configuration"
	"github.com/apecloud/kubeblocks/pkg/controller/render"
	intctrlutil "github.com/apecloud/kubeblocks/pkg/controllerutil"
)

type reconfigureRelatedResource struct {
	ctx        context.Context
	client     client.Client
	configSpec *appsv1.ComponentFileTemplate

	clusterName   string
	componentName string

	configMapObj *corev1.ConfigMap
}

func prepareRelatedResource(reqCtx intctrlutil.RequestCtx, client client.Client, cm *corev1.ConfigMap) (*reconfigureRelatedResource, error) {
	configResources := reconfigureRelatedResource{
		configMapObj:  cm,
		ctx:           reqCtx.Ctx,
		client:        client,
		clusterName:   cm.Labels[constant.AppInstanceLabelKey],
		componentName: cm.Labels[constant.KBAppComponentLabelKey],
	}

	fetcher := configctrl.NewResourceFetcher(&render.ResourceCtx{
		Context:       reqCtx.Ctx,
		Client:        client,
		Namespace:     cm.Namespace,
		ClusterName:   configResources.clusterName,
		ComponentName: configResources.componentName,
	})
	if fetcher.ComponentParameter(); fetcher.Err != nil {
		return nil, fetcher.Err
	}
	if fetcher.ComponentParameterObj == nil {
		return nil, fmt.Errorf("not found configuration object for configmap: %s", cm.Name)
	}
	if err := resolveComponentTemplateSpec(&configResources, fetcher, cm); err != nil {
		return nil, fetcher.Err
	}
	return &configResources, nil
}

func resolveComponentTemplateSpec(resources *reconfigureRelatedResource, fetcher *configctrl.Fetcher, cm *corev1.ConfigMap) error {
	configSpecName, ok := cm.Labels[constant.CMConfigurationSpecProviderLabelKey]
	if !ok {
		return nil
	}

	configSpec := intctrlutil.GetConfigTemplateItem(&fetcher.ComponentParameterObj.Spec, configSpecName)
	if configSpec == nil {
		return fmt.Errorf("not found config spec: %s in configuration[%s]", configSpecName, fetcher.ComponentParameterObj.Name)
	}
	resources.configSpec = configSpec.ConfigSpec
	return nil
}

func (r *reconfigureRelatedResource) componentMatchLabels() map[string]string {
	return map[string]string{
		constant.AppInstanceLabelKey:    r.clusterName,
		constant.KBAppComponentLabelKey: r.componentName,
	}
}
