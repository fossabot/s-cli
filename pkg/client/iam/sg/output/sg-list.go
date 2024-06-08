package output

import (
	"fmt"

	"skynx.io/s-api-go/grpc/resources/iam"
	"skynx.io/s-cli/pkg/output"
	"skynx.io/s-cli/pkg/output/table"
	"skynx.io/s-lib/pkg/utils/colors"
)

func (api *API) List(securityGroups map[string]*iam.SecurityGroup) {
	output.SectionHeader("IAM: Security Groups")
	output.TitleT1("Security Group List")

	t := table.New()
	t.Header(colors.Black("ACCOUNT ID"), colors.Black("SECURITY GROUP"))

	for _, sg := range securityGroups {
		t.AddRow(sg.AccountID, colors.DarkWhite(sg.SecurityGroupID))
	}

	t.Render()
	fmt.Println()
}
