package tenant

import (
	"context"

	"github.com/AlecAivazis/survey/v2"
	"skynx.io/s-api-go/grpc/resources/tenant"
	"skynx.io/s-cli/pkg/client/account"
	"skynx.io/s-cli/pkg/grpc"
	"skynx.io/s-cli/pkg/input"
	"skynx.io/s-cli/pkg/output"
	"skynx.io/s-cli/pkg/status"
)

func (api *API) New() {
	a := account.GetAccount()

	ntr := &tenant.NewTenantRequest{
		AccountID:   a.AccountID,
		Name:        input.GetInput("Tenant Name:", "", "", validTenantName),
		Description: input.GetInput("Description:", "", "", survey.Required),
	}

	s := output.Spinner()

	nxc, grpcConn := grpc.GetTenantAPIClient()
	defer grpcConn.Close()

	t, err := nxc.CreateTenant(context.TODO(), ntr)
	if err != nil {
		s.Stop()
		status.Error(err, "Unable to set tenant")
	}

	s.Stop()

	// output.Show(t)
	Output().Show(t)
}
