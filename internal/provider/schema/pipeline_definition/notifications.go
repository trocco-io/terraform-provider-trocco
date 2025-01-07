package pipeline_definition

import (
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
)

func Notifications() schema.Attribute {
	return schema.ListNestedAttribute{
		MarkdownDescription: "The notifications of the pipeline definition",
		Optional:            true,
		NestedObject: schema.NestedAttributeObject{
			Attributes: map[string]schema.Attribute{
				"type": schema.StringAttribute{
					MarkdownDescription: "The type of the notification",
					Optional:            true,
				},
				"email_config": schema.SingleNestedAttribute{
					MarkdownDescription: "The email configuration of the notification",
					Optional:            true,
					Attributes: map[string]schema.Attribute{
						"notification_id": schema.Int64Attribute{
							MarkdownDescription: "The notification id",
							Required:            true,
						},
						"notify_when": schema.StringAttribute{
							MarkdownDescription: "When to notify",
							Required:            true,
						},
						"message": schema.StringAttribute{
							MarkdownDescription: "The message of the notification",
							Required:            true,
						},
					},
				},
				"slack_config": schema.SingleNestedAttribute{
					MarkdownDescription: "The slack configuration of the notification",
					Optional:            true,
					Attributes: map[string]schema.Attribute{
						"notification_id": schema.Int64Attribute{
							MarkdownDescription: "The notification id",
							Required:            true,
						},
						"notify_when": schema.StringAttribute{
							MarkdownDescription: "When to notify",
							Required:            true,
						},
						"message": schema.StringAttribute{
							MarkdownDescription: "The message of the notification",
							Required:            true,
						},
					},
				},
			},
		},
	}
}
