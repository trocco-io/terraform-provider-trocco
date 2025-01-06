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
					Optional: true,
				},
				"email_config": schema.SingleNestedAttribute{
					Optional: true,
					Attributes: map[string]schema.Attribute{
						"notification_id": schema.Int64Attribute{
							Required: true,
						},
						"notify_when": schema.StringAttribute{
							Required: true,
						},
						"message": schema.StringAttribute{
							Required: true,
						},
					},
				},
				"slack_config": schema.SingleNestedAttribute{
					Optional: true,
					Attributes: map[string]schema.Attribute{
						"notification_id": schema.Int64Attribute{
							Required: true,
						},
						"notify_when": schema.StringAttribute{
							Required: true,
						},
						"message": schema.StringAttribute{
							Required: true,
						},
					},
				},
			},
		},
	}
}
