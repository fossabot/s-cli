package network

import (
	"context"

	"github.com/AlecAivazis/survey/v2"
	"skynx.io/s-api-go/grpc/resources/topology"
	"skynx.io/s-cli/pkg/grpc"
	"skynx.io/s-cli/pkg/input"
	"skynx.io/s-cli/pkg/output"
	"skynx.io/s-cli/pkg/status"
)

func (api *API) Update() {
	n := GetNetwork(false)

	desc := input.GetInput("Description:", "", n.Description, survey.Required)
	rs := input.GetConfirm("Route this network's subnets each other?", n.RoutedSubnets)

	s := output.Spinner()

	nxc, grpcConn := grpc.GetTopologyAPIClient()
	defer grpcConn.Close()

	unr := &topology.UpdateNetworkRequest{
		AccountID:     n.AccountID,
		TenantID:      n.TenantID,
		NetID:         n.NetID,
		Description:   desc,
		RoutedSubnets: rs,
	}

	n, err := nxc.UpdateNetwork(context.TODO(), unr)
	if err != nil {
		s.Stop()
		status.Error(err, "Unable to update network")
	}

	s.Stop()

	// output.Show(n)
	Output().Show(n)
}
