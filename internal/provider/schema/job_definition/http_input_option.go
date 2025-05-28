package job_definition

import (
	troccoPlanModifier "terraform-provider-trocco/internal/provider/planmodifier"
	"terraform-provider-trocco/internal/provider/schema/job_definition/parser"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

func HttpInputOptionSchema() schema.Attribute {
	return schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "Attributes about source HTTP",
		Attributes: map[string]schema.Attribute{
			"url": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "URL to fetch data from",
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtLeast(1),
				},
			},
			"method": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "HTTP method (GET or POST)",
				Validators: []validator.String{
					stringvalidator.OneOf("GET", "POST"),
				},
			},
			"user_agent": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "User agent string to use for requests",
			},
			"charset": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("UTF-8"),
				MarkdownDescription: "Character set of the response",
			},
			"pager_type": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("disable"),
				MarkdownDescription: "Type of pagination (offset, cursor, disable)",
				Validators: []validator.String{
					stringvalidator.OneOf("offset", "cursor", "disable"),
				},
			},
			"pager_from_param": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Parameter name for the starting offset/page",
			},
			"pager_to_param": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Parameter name for the ending offset/page",
			},
			"pager_pages": schema.Int64Attribute{
				Optional:            true,
				MarkdownDescription: "Number of pages to fetch",
			},
			"pager_start": schema.Int64Attribute{
				Optional:            true,
				MarkdownDescription: "Starting page number",
			},
			"pager_step": schema.Int64Attribute{
				Optional:            true,
				MarkdownDescription: "Step size for pagination",
			},
			"cursor_request_parameter_cursor_name": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Parameter name for cursor-based pagination",
			},
			"cursor_response_parameter_cursor_json_path": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "JSONPath to extract cursor value from response",
			},
			"cursor_request_parameter_limit_name": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Parameter name for limit in cursor-based pagination",
			},
			"cursor_request_parameter_limit_value": schema.Int64Attribute{
				Optional:            true,
				MarkdownDescription: "Value for limit parameter in cursor-based pagination",
			},
			"request_params": schema.SetNestedAttribute{
				Optional:            true,
				MarkdownDescription: "Request parameters to include in the URL",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"key": schema.StringAttribute{
							Required:            true,
							MarkdownDescription: "Parameter key",
						},
						"value": schema.StringAttribute{
							Required:            true,
							Sensitive:           true,
							MarkdownDescription: "Parameter value",
						},
						"masking": schema.BoolAttribute{
							Optional:            true,
							Computed:            true,
							Default:             booldefault.StaticBool(false),
							MarkdownDescription: "Whether to mask this parameter",
						},
					},
				},
			},
			"request_body": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Request body for POST/PUT requests",
			},
			"request_headers": schema.SetNestedAttribute{
				Optional:            true,
				MarkdownDescription: "HTTP headers to include in the request",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"key": schema.StringAttribute{
							Required:            true,
							MarkdownDescription: "Header key",
						},
						"value": schema.StringAttribute{
							Required:            true,
							Sensitive:           true,
							MarkdownDescription: "Header value",
						},
						"masking": schema.BoolAttribute{
							Optional:            true,
							Computed:            true,
							Default:             booldefault.StaticBool(false),
							MarkdownDescription: "Whether to mask this header",
						},
					},
				},
			},
			"success_code": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("200"),
				MarkdownDescription: "HTTP status code that indicate success",
			},
			"open_timeout": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(2000),
				MarkdownDescription: "Timeout for opening connection in seconds",
				Validators: []validator.Int64{
					int64validator.AtLeast(1),
				},
			},
			"read_timeout": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(10000),
				MarkdownDescription: "Timeout for reading response in seconds",
				Validators: []validator.Int64{
					int64validator.AtLeast(1),
				},
			},
			"max_retries": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(5),
				MarkdownDescription: "Maximum number of retry attempts",
				Validators: []validator.Int64{
					int64validator.AtLeast(0),
				},
			},
			"retry_interval": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(10000),
				MarkdownDescription: "Interval between retry attempts in seconds",
				Validators: []validator.Int64{
					int64validator.AtLeast(0),
				},
			},
			"request_interval": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(0),
				MarkdownDescription: "Interval between requests in seconds",
				Validators: []validator.Int64{
					int64validator.AtLeast(0),
				},
			},
			"jsonpath_parser": parser.JsonpathParserSchema(),
			"xml_parser":      parser.XmlParserSchema(),
			"excel_parser":    parser.ExcelParserSchema(),
			"ltsv_parser":     parser.LtsvParserSchema(),
			"jsonl_parser":    parser.JsonlParserSchema(),
			"csv_parser":      parser.CsvParserSchema(),
			// Unsupported: Parquet parser stub (here only to keep FileParserPlanModifier quiet).
			"parquet_parser":           parser.ParquetParserSchema(),
			"custom_variable_settings": CustomVariableSettingsSchema(),
		},
		PlanModifiers: []planmodifier.Object{
			&troccoPlanModifier.FileParserPlanModifier{},
		},
	}
}
