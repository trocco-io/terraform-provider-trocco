package filters

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

func FilterColumnsSchema() schema.Attribute {
	return schema.ListNestedAttribute{
		Required: true,
		NestedObject: schema.NestedAttributeObject{
			Attributes: map[string]schema.Attribute{
				"name": schema.StringAttribute{
					Required: true,
					Validators: []validator.String{
						stringvalidator.UTF8LengthAtLeast(1),
					},
					MarkdownDescription: "Column name",
				},
				"src": schema.StringAttribute{
					Required:            true,
					MarkdownDescription: "Column name in source",
				},
				"type": schema.StringAttribute{
					Required: true,
					Validators: []validator.String{
						stringvalidator.OneOf("string", "long", "timestamp", "double", "boolean", "json"),
					},
					MarkdownDescription: "column type",
				},
				"default": schema.StringAttribute{
					Optional:            true,
					MarkdownDescription: "Default value. For existing columns, this value will be inserted only if input is null. For new columns, this value is inserted for all.",
				},
				"format": schema.StringAttribute{
					Optional:            true,
					MarkdownDescription: "date/time format",
				},
				"json_expand_enabled": schema.BoolAttribute{
					Required:            true,
					MarkdownDescription: "Flag whether to expand JSON",
				},
				"json_expand_keep_base_column": schema.BoolAttribute{
					Required:            true,
					MarkdownDescription: "Flag whether to keep the base column",
				},
				"json_expand_columns": schema.ListNestedAttribute{
					Required: true,
					NestedObject: schema.NestedAttributeObject{
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Required: true,
								Validators: []validator.String{
									stringvalidator.UTF8LengthAtLeast(1),
								},
								MarkdownDescription: "Column name",
							},
							"json_path": schema.StringAttribute{
								Required: true,
								Validators: []validator.String{
									stringvalidator.UTF8LengthAtLeast(1),
								},
								MarkdownDescription: "JSON path. To extract id and age from a JSON column such as {'{“id”: 10, “person”: {“age”: 30}}'}, specify id and person.age in the JSON path, respectively.",
							},
							"type": schema.StringAttribute{
								Required: true,
								Validators: []validator.String{
									stringvalidator.OneOf("boolean", "long", "timestamp", "string"),
								},
								MarkdownDescription: "Column type",
							},
							"format": schema.StringAttribute{
								Optional:            true,
								MarkdownDescription: "date/time format",
							},
							"timezone": schema.StringAttribute{
								Optional:            true,
								MarkdownDescription: "time zone",
							},
						},
					},
				},
			},
		},
	}
}
