package tasklog

import (
	"context"
	"fmt"
	"sort"
	"time"

	"github.com/AlecAivazis/survey/v2"
	"skynx.io/s-api-go/grpc/resources/ops"
	"skynx.io/s-api-go/grpc/resources/resource"
	"skynx.io/s-cli/pkg/client/ops/workflow"
	"skynx.io/s-cli/pkg/grpc"
	"skynx.io/s-cli/pkg/input"
	"skynx.io/s-cli/pkg/output"
	"skynx.io/s-cli/pkg/status"
)

func GetTaskLog() *ops.TaskLog {
	var taskLogID string

	var taskLogOpts []string
	taskLogList := make(map[string]*ops.TaskLog)

	for _, tl := range taskLogs() {
		tm := time.UnixMilli(tl.Status.Timestamp).String()
		taskLogID = fmt.Sprintf("Task: %s | Timestamp: %s\n  Target: %s\n", tl.TaskName, tm, tl.NodeName)
		taskLogOpts = append(taskLogOpts, taskLogID)
		taskLogList[taskLogID] = tl
	}

	sort.Strings(taskLogOpts)

	taskLogID = input.GetSelect("TaskLog:", "", taskLogOpts, survey.Required)

	tl := taskLogList[taskLogID]

	s := output.Spinner()

	nxc, grpcConn := grpc.GetOpsAPIClient()
	defer grpcConn.Close()

	tlr := &ops.TaskLogReq{
		AccountID:  tl.AccountID,
		TenantID:   tl.TenantID,
		ProjectID:  tl.ProjectID,
		WorkflowID: tl.WorkflowID,
		TaskLogID:  tl.TaskLogID,
	}

	tl, err := nxc.GetTaskLog(context.TODO(), tlr)
	if err != nil {
		s.Stop()
		status.Error(err, "Unable to get tasklog")
	}

	s.Stop()

	return tl
}

func taskLogs() []*ops.TaskLog {
	wf := workflow.GetWorkflow()

	s := output.Spinner()
	defer s.Stop()

	nxc, grpcConn := grpc.GetOpsAPIClient()
	defer grpcConn.Close()

	lr := &ops.ListTaskLogsRequest{
		Meta: &resource.ListRequest{},
		Workflow: &ops.WorkflowReq{
			AccountID:  wf.AccountID,
			TenantID:   wf.TenantID,
			ProjectID:  wf.ProjectID,
			WorkflowID: wf.WorkflowID,
		},
	}

	var taskLogs []*ops.TaskLog

	for {
		tll, err := nxc.ListTaskLogs(context.TODO(), lr)
		if err != nil {
			s.Stop()
			status.Error(err, "Unable to list operations")
		}

		taskLogs = append(taskLogs, tll.TaskLogs...)

		if len(tll.Meta.NextPageToken) > 0 {
			lr.Meta.PageToken = tll.Meta.NextPageToken
		} else {
			break
		}
	}

	return taskLogs
}
