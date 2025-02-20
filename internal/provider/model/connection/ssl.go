package connection

import "github.com/hashicorp/terraform-plugin-framework/types"

type SSL struct {
	CA      types.String `tfsdk:"ca"`
	Cert    types.String `tfsdk:"cert"`
	Key     types.String `tfsdk:"key"`
	SSLMode types.String `tfsdk:"ssl_mode"`
}
