package pipeline_definition

import (
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
)

func CustomVariables() schema.Attribute {
	return schema.SetNestedAttribute{
		MarkdownDescription: "The custom variables of the pipeline definition",
		Optional:            true,
		NestedObject: schema.NestedAttributeObject{
			Attributes: map[string]schema.Attribute{
				"name": schema.StringAttribute{
					MarkdownDescription: "The name of the custom variable",
					Required:            true,
				},
				"type": schema.StringAttribute{
					MarkdownDescription: "The type of the custom variable",
					Required:            true,
				},
				"value": schema.StringAttribute{
					MarkdownDescription: "The value of the custom variable",
					Optional:            true,
				},
				"quantity": schema.Int64Attribute{
					MarkdownDescription: "The quantity of the custom variable",
					Optional:            true,
				},
				"unit": schema.StringAttribute{
					MarkdownDescription: "The unit of the custom variable",
					Optional:            true,
				},
				"direction": schema.StringAttribute{
					MarkdownDescription: "The direction of the custom variable",
					Optional:            true,
				},
				"format": schema.StringAttribute{
					MarkdownDescription: "The format of the custom variable",
					Optional:            true,
				},
				"time_zone": schema.StringAttribute{
					MarkdownDescription: "The time zone of the custom variable",
					Optional:            true,
				},
			},
		},
	}
}
