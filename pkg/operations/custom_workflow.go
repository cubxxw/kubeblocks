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

package operations

import (
	"fmt"

	"sigs.k8s.io/controller-runtime/pkg/client"

	appsv1 "github.com/apecloud/kubeblocks/apis/apps/v1"
	opsv1alpha1 "github.com/apecloud/kubeblocks/apis/operations/v1alpha1"
	intctrlutil "github.com/apecloud/kubeblocks/pkg/controllerutil"
	"github.com/apecloud/kubeblocks/pkg/operations/custom"
)

type WorkflowStatus struct {
	IsCompleted    bool
	ExistFailure   bool
	CompletedCount int
}

type WorkflowContext struct {
	reqCtx intctrlutil.RequestCtx
	Cli    client.Client
	OpsRes *OpsResource
}

func NewWorkflowContext(
	ctx intctrlutil.RequestCtx,
	cli client.Client,
	opsRes *OpsResource) *WorkflowContext {
	return &WorkflowContext{
		reqCtx: ctx,
		Cli:    cli,
		OpsRes: opsRes,
	}
}

// Run actions execution layer.
func (w *WorkflowContext) Run(compCustomSpec *opsv1alpha1.CustomOpsComponent) (*WorkflowStatus, error) {
	var (
		err            error
		actionStatus   *custom.ActionStatus
		compStatus     = w.OpsRes.OpsRequest.Status.Components[compCustomSpec.ComponentName]
		workflowStatus = &WorkflowStatus{}
		actions        = w.OpsRes.OpsDef.Spec.Actions
		compSpec       = getComponentSpecOrShardingTemplate(w.OpsRes.Cluster, compCustomSpec.ComponentName)
	)
	defer func() {
		if intctrlutil.IsTargetError(err, intctrlutil.ErrorTypeFatal) {
			// if the error is Fatal, mark the workflow is Failed.
			compStatus.Message = err.Error()
			workflowStatus.IsCompleted = true
			workflowStatus.ExistFailure = true
		}
		w.OpsRes.OpsRequest.Status.Components[compCustomSpec.ComponentName] = compStatus
	}()
	setSucceedWorkflowStatus := func(actionIndex int) {
		workflowStatus.CompletedCount += 1
		if actionIndex == len(actions)-1 {
			workflowStatus.IsCompleted = true
		}
	}
steps:
	for i := range actions {
		actionProgress := findActionProgress(compStatus.ProgressDetails, actions[i].Name)
		if actionProgress == nil {
			err = intctrlutil.NewFatalError("can not find the action progress for action " + actions[i].Name)
			return nil, err
		}
		switch actionProgress.Status {
		case opsv1alpha1.PendingProgressStatus:
			// execute action and set status progress
			progressDetail := *actionProgress
			ac := w.getAction(actions[i], compCustomSpec, compSpec, progressDetail)
			if ac == nil {
				err = intctrlutil.NewFatalError("the action type is not implement for action " + actions[i].Name)
				return nil, err
			}
			actionStatus, err = ac.Execute(custom.ActionContext{ReqCtx: w.reqCtx, Client: w.Cli, Action: &actions[i]})
			if err != nil {
				return nil, err
			}
			progressDetail.ActionTasks = actionStatus.ActionTasks
			progressDetail.SetStatusAndMessage(opsv1alpha1.ProcessingProgressStatus,
				fmt.Sprintf(`Start to processing action "%s" of the component %s`, actions[i].Name, compCustomSpec.ComponentName))
			setComponentStatusProgressDetail(w.reqCtx.Recorder, w.OpsRes.OpsRequest, &compStatus.ProgressDetails, progressDetail)
			break steps
		case opsv1alpha1.ProcessingProgressStatus:
			// check action status and set status progress
			progressDetail := *actionProgress
			ac := w.getAction(actions[i], compCustomSpec, compSpec, progressDetail)
			if ac == nil {
				err = intctrlutil.NewFatalError("the action type is not implement for action " + actions[i].Name)
				return nil, err
			}
			actionStatus, err = ac.CheckStatus(custom.ActionContext{ReqCtx: w.reqCtx, Client: w.Cli, Action: &actions[i]})
			if err != nil {
				return nil, err
			}
			progressDetail.ActionTasks = actionStatus.ActionTasks
			if actionStatus.IsCompleted {
				if actionStatus.ExistFailure {
					progressDetail.Status = opsv1alpha1.FailedProgressStatus
				} else {
					progressDetail.Status = opsv1alpha1.SucceedProgressStatus
				}
				progressDetail.Message = fmt.Sprintf(`the action "%s" of the component "%s" is %s`,
					actions[i].Name, compCustomSpec.ComponentName, progressDetail.Status)
			}
			setComponentStatusProgressDetail(w.reqCtx.Recorder, w.OpsRes.OpsRequest, &compStatus.ProgressDetails, progressDetail)
			break steps
		case opsv1alpha1.FailedProgressStatus:
			if actions[i].FailurePolicy == opsv1alpha1.FailurePolicyFail {
				workflowStatus.ExistFailure = true
				workflowStatus.CompletedCount += 1
				workflowStatus.IsCompleted = true
				break steps
			} else {
				// if the action is final action and ignore Failure, mark workflow to succeed
				setSucceedWorkflowStatus(i)
			}
		case opsv1alpha1.SucceedProgressStatus:
			// if the action is final action, mark workflow to succeed
			setSucceedWorkflowStatus(i)
		}
	}
	return workflowStatus, nil
}

func (w *WorkflowContext) getAction(action opsv1alpha1.OpsAction,
	compCustomItem *opsv1alpha1.CustomOpsComponent,
	compSpec *appsv1.ClusterComponentSpec,
	progressDetail opsv1alpha1.ProgressStatusDetail) custom.OpsAction {
	switch {
	case action.Workload != nil:
		return custom.NewWorkloadAction(w.OpsRes.OpsRequest, w.OpsRes.Cluster,
			w.OpsRes.OpsDef, compCustomItem, compSpec, progressDetail)
	case action.Exec != nil:
		return custom.NewExecAction(w.OpsRes.OpsRequest, w.OpsRes.Cluster,
			w.OpsRes.OpsDef, compCustomItem, compSpec, progressDetail)
	case action.ResourceModifier != nil:
		// TODO: implement it.
		return nil
	default:
		return nil
	}
}
