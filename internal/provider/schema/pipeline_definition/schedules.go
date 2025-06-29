package pipeline_definition

import (
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
)

func SchedulesSchema() schema.Attribute {
	return schema.SetNestedAttribute{
		MarkdownDescription: "The schedules of the pipeline definition",
		Optional:            true,
		NestedObject: schema.NestedAttributeObject{
			Attributes: map[string]schema.Attribute{
				"frequency": schema.StringAttribute{
					MarkdownDescription: "The frequency of the schedule",
					Required:            true,
				},
				"time_zone": schema.StringAttribute{
					MarkdownDescription: "The time zone of the schedule",
					Required:            true,
				},
				"minute": schema.Int64Attribute{
					MarkdownDescription: "The minute of the schedule",
					Required:            true,
				},
				"day": schema.Int64Attribute{
					MarkdownDescription: "The day of the schedule",
					Optional:            true,
				},
				"day_of_week": schema.Int64Attribute{
					MarkdownDescription: "The day of the week of the schedule",
					Optional:            true,
				},
				"hour": schema.Int64Attribute{
					MarkdownDescription: "The hour of the schedule",
					Optional:            true,
				},
			},
		},
	}
}
