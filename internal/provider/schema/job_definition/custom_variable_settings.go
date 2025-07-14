package job_definition

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	planModifier "terraform-provider-trocco/internal/provider/planmodifier"
	validatorHelpers "terraform-provider-trocco/internal/provider/validator"
)

func CustomVariableSettingsSchema() schema.Attribute {
	return schema.ListNestedAttribute{
		Optional: true,
		NestedObject: schema.NestedAttributeObject{
			Attributes: map[string]schema.Attribute{
				"name": schema.StringAttribute{
					Required: true,
					Validators: []validator.String{
						validatorHelpers.WrappingDollarValidator{},
					},
					MarkdownDescription: "Custom variable name. It must start and end with `$`",
				},
				"type": schema.StringAttribute{
					Required: true,
					Validators: []validator.String{
						stringvalidator.OneOf("string", "timestamp", "timestamp_runtime"),
					},
					MarkdownDescription: "Custom variable type. The following types are supported: `string`, `timestamp`, `timestamp_runtime`",
				},
				"value": schema.StringAttribute{
					Optional:            true,
					MarkdownDescription: "Fixed string which will replace variables at runtime. Required in `string` type",
				},
				"quantity": schema.Int64Attribute{
					Optional: true,
					Validators: []validator.Int64{
						int64validator.AtLeast(0),
					},
					MarkdownDescription: "Quantity used to calculate diff from context_time. Required in `timestamp` and `timestamp_runtime` types",
				},
				"unit": schema.StringAttribute{
					Optional: true,
					Validators: []validator.String{
						stringvalidator.OneOf("hour", "date", "month"),
					},
					MarkdownDescription: "Time unit used to calculate diff from context_time. The following units are supported: `hour`, `date`, `month`. Required in `timestamp` and `timestamp_runtime` types",
				},
				"direction": schema.StringAttribute{
					Optional: true,
					Validators: []validator.String{
						stringvalidator.OneOf("ago", "later"),
					},
					MarkdownDescription: "Direction of the diff from context_time. The following directions are supported: `ago`, `later`. Required in `timestamp` and `timestamp_runtime` types",
				},
				"format": schema.StringAttribute{
					Optional:            true,
					MarkdownDescription: "Format used to replace variables. Required in `timestamp` and `timestamp_runtime` types",
				},
				"time_zone": schema.StringAttribute{
					Optional:            true,
					MarkdownDescription: "Time zone used to format the timestamp. Required in `timestamp` and `timestamp_runtime` types",
				},
			},
			PlanModifiers: []planmodifier.Object{
				&planModifier.CustomVariableSettingPlanModifier{},
			},
		},
	}
}
