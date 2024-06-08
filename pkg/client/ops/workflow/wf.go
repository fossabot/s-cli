package workflow

import (
	"context"
	"os"
	"sort"

	"github.com/AlecAivazis/survey/v2"
	"skynx.io/s-api-go/grpc/resources/ops"
	"skynx.io/s-api-go/grpc/resources/resource"
	"skynx.io/s-cli/pkg/client/ops/project"
	"skynx.io/s-cli/pkg/grpc"
	"skynx.io/s-cli/pkg/input"
	"skynx.io/s-cli/pkg/output"
	"skynx.io/s-cli/pkg/status"
	"skynx.io/s-cli/pkg/vars"
	"skynx.io/s-lib/pkg/utils/msg"
)

func GetWorkflow() *ops.Workflow {
	wfl := workflows()

	if len(wfl) == 0 {
		msg.Info("No workflows found")
		os.Exit(1)
	}

	var workflowOptID string
	workflowsOpts := make([]string, 0)
	workflows := make(map[string]*ops.Workflow)

	for _, wf := range wfl {
		workflowOptID = wf.Name
		workflowsOpts = append(workflowsOpts, workflowOptID)
		workflows[workflowOptID] = wf
	}

	sort.Strings(workflowsOpts)

	workflowOptID = input.GetSelect("Workflow:", "", workflowsOpts, survey.Required)

	vars.WorkflowID = workflows[workflowOptID].WorkflowID

	return workflows[workflowOptID]
}

func workflows() map[string]*ops.Workflow {
	p := project.GetProject()

	s := output.Spinner()
	defer s.Stop()

	nxc, grpcConn := grpc.GetOpsAPIClient()
	defer grpcConn.Close()

	pr := &ops.ProjectReq{
		AccountID: p.AccountID,
		TenantID:  p.TenantID,
		ProjectID: p.ProjectID,
	}

	lr := &ops.ListWorkflowsRequest{
		Meta:    &resource.ListRequest{},
		Project: pr,
	}

	workflows := make(map[string]*ops.Workflow)

	for {
		wfl, err := nxc.ListWorkflows(context.TODO(), lr)
		if err != nil {
			s.Stop()
			status.Error(err, "Unable to list workflows")
		}

		for _, wf := range wfl.Workflows {
			workflows[wf.WorkflowID] = wf
		}

		if len(wfl.Meta.NextPageToken) > 0 {
			lr.Meta.PageToken = wfl.Meta.NextPageToken
		} else {
			break
		}
	}

	return workflows
}
