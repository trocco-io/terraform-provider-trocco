package filters

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

func FilterUnixTimeConversionsSchema() schema.Attribute {
	return schema.ListNestedAttribute{
		Optional:            true,
		MarkdownDescription: "UNIX time conversion",
		NestedObject: schema.NestedAttributeObject{
			Attributes: map[string]schema.Attribute{
				"column_name": schema.StringAttribute{
					Required: true,
					Validators: []validator.String{
						stringvalidator.UTF8LengthAtLeast(1),
					},
					MarkdownDescription: "Target column name",
				},
				"kind": schema.StringAttribute{
					Required: true,
					Validators: []validator.String{
						stringvalidator.OneOf("unixtime_to_timestamp", "unixtime_to_string", "timestamp_to_unixtime", "string_to_unixtime"),
					},
					MarkdownDescription: "Conversion Type",
				},
				"unixtime_unit": schema.StringAttribute{
					Required: true,
					Validators: []validator.String{
						stringvalidator.OneOf("second", "millisecond", "microsecond", "nanosecond"),
					},
					MarkdownDescription: "UNIX time units before conversion",
				},
				"datetime_format": schema.StringAttribute{
					Required:            true,
					MarkdownDescription: "Date and tim format after conversion",
				},
				"datetime_timezone": schema.StringAttribute{
					Required:            true,
					MarkdownDescription: "Time zon after conversion",
				},
			},
		},
	}
}
