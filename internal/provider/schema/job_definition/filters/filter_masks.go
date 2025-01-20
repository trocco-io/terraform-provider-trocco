package filters

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

func FilterMasksSchema() schema.Attribute {
	return schema.ListNestedAttribute{
		Optional:            true,
		MarkdownDescription: "Filter masks to be attached to the job definition",
		NestedObject: schema.NestedAttributeObject{
			Attributes: map[string]schema.Attribute{
				"name": schema.StringAttribute{
					Required: true,
					Validators: []validator.String{
						stringvalidator.UTF8LengthAtLeast(1),
					},
					MarkdownDescription: "Target column name",
				},
				"mask_type": schema.StringAttribute{
					Required: true,
					Validators: []validator.String{
						stringvalidator.OneOf("all", "email", "regex", "substring"),
					},
					MarkdownDescription: "Masking type",
				},
				"length": schema.Int64Attribute{
					Optional:            true,
					MarkdownDescription: "Number of mask symbols",
				},
				"pattern": schema.StringAttribute{
					Optional:            true,
					MarkdownDescription: "regular expression pattern",
				},
				"start_index": schema.Int64Attribute{
					Optional:            true,
					MarkdownDescription: "Mask start position",
				},
				"end_index": schema.Int64Attribute{
					Optional:            true,
					MarkdownDescription: "Mask end position",
				},
			},
		},
	}
}
