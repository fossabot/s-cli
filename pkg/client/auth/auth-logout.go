package auth

import (
	"context"
	"os"

	auth_pb "skynx.io/s-api-go/grpc/resources/iam/auth"
	"skynx.io/s-cli/pkg/auth"
	"skynx.io/s-cli/pkg/grpc"
	"skynx.io/s-cli/pkg/output"
	"skynx.io/s-cli/pkg/status"
)

func (api *API) Logout() {
	apiKeyFile, err := auth.GetAPIKeyFile()
	if err != nil {
		status.Error(err, "Unable to find API key")
	}

	if _, err := os.Stat(apiKeyFile); os.IsNotExist(err) {
		output.Logout()
		return
	}

	s := output.Spinner()

	nxc, grpcConn := grpc.GetManagerAPIClient(true)
	defer grpcConn.Close()

	accountID, err := auth.GetAccountID()
	if err != nil {
		s.Stop()
		status.Error(err, "Unable to find auth accountID")
	}

	userID, err := auth.GetUserID()
	if err != nil {
		s.Stop()
		status.Error(err, "Unable to find auth userID")
	}

	req := &auth_pb.SignoutRequest{
		AccountID: accountID,
		UserID:    userID,
		// SessionID:
	}

	resp, err := nxc.Signout(context.TODO(), req)
	if err != nil {
		grpcConn.Close()
		s.Stop()
		status.Error(err, "Unable to signout")
	}

	if resp.Result != auth_pb.SignoutResult_SIGNOUT_SUCCESSFUL {
		grpcConn.Close()
		s.Stop()
		status.Error(err, "Unable to signout")
	}

	if err := os.RemoveAll(apiKeyFile); err != nil {
		s.Stop()
		status.Error(err, "Unable to logout")
	}

	s.Stop()

	output.Logout()
}
