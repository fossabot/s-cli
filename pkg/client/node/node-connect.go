package node

import (
	"context"

	"skynx.io/s-api-go/grpc/resources/topology"
	"skynx.io/s-cli/pkg/client/subnet"
	"skynx.io/s-cli/pkg/grpc"
	"skynx.io/s-cli/pkg/output"
	"skynx.io/s-cli/pkg/status"
	"skynx.io/s-lib/pkg/sx"
)

func (api *API) Connect() {
	n := GetNodeByTenant(false, nil)

	s := subnet.GetSubnet(false)

	ss := output.Spinner()

	nxc, grpcConn := grpc.GetTopologyAPIClient()
	defer grpcConn.Close()

	unr := &topology.UpdateNodeNetworkingRequest{
		NodeReq: &topology.NodeReq{
			AccountID: n.AccountID,
			TenantID:  n.TenantID,
			NodeID:    n.NodeID,
		},
		NetID:    s.NetID,
		SubnetID: s.SubnetID,
	}

	sr, err := nxc.UpdateNodeNetworking(context.TODO(), unr)
	if err != nil {
		ss.Stop()
		status.Error(err, "Unable to update node network configuration")
	}

	ss.Stop()

	// output.Show(sr)
	Output().Show(sr)
}

func (api *API) Disconnect() {
	n := GetNodeByTenant(false, sx.Bool(true))

	s := output.Spinner()

	nxc, grpcConn := grpc.GetTopologyAPIClient()
	defer grpcConn.Close()

	unr := &topology.UpdateNodeNetworkingRequest{
		NodeReq: &topology.NodeReq{
			AccountID: n.AccountID,
			TenantID:  n.TenantID,
			NodeID:    n.NodeID,
		},
		NetID:    "",
		SubnetID: "",
	}

	sr, err := nxc.UpdateNodeNetworking(context.TODO(), unr)
	if err != nil {
		s.Stop()
		status.Error(err, "Unable to update node network configuration")
	}

	s.Stop()

	// output.Show(sr)
	Output().Show(sr)
}
