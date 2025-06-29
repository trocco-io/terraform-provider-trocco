package pipeline_definition

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

func SnowflakeDatamartTaskConfigSchema() schema.Attribute {
	return schema.SingleNestedAttribute{
		MarkdownDescription: "The task configuration for the datamart task.",
		Optional:            true,
		Attributes: map[string]schema.Attribute{
			"definition_id": schema.Int64Attribute{
				MarkdownDescription: "The definition id to use for the datamart task",
				Required:            true,
			},
			"custom_variable_loop": CustomVariableLoopSchema(),
		},
	}
}
