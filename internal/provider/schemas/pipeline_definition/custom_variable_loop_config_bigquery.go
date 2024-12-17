package workflow

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func NewBigqueryCustomVariableLoopConfigAttribute() schema.Attribute {
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
