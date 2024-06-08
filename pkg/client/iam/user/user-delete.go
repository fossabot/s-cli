package user

import (
	"context"

	"skynx.io/s-api-go/grpc/resources/iam"
	"skynx.io/s-cli/pkg/grpc"
	"skynx.io/s-cli/pkg/output"
	"skynx.io/s-cli/pkg/status"
)

func (api *API) Delete() {
	u := GetUser(false)

	nxc, grpcConn := grpc.GetIAMAPIClient()
	defer grpcConn.Close()

	output.ConfirmDeletion()

	ur := &iam.UserReq{
		AccountID: u.AccountID,
		UserID:    u.UserID,
	}

	s := output.Spinner()

	sr, err := nxc.DeleteUser(context.TODO(), ur)
	if err != nil {
		s.Stop()
		status.Error(err, "Unable to delete user")
	}

	s.Stop()

	output.Show(sr)
}
