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
	"os"

	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	appsv1alpha1 "github.com/apecloud/kubeblocks/apis/apps/v1alpha1"
	"github.com/apecloud/kubeblocks/pkg/constant"
	podutil "github.com/apecloud/kubeblocks/pkg/controllerutil"
	viper "github.com/apecloud/kubeblocks/pkg/viperx"
)

const (
	// StatefulSetSpec.Spec.MinReadySeconds
	// units: s
	defaultMinReadySeconds = 10
)

type rollingUpgradePolicy struct {
}

func init() {
	RegisterPolicy(appsv1alpha1.RollingPolicy, &rollingUpgradePolicy{})
	if err := viper.BindEnv(constant.PodMinReadySecondsEnv); err != nil {
		os.Exit(-1)
	}
	viper.SetDefault(constant.PodMinReadySecondsEnv, defaultMinReadySeconds)
}

func (r *rollingUpgradePolicy) Upgrade(params reconfigureParams) (ReturnedStatus, error) {
	return performRollingUpgrade(params, GetInstanceSetRollingUpgradeFuncs())
}

func (r *rollingUpgradePolicy) GetPolicyName() string {
	return string(appsv1alpha1.RollingPolicy)
}

func canPerformUpgrade(pods []corev1.Pod, params reconfigureParams) bool {
	target := params.getTargetReplicas()
	// TODO(xingran&zhangtao): review this logic
	return len(pods) == target

	/*
		if len(pods) < target {
			params.Ctx.Log.Info(fmt.Sprintf("component pods are not all ready, %d pods are ready, which is less than the expected replicas(%d).", len(pods), target))
			return false
		}
		return true
	*/
}

func performRollingUpgrade(params reconfigureParams, funcs RollingUpgradeFuncs) (ReturnedStatus, error) {
	pods, err := funcs.GetPodsFunc(params)
	if err != nil {
		return makeReturnedStatus(ESFailedAndRetry), err
	}

	var (
		rollingReplicas = params.maxRollingReplicas()
		configKey       = params.getConfigKey()
		configVersion   = params.getTargetVersionHash()
	)

	if !canPerformUpgrade(pods, params) {
		return makeReturnedStatus(ESRetry), nil
	}

	podStats := staticPodStats(pods, params.getTargetReplicas(), params.podMinReadySeconds())
	podWins := markDynamicCursor(pods, podStats, configKey, configVersion, rollingReplicas)
	if !validPodState(podWins) {
		params.Ctx.Log.Info("wait for pod stat ready.")
		return makeReturnedStatus(ESRetry), nil
	}

	waitRollingPods := podWins.getWaitRollingPods()
	if len(waitRollingPods) == 0 {
		return makeReturnedStatus(ESNone, withSucceed(int32(podStats.targetReplica)), withExpected(int32(podStats.targetReplica))), nil
	}

	for _, pod := range waitRollingPods {
		if podStats.isUpdating(&pod) {
			params.Ctx.Log.Info("pod is in rolling update.", "pod name", pod.Name)
			continue
		}
		if err := funcs.RestartContainerFunc(&pod, params.Ctx.Ctx, params.ContainerNames, params.ReconfigureClientFactory); err != nil {
			return makeReturnedStatus(ESFailedAndRetry), err
		}
		if err := updatePodLabelsWithConfigVersion(&pod, configKey, configVersion, params.Client, params.Ctx.Ctx); err != nil {
			return makeReturnedStatus(ESFailedAndRetry), err
		}
	}

	return makeReturnedStatus(ESRetry,
		withExpected(int32(podStats.targetReplica)),
		withSucceed(int32(len(podStats.updated)+len(podStats.updating)))), nil
}

func validPodState(wind switchWindow) bool {
	for i := 0; i < wind.begin; i++ {
		pod := &wind.pods[i]
		if !wind.isReady(pod) {
			return false
		}
	}
	return true
}

func markDynamicCursor(pods []corev1.Pod, podsStats *componentPodStats, configKey, currentVersion string, rollingReplicas int32) switchWindow {
	podWindows := switchWindow{
		end:               0,
		begin:             len(pods),
		pods:              pods,
		componentPodStats: podsStats,
	}

	// find update last
	for i := podsStats.targetReplica - 1; i >= 0; i-- {
		pod := &pods[i]
		if !podutil.IsMatchConfigVersion(pod, configKey, currentVersion) {
			podWindows.end = i + 1
			break
		}
		if !podsStats.isAvailable(pod) {
			podsStats.updating[pod.Name] = pod
			podWindows.end = i + 1
			break
		}
		podsStats.updated[pod.Name] = pod
	}

	podWindows.begin = max(podWindows.end-int(rollingReplicas), 0)
	for i := podWindows.begin; i < podWindows.end; i++ {
		pod := &pods[i]
		if podutil.IsMatchConfigVersion(pod, configKey, currentVersion) {
			podsStats.updating[pod.Name] = pod
		}
	}
	return podWindows
}

func staticPodStats(pods []corev1.Pod, targetReplicas int, minReadySeconds int32) *componentPodStats {
	podsStats := &componentPodStats{
		updated:       make(map[string]*corev1.Pod),
		updating:      make(map[string]*corev1.Pod),
		available:     make(map[string]*corev1.Pod),
		ready:         make(map[string]*corev1.Pod),
		targetReplica: targetReplicas,
	}

	for i := 0; i < len(pods); i++ {
		pod := &pods[i]
		switch {
		case podutil.IsAvailable(pod, minReadySeconds):
			podsStats.available[pod.Name] = pod
		case podutil.PodIsReady(pod):
			podsStats.ready[pod.Name] = pod
		default:
		}
	}
	return podsStats
}

type componentPodStats struct {
	// failed to start pod
	ready     map[string]*corev1.Pod
	available map[string]*corev1.Pod

	// updated pod count
	updated  map[string]*corev1.Pod
	updating map[string]*corev1.Pod

	// expected pod
	targetReplica int
}

func (s *componentPodStats) isAvailable(pod *corev1.Pod) bool {
	_, ok := s.available[pod.Name]
	return ok
}

func (s *componentPodStats) isReady(pod *corev1.Pod) bool {
	_, ok := s.ready[pod.Name]
	return ok || s.isAvailable(pod)
}

func (s *componentPodStats) isUpdating(pod *corev1.Pod) bool {
	_, ok := s.updating[pod.Name]
	return ok
}

type switchWindow struct {
	begin int
	end   int

	pods []corev1.Pod
	*componentPodStats
}

func (w *switchWindow) getWaitRollingPods() []corev1.Pod {
	return w.pods[w.begin:w.end]
}

func updatePodLabelsWithConfigVersion(pod *corev1.Pod, labelKey, configVersion string, cli client.Client, ctx context.Context) error {
	patch := client.MergeFrom(pod.DeepCopy())
	if pod.Labels == nil {
		pod.Labels = make(map[string]string, 1)
	}
	pod.Labels[labelKey] = configVersion
	return cli.Patch(ctx, pod, patch)
}
