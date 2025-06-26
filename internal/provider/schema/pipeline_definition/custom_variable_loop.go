package pipeline_definition

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

func CustomVariableLoop() schema.Attribute {
	return schema.SingleNestedAttribute{
		MarkdownDescription: "The custom variable loop of the pipeline definition",
		Optional:            true,
		Attributes: map[string]schema.Attribute{
			"type": schema.StringAttribute{
				MarkdownDescription: "The type of the custom variable loop. Allowed values: \"string\", \"period\", \"bigquery\", \"snowflake\", \"redshift\".",
				Required:            true,
			},
			"is_parallel_execution_allowed": schema.BoolAttribute{
				MarkdownDescription: "Whether parallel execution is allowed",
				Optional:            true,
				Computed:            true,
			},
			"is_stopped_on_errors": schema.BoolAttribute{
				MarkdownDescription: "Whether the loop is stopped on errors",
				Optional:            true,
				Computed:            true,
			},
			"max_errors": schema.Int64Attribute{
				MarkdownDescription: "The maximum number of errors",
				Optional:            true,
				Computed:            true,
			},
			"string_config":    StringCustomVariableLoopConfig(),
			"period_config":    PeriodCustomVariableLoopConfig(),
			"bigquery_config":  BigqueryCustomVariableLoopConfig(),
			"snowflake_config": SnowflakeCustomVariableLoopConfig(),
			"redshift_config":  RedshiftCustomVariableLoopConfig(),
		},
	}
}
