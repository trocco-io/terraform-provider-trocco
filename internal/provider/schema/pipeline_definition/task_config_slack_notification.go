package pipeline_definition

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

func SlackNotificationTaskConfigSchema() schema.Attribute {
	return schema.SingleNestedAttribute{
		MarkdownDescription: "The task configuration for the slack notification task.",
		Optional:            true,
		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				MarkdownDescription: "The name of the task",
				Required:            true,
			},
			"connection_id": schema.Int64Attribute{
				MarkdownDescription: "The connection id to use for the task",
				Required:            true,
			},
			"message": schema.StringAttribute{
				MarkdownDescription: "The message to send",
				Required:            true,
			},
			"ignore_error": schema.BoolAttribute{
				MarkdownDescription: "Whether to ignore errors",
				Required:            true,
			},
		},
	}
}
