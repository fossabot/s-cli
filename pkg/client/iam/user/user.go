package user

import (
	"context"
	"os"
	"sort"

	"github.com/AlecAivazis/survey/v2"
	"skynx.io/s-api-go/grpc/resources/iam"
	"skynx.io/s-api-go/grpc/resources/resource"
	"skynx.io/s-cli/pkg/client/account"
	"skynx.io/s-cli/pkg/grpc"
	"skynx.io/s-cli/pkg/input"
	"skynx.io/s-cli/pkg/output"
	"skynx.io/s-cli/pkg/status"
	"skynx.io/s-cli/pkg/vars"
	"skynx.io/s-lib/pkg/utils/msg"
)

func GetUser(edit bool) *iam.User {
	ul := users()

	if len(ul) == 0 {
		msg.Info("No objects found")
		os.Exit(1)
	}

	var userOptID string
	usersOpts := make([]string, 0)
	users := make(map[string]*iam.User)

	for _, u := range ul {
		userOptID = u.UserID
		usersOpts = append(usersOpts, userOptID)
		users[userOptID] = u
	}

	sort.Strings(usersOpts)

	if edit {
		usersOpts = append(usersOpts, input.NewResource)
	}

	userOptID = input.GetSelect("User:", "", usersOpts, survey.Required)

	if userOptID == input.NewResource {
		return nil
	}

	vars.UserID = users[userOptID].UserID

	return users[userOptID]
}

func users() map[string]*iam.User {
	a := account.GetAccount()

	s := output.Spinner()
	defer s.Stop()

	nxc, grpcConn := grpc.GetIAMAPIClient()
	defer grpcConn.Close()

	lr := &iam.ListUsersRequest{
		Meta:      &resource.ListRequest{},
		AccountID: a.AccountID,
	}

	users := make(map[string]*iam.User)

	for {
		ul, err := nxc.ListUsers(context.TODO(), lr)
		if err != nil {
			status.Error(err, "Unable to list users")
		}

		for _, u := range ul.Users {
			users[u.UserID] = u
		}

		if len(ul.Meta.NextPageToken) > 0 {
			lr.Meta.PageToken = ul.Meta.NextPageToken
		} else {
			break
		}
	}

	return users
}
