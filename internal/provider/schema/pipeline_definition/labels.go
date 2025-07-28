package pipeline_definition

import (
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func LabelsSchema() schema.Attribute {
	return schema.SetAttribute{
		MarkdownDescription: "Labels to be attached to the pipeline definition",
		Optional:            true,
		ElementType:         types.StringType,
	}
}
