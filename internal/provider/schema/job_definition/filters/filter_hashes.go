package filters

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

func FilterHashesSchema() schema.Attribute {
	return schema.ListNestedAttribute{
		Optional:            true,
		MarkdownDescription: "Column hashing",
		NestedObject: schema.NestedAttributeObject{
			Attributes: map[string]schema.Attribute{
				"name": schema.StringAttribute{
					Required: true,
					Validators: []validator.String{
						stringvalidator.UTF8LengthAtLeast(1),
					},
					MarkdownDescription: "Target column name. Replaces the string in the set column with a hashed version using SHA-256.",
				},
			},
		},
	}
}
