package filters

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

func FilterGsubSchema() schema.Attribute {
	return schema.ListNestedAttribute{
		Optional:            true,
		MarkdownDescription: "String Regular Expression Replacement",
		NestedObject: schema.NestedAttributeObject{
			Attributes: map[string]schema.Attribute{
				"column_name": schema.StringAttribute{
					Required: true,
					Validators: []validator.String{
						stringvalidator.UTF8LengthAtLeast(1),
					},
					MarkdownDescription: "Target column name",
				},
				"pattern": schema.StringAttribute{
					Required: true,
					Validators: []validator.String{
						stringvalidator.UTF8LengthAtLeast(1),
					},
					MarkdownDescription: "Regular expression pattern",
				},
				"to": schema.StringAttribute{
					Required: true,
					Validators: []validator.String{
						stringvalidator.UTF8LengthAtLeast(1),
					},
					MarkdownDescription: "String to be replaced",
				},
			},
		},
	}
}
