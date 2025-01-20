package pipeline_definition

import (
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func Labels() schema.Attribute {
	return schema.SetAttribute{
		MarkdownDescription: "The labels of the pipeline definition",
		Optional:            true,
		ElementType:         types.StringType,
	}
}
