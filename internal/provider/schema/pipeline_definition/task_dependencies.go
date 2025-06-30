package pipeline_definition

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

func TaskDependenciesSchema() schema.Attribute {
	return schema.SetNestedAttribute{
		MarkdownDescription: "The task dependencies of the workflow.",
		Optional:            true,
		NestedObject: schema.NestedAttributeObject{
			Attributes: map[string]schema.Attribute{
				"source": schema.StringAttribute{
					MarkdownDescription: "The source task key.",
					Required:            true,
				},
				"destination": schema.StringAttribute{
					MarkdownDescription: "The destination task key.",
					Required:            true,
				},
			},
		},
	}
}
