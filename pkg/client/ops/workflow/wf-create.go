package workflow

import (
	"context"

	"skynx.io/s-api-go/grpc/resources/ops"
	"skynx.io/s-cli/pkg/client/ops/project"
	"skynx.io/s-cli/pkg/grpc"
	"skynx.io/s-cli/pkg/output"
	"skynx.io/s-cli/pkg/status"
	"skynx.io/s-lib/pkg/utils"
)

func (api *API) Create(yamlFile string) {
	p := project.GetProject()

	s := output.Spinner()

	nxc, grpcConn := grpc.GetOpsAPIClient()
	defer grpcConn.Close()

	wf := &ops.Workflow{}

	if err := utils.FileParser(yamlFile, wf); err != nil {
		s.Stop()
		status.Error(err, "Unable to parse YAML file")
	}

	wf.AccountID = p.AccountID
	wf.TenantID = p.TenantID
	wf.ProjectID = p.ProjectID

	wf, err := nxc.CreateWorkflow(context.TODO(), wf)
	if err != nil {
		s.Stop()
		status.Error(err, "Unable to create workflow")
	}

	s.Stop()

	Output().Show(wf)
}
