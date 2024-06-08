package node

import (
	"context"

	"skynx.io/s-api-go/grpc/resources/topology"
	"skynx.io/s-cli/pkg/grpc"
	"skynx.io/s-cli/pkg/output"
	"skynx.io/s-cli/pkg/status"
)

func (api *API) Delete() {
	n := GetNodeByTenant(false, nil)

	output.ConfirmDeletion()

	s := output.Spinner()

	nxc, grpcConn := grpc.GetTopologyAPIClient()
	defer grpcConn.Close()

	nr := &topology.NodeReq{
		AccountID: n.AccountID,
		TenantID:  n.TenantID,
		NodeID:    n.NodeID,
	}

	sr, err := nxc.DeleteNode(context.TODO(), nr)
	if err != nil {
		s.Stop()
		status.Error(err, "Unable to delete node")
	}

	s.Stop()

	output.Show(sr)
}
