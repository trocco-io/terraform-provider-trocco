package connection

import "github.com/hashicorp/terraform-plugin-framework/types"

type AWSAssumeRole struct {
	AccountID       types.String `tfsdk:"account_id"`
	AccountRoleName types.String `tfsdk:"role_name"`
}

func (r *AWSAssumeRole) Copy() *AWSAssumeRole {
	if r == nil {
		return nil
	}
	return &AWSAssumeRole{
		AccountID:       types.StringPointerValue(r.AccountID.ValueStringPointer()),
		AccountRoleName: types.StringPointerValue(r.AccountRoleName.ValueStringPointer()),
	}
}
