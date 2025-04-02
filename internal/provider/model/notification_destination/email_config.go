package notification_destination

import "github.com/hashicorp/terraform-plugin-framework/types"

type EmailConfig struct {
	Email types.String `tfsdk:"email"`
}
