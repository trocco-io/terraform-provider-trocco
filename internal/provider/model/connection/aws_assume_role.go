package connection

import (
	"terraform-provider-trocco/internal/client"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AWSAssumeRole struct {
	AccountID       types.String `tfsdk:"account_id"`
	AccountRoleName types.String `tfsdk:"role_name"`
}

func NewAWSAssumeRole(c *client.Connection) *AWSAssumeRole {
	if c.AWSAuthType == nil || *c.AWSAuthType != "assume_role" {
		return nil
	}
	return &AWSAssumeRole{
		AccountID:       types.StringPointerValue(c.AWSAssumeRoleAccountID),
		AccountRoleName: types.StringPointerValue(c.AWSAssumeRoleName),
	}
}
