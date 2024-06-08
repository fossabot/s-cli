package policy

import (
	"context"

	"skynx.io/s-api-go/grpc/resources/topology"
	"skynx.io/s-cli/pkg/client/subnet"
	"skynx.io/s-cli/pkg/grpc"
	"skynx.io/s-cli/pkg/output"
	"skynx.io/s-cli/pkg/status"
)

func (api *API) SetDefault() {
	s := subnet.GetSubnet(false)

	nxc, grpcConn := grpc.GetTopologyAPIClient()
	defer grpcConn.Close()

	usr := &topology.UpdateSubnetRequest{
		AccountID:     s.AccountID,
		TenantID:      s.TenantID,
		NetID:         s.NetID,
		SubnetID:      s.SubnetID,
		Description:   s.Description,
		DefaultPolicy: subnet.GetSecurityPolicy("Default Security Policy:"),
	}

	ss := output.Spinner()

	s, err := nxc.UpdateSubnet(context.TODO(), usr)
	if err != nil {
		ss.Stop()
		status.Error(err, "Unable to set network policy")
	}

	ss.Stop()

	// output.Show(s)
	Output().Show(s)
}
