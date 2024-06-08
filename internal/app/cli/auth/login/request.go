package login

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
	"skynx.io/s-api-go/grpc/resources/iam/auth"
	"skynx.io/s-cli/pkg/input"
	"skynx.io/s-cli/pkg/output"
	"skynx.io/s-lib/pkg/utils/msg"
	"skynx.io/s-lib/pkg/version"
)

func userAgent() string {
	return fmt.Sprintf("%s/%s", version.CLI_NAME, version.GetVersion())
}

func NewRequestWithOTP() *auth.OTPSigninRequest {
	req := &auth.OTPSigninRequest{
		UserEmail: viper.GetString("user.email"),
		Method:    auth.SigninMethod_SIGNIN_BY_EMAIL,
		UserAgent: userAgent(),
	}

	if len(req.UserEmail) == 0 {
		output.AuthenticationRequired()
		req.UserEmail = input.GetInput("Email:", "", "", input.ValidEmail)
	}

	return req
}

func NewRequestWithToken() *auth.LoginRequest {
	userToken := viper.GetString("token")

	if len(userToken) == 0 {
		msg.Error("Authorization token not found")
		os.Exit(1)
	}

	return &auth.LoginRequest{
		UserToken: userToken,
	}
}
