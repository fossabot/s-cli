package account

import (
	"context"
	"fmt"
	"os"

	"skynx.io/s-api-go/grpc/resources/account"
	"skynx.io/s-cli/pkg/grpc"
	"skynx.io/s-cli/pkg/input"
	"skynx.io/s-cli/pkg/output"
	"skynx.io/s-cli/pkg/status"
	"skynx.io/s-lib/pkg/utils/msg"
)

func (api *API) Cancel() {
	a := FetchAccount()

	nxc, grpcConn := grpc.GetAccountAPIClient(true)
	defer grpcConn.Close()

	confirmCancelation()

	s := output.Spinner()

	ar := &account.AccountReq{
		AccountID: a.AccountID,
	}

	sr, err := nxc.CancelAccount(context.TODO(), ar)
	if err != nil {
		s.Stop()
		status.Error(err, "Unable to cancel account")
	}

	s.Stop()

	output.Show(sr)
}

func confirmCancelation() {
	msg.Alert("You are about to cancel your skynx account.")
	msg.Alert("All its remaining resources or configurations (rbac, users, etc) will be deleted.")
	msg.Alert("This action is irreversible, please, double check.")

	if !input.GetConfirm("Confirm account cancelation?", false) {
		fmt.Println()
		os.Exit(0)
	}

	if !input.GetConfirm("Last chance. Confirm account cancelation?", false) {
		fmt.Println()
		os.Exit(0)
	}
}
