package workflow

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

func NewPeriodCustomVariableLoopConfigAttribute() schema.Attribute {
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
