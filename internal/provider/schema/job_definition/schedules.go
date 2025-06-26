package job_definition

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	planmodifier2 "terraform-provider-trocco/internal/provider/planmodifier"
)

func SchedulesSchema() schema.Attribute {
	return schema.SetNestedAttribute{
		Optional: true,
		NestedObject: schema.NestedAttributeObject{
			Attributes: map[string]schema.Attribute{
				"frequency": schema.StringAttribute{
					Required: true,
					Validators: []validator.String{
						stringvalidator.OneOf("hourly", "daily", "weekly", "monthly"),
					},
					MarkdownDescription: "Frequency of automatic execution. The following frequencies are supported: `hourly`, `daily`, `weekly`, `monthly`",
				},
				"minute": schema.Int64Attribute{
					Required: true,
					Validators: []validator.Int64{
						int64validator.Between(0, 59),
					},
					MarkdownDescription: "Value of minute. Required for all schedules",
				},
				"hour": schema.Int64Attribute{
					Optional: true,
					Validators: []validator.Int64{
						int64validator.Between(0, 23),
					},
					MarkdownDescription: "Value of hour. Required in `daily`, `weekly`, and `monthly` schedules",
				},
				"day_of_week": schema.Int64Attribute{
					Optional: true,
					Validators: []validator.Int64{
						int64validator.Between(0, 6),
					},
					MarkdownDescription: "Value of day of week. Sunday - Saturday is represented as 0 - 6. Required in `weekly` schedule",
				},
				"day": schema.Int64Attribute{
					Optional: true,
					Validators: []validator.Int64{
						int64validator.Between(1, 31),
					},
					MarkdownDescription: "Value of day. Required in `monthly` schedule",
				},
				"time_zone": schema.StringAttribute{
					Required:            true,
					MarkdownDescription: "Time zone to be used for calculation",
				},
			},
			PlanModifiers: []planmodifier.Object{
				&planmodifier2.SchedulePlanModifier{},
			},
		},
		MarkdownDescription: "Schedules to be attached to the job definition",
	}
}
