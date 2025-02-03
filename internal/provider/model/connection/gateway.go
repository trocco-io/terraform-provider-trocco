package connection

import "github.com/hashicorp/terraform-plugin-framework/types"

type Gateway struct {
	Host          types.String `tfsdk:"host"`
	Port          types.Int64  `tfsdk:"port"`
	UserName      types.String `tfsdk:"user_name"`
	Password      types.String `tfsdk:"password"`
	Key           types.String `tfsdk:"key"`
	KeyPassphrase types.String `tfsdk:"key_passphrase"`
}
