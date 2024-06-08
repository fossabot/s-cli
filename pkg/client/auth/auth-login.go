package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/skratchdot/open-golang/open"
	"github.com/spf13/viper"
	auth_pb "skynx.io/s-api-go/grpc/resources/iam/auth"
	"skynx.io/s-cli/pkg/auth"
	"skynx.io/s-cli/pkg/grpc"
	"skynx.io/s-cli/pkg/input"
	"skynx.io/s-cli/pkg/output"
	"skynx.io/s-cli/pkg/status"
	"skynx.io/s-lib/pkg/errors"
	"skynx.io/s-lib/pkg/utils/colors"
	"skynx.io/s-lib/pkg/utils/msg"
)

func (api *API) OTPSignin(req *auth_pb.OTPSigninRequest, verbose bool) {
	nxc, grpcConn := grpc.GetManagerAPIClient(false)
	defer grpcConn.Close()

	otpResp, err := nxc.Signin(context.TODO(), req)
	if err != nil {
		grpcConn.Close()
		status.Error(err, "Unable to signin")
	}

	fmt.Printf("\n  %s\n\n",
		colors.Black(`A one-time code to log in to your account has been sent to your email.
  Your code expires in 10 minutes.
  Please, check your inbox (and spam/junk folders, just in case).`))

	authReq := &auth_pb.OTPAuthenticationRequest{
		MethodID: otpResp.MethodID,
		Code:     input.GetInput("Code:", "", "", survey.Required),
	}

	s := output.Spinner()

	resp, err := nxc.OTPAuthenticate(context.TODO(), authReq)
	if err != nil {
		grpcConn.Close()
		s.Stop()
		status.Error(err, "Unable to authenticate")
	}

	s.Stop()

	if otpResp.UserCreated {
		fmt.Printf("\n  %s\n",
			colors.Black(`Welcome to skynx.com!

  Your user does not have an account registered, but you can now create one for free
  at https://skynx.com/signin to start enjoying the platform in less than 2 minutes.

Thanks for using skynx.com ;)`))

		// go to signin URL
		signinURL := "https://cloud.skynx.com/signin"

		if err := open.Start(signinURL); err != nil {
			status.Error(err, "Unable to open URL in your browser")
		}

		fmt.Printf("\n%s %s\n\n", colors.DarkWhite("->"), colors.Black("Opening signin URL in your browser..."))

		return
	}

	authResponse(resp, verbose)
}

func (api *API) LoginWithToken(req *auth_pb.LoginRequest, verbose bool) {
	nxc, grpcConn := grpc.GetManagerAPIClient(false)
	defer grpcConn.Close()

	resp, err := nxc.Login(context.TODO(), req)
	if err != nil {
		grpcConn.Close()
		status.Error(err, "Unable to login")
	}

	authResponse(resp, verbose)
}

func authResponse(resp *auth_pb.AuthenticationResponse, verbose bool) {
	switch resp.Result {
	case auth_pb.AuthenticationResult_AUTHENTICATION_FAILED:
		msg.Error("Login failed")
	case auth_pb.AuthenticationResult_ACCOUNT_DISABLED:
		msg.Error("Account disabled, please contact customer service")
	case auth_pb.AuthenticationResult_USER_DISABLED:
		msg.Error("User disabled, please contact your skynx account administrator")
	}

	if resp.Result != auth_pb.AuthenticationResult_AUTHENTICATION_SUCCESSFUL {
		os.Exit(1)
	}

	ac := &auth.Credentials{
		AccountID:    resp.AccountID,
		LocationID:   resp.LocationID,
		FederationID: resp.FederationID,
		Key:          resp.AuthKey.Key,
		ExpiresAt:    resp.AuthExpiresAt,
		UserID:       resp.UserID,
	}

	if err := setAPIKey(ac); err != nil {
		status.Error(err, "Unable to set apiKey")
	}

	viper.Set("user.accountID", resp.AccountID)
	// viper.Set("user.locationID", resp.LocationID)
	viper.Set("user.federationID", resp.FederationID)
	// viper.Set("user.userID", resp.UserID)
	// viper.Set("user.authExpiresAt", resp.AuthExpiresAt)
	viper.Set("user.isAdmin", resp.IsAdmin)

	if verbose {
		fmt.Println()
		output.Authenticated()
	}
}

func setAPIKey(ac *auth.Credentials) error {
	jsonData, err := json.Marshal(ac)
	if err != nil {
		return errors.Wrapf(err, "[%v] function json.Marshal()", errors.Trace())
	}

	apiKeyFile, err := auth.GetAPIKeyFile()
	if err != nil {
		return errors.Wrapf(err, "[%v] function getAPIKeyFile()", errors.Trace())
	}

	if err := os.WriteFile(apiKeyFile, jsonData, 0600); err != nil {
		return errors.Wrapf(err, "[%v] function os.WriteFile()", errors.Trace())
	}

	return nil
}
