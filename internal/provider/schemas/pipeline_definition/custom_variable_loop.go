package workflow

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
			"string_config":    NewStringCustomVariableLoopConfigAttribute(),
			"period_config":    NewPeriodCustomVariableLoopConfigAttribute(),
			"bigquery_config":  NewBigqueryCustomVariableLoopConfigAttribute(),
			"snowflake_config": NewSnowflakeCustomVariableLoopConfigAttribute(),
			"redshift_config":  NewRedshiftCustomVariableLoopConfigAttribute(),
		},
	}
}
