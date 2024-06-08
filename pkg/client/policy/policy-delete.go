package policy

import (
	"context"

	"skynx.io/s-api-go/grpc/resources/topology"
	"skynx.io/s-cli/pkg/client/subnet"
	"skynx.io/s-cli/pkg/grpc"
	"skynx.io/s-cli/pkg/output"
	"skynx.io/s-cli/pkg/status"
)

func (api *API) Delete() {
	s := subnet.GetSubnet(false)

	nxc, grpcConn := grpc.GetTopologyAPIClient()
	defer grpcConn.Close()

	output.ConfirmDeletion()

	ss := output.Spinner()

	sr := &topology.SubnetReq{
		AccountID: s.AccountID,
		TenantID:  s.TenantID,
		NetID:     s.NetID,
		SubnetID:  s.SubnetID,
	}

	np, err := nxc.DeleteNetworkPolicy(context.TODO(), sr)
	if err != nil {
		ss.Stop()
		status.Error(err, "Unable to delete network policy")
	}

	s.NetworkPolicy = np

	ss.Stop()

	// output.Show(s)
	Output().Show(s)
}
