package output

import (
	"fmt"

	"skynx.io/s-api-go/grpc/resources/iam"
	"skynx.io/s-cli/pkg/output"
	"skynx.io/s-cli/pkg/output/table"
	"skynx.io/s-lib/pkg/utils/colors"
)

func (api *API) List(roles map[string]*iam.Role) {
	output.SectionHeader("IAM: Roles")
	output.TitleT1("Role List")

	t := table.New()
	t.Header(colors.Black("ACCOUNT ID"), colors.Black("ROLE"))

	for _, r := range roles {
		t.AddRow(r.AccountID, colors.DarkWhite(r.RoleID))
	}

	t.Render()
	fmt.Println()
}
