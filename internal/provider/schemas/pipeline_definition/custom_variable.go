package pipeline_definition

import (
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
)

func NewCustomVariableAttribute() schema.Attribute {
	return schema.SetNestedAttribute{
		Optional: true,
		NestedObject: schema.NestedAttributeObject{
			Attributes: map[string]schema.Attribute{
				"name": schema.StringAttribute{
					Required: true,
				},
				"type": schema.StringAttribute{
					Required: true,
				},
				"value": schema.StringAttribute{
					Optional: true,
				},
				"quantity": schema.Int64Attribute{
					Optional: true,
				},
				"unit": schema.StringAttribute{
					Optional: true,
				},
				"direction": schema.StringAttribute{
					Optional: true,
				},
				"format": schema.StringAttribute{
					Optional: true,
				},
				"time_zone": schema.StringAttribute{
					Optional: true,
				},
			},
		},
	}
}
