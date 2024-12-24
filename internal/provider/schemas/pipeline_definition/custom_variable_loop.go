package pipeline_definition

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

func CustomVariableLoop() schema.Attribute {
	return schema.SingleNestedAttribute{
		Optional: true,
		Attributes: map[string]schema.Attribute{
			"type": schema.StringAttribute{
				Required: true,
			},
			"is_parallel_execution_allowed": schema.BoolAttribute{
				Optional: true,
				Computed: true,
			},
			"is_stopped_on_errors": schema.BoolAttribute{
				Optional: true,
				Computed: true,
			},
			"max_errors": schema.Int64Attribute{
				Optional: true,
				Computed: true,
			},
			"string_config":    StringCustomVariableLoopConfig(),
			"period_config":    PeriodCustomVariableLoopConfig(),
			"bigquery_config":  BigqueryCustomVariableLoopConfig(),
			"snowflake_config": SnowflakeCustomVariableLoopConfig(),
			"redshift_config":  RedshiftCustomVariableLoopConfig(),
		},
	}
}
