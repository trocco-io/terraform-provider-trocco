package job_definition

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func RedshiftOutputOptionSchema() schema.Attribute {
	return schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "Attributes of destination Redshift settings",
		Attributes: map[string]schema.Attribute{
			"redshift_connection_id": schema.Int64Attribute{
				Required: true,
				Validators: []validator.Int64{
					int64validator.AtLeast(1),
				},
				MarkdownDescription: "ID of Redshift connection",
			},
			"database": schema.StringAttribute{
				Required: true,
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtLeast(1),
				},
				MarkdownDescription: "Database name",
			},
			"schema": schema.StringAttribute{
				Required: true,
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtLeast(1),
				},
				MarkdownDescription: "Schema name",
			},
			"table": schema.StringAttribute{
				Required: true,
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtLeast(1),
				},
				MarkdownDescription: "Table name",
			},
			"create_table_constraint": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Constraint added to CREATE TABLE statement",
			},
			"create_table_option": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Option added to CREATE TABLE statement",
			},
			"s3_bucket": schema.StringAttribute{
				Required: true,
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtLeast(1),
				},
				MarkdownDescription: "S3 bucket for temporary data",
			},
			"s3_key_prefix": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "S3 key prefix for temporary data",
			},
			"delete_s3_temp_file": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
				MarkdownDescription: "Whether to delete temporary S3 files after transfer. Default is true.",
			},
			"copy_iam_role_name": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "IAM role name for COPY command",
			},
			"retry_limit": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				Validators: []validator.Int64{
					int64validator.AtLeast(0),
				},
				Default:             int64default.StaticInt64(12),
				MarkdownDescription: "Maximum number of retries. Default is 12.",
			},
			"retry_wait": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				Validators: []validator.Int64{
					int64validator.AtLeast(0),
				},
				Default:             int64default.StaticInt64(1000),
				MarkdownDescription: "Retry wait time in milliseconds. Default is 1000.",
			},
			"max_retry_wait": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				Validators: []validator.Int64{
					int64validator.AtLeast(0),
				},
				Default:             int64default.StaticInt64(1800000),
				MarkdownDescription: "Maximum retry wait time in milliseconds. Default is 1800000.",
			},
			"mode": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Validators: []validator.String{
					stringvalidator.OneOf("insert", "insert_direct", "truncate_insert", "replace", "merge"),
				},
				Default:             stringdefault.StaticString("insert"),
				MarkdownDescription: "Transfer mode. One of `insert`, `insert_direct`, `truncate_insert`, `replace`, `merge`",
			},
			"default_time_zone": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtLeast(1),
				},
				Default:             stringdefault.StaticString("UTC"),
				MarkdownDescription: "Default time zone. Default is UTC.",
			},
			"before_load": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "SQL statement to execute before data transfer",
			},
			"after_load": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "SQL statement to execute after data transfer",
			},
			"batch_size": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				Validators: []validator.Int64{
					int64validator.AtLeast(0),
				},
				Default:             int64default.StaticInt64(16384),
				MarkdownDescription: "Batch size in KB. Default is 16384.",
			},
			"redshift_output_option_column_options": schema.ListNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Required:            true,
							MarkdownDescription: "Column name",
						},
						"type": schema.StringAttribute{
							Optional: true,
							Validators: []validator.String{
								stringvalidator.OneOf(
									"BIGINT", "VARCHAR", "BOOLEAN", "DOUBLE PRECISION",
									"CLOB", "TIMESTAMP", "TIME", "DATE",
								),
							},
							MarkdownDescription: "Column type. One of `BIGINT`, `VARCHAR`, `BOOLEAN`, `DOUBLE PRECISION`, `CLOB`, `TIMESTAMP`, `TIME`, `DATE`",
						},
						"value_type": schema.StringAttribute{
							Optional: true,
							Validators: []validator.String{
								stringvalidator.OneOf(
									"byte", "short", "int", "long", "double", "float",
									"boolean", "string", "nstring", "date", "time",
									"timestamp", "decimal", "json", "null", "pass", "coalesce",
								),
							},
							MarkdownDescription: "Value type",
						},
						"timestamp_format": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Timestamp format (required when type is TIMESTAMP, TIME, or DATE and value_type is string or nstring)",
						},
						"timezone": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Timezone (applicable when type is TIMESTAMP)",
						},
					},
				},
			},
			"redshift_output_option_merge_keys": schema.SetAttribute{
				Optional:            true,
				ElementType:         types.StringType,
				MarkdownDescription: "Merge keys (only applicable if mode is 'merge')",
			},
			"custom_variable_settings": CustomVariableSettingsSchema(),
		},
	}
}
