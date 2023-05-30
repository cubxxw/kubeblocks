/*
Copyright (C) 2022-2023 ApeCloud Co., Ltd

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

package replication

import (
	"context"
	"fmt"
	"strings"

	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	appsv1alpha1 "github.com/apecloud/kubeblocks/apis/apps/v1alpha1"
	componentutil "github.com/apecloud/kubeblocks/controllers/apps/components/util"
	"github.com/apecloud/kubeblocks/internal/constant"
	intctrlcomputil "github.com/apecloud/kubeblocks/internal/controller/component"
	intctrlutil "github.com/apecloud/kubeblocks/internal/controllerutil"
)

// ProbeDetectManager implements the SwitchDetectManager interface with KubeBlocks Probe.
type ProbeDetectManager struct{}

// SwitchActionWithJobHandler implements the SwitchActionHandler interface with executing switch commands by k8s Job.
type SwitchActionWithJobHandler struct{}

// SwitchElectionRoleFilter implements the SwitchElectionFilter interface and is used to filter the instances which role cannot be elected as candidate primary.
type SwitchElectionRoleFilter struct{}

// SwitchElectionHealthFilter implements the SwitchElectionFilter interface and is used to filter unhealthy instances that cannot be selected as candidate primary.
type SwitchElectionHealthFilter struct{}

// SwitchRoleInfoList is a sort.Interface that Sorts a list of SwitchRoleInfo based on LagDetectInfo value.
type SwitchRoleInfoList []*SwitchRoleInfo

const (
	SwitchElectionRoleFilterName   = "SwitchElectionRoleFilter"
	SwitchElectionHealthFilterName = "SwitchElectionHealthFilter"
)

// Environment names for switchStatements
const (
	KBSwitchPromoteStmtEnvName = "KB_SWITCH_PROMOTE_STATEMENT"
	KBSwitchDemoteStmtEnvName  = "KB_SWITCH_DEMOTE_STATEMENT"
	KBSwitchFollowStmtEnvName  = "KB_SWITCH_FOLLOW_STATEMENT"

	KBSwitchOldPrimaryRoleName = "KB_OLD_PRIMARY_ROLE_NAME"
	KBSwitchNewPrimaryRoleName = "KB_NEW_PRIMARY_ROLE_NAME"

	KBSwitchRoleEndPoint = "KB_SWITCH_ROLE_ENDPOINT"
)

const (
	KBSwitchJobLabelKey      = "kubeblocks.io/switch-job"
	KBSwitchJobLabelValue    = "kb-switch-job"
	KBSwitchJobNamePrefix    = "kb-switch-job"
	KBSwitchJobContainerName = "switch-job-container"
)

var _ SwitchDetectManager = &ProbeDetectManager{}

var _ SwitchActionHandler = &SwitchActionWithJobHandler{}

var defaultSwitchElectionFilters = []func() SwitchElectionFilter{
	newSwitchElectionHealthFilter,
	newSwitchElectionRoleFilter,
}

// HandleReplicationSetHASwitch handles high-availability switching of a single replication workload under current cluster.
func HandleReplicationSetHASwitch(ctx context.Context,
	cli client.Client,
	cluster *appsv1alpha1.Cluster,
	clusterCompSpec *appsv1alpha1.ClusterComponentSpec) error {
	if clusterCompSpec == nil {
		return fmt.Errorf("cluster componentSpec can not be nil")
	}

	compDef, err := filterReplicationWorkload(ctx, cli, cluster, clusterCompSpec.Name)
	if err != nil {
		return err
	}
	if compDef == nil {
		return nil
	}

	candidateInstanceChanged, currentPrimaryInstanceName, err := componentutil.CheckCandidateInstanceChanged(ctx, cli, cluster, clusterCompSpec.Name)
	if err != nil {
		return err
	}
	// there is no need to perform HA operation when candidateInstance has not changed
	if !candidateInstanceChanged {
		return nil
	}

	// create a new Switch object
	s := newSwitch(ctx, cli, cluster, compDef, clusterCompSpec, nil, nil, nil, nil, nil)

	// initialize switchInstance according to the candidateInstance
	candidateInstanceName := fmt.Sprintf("%s-%s-%d", cluster.Name, clusterCompSpec.Name, clusterCompSpec.CandidateInstance.Index)
	if err := s.initSwitchInstance(currentPrimaryInstanceName, candidateInstanceName); err != nil {
		return err
	}

	// health detection, role detection, delay detection of oldPrimaryInstance and candidateInstance
	s.detection(true)
	if err := checkSwitchStatus(s.SwitchStatus); err != nil {
		return err
	}

	// make switch decision, if returns true, then start to do switch action, otherwise returns fail
	if s.decision() {
		if err := s.doSwitch(); err != nil {
			return err
		}
	} else {
		return checkSwitchStatus(s.SwitchStatus)
	}

	// switch succeed, update role labels
	if err := s.updateRoleLabel(); err != nil {
		return err
	}

	// clean job if execute switch commands by k8s job.
	if err := cleanSwitchCmdJobs(s); err != nil {
		return err
	}

	return nil
}

// Len is the implementation of the sort.Interface, calculate the length of the list of SwitchRoleInfoList.
func (sl SwitchRoleInfoList) Len() int {
	return len(sl)
}

// Swap is the implementation of the sort.Interface, exchange two items in SwitchRoleInfoList.
func (sl SwitchRoleInfoList) Swap(i, j int) {
	sl[i], sl[j] = sl[j], sl[i]
}

// Less is the implementation of the sort.Interface, sort the SwitchRoleInfo with LagDetectInfo.
func (sl SwitchRoleInfoList) Less(i, j int) bool {
	return *sl[i].LagDetectInfo < *sl[j].LagDetectInfo
}

func (f *SwitchElectionRoleFilter) name() string {
	return SwitchElectionRoleFilterName
}

// filter is used to filter the instance which role cannot be elected as candidate primary.
func (f *SwitchElectionRoleFilter) filter(roleInfoList []*SwitchRoleInfo) ([]*SwitchRoleInfo, error) {
	var filterRoles []*SwitchRoleInfo
	for _, roleInfo := range roleInfoList {
		if roleInfo.RoleDetectInfo == nil {
			// REVIEW/TODO: need avoid using dynamic error string, this is bad for
			// error type checking (errors.Is)
			return nil, fmt.Errorf("pod %s RoleDetectInfo is nil, pls check", roleInfo.Pod.Name)
		}
		isPrimaryPod, err := checkObjRoleLabelIsPrimary(roleInfo.Pod)
		if err != nil {
			return filterRoles, err
		}
		if string(*roleInfo.RoleDetectInfo) != constant.Primary && !isPrimaryPod {
			filterRoles = append(filterRoles, roleInfo)
		}
	}
	return filterRoles, nil
}

// newSwitchElectionRoleFilter initializes a SwitchElectionRoleFilter and returns it.
func newSwitchElectionRoleFilter() SwitchElectionFilter {
	return &SwitchElectionHealthFilter{}
}

func (f *SwitchElectionHealthFilter) name() string {
	return SwitchElectionHealthFilterName
}

// filter is used to filter unhealthy instances that cannot be selected as candidate primary.
func (f *SwitchElectionHealthFilter) filter(roleInfoList []*SwitchRoleInfo) ([]*SwitchRoleInfo, error) {
	var filterRoles []*SwitchRoleInfo
	for _, roleInfo := range roleInfoList {
		if roleInfo.HealthDetectInfo == nil {
			// REVIEW/TODO: need avoid using dynamic error string, this is bad for
			// error type checking (errors.Is)
			return nil, fmt.Errorf("pod %s HealthDetectInfo is nil, pls check", roleInfo.Pod.Name)
		}
		if *roleInfo.HealthDetectInfo {
			filterRoles = append(filterRoles, roleInfo)
		}
	}
	return filterRoles, nil
}

// newSwitchElectionHealthFilter initializes a SwitchElectionHealthFilter and returns it.
func newSwitchElectionHealthFilter() SwitchElectionFilter {
	return &SwitchElectionHealthFilter{}
}

// buildExecSwitchCommandEnvs builds a series of envs for subsequent switching actions.
func (handler *SwitchActionWithJobHandler) buildExecSwitchCommandEnvs(s *Switch) ([]corev1.EnvVar, error) {
	var switchEnvs []corev1.EnvVar

	// replace secret env and merge envs defined in switchCmdExecutorConfig
	replaceSwitchCmdExecutorConfigEnv(s.SwitchResource.Cluster.Name, s.SwitchResource.CompDef.ReplicationSpec.SwitchCmdExecutorConfig)
	switchEnvs = append(switchEnvs, s.SwitchResource.CompDef.ReplicationSpec.SwitchCmdExecutorConfig.Env...)

	// inject the new primary info into the environment variable
	svcName := strings.Join([]string{s.SwitchResource.Cluster.Name, s.SwitchResource.CompSpec.Name, "headless"}, "-")
	primaryEnvs := []corev1.EnvVar{
		{
			Name:  KBSwitchOldPrimaryRoleName,
			Value: fmt.Sprintf("%s.%s", s.SwitchInstance.OldPrimaryRole.Pod.Name, svcName),
		},
		{
			Name:  KBSwitchNewPrimaryRoleName,
			Value: fmt.Sprintf("%s.%s", s.SwitchInstance.CandidatePrimaryRole.Pod.Name, svcName),
		},
	}
	switchEnvs = append(switchEnvs, primaryEnvs...)

	// inject switchStatements as env variables
	switchStatements, err := getSwitchStatementsBySwitchPolicyType(s.SwitchResource.CompSpec.SwitchPolicy.Type, s.SwitchResource.CompDef.ReplicationSpec)
	if err != nil {
		return nil, err
	}
	promoteStmtEnv := corev1.EnvVar{
		Name:  KBSwitchPromoteStmtEnvName,
		Value: strings.Join(switchStatements.Promote, " "),
	}
	demoteStmtEnv := corev1.EnvVar{
		Name:  KBSwitchDemoteStmtEnvName,
		Value: strings.Join(switchStatements.Demote, " "),
	}
	followStmtEnv := corev1.EnvVar{
		Name:  KBSwitchFollowStmtEnvName,
		Value: strings.Join(switchStatements.Follow, " "),
	}
	switchEnvs = append(switchEnvs, promoteStmtEnv, demoteStmtEnv, followStmtEnv)

	return switchEnvs, nil
}

// execSwitchCommands executes switch commands with k8s job.
func (handler *SwitchActionWithJobHandler) execSwitchCommands(s *Switch, switchEnvs []corev1.EnvVar) error {
	if s.SwitchResource.CompDef.ReplicationSpec.SwitchCmdExecutorConfig == nil {
		return fmt.Errorf("switchCmdExecutorConfig and SwitchSteps can not be nil")
	}
	for i, switchStep := range s.SwitchResource.CompDef.ReplicationSpec.SwitchCmdExecutorConfig.SwitchSteps {
		cmdJobs, err := renderAndCreateSwitchCmdJobs(s, switchEnvs, switchStep, i)
		if err != nil {
			return err
		}
		if err := checkSwitchCmdJobSucceed(s, cmdJobs); err != nil {
			return err
		}
	}
	return nil
}

// healthDetect is the implementation of the SwitchDetectManager interface, which gets health detection information by actively calling the API provided by the probe
// TODO(xingran) Wait for the probe interface to be ready before implementation
func (pdm *ProbeDetectManager) healthDetect(pod *corev1.Pod) (*HealthDetectResult, error) {
	var res HealthDetectResult = true
	return &res, nil
}

// roleDetect is the implementation of the SwitchDetectManager interface, which gets role detection information by actively calling the API provided by the probe
// TODO(xingran) Wait for the probe interface to be ready before implementation
func (pdm *ProbeDetectManager) roleDetect(pod *corev1.Pod) (*RoleDetectResult, error) {
	var res RoleDetectResult
	role := pod.Labels[constant.RoleLabelKey]
	res = DetectRoleSecondary
	if role == constant.Primary {
		res = DetectRolePrimary
	}
	return &res, nil
}

// lagDetect is the implementation of the SwitchDetectManager interface, which gets replication lag detection information by actively calling the API provided by the probe
// TODO(xingran) Wait for the probe interface to be ready before implementation
func (pdm *ProbeDetectManager) lagDetect(pod *corev1.Pod) (*LagDetectResult, error) {
	var res LagDetectResult = 0
	return &res, nil
}

// getSwitchStatementsBySwitchPolicyType gets the SwitchStatements corresponding to switchPolicyType
func getSwitchStatementsBySwitchPolicyType(switchPolicyType appsv1alpha1.SwitchPolicyType,
	replicationSpec *appsv1alpha1.ReplicationSetSpec) (*appsv1alpha1.SwitchStatements, error) {
	if replicationSpec == nil || len(replicationSpec.SwitchPolicies) == 0 {
		return nil, fmt.Errorf("replicationSpec and replicationSpec.SwitchPolicies can not be nil")
	}
	for _, switchPolicy := range replicationSpec.SwitchPolicies {
		if switchPolicy.Type == switchPolicyType {
			return switchPolicy.SwitchStatements, nil
		}
	}
	return nil, fmt.Errorf("cannot find mapping switchStatements of switchPolicyType %s", switchPolicyType)
}

// replaceSwitchCmdExecutorConfigEnv replaces switch execute config secret env.
func replaceSwitchCmdExecutorConfigEnv(clusterName string, switchCmdExecuteConfig *appsv1alpha1.SwitchCmdExecutorConfig) {
	namedValuesMap := intctrlcomputil.GetEnvReplacementMapForConnCredential(clusterName)
	if switchCmdExecuteConfig != nil {
		switchCmdExecuteConfig.Env = intctrlcomputil.ReplaceSecretEnvVars(namedValuesMap, switchCmdExecuteConfig.Env)
	}
}

// checkSwitchStatus checks the status of every phase of Switch
func checkSwitchStatus(status *SwitchStatus) error {
	if status.SwitchPhaseStatus != SwitchPhaseStatusSucceed {
		return fmt.Errorf(status.Reason)
	}
	return nil
}

// renderAndCreateSwitchCmdJobs renders and creates jobs to execute the switch command.
func renderAndCreateSwitchCmdJobs(s *Switch, switchEnvs []corev1.EnvVar,
	switchStep appsv1alpha1.SwitchStep, switchStepIndex int) ([]*batchv1.Job, error) {
	var enginePods []*corev1.Pod
	var cmdJobs []*batchv1.Job
	switch switchStep.Role {
	case appsv1alpha1.NewPrimary:
		enginePods = append(enginePods, s.SwitchInstance.CandidatePrimaryRole.Pod)
	case appsv1alpha1.OldPrimary:
		enginePods = append(enginePods, s.SwitchInstance.OldPrimaryRole.Pod)
	case appsv1alpha1.Secondaries:
		for _, pod := range s.SwitchInstance.SecondariesRole {
			enginePods = append(enginePods, pod.Pod)
		}
	}

	renderJob := func(jobName string, switchEnvs []corev1.EnvVar) *batchv1.Job {
		job := &batchv1.Job{
			ObjectMeta: metav1.ObjectMeta{
				Namespace: s.SwitchResource.Cluster.Namespace,
				Name:      jobName,
				Labels:    getSwitchCmdJobLabel(s.SwitchResource.Cluster.Name, s.SwitchResource.CompSpec.Name),
			},
			Spec: batchv1.JobSpec{
				Template: corev1.PodTemplateSpec{
					ObjectMeta: metav1.ObjectMeta{
						Namespace: s.SwitchResource.Cluster.Namespace,
						Name:      jobName},
					Spec: corev1.PodSpec{
						RestartPolicy: corev1.RestartPolicyNever,
						Containers: []corev1.Container{
							{
								Name:            KBSwitchJobContainerName,
								Image:           s.SwitchResource.CompDef.ReplicationSpec.SwitchCmdExecutorConfig.Image,
								ImagePullPolicy: corev1.PullIfNotPresent,
								Command:         switchStep.Command,
								Args:            switchStep.Args,
								Env:             switchEnvs,
							},
						},
					},
				},
			},
		}
		if len(s.SwitchResource.Cluster.Spec.Tolerations) > 0 {
			job.Spec.Template.Spec.Tolerations = s.SwitchResource.Cluster.Spec.Tolerations
		}
		return job
	}

	for index, enginePod := range enginePods {
		jobName := fmt.Sprintf("%s-%s-%s-%d-%d", KBSwitchJobNamePrefix,
			s.SwitchResource.CompSpec.Name, strings.ToLower(string(switchStep.Role)), switchStepIndex, index)
		svcName := strings.Join([]string{s.SwitchResource.Cluster.Name, s.SwitchResource.CompSpec.Name, "headless"}, "-")
		switchEnvs = append(switchEnvs, corev1.EnvVar{
			Name:  KBSwitchRoleEndPoint,
			Value: fmt.Sprintf("%s.%s", enginePod.Name, svcName),
		})
		job := renderJob(jobName, switchEnvs)
		cmdJobs = append(cmdJobs, job)

		key := types.NamespacedName{Namespace: s.SwitchResource.Cluster.Namespace, Name: jobName}
		exists, _ := intctrlutil.CheckResourceExists(s.SwitchResource.Ctx, s.SwitchResource.Cli, key, &batchv1.Job{})
		if exists {
			continue
		}

		// if job not exist, create a job
		if err := s.SwitchResource.Cli.Create(s.SwitchResource.Ctx, job); err != nil {
			return nil, err
		}
	}
	return cmdJobs, nil
}

// checkSwitchCmdJobSucceed checks the result of switch command job execution.
func checkSwitchCmdJobSucceed(s *Switch, cmdJobs []*batchv1.Job) error {
	for _, cmdJob := range cmdJobs {
		if err := componentutil.CheckJobSucceed(s.SwitchResource.Ctx, s.SwitchResource.Cli, s.SwitchResource.Cluster, cmdJob); err != nil {
			return err
		}
	}
	return nil
}

// cleanSwitchCmdJobs cleans up the job tasks that execute the switch commands.
func cleanSwitchCmdJobs(s *Switch) error {
	return componentutil.CleanJobWithLabels(s.SwitchResource.Ctx, s.SwitchResource.Cli,
		s.SwitchResource.Cluster, getSwitchCmdJobLabel(s.SwitchResource.Cluster.Name, s.SwitchResource.CompSpec.Name))
}

// getSwitchCmdJobLabel gets the labels for job that executes the switch commands.
func getSwitchCmdJobLabel(clusterName, componentName string) map[string]string {
	return map[string]string{
		constant.AppInstanceLabelKey:    clusterName,
		constant.KBAppComponentLabelKey: componentName,
		constant.AppManagedByLabelKey:   constant.AppName,
		KBSwitchJobLabelKey:             KBSwitchJobLabelValue,
	}
}
