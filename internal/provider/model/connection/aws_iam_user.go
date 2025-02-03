package connection

import "github.com/hashicorp/terraform-plugin-framework/types"

type AWSIAMUser struct {
	AccessKeyID     types.String `tfsdk:"access_key_id"`
	SecretAccessKey types.String `tfsdk:"secret_access_key"`
}
