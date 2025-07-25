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

package apps

import (
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	kbappsv1 "github.com/apecloud/kubeblocks/apis/apps/v1"
	workloads "github.com/apecloud/kubeblocks/apis/workloads/v1"
	"github.com/apecloud/kubeblocks/pkg/constant"
)

type MockInstanceSetFactory struct {
	BaseFactory[workloads.InstanceSet, *workloads.InstanceSet, MockInstanceSetFactory]
}

func NewInstanceSetFactory(namespace, name string, clusterName string, componentName string) *MockInstanceSetFactory {
	f := &MockInstanceSetFactory{}
	f.Init(namespace, name,
		&workloads.InstanceSet{
			ObjectMeta: metav1.ObjectMeta{
				Labels: map[string]string{
					constant.AppInstanceLabelKey:    clusterName,
					constant.KBAppComponentLabelKey: componentName,
					constant.AppManagedByLabelKey:   constant.AppName,
				},
			},
			Spec: workloads.InstanceSetSpec{
				Selector: &metav1.LabelSelector{
					MatchLabels: map[string]string{
						constant.AppInstanceLabelKey:    clusterName,
						constant.KBAppComponentLabelKey: componentName,
						constant.AppManagedByLabelKey:   constant.AppName,
					},
				},
				Template: corev1.PodTemplateSpec{
					ObjectMeta: metav1.ObjectMeta{
						Labels: map[string]string{
							constant.AppInstanceLabelKey:    clusterName,
							constant.KBAppComponentLabelKey: componentName,
							constant.AppManagedByLabelKey:   constant.AppName,
						},
					},
				},
				InstanceUpdateStrategy: &workloads.InstanceUpdateStrategy{
					Type: kbappsv1.OnDeleteStrategyType,
				},
			},
		}, f)
	return f
}

func (factory *MockInstanceSetFactory) SetReplicas(replicas int32) *MockInstanceSetFactory {
	factory.Get().Spec.Replicas = &replicas
	return factory
}

func (factory *MockInstanceSetFactory) SetRoles(roles []workloads.ReplicaRole) *MockInstanceSetFactory {
	factory.Get().Spec.Roles = roles
	return factory
}

func (factory *MockInstanceSetFactory) AddVolume(volume corev1.Volume) *MockInstanceSetFactory {
	volumes := &factory.Get().Spec.Template.Spec.Volumes
	*volumes = append(*volumes, volume)
	return factory
}

func (factory *MockInstanceSetFactory) AddConfigmapVolume(volumeName string, configmapName string) *MockInstanceSetFactory {
	volume := corev1.Volume{
		Name: volumeName,
		VolumeSource: corev1.VolumeSource{
			ConfigMap: &corev1.ConfigMapVolumeSource{
				LocalObjectReference: corev1.LocalObjectReference{Name: configmapName},
			},
		},
	}
	factory.AddVolume(volume)
	return factory
}

func (factory *MockInstanceSetFactory) AddVolumeClaimTemplate(pvc corev1.PersistentVolumeClaim) *MockInstanceSetFactory {
	volumeClaimTpls := &factory.Get().Spec.VolumeClaimTemplates
	*volumeClaimTpls = append(*volumeClaimTpls, pvc)
	return factory
}

func (factory *MockInstanceSetFactory) SetPVCRetentionPolicy(retentionPolicy *workloads.PersistentVolumeClaimRetentionPolicy) *MockInstanceSetFactory {
	factory.Get().Spec.PersistentVolumeClaimRetentionPolicy = retentionPolicy
	return factory
}

func (factory *MockInstanceSetFactory) AddContainer(container corev1.Container) *MockInstanceSetFactory {
	containers := &factory.Get().Spec.Template.Spec.Containers
	*containers = append(*containers, container)
	return factory
}

func (factory *MockInstanceSetFactory) AddConfigs(config ...workloads.ConfigTemplate) *MockInstanceSetFactory {
	configs := &factory.Get().Spec.Configs
	*configs = append(*configs, config...)
	return factory
}

func (factory *MockInstanceSetFactory) SetInstanceUpdateStrategy(instanceUpdateStrategy *workloads.InstanceUpdateStrategy) *MockInstanceSetFactory {
	factory.Get().Spec.InstanceUpdateStrategy = instanceUpdateStrategy
	return factory
}

func (factory *MockInstanceSetFactory) SetFlatInstanceOrdinal(flatInstanceOrdinal bool) *MockInstanceSetFactory {
	factory.Get().Spec.FlatInstanceOrdinal = flatInstanceOrdinal
	return factory
}

func (factory *MockInstanceSetFactory) SetPodManagementPolicy(podManagementPolicy appsv1.PodManagementPolicyType) *MockInstanceSetFactory {
	factory.Get().Spec.PodManagementPolicy = podManagementPolicy
	return factory
}
