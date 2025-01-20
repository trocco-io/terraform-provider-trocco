package filters

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

func FilterRowsSchema() schema.Attribute {
	return schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "Filter settings",
		Attributes: map[string]schema.Attribute{
			"condition": schema.StringAttribute{
				Required: true,
				Validators: []validator.String{
					stringvalidator.OneOf("and", "or"),
				},
				MarkdownDescription: "Conditions for applying multiple filtering",
			},
			"filter_row_conditions": schema.ListNestedAttribute{
				Required: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"column": schema.StringAttribute{
							Required: true,
							Validators: []validator.String{
								stringvalidator.UTF8LengthAtLeast(1),
							},
							MarkdownDescription: "Target column name",
						},
						"operator": schema.StringAttribute{
							Required: true,
							Validators: []validator.String{
								stringvalidator.OneOf("greater", "greater_equal", "less", "less_equal", "equal", "not_equal", "start_with", "end_with", "include", "is_null", "is_not_null", "regexp"),
							},
							MarkdownDescription: "Operator",
						},
						"argument": schema.StringAttribute{
							Required: true,
							Validators: []validator.String{
								stringvalidator.UTF8LengthAtLeast(1),
							},
							MarkdownDescription: "Argument",
						},
					},
				},
			},
		},
	}
}
