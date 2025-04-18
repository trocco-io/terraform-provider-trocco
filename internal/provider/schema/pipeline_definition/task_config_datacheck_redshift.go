package pipeline_definition

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"

	troccoPlanModifier "terraform-provider-trocco/internal/provider/planmodifier"
)

func RedshiftDatacheckTaskConfig() schema.Attribute {
	return schema.SingleNestedAttribute{
		MarkdownDescription: "The task configuration for the datacheck task.",
		Optional:            true,
		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				MarkdownDescription: "The name of the datacheck task",
				Required:            true,
			},
			"connection_id": schema.Int64Attribute{
				MarkdownDescription: "The connection id to use for the datacheck task",
				Required:            true,
			},
			"query": schema.StringAttribute{
				MarkdownDescription: "The query to run for the datacheck task",
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					troccoPlanModifier.NormalizeSQLQuery(),
				},
			},
			"operator": schema.StringAttribute{
				MarkdownDescription: "The operator to use for the datacheck task",
				Optional:            true,
			},
			"query_result": schema.Int64Attribute{
				MarkdownDescription: "The query result to use for the datacheck task",
				Optional:            true,
			},
			"accepts_null": schema.BoolAttribute{
				MarkdownDescription: "Whether the datacheck task accepts null values",
				Optional:            true,
			},
			"database": schema.StringAttribute{
				MarkdownDescription: "The database to use for the datacheck task",
				Optional:            true,
			},
			"custom_variables": CustomVariables(),
		},
	}
}
