package job_definition

import (
	planmodifier2 "terraform-provider-trocco/internal/provider/planmodifier"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

func PostgresqlInputOptionSchema() schema.Attribute {
	return schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "Attributes of source postgresql",
		Attributes: map[string]schema.Attribute{
			"database": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "database name",
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtLeast(1),
				},
			},
			"schema": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "schema name",
				Default:             stringdefault.StaticString("public"),
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtLeast(1),
				},
			},
			"query": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "If you want to use all record loading, specify it.",
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtLeast(1),
				},
			},
			"incremental_loading_enabled": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
				MarkdownDescription: "If it is true, to be incremental loading. If it is false, to be all record loading",
			},
			"table": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "table name. If you want to use incremental loading, specify it.",
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtLeast(1),
				},
			},
			"incremental_columns": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Columns to determine incremental data",
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtLeast(1),
				},
			},
			"last_record": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Last record transferred. The value of the column specified here is stored in “Last Transferred Record” for each transfer, and for the second and subsequent transfers, only records for which the value of the “Column for Determining Incremental Data” is greater than the value of the previous transfer (= “Last Transferred Record”) are transferred. If you wish to specify multiple columns, specify them separated by commas. If not specified, the primary key is used.",
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtLeast(1),
				},
			},
			"fetch_rows": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(10000),
				MarkdownDescription: "Number of records processed by the cursor at one time",
				Validators: []validator.Int64{
					int64validator.AtLeast(1),
				},
			},
			"connect_timeout": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				Default:  int64default.StaticInt64(300),
				Validators: []validator.Int64{
					int64validator.AtLeast(1),
				},
				MarkdownDescription: "Connection timeout (sec)",
			},
			"socket_timeout": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				Default:  int64default.StaticInt64(1800),
				Validators: []validator.Int64{
					int64validator.AtLeast(1),
				},
				MarkdownDescription: "Socket timeout (seconds)",
			},
			"default_time_zone": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Default time zone. enter the server-side time zone setting for PostgreSQL. If the time zone is set to Japan, enter “Asia/Tokyo”.",
				Default:             stringdefault.StaticString("UTC"),
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtLeast(1),
				},
			},
			"postgresql_connection_id": schema.Int64Attribute{
				Required: true,
				Validators: []validator.Int64{
					int64validator.AtLeast(1),
				},
				MarkdownDescription: "ID of Postgresql connection",
			},
			"input_option_columns": schema.ListNestedAttribute{
				Required:            true,
				MarkdownDescription: "List of columns to be retrieved and their types",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Required: true,
							Validators: []validator.String{
								stringvalidator.UTF8LengthAtLeast(1),
							},
							MarkdownDescription: "Column name",
						},
						"type": schema.StringAttribute{
							Required:            true,
							MarkdownDescription: "Column type",
							Validators: []validator.String{
								stringvalidator.OneOf("boolean", "long", "timestamp", "double", "string", "json"),
							},
						},
					},
				},
				Validators: []validator.List{
					listvalidator.SizeAtLeast(1),
				},
			},
			"input_option_column_options": schema.ListNestedAttribute{
				Optional:            true,
				MarkdownDescription: "List of unsupported data types and their convertible types",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"column_name": schema.StringAttribute{
							Required: true,
							Validators: []validator.String{
								stringvalidator.UTF8LengthAtLeast(1),
							},
							MarkdownDescription: "Column name",
						},
						"column_value_type": schema.StringAttribute{
							Required:            true,
							MarkdownDescription: "Column value type",
							Validators: []validator.String{
								stringvalidator.OneOf("boolean", "long", "timestamp", "double", "string", "json"),
							},
						},
					},
				},
			},
			"custom_variable_settings": CustomVariableSettingsSchema(),
		},
		PlanModifiers: []planmodifier.Object{
			&planmodifier2.PostgresqlInputOptionPlanModifier{},
		},
	}
}
