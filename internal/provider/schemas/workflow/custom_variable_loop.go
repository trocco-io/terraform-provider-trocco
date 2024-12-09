package workflow

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func NewCustomVariableLoopAttribute() schema.Attribute {
	return schema.SingleNestedAttribute{
		Optional: true,
		Attributes: map[string]schema.Attribute{
			"type": schema.StringAttribute{
				Required: true,
			},
			"string_config":    newStringCustomVariableLoopConfigAttribute(),
			"period_config":    newPeriodCustomVariableLoopConfigAttribute(),
			"bigquery_config":  newBigqueryCustomVariableLoopConfigAttribute(),
			"snowflake_config": newSnowflakeCustomVariableLoopConfigAttribute(),
			"redshift_config":  newRedshiftCustomVariableLoopConfigAttribute(),
		},
	}
}

func newStringCustomVariableLoopConfigAttribute() schema.Attribute {
	return schema.SingleNestedAttribute{
		Optional: true,
		Attributes: map[string]schema.Attribute{
			"variables": schema.ListNestedAttribute{
				Required: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Required: true,
						},
						"values": schema.ListAttribute{
							Required:    true,
							ElementType: types.StringType,
						},
					},
				},
			},
		},
	}
}

func newPeriodCustomVariableLoopConfigAttribute() schema.Attribute {
	return schema.SingleNestedAttribute{
		Optional: true,
		Attributes: map[string]schema.Attribute{
			"interval": schema.StringAttribute{
				Required: true,
			},
			"time_zone": schema.StringAttribute{
				Required: true,
			},
			"from": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"value": schema.Int64Attribute{
						Required: true,
					},
					"unit": schema.StringAttribute{
						Required: true,
					},
				},
			},
			"to": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"value": schema.Int64Attribute{
						Required: true,
					},
					"unit": schema.StringAttribute{
						Required: true,
					},
				},
			},
			"variables": schema.ListNestedAttribute{
				Required: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Required: true,
						},
						"offset": schema.SingleNestedAttribute{
							Required: true,
							Attributes: map[string]schema.Attribute{
								"value": schema.Int64Attribute{
									Required: true,
								},
								"unit": schema.StringAttribute{
									Required: true,
								},
							},
						},
					},
				},
			},
		},
	}
}

func newBigqueryCustomVariableLoopConfigAttribute() schema.Attribute {
	return schema.SingleNestedAttribute{
		Optional: true,
		Attributes: map[string]schema.Attribute{
			"connection_id": schema.Int64Attribute{
				Required: true,
			},
			"query": schema.StringAttribute{
				Required: true,
			},
			"variables": schema.ListAttribute{
				Required:    true,
				ElementType: types.StringType,
			},
		},
	}
}

func newSnowflakeCustomVariableLoopConfigAttribute() schema.Attribute {
	return schema.SingleNestedAttribute{
		Optional: true,
		Attributes: map[string]schema.Attribute{
			"connection_id": schema.Int64Attribute{
				Required: true,
			},
			"query": schema.StringAttribute{
				Required: true,
			},
			"warehouse": schema.StringAttribute{
				Required: true,
			},
			"variables": schema.ListAttribute{
				Required:    true,
				ElementType: types.StringType,
			},
		},
	}
}

func newRedshiftCustomVariableLoopConfigAttribute() schema.Attribute {
	return schema.SingleNestedAttribute{
		Optional: true,
		Attributes: map[string]schema.Attribute{
			"connection_id": schema.Int64Attribute{
				Required: true,
			},
			"query": schema.StringAttribute{
				Required: true,
			},
			"database": schema.StringAttribute{
				Required: true,
			},
			"variables": schema.ListAttribute{
				Required:    true,
				ElementType: types.StringType,
			},
		},
	}
}
