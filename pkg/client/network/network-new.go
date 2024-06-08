package network

import (
	"context"

	"github.com/AlecAivazis/survey/v2"
	"skynx.io/s-api-go/grpc/resources/topology"
	"skynx.io/s-cli/pkg/client/tenant"
	"skynx.io/s-cli/pkg/grpc"
	"skynx.io/s-cli/pkg/input"
	"skynx.io/s-cli/pkg/output"
	"skynx.io/s-cli/pkg/status"
)

func (api *API) New() {
	t := tenant.GetTenant()

	nnr := &topology.NewNetworkRequest{
		AccountID: t.AccountID,
		TenantID:  t.TenantID,
	}

	helpMsg := "A valid /16 network with format 'n.n.0.0/16' is required"
	nnr.NetworkCIDR = input.GetInput("Network CIDR:", helpMsg, "", input.ValidNetwork)

	nnr.Description = input.GetInput("Description:", "", "", survey.Required)

	nnr.RoutedSubnets = input.GetConfirm("Route this network's subnets each other?", false)

	nnr.LocationID = GetConnectivityZone().LocationID

	s := output.Spinner()

	nxc, grpcConn := grpc.GetTopologyAPIClient()
	defer grpcConn.Close()

	n, err := nxc.CreateNetwork(context.TODO(), nnr)
	if err != nil {
		s.Stop()
		status.Error(err, "Unable to create network")
	}

	s.Stop()

	// output.Show(n)
	Output().Show(n)
}
