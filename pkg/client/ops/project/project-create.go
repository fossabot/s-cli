package project

import (
	"context"

	"github.com/AlecAivazis/survey/v2"
	"skynx.io/s-api-go/grpc/resources/ops"
	"skynx.io/s-cli/pkg/client/tenant"
	"skynx.io/s-cli/pkg/grpc"
	"skynx.io/s-cli/pkg/input"
	"skynx.io/s-cli/pkg/output"
	"skynx.io/s-cli/pkg/status"
)

func (api *API) Create() {
	t := tenant.GetTenant()

	npr := &ops.NewProjectRequest{
		AccountID:        t.AccountID,
		TenantID:         t.TenantID,
		Name:             input.GetInput("Name:", "", "", input.ValidID),
		Description:      input.GetInput("Description:", "", "", survey.Required),
		ReviewRequired:   input.GetConfirm("Enable workflow-required reviews?", false),
		ApprovalRequired: input.GetConfirm("Enable workflow-required approvals?", false),
	}

	s := output.Spinner()

	nxc, grpcConn := grpc.GetOpsAPIClient()
	defer grpcConn.Close()

	p, err := nxc.CreateProject(context.TODO(), npr)
	if err != nil {
		s.Stop()
		status.Error(err, "Unable to create project")
	}

	s.Stop()

	Output().Show(p)
}
