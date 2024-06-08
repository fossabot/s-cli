package project

import (
	"context"

	"github.com/AlecAivazis/survey/v2"
	"skynx.io/s-cli/pkg/grpc"
	"skynx.io/s-cli/pkg/input"
	"skynx.io/s-cli/pkg/output"
	"skynx.io/s-cli/pkg/status"
)

func (api *API) Update() {
	p := GetProject()

	p.Name = input.GetInput("Name:", "", p.Name, input.ValidID)
	p.Description = input.GetInput("Description:", "", p.Description, survey.Required)
	p.ReviewRequired = input.GetConfirm("Enable workflow-required reviews?", p.ReviewRequired)
	p.ApprovalRequired = input.GetConfirm("Enable workflow-required approvals?", p.ApprovalRequired)

	s := output.Spinner()

	nxc, grpcConn := grpc.GetOpsAPIClient()
	defer grpcConn.Close()

	p, err := nxc.SetProject(context.TODO(), p)
	if err != nil {
		s.Stop()
		status.Error(err, "Unable to set project")
	}

	s.Stop()

	Output().Show(p)
}
