package cmd

import (
	"fmt"
	"os"

	"skynx.io/s-cli/internal/app/cli/auth/login"
	"skynx.io/s-cli/pkg/client"
	"skynx.io/s-cli/pkg/output"
	"skynx.io/s-lib/pkg/utils/colors"
	"skynx.io/s-lib/pkg/version"
)

func header() {
	fmt.Println(colors.Black(version.CLI_NAME + " " + version.GetVersion()))
	output.AppHeader(false)
}

func appHeader(str string) string {
	h1 := colors.Black(version.CLI_NAME + " " + version.GetVersion())
	h2 := output.AppHeader(true)

	return fmt.Sprintf("%s\n%s%s", h1, h2, str)
}

func preflightNoLogin() {
	header()

	// if !isConfigured {
	// 	notConfigured()
	// 	os.Exit(0)
	// }

	// if _, err := auth.GetAccountID(); err != nil {
	// 	status.Error(fmt.Errorf("missing accountID"), "Unable to get account")
	// }
}

func preflight() {
	header()

	if isConfigured {
		// silent auto-login
		autoLogin()
	} else {
		// notConfigured()
		os.Exit(0)
	}
}

func autoLogin() {
	if client.Auth().LoginRequired() {
		client.Auth().OTPSignin(login.NewRequestWithOTP(), true)
		// client.Auth().LoginWithToken(login.NewRequestWithToken(), true)
	}
}

/*
func notConfigured() {
	msg.Error("Configuration not detected")

	fmt.Printf("%s\n", colors.DarkBlue("_"))
	cmd := colors.DarkWhite(fmt.Sprintf("%s setup", version.CLI_NAME))
	q := colors.DarkBlue("'")
	msg := fmt.Sprintf("%s %s%s%s", colors.Black("Please configure the client with"), q, cmd, q)
	fmt.Printf("%s %s\n\n", colors.Cyan("🢂"), msg)
}
*/
