package output

import (
	"fmt"

	"skynx.io/s-api-go/grpc/resources/tenant"
	"skynx.io/s-cli/pkg/output"
	"skynx.io/s-cli/pkg/output/table"
	"skynx.io/s-lib/pkg/utils/colors"
)

func (api *API) Show(t *tenant.Tenant) {
	output.SectionHeader("Tenant Details")
	output.TitleT1("Tenant Information")

	tbl := table.New()

	// tbl.AddRow(colors.Black("Account ID"), colors.DarkWhite(t.AccountID))
	tbl.AddRow(colors.Black("Tenant ID"), colors.DarkWhite(t.TenantID))
	tbl.AddRow(colors.Black("Name"), colors.DarkWhite(t.Name))
	tbl.AddRow(colors.Black("Description"), colors.DarkWhite(t.Description))

	tbl.Render()
	fmt.Println()
}
