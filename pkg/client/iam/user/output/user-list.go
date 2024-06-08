package output

import (
	"fmt"

	"skynx.io/s-api-go/grpc/resources/iam"
	"skynx.io/s-cli/pkg/output"
	"skynx.io/s-cli/pkg/output/table"
	"skynx.io/s-lib/pkg/utils/colors"
)

func (api *API) List(users map[string]*iam.User) {
	output.SectionHeader("IAM: Users")
	output.TitleT1("User List")

	t := table.New()
	t.Header(colors.Black("ACCOUNT ID"), colors.Black("USER EMAIL"))

	for _, u := range users {
		t.AddRow(u.AccountID, colors.DarkWhite(u.Email))
	}

	t.Render()
	fmt.Println()
}
