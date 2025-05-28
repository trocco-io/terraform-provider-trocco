package pipeline_definition

import (
	"terraform-provider-trocco/internal/provider/custom_type"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

func Notifications() schema.Attribute {
	return schema.ListNestedAttribute{
		MarkdownDescription: "The notifications of the pipeline definition",
		Optional:            true,
		NestedObject: schema.NestedAttributeObject{
			Attributes: map[string]schema.Attribute{
				"type": schema.StringAttribute{
					MarkdownDescription: "The type of the notification",
					Required:            true,
				},
				"destination_type": schema.StringAttribute{
					MarkdownDescription: "The destination type of the notification",
					Required:            true,
				},
				"notify_when": schema.StringAttribute{
					MarkdownDescription: "When to notify",
					Optional:            true,
				},
				"time": schema.Int64Attribute{
					MarkdownDescription: "The time of the notification",
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
						"message": schema.StringAttribute{
							MarkdownDescription: "The message of the notification",
							Required:            true,
							CustomType:          custom_type.TrimmedStringType{},
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
						"message": schema.StringAttribute{
							MarkdownDescription: "The message of the notification",
							Required:            true,
							CustomType:          custom_type.TrimmedStringType{},
						},
					},
				},
			},
		},
	}
}
