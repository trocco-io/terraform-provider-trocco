package pipeline_definition

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

func NewCustomVariableLoopAttribute() schema.Attribute {
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
			"string_config":    NewStringCustomVariableLoopConfigAttribute(),
			"period_config":    NewPeriodCustomVariableLoopConfigAttribute(),
			"bigquery_config":  NewBigqueryCustomVariableLoopConfigAttribute(),
			"snowflake_config": NewSnowflakeCustomVariableLoopConfigAttribute(),
			"redshift_config":  NewRedshiftCustomVariableLoopConfigAttribute(),
		},
	}
}
