package pipeline_definition

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func NewRedshiftCustomVariableLoopConfigAttribute() schema.Attribute {
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
