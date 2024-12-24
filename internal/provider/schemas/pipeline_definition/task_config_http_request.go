package pipeline_definition

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

func HTTPRequestTaskConfig() schema.Attribute {
	return schema.SingleNestedAttribute{
		Optional: true,
		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				Required: true,
			},
			"connection_id": schema.Int64Attribute{
				Optional: true,
			},
			"http_method": schema.StringAttribute{
				Required: true,
			},
			"url": schema.StringAttribute{
				Required: true,
			},
			"request_body": schema.StringAttribute{
				Optional: true,
			},
			"request_headers": schema.ListNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"key": schema.StringAttribute{
							Required: true,
						},
						"value": schema.StringAttribute{
							Required: true,
						},
						"masking": schema.BoolAttribute{
							Optional: true,
						},
					},
				},
			},
			"request_parameters": schema.ListNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"key": schema.StringAttribute{
							Required: true,
						},
						"value": schema.StringAttribute{
							Required: true,
						},
						"masking": schema.BoolAttribute{
							Optional: true,
						},
					},
				},
			},
			"custom_variables": CustomVariables(),
		},
	}
}
