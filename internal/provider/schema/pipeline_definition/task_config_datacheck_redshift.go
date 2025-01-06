package pipeline_definition

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

func RedshiftDatacheckTaskConfig() schema.Attribute {
	return schema.SingleNestedAttribute{
		Optional: true,
		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				Required: true,
			},
			"connection_id": schema.Int64Attribute{
				Required: true,
			},
			"query": schema.StringAttribute{
				Optional: true,
			},
			"operator": schema.StringAttribute{
				Optional: true,
			},
			"query_result": schema.Int64Attribute{
				Optional: true,
			},
			"accepts_null": schema.BoolAttribute{
				Optional: true,
			},
			"database": schema.StringAttribute{
				Optional: true,
			},
			"custom_variables": CustomVariables(),
		},
	}
}
