package workflow

import (
	"context"
	"fmt"

	"skynx.io/s-api-go/grpc/resources/ops"
	"skynx.io/s-cli/pkg/grpc"
	"skynx.io/s-cli/pkg/output"
	"skynx.io/s-cli/pkg/status"
	"skynx.io/s-lib/pkg/utils"
)

func (api *API) Update(yamlFile string) {
	wf := GetWorkflow()

	s := output.Spinner()

	nxc, grpcConn := grpc.GetOpsAPIClient()
	defer grpcConn.Close()

	nwf := &ops.Workflow{}

	if err := utils.FileParser(yamlFile, nwf); err != nil {
		s.Stop()
		status.Error(err, "Unable to parse YAML file")
	}

	nwf.AccountID = wf.AccountID
	nwf.TenantID = wf.TenantID
	nwf.ProjectID = wf.ProjectID

	if nwf.WorkflowID != wf.WorkflowID {
		s.Stop()
		err := fmt.Errorf("workflowID mismatch")
		status.Error(err, "WorkflowID is not valid")
	}

	wf, err := nxc.SetWorkflow(context.TODO(), nwf)
	if err != nil {
		s.Stop()
		status.Error(err, "Unable to set workflow")
	}

	s.Stop()

	Output().Show(wf)
}
