package policy

/*
import (
	"context"

	"skynx.io/s-api-go/grpc/resources/topology"
	"skynx.io/s-cli/pkg/client/account"
	"skynx.io/s-cli/pkg/grpc"
	"skynx.io/s-cli/pkg/output"
	"skynx.io/s-cli/pkg/status"
	"skynx.io/s-lib/pkg/utils"
)

func (api *API) Import(yamlFile string) {
	a := account.GetAccount()

	nxc, grpcConn := grpc.GetTopologyAPIClient()
	defer grpcConn.Close()

	s := output.Spinner()

	snpr := &topology.SetNetworkPolicyRequest{}

	if err := utils.FileParser(yamlFile, snpr); err != nil {
		s.Stop()
		status.Error(err, "Unable to parse YAML file")
	}

	snpr.AccountID = a.AccountID

	np, err := nxc.SetNetworkPolicy(context.TODO(), snpr)
	if err != nil {
		s.Stop()
		status.Error(err, "Unable to set network policy")
	}

	s.Stop()

	// output.Show(np)
	Output().Show(snpr.TenantID, snpr.NetID, snpr.SubnetID, np)
}
*/
