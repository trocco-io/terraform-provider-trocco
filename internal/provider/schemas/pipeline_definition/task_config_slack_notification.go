package pipeline_definition

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

func SlackNotificationTaskConfig() schema.Attribute {
	return schema.SingleNestedAttribute{
		Optional: true,
		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				Required: true,
			},
			"connection_id": schema.Int64Attribute{
				Required: true,
			},
			"message": schema.StringAttribute{
				Required: true,
			},
			"ignore_error": schema.BoolAttribute{
				Required: true,
			},
		},
	}
}