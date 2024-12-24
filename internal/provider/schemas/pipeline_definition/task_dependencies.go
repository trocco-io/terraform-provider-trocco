package pipeline_definition

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

func TaskDependencies() schema.Attribute {
	return schema.SetNestedAttribute{
		MarkdownDescription: "The task dependencies of the workflow.",
		Required:            true,
		NestedObject: schema.NestedAttributeObject{
			Attributes: map[string]schema.Attribute{
				"source": schema.StringAttribute{
					Required: true,
				},
				"destination": schema.StringAttribute{
					Required: true,
				},
			},
		},
	}
}
