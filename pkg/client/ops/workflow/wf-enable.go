package workflow

import (
	"context"

	"skynx.io/s-cli/pkg/grpc"
	"skynx.io/s-cli/pkg/output"
	"skynx.io/s-cli/pkg/status"
)

func (api *API) Enable() {
	wfEnable(true)
}

func (api *API) Disable() {
	wfEnable(false)
}

func wfEnable(enabled bool) {
	wf := GetWorkflow()

	nxc, grpcConn := grpc.GetOpsAPIClient()
	defer grpcConn.Close()

	s := output.Spinner()

	wf.Enabled = enabled

	workflow, err := nxc.SetWorkflow(context.TODO(), wf)
	if err != nil {
		s.Stop()
		status.Error(err, "Unable to set workflow")
	}

	s.Stop()

	Output().Show(workflow)
}
