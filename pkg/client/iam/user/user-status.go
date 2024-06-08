package user

import (
	"context"

	"skynx.io/s-api-go/grpc/resources/iam"
	"skynx.io/s-cli/pkg/grpc"
	"skynx.io/s-cli/pkg/output"
	"skynx.io/s-cli/pkg/status"
)

func (api *API) Enable() {
	setStatus(true)
}

func (api *API) Disable() {
	setStatus(false)
}

func setStatus(enabled bool) {
	u := GetUser(false)

	nxc, grpcConn := grpc.GetIAMAPIClient()
	defer grpcConn.Close()

	ur := &iam.UserReq{
		AccountID: u.AccountID,
		UserID:    u.UserID,
	}

	s := output.Spinner()

	var err error
	if enabled {
		u, err = nxc.EnableUser(context.TODO(), ur)
	} else {
		u, err = nxc.DisableUser(context.TODO(), ur)
	}
	if err != nil {
		s.Stop()
		status.Error(err, "Unable to set user")
	}

	s.Stop()

	Output().Show(u)
}
