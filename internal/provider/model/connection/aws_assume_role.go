package connection

import "github.com/hashicorp/terraform-plugin-framework/types"

type AWSAssumeRole struct {
	AccountID       types.String `tfsdk:"account_id"`
	AccountRoleName types.String `tfsdk:"role_name"`
}
