package pipeline_definition

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

func HTTPRequestTaskConfig() schema.Attribute {
	return schema.SingleNestedAttribute{
		MarkdownDescription: "The task configuration for the HTTP request task.",
		Optional:            true,
		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				MarkdownDescription: "The name of the task",
				Required:            true,
			},
			"connection_id": schema.Int64Attribute{
				MarkdownDescription: "The connection id to use for the task",
				Optional:            true,
			},
			"http_method": schema.StringAttribute{
				MarkdownDescription: "The HTTP method to use for the request",
				Required:            true,
			},
			"url": schema.StringAttribute{
				MarkdownDescription: "The URL to send the request to",
				Required:            true,
			},
			"request_body": schema.StringAttribute{
				MarkdownDescription: "The body of the request",
				Optional:            true,
			},
			"request_headers": schema.ListNestedAttribute{
				MarkdownDescription: "The headers to send with the request",
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"key": schema.StringAttribute{
							MarkdownDescription: "The key of the header",
							Required:            true,
						},
						"value": schema.StringAttribute{
							MarkdownDescription: "The value of the header",
							Required:            true,
							Sensitive:           true,
						},
						"masking": schema.BoolAttribute{
							MarkdownDescription: "Whether to mask the value of the header",
							Optional:            true,
						},
					},
				},
			},
			"request_parameters": schema.ListNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"key": schema.StringAttribute{
							MarkdownDescription: "The key of the parameter",
							Required:            true,
						},
						"value": schema.StringAttribute{
							MarkdownDescription: "The value of the parameter",
							Required:            true,
							Sensitive:           true,
						},
						"masking": schema.BoolAttribute{
							MarkdownDescription: "Whether to mask the value of the parameter",
							Optional:            true,
						},
					},
				},
			},
			"custom_variables": CustomVariables(),
		},
	}
}
