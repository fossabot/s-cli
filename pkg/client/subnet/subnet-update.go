package subnet

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
	s := GetSubnet(false)

	desc := input.GetInput("Subnet Description:", "", s.Description, survey.Required)
	secPolicy := GetSecurityPolicy("Default Security Policy:")

	ss := output.Spinner()

	nxc, grpcConn := grpc.GetTopologyAPIClient()
	defer grpcConn.Close()

	usr := &topology.UpdateSubnetRequest{
		AccountID:     s.AccountID,
		TenantID:      s.TenantID,
		NetID:         s.NetID,
		SubnetID:      s.SubnetID,
		Description:   desc,
		DefaultPolicy: secPolicy,
	}

	s, err := nxc.UpdateSubnet(context.TODO(), usr)
	if err != nil {
		ss.Stop()
		status.Error(err, "Unable to update subnet")
	}

	ss.Stop()

	// output.Show(s)
	Output().Show(s)
}
