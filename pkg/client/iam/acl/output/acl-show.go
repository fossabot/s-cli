package output

import (
	"fmt"

	"skynx.io/s-api-go/grpc/resources/iam"
	"skynx.io/s-cli/pkg/output"
	"skynx.io/s-cli/pkg/output/table"
	"skynx.io/s-lib/pkg/utils/colors"
)

func (api *API) Show(acl *iam.ACL) {
	output.SectionHeader("IAM: ACL Details")
	output.TitleT1("ACL Information")

	t := table.New()

	t.AddRow(colors.Black("Account ID"), colors.DarkWhite(acl.AccountID))
	t.AddRow(colors.Black("ACL"), colors.DarkWhite(acl.ACLID))

	t.Render()
	fmt.Println()

	if len(acl.Tags) > 0 {
		output.SubTitleT2("Allowed Tags")

		for _, tag := range acl.Tags {
			fmt.Printf(" ■ %s\n", colors.DarkGreen(tag))
		}
		fmt.Println()
	}

	if len(acl.Users) > 0 {
		output.SubTitleT2("Users")

		for u := range acl.Users {
			fmt.Printf(" ■ %s\n", colors.DarkGreen(u))
		}
		fmt.Println()
	}
}
