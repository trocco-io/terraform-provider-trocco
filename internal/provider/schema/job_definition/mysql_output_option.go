package job_definition

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

func MysqlOutputOptionSchema() schema.Attribute {
	return schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "Attributes of destination MySQL settings",
		Attributes: map[string]schema.Attribute{
			"mysql_connection_id": schema.Int64Attribute{
				Required: true,
				Validators: []validator.Int64{
					int64validator.AtLeast(1),
				},
				MarkdownDescription: "ID of MySQL connection",
			},
			"database": schema.StringAttribute{
				Required: true,
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtLeast(1),
				},
				MarkdownDescription: "Database name",
			},
			"table": schema.StringAttribute{
				Required: true,
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtLeast(1),
				},
				MarkdownDescription: "Table name",
			},
			"mode": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Validators: []validator.String{
					stringvalidator.OneOf("insert", "insert_direct", "truncate_insert", "replace", "merge", "merge_direct"),
				},
				Default:             stringdefault.StaticString("insert"),
				MarkdownDescription: "Transfer mode. One of `insert`, `insert_direct`, `truncate_insert`, `replace`, `merge`, `merge_direct`",
			},
			"retry_limit": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				Validators: []validator.Int64{
					int64validator.AtLeast(0),
				},
				Default:             int64default.StaticInt64(12),
				MarkdownDescription: "Maximum number of retries",
			},
			"retry_wait": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				Validators: []validator.Int64{
					int64validator.AtLeast(0),
				},
				Default:             int64default.StaticInt64(1000),
				MarkdownDescription: "Retry wait time (milliseconds)",
			},
			"max_retry_wait": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				Validators: []validator.Int64{
					int64validator.AtLeast(0),
				},
				Default:             int64default.StaticInt64(1800000),
				MarkdownDescription: "Maximum retry wait time (milliseconds)",
			},
			"default_time_zone": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtLeast(1),
				},
				Default:             stringdefault.StaticString("UTC"),
				MarkdownDescription: "Default time zone",
			},
			"before_load": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "SQL statement to execute before data transfer",
			},
			"after_load": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "SQL statement to execute after data transfer",
			},
			"mysql_output_option_column_options": schema.ListNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Required:            true,
							MarkdownDescription: "Column name",
						},
						"type": schema.StringAttribute{
							Required: true,
							Validators: []validator.String{
								stringvalidator.OneOf("TINYTEXT", "TEXT", "MEDIUMTEXT", "LONGTEXT", "DECIMAL"),
							},
							MarkdownDescription: "Column type. One of `TINYTEXT`, `TEXT`, `MEDIUMTEXT`, `LONGTEXT`, `DECIMAL`",
						},
						"scale": schema.Int64Attribute{
							Optional: true,
							Computed: true,
							Validators: []validator.Int64{
								int64validator.AtLeast(0),
							},
							MarkdownDescription: "Number of decimal places for DECIMAL type (required when type is DECIMAL)",
						},
						"precision": schema.Int64Attribute{
							Optional: true,
							Computed: true,
							Validators: []validator.Int64{
								int64validator.AtLeast(0),
							},
							MarkdownDescription: "Total number of digits for DECIMAL type (required when type is DECIMAL)",
						},
					},
				},
			},
			"custom_variable_settings": CustomVariableSettingsSchema(),
		},
	}
}
