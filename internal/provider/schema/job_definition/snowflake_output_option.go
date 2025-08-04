package job_definition

import (
	planModifier "terraform-provider-trocco/internal/provider/planmodifier"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func SnowflakeOutputOptionSchema() schema.Attribute {
	return schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "Attributes of destination Snowflake settings",
		Attributes: map[string]schema.Attribute{
			"warehouse": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "Warehouse name",
			},
			"database": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "Database name",
			},
			"schema": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "Schema name",
			},
			"table": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "Table name",
			},
			"mode": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Default:  stringdefault.StaticString("insert"),
				Validators: []validator.String{
					stringvalidator.OneOf("insert", "insert_direct", "truncate_insert", "replace", "merge"),
				},
				MarkdownDescription: "Transfer mode",
			},
			"empty_field_as_null": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
				MarkdownDescription: "Replace empty string with NULL",
			},
			"delete_stage_on_error": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
				MarkdownDescription: "Delete temporary stage on error",
			},
			"batch_size": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				Default:  int64default.StaticInt64(50),
				Validators: []validator.Int64{
					int64validator.AtLeast(1),
				},
				MarkdownDescription: "Batch size (MB)",
			},
			"retry_limit": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				Default:  int64default.StaticInt64(12),
				Validators: []validator.Int64{
					int64validator.AtLeast(1),
				},
				MarkdownDescription: "Maximum retry limit",
			},
			"retry_wait": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				Default:  int64default.StaticInt64(1000),
				Validators: []validator.Int64{
					int64validator.AtLeast(1),
				},
				MarkdownDescription: "Retry wait time (milliseconds)",
			},
			"max_retry_wait": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				Default:  int64default.StaticInt64(1800000),
				Validators: []validator.Int64{
					int64validator.AtLeast(1),
				},
				MarkdownDescription: "Maximum retry wait time (milliseconds)",
			},
			"default_time_zone": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("UTC"),
				MarkdownDescription: "Default time zone",
			},
			"snowflake_connection_id": schema.Int64Attribute{
				Required:            true,
				MarkdownDescription: "Snowflake connection ID",
			},
			"snowflake_output_option_column_options": schema.ListNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Required:            true,
							MarkdownDescription: "Column name",
						},
						"type": schema.StringAttribute{
							Required:            true,
							MarkdownDescription: "Data type",
						},
						"value_type": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Value type",
						},
						"timestamp_format": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Timestamp format",
						},
						"timezone": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Time zone",
						},
					},
					PlanModifiers: []planmodifier.Object{
						&planModifier.SnowflakeOutputOptionColumnPlanModifier{},
					},
				},
			},
			"snowflake_output_option_merge_keys": schema.SetAttribute{
				Optional:            true,
				ElementType:         types.StringType,
				MarkdownDescription: "Merge keys (only applicable if mode is 'merge')",
			},
			"custom_variable_settings": CustomVariableSettingsSchema(),
		},
		PlanModifiers: []planmodifier.Object{
			&planModifier.SnowflakeOutputOptionPlanModifier{},
		},
	}
}
