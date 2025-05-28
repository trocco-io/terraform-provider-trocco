package job_definition

import (
	"terraform-provider-trocco/internal/provider/custom_type"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

func NotificationsSchema() schema.Attribute {
	return schema.SetNestedAttribute{
		Optional: true,
		NestedObject: schema.NestedAttributeObject{
			Attributes: map[string]schema.Attribute{
				"destination_type": schema.StringAttribute{
					Required: true,
					Validators: []validator.String{
						stringvalidator.OneOf("slack", "email"),
					},
					MarkdownDescription: "Destination service where the notification will be sent. The following types are supported: `slack`, `email`",
				},
				"slack_channel_id": schema.Int64Attribute{
					Optional: true,
					Validators: []validator.Int64{
						int64validator.AtLeast(1),
					},
					MarkdownDescription: "ID of the slack channel used to send notifications. Required when `destination_type` is `slack`",
				},
				"email_id": schema.Int64Attribute{
					Optional: true,
					Validators: []validator.Int64{
						int64validator.AtLeast(1),
					},
					MarkdownDescription: "ID of the email used to send notifications. Required when `destination_type` is `email`",
				},
				"notification_type": schema.StringAttribute{
					Required: true,
					Validators: []validator.String{
						stringvalidator.OneOf("job", "record", "exec_time"),
					},
					MarkdownDescription: "Category of condition. The following types are supported: `job`, `record`, `exec_time`",
				},
				"notify_when": schema.StringAttribute{
					Optional: true,
					Validators: []validator.String{
						stringvalidator.OneOf("finished", "failed"),
					},
					MarkdownDescription: "Specifies the job status that trigger a notification. The following types are supported: `finished`, `failed`. Required when `notification_type` is `job`",
				},
				"record_count": schema.Int64Attribute{
					Optional:            true,
					MarkdownDescription: "The number of records to be used for condition. Required when `notification_type` is `record`",
				},
				"record_operator": schema.StringAttribute{
					Optional: true,
					Validators: []validator.String{
						stringvalidator.OneOf("above", "below"),
					},
					MarkdownDescription: "Operator to be used for condition. The following operators are supported: `above`, `below`. Required when `notification_type` is `record`",
				},
				"record_type": schema.StringAttribute{
					Optional: true,
					Validators: []validator.String{
						stringvalidator.OneOf("transfer", "skipped"),
					},
					MarkdownDescription: "Condition for number of records to be notified",
				},
				"message": schema.StringAttribute{
					Required:            true,
					MarkdownDescription: "The message to be sent with the notification",
					CustomType:          custom_type.TrimmedStringType{},
				},
				"minutes": schema.Int64Attribute{
					Optional: true,
				},
			},
		},
		MarkdownDescription: "Notifications to be attached to the job definition",
	}
}
