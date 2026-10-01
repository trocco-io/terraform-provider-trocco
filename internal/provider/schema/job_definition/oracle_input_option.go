package job_definition

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

func OracleInputOptionSchema() schema.Attribute {
	return schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "Attributes of source Oracle Database",
		Attributes: map[string]schema.Attribute{
			"oracle_connection_id": schema.Int64Attribute{
				Required:            true,
				MarkdownDescription: "ID of Oracle connection",
				Validators: []validator.Int64{
					int64validator.AtLeast(1),
				},
			},
			"database": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Database name (SID or service name). Required when using host/port connection. Ignored when using TNS naming.",
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtLeast(1),
				},
			},
			"connection_type": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("sid"),
				MarkdownDescription: "How to interpret `database` field: `sid` (Oracle SID) or `service_name` (Oracle Service Name). Default: `sid`. Ignored when using TNS naming.",
				Validators: []validator.String{
					stringvalidator.OneOf("sid", "service_name"),
				},
			},
			"net_service_name": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Net service name from tnsnames.ora. Required when using TNS naming connection.",
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtLeast(1),
				},
			},
			"schema": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Schema name",
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtLeast(1),
				},
			},
			"query": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "SQL query to fetch data. Required when `incremental_loading_enabled` is false.",
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtLeast(1),
				},
			},
			"incremental_loading_enabled": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
				MarkdownDescription: "Whether to use incremental loading (true) or query-based transfer (false). Default: false.",
			},
			"table": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Table name. Required when `incremental_loading_enabled` is true.",
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtLeast(1),
				},
			},
			"incremental_columns": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Comma-separated column names to identify incremental records. If omitted, primary key is used.",
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtLeast(1),
				},
			},
			"source_time_zone": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Timezone (IANA format, e.g., 'Asia/Tokyo') used to interpret DATE/TIMESTAMP columns without timezone information. If omitted, the timezone of the machine running the transfer is used.",
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtLeast(1),
				},
			},
			"input_option_columns": schema.ListNestedAttribute{
				Required:            true,
				MarkdownDescription: "Column definitions. When updated, fully replaces the existing list (id ascending order).",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Required:            true,
							MarkdownDescription: "Column name",
							Validators: []validator.String{
								stringvalidator.UTF8LengthAtLeast(1),
							},
						},
						"type": schema.StringAttribute{
							Required:            true,
							MarkdownDescription: "Column type: `boolean`, `long`, `double`, `string`, `json`, or `timestamp`",
							Validators: []validator.String{
								stringvalidator.OneOf("boolean", "long", "double", "string", "json", "timestamp"),
							},
						},
						"format": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Timestamp format (strftime). Required when `type` is `timestamp`.",
							Validators: []validator.String{
								stringvalidator.UTF8LengthAtLeast(1),
							},
						},
						"timezone": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Column-level timezone (IANA format). If omitted, `source_time_zone` is used. Only for `type: timestamp`.",
							Validators: []validator.String{
								stringvalidator.UTF8LengthAtLeast(1),
							},
						},
					},
				},
			},
			"input_option_column_options": schema.ListNestedAttribute{
				Optional:            true,
				MarkdownDescription: "Per-column options (e.g., NUMBER-to-string conversion). When updated, fully replaces the existing list.",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"column_name": schema.StringAttribute{
							Required:            true,
							MarkdownDescription: "Column name to apply option to",
							Validators: []validator.String{
								stringvalidator.UTF8LengthAtLeast(1),
							},
						},
						"column_value_type": schema.StringAttribute{
							Required:            true,
							MarkdownDescription: "Value type override. Currently only `string` is supported.",
							Validators: []validator.String{
								stringvalidator.OneOf("string"),
							},
						},
					},
				},
			},
			"custom_variable_settings": CustomVariableSettingsSchema(),
		},
	}
}
