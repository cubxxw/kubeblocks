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

package controllerutil

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	appsv1 "github.com/apecloud/kubeblocks/apis/apps/v1"
	"github.com/apecloud/kubeblocks/pkg/constant"
	viper "github.com/apecloud/kubeblocks/pkg/viperx"
)

type createVolumeFn func(volumeName string) corev1.Volume

func CreateVolumeIfNotExist(volumes []corev1.Volume, volumeName string, createFn createVolumeFn) []corev1.Volume {
	for _, vol := range volumes {
		if vol.Name == volumeName {
			return volumes
		}
	}
	return append(volumes, createFn(volumeName))
}

func ToCoreV1PVCs(vcts []appsv1.ClusterComponentVolumeClaimTemplate) []corev1.PersistentVolumeClaim {
	storageClassName := func(spec appsv1.PersistentVolumeClaimSpec, defaultStorageClass string) *string {
		if spec.StorageClassName != nil && *spec.StorageClassName != "" {
			return spec.StorageClassName
		}
		if defaultStorageClass != "" {
			return &defaultStorageClass
		}
		return nil
	}
	var pvcs []corev1.PersistentVolumeClaim
	for _, v := range vcts {
		pvcs = append(pvcs, corev1.PersistentVolumeClaim{
			ObjectMeta: metav1.ObjectMeta{
				Name:        v.Name,
				Labels:      v.Labels,
				Annotations: v.Annotations,
			},
			Spec: corev1.PersistentVolumeClaimSpec{
				AccessModes:               v.Spec.AccessModes,
				Resources:                 v.Spec.Resources,
				StorageClassName:          storageClassName(v.Spec, viper.GetString(constant.CfgKeyDefaultStorageClass)),
				VolumeMode:                v.Spec.VolumeMode,
				VolumeAttributesClassName: v.Spec.VolumeAttributesClassName,
			},
		})
	}
	return pvcs
}

func ToCoreV1PVCTs(vcts []appsv1.ClusterComponentVolumeClaimTemplate) []corev1.PersistentVolumeClaimTemplate {
	pvcs := ToCoreV1PVCs(vcts)
	pvct := func(i int) corev1.PersistentVolumeClaimTemplate {
		return corev1.PersistentVolumeClaimTemplate{
			ObjectMeta: pvcs[i].ObjectMeta,
			Spec:       pvcs[i].Spec,
		}
	}
	var pvcts []corev1.PersistentVolumeClaimTemplate
	for i := range pvcs {
		pvcts = append(pvcts, pvct(i))
	}
	return pvcts
}
