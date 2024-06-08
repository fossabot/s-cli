package workflow

import (
	"context"

	"skynx.io/s-api-go/grpc/resources/ops"
	"skynx.io/s-cli/pkg/grpc"
	"skynx.io/s-cli/pkg/output"
	"skynx.io/s-cli/pkg/status"
)

func (api *API) Delete() {
	wf := GetWorkflow()

	output.ConfirmDeletion()

	s := output.Spinner()

	nxc, grpcConn := grpc.GetOpsAPIClient()
	defer grpcConn.Close()

	wfr := &ops.WorkflowReq{
		AccountID:  wf.AccountID,
		TenantID:   wf.TenantID,
		ProjectID:  wf.ProjectID,
		WorkflowID: wf.WorkflowID,
	}

	sr, err := nxc.DeleteWorkflow(context.TODO(), wfr)
	if err != nil {
		s.Stop()
		status.Error(err, "Unable to delete workflow")
	}

	s.Stop()

	output.Show(sr)
}
