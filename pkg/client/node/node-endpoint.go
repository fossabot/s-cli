package node

import (
	"context"

	"skynx.io/s-api-go/grpc/resources/topology"
	"skynx.io/s-cli/pkg/grpc"
	"skynx.io/s-cli/pkg/output"
	"skynx.io/s-cli/pkg/status"
	"skynx.io/s-lib/pkg/sx"
)

func (api *API) ShowEndpoint() {
	n := GetNodeByTenant(false, sx.Bool(true))
	eID := GetEndpoint(n).EndpointID

	for endpointID, e := range n.Endpoints {
		if endpointID == eID {
			// output.Show(e)
			Output().ShowEndpoint(e)
		}
	}
}

func (api *API) DeleteEndpoint() {
	n := GetNodeByTenant(false, sx.Bool(true))
	eID := GetEndpoint(n).EndpointID

	s := output.Spinner()

	nxc, grpcConn := grpc.GetTopologyAPIClient()
	defer grpcConn.Close()

	er := &topology.EndpointRequest{
		NodeReq: &topology.NodeReq{
			AccountID: n.AccountID,
			TenantID:  n.TenantID,
			NodeID:    n.NodeID,
		},
		EndpointID: eID,
	}

	sr, err := nxc.DeleteNetworkEndpoint(context.TODO(), er)
	if err != nil {
		s.Stop()
		status.Error(err, "Unable to delete node network endpoint")
	}

	s.Stop()

	output.Show(sr)
}
