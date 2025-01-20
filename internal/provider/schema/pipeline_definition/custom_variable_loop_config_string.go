package pipeline_definition

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func StringCustomVariableLoopConfig() schema.Attribute {
	return schema.SingleNestedAttribute{
		MarkdownDescription: "String custom variable loop configuration",
		Optional:            true,
		Attributes: map[string]schema.Attribute{
			"variables": schema.ListNestedAttribute{
				MarkdownDescription: "Custom variables",
				Required:            true,
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
