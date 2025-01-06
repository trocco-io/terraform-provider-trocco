package pipeline_definition

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

func PeriodCustomVariableLoopConfig() schema.Attribute {
	return schema.SingleNestedAttribute{
		MarkdownDescription: "Period custom variable loop configuration",
		Optional:            true,
		Attributes: map[string]schema.Attribute{
			"interval": schema.StringAttribute{
				MarkdownDescription: "Interval of the loop",
				Required:            true,
			},
			"time_zone": schema.StringAttribute{
				MarkdownDescription: "Timezone of the configuration",
				Required:            true,
			},
			"from": schema.SingleNestedAttribute{
				MarkdownDescription: "Start of the loop",
				Required:            true,
				Attributes: map[string]schema.Attribute{
					"value": schema.Int64Attribute{
						MarkdownDescription: "Value",
						Required:            true,
					},
					"unit": schema.StringAttribute{
						MarkdownDescription: "Unit",
						Required:            true,
					},
				},
			},
			"to": schema.SingleNestedAttribute{
				MarkdownDescription: "End of the loop",
				Required:            true,
				Attributes: map[string]schema.Attribute{
					"value": schema.Int64Attribute{
						MarkdownDescription: "Value",
						Required:            true,
					},
					"unit": schema.StringAttribute{
						MarkdownDescription: "Unit",
						Required:            true,
					},
				},
			},
			"variables": schema.ListNestedAttribute{
				MarkdownDescription: "Custom variables to be expanded",
				Required:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							MarkdownDescription: "Name of custom variable",
							Required:            true,
						},
						"offset": schema.SingleNestedAttribute{
							MarkdownDescription: "Offset on custom variable expanded",
							Required:            true,
							Attributes: map[string]schema.Attribute{
								"value": schema.Int64Attribute{
									MarkdownDescription: "Value",
									Required:            true,
								},
								"unit": schema.StringAttribute{
									MarkdownDescription: "Unit",
									Required:            true,
								},
							},
						},
					},
				},
			},
		},
	}
}
