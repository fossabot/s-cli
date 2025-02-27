package vs

import (
	"context"
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"skynx.io/s-api-go/grpc/resources/topology"
	"skynx.io/s-cli/pkg/client/node"
	"skynx.io/s-cli/pkg/grpc"
	"skynx.io/s-cli/pkg/input"
	"skynx.io/s-cli/pkg/output"
	"skynx.io/s-cli/pkg/status"
	"skynx.io/s-lib/pkg/sx"
)

func (api *API) AddVSAppSvc() {
	vs := getVS(false)

	avsasr := &topology.AddVSAppSvcRequest{
		AccountID: vs.AccountID,
		VSID:      vs.VSID,
	}

	ok := input.GetConfirm("Add node appSvc to virtual server?", true)
	if !ok {
		Output().Show(vs)
		return
	}

	n := node.GetNodeByTenant(false, sx.Bool(true))

	if vs.TenantID != n.TenantID {
		status.Error(fmt.Errorf("vs/node tenantID mismatch"), "Unable to modifiy virtual server")
	}

	if vs.NetID != n.Cfg.NetID {
		status.Error(fmt.Errorf("vs/node netID mismatch"), "Unable to modifiy virtual server")
	}

	protocols := []string{"HTTP", "HTTPS"}
	var defaultPort string
	var rsProto topology.VSProto

	proto := input.GetSelect("Protocol:", "", protocols, survey.Required)

	switch proto {
	case "HTTP":
		rsProto = topology.VSProto_PROTO_TCP_HTTP
		defaultPort = "8080"
	case "HTTPS":
		rsProto = topology.VSProto_PROTO_TCP_HTTPS
		defaultPort = "443"
	}

	avsasr.NodeAppSvcReq = &topology.NodeAppSvcReq{
		TenantID: n.TenantID,
		NetID:    n.Cfg.NetID,
		SubnetID: n.Cfg.SubnetID,
		NodeID:   n.NodeID,
		NodeName: n.Cfg.NodeName,
		// AppSvcName:        input.GetInput("App Svc Name:", "", "", input.ValidName),
		// AppSvcDescription: input.GetInput("Description:", "", "", nil),
		Proto:  rsProto,
		RSPort: getAppSvcPort(defaultPort),
	}

	s := output.Spinner()

	nxc, grpcConn := grpc.GetTopologyAPIClient()
	defer grpcConn.Close()

	vs, err := nxc.AddVSAppSvc(context.TODO(), avsasr)
	if err != nil {
		s.Stop()
		status.Error(err, "Unable to modifiy virtual server")
	}

	s.Stop()

	// output.Show(n)
	Output().Show(vs)
}
