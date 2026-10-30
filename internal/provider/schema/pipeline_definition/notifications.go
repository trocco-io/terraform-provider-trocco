package pipeline_definition

import (
	"terraform-provider-trocco/internal/provider/custom_type"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

func NotificationsSchema() schema.Attribute {
	return schema.ListNestedAttribute{
		MarkdownDescription: "The notifications of the pipeline definition",
		Optional:            true,
		NestedObject: schema.NestedAttributeObject{
			Attributes: map[string]schema.Attribute{
				"id": schema.Int64Attribute{
					Computed:            true,
					MarkdownDescription: "Server-assigned ID of the notification. Unique within `(type, destination_type)` for matching across API responses.",
				},
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
				"http_config": schema.SingleNestedAttribute{
					MarkdownDescription: "The HTTP configuration of the notification. Used when `destination_type` is `http`",
					Optional:            true,
					Attributes: map[string]schema.Attribute{
						"notification_id": schema.Int64Attribute{
							MarkdownDescription: "The ID of the HTTP notification destination",
							Required:            true,
						},
						"message": schema.StringAttribute{
							MarkdownDescription: "The message of the notification. It is sent as the request body and must be a valid JSON string",
							Required:            true,
							CustomType:          custom_type.TrimmedStringType{},
						},
					},
				},
			},
		},
	}
}
