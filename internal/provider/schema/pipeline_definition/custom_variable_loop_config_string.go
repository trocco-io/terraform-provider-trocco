package pipeline_definition

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func StringCustomVariableLoopConfigSchema() schema.Attribute {
	return schema.SingleNestedAttribute{
		MarkdownDescription: "String custom variable loop configuration",
		Optional:            true,
		Attributes: map[string]schema.Attribute{
			"variables": schema.ListNestedAttribute{
				MarkdownDescription: "Custom variables",
				Required:            true,
				Validators: []validator.List{
					listvalidator.SizeAtLeast(1),
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							MarkdownDescription: "Custom variable name",
							Required:            true,
						},
						"values": schema.ListAttribute{
							MarkdownDescription: "Custom variable values",
							Required:            true,
							ElementType:         types.StringType,
						},
					},
				},
			},
		},
	}
}
