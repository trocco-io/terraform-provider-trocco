package pipeline_definition

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

func TroccoDBTTaskConfigSchema() schema.Attribute {
	return schema.SingleNestedAttribute{
		MarkdownDescription: "The task configuration for the trocco dbt task.",
		Optional:            true,
		Attributes: map[string]schema.Attribute{
			"definition_id": schema.Int64Attribute{
				MarkdownDescription: "The definition id to use for the trocco dbt task",
				Required:            true,
			},
		},
	}
}
