package filters

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

func FilterStringTransformsSchema() schema.Attribute {
	return schema.ListNestedAttribute{
		Optional:            true,
		MarkdownDescription: "Character string conversion",
		NestedObject: schema.NestedAttributeObject{
			Attributes: map[string]schema.Attribute{
				"column_name": schema.StringAttribute{
					Required: true,
					Validators: []validator.String{
						stringvalidator.UTF8LengthAtLeast(1),
					},
					MarkdownDescription: "Column name",
				},
				"type": schema.StringAttribute{
					Required: true,
					Validators: []validator.String{
						stringvalidator.OneOf("normalize_nfkc"),
					},
					MarkdownDescription: "Transformation type",
				},
			},
		},
	}
}
