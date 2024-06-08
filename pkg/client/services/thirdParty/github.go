package thirdParty

import (
	"fmt"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"skynx.io/s-api-go/grpc/resources/services/thirdParty"
	"skynx.io/s-cli/pkg/input"
	"skynx.io/s-cli/pkg/output"
)

func (a *API) GitHub(gh *thirdParty.GitHub) *thirdParty.GitHub {
	fmt.Println()

	output.Header("Github Setup")

	if gh == nil {
		gh = &thirdParty.GitHub{}
	}

	if gh.Enabled {
		if input.GetConfirm("Delete GitHub integration?", false) {
			gh = &thirdParty.GitHub{}
		}
	}

	if input.GetConfirm("Configure GitHub integration?", true) {
		fmt.Println()

		setGitHubConfig(gh)
	}

	return gh
}

func setGitHubConfig(gh *thirdParty.GitHub) {
	gh.AccessToken = strings.TrimSpace(input.GetInput(
		"Github Access Token:",
		"",
		gh.AccessToken,
		survey.Required,
	))
	gh.WebhooksSecret = strings.TrimSpace(input.GetInput(
		"Github Webhooks Secret:",
		"",
		gh.WebhooksSecret,
		survey.Required,
	))
	gh.Enabled = true
}
