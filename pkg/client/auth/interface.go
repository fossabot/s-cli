package auth

import "skynx.io/s-api-go/grpc/resources/iam/auth"

type Interface interface {
	LoginRequired() bool
	// AutoLogin(req *auth.LoginRequest)
	OTPSignin(req *auth.OTPSigninRequest, verbose bool)
	LoginWithToken(req *auth.LoginRequest, verbose bool)
	Logout()
	// PasswordReset()
	// ConfirmationMailResend()
	// Token()
}
type API struct{}

func Resource() Interface {
	return &API{}
}
