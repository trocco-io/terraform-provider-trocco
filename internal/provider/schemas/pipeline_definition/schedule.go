package pipeline_definition

import (
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
)

func Schedule() schema.Attribute {
	return schema.SetNestedAttribute{
		Optional: true,
		NestedObject: schema.NestedAttributeObject{
			Attributes: map[string]schema.Attribute{
				"frequency": schema.StringAttribute{
					Required: true,
				},
				"time_zone": schema.StringAttribute{
					Required: true,
				},
				"minute": schema.Int64Attribute{
					Required: true,
				},
				"day": schema.Int64Attribute{
					Optional: true,
				},
				"day_of_week": schema.Int64Attribute{
					Optional: true,
				},
				"hour": schema.Int64Attribute{
					Optional: true,
				},
			},
		},
	}
}
