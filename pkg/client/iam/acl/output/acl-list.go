package output

import (
	"fmt"

	"skynx.io/s-api-go/grpc/resources/iam"
	"skynx.io/s-cli/pkg/output"
	"skynx.io/s-cli/pkg/output/table"
	"skynx.io/s-lib/pkg/utils/colors"
)

func (api *API) List(acls map[string]*iam.ACL) {
	output.SectionHeader("IAM: ACLs")
	output.TitleT1("ACL List")

	t := table.New()
	t.Header(colors.Black("ACCOUNT ID"), colors.Black("ACL"))

	for _, a := range acls {
		t.AddRow(a.AccountID, colors.DarkWhite(a.ACLID))
	}

	t.Render()
	fmt.Println()
}
