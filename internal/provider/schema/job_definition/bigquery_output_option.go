package job_definition

import (
	planmodifier2 "terraform-provider-trocco/internal/provider/planmodifier"
	mergeKeysValidator "terraform-provider-trocco/internal/provider/validator/job_definition/output_option/bigquery_output_option"

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

func BigqueryOutputOptionSchema() schema.Attribute {
	return schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "Attributes of destination BigQuery settings",
		Attributes: map[string]schema.Attribute{
			"dataset": schema.StringAttribute{
				Required: true,
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtLeast(1),
				},
				MarkdownDescription: "Dataset name",
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
				Default:  stringdefault.StaticString("append"),
				Validators: []validator.String{
					stringvalidator.OneOf("append", "append_direct", "replace", "delete_in_advance", "merge"),
				},
				MarkdownDescription: "Transfer mode",
			},
			"auto_create_dataset": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
				MarkdownDescription: "Option for automatic data set generation",
			},
			"open_timeout_sec": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				Default:  int64default.StaticInt64(300),
				Validators: []validator.Int64{
					int64validator.AtLeast(1),
				},
				MarkdownDescription: "Timeout to start connection (seconds)",
			},
			"timeout_sec": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				Default:  int64default.StaticInt64(300),
				Validators: []validator.Int64{
					int64validator.AtLeast(1),
				},
				MarkdownDescription: "Time out (seconds)",
			},
			"send_timeout_sec": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				Default:  int64default.StaticInt64(300),
				Validators: []validator.Int64{
					int64validator.AtLeast(1),
				},
				MarkdownDescription: "Transmission timeout (sec)",
			},
			"read_timeout_sec": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				Default:  int64default.StaticInt64(300),
				Validators: []validator.Int64{
					int64validator.AtLeast(1),
				},
				MarkdownDescription: "Read timeout (seconds)",
			},
			"retries": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				Default:  int64default.StaticInt64(5),
				Validators: []validator.Int64{
					int64validator.AtLeast(1),
				},
				MarkdownDescription: "Number of retries",
			},
			"partitioning_type": schema.StringAttribute{
				Optional: true,
				Validators: []validator.String{
					stringvalidator.OneOf("ingestion_time", "time_unit_column"),
				},
				MarkdownDescription: "Partitioning type. If params is null, No partitions. ingestion_time: Partitioning by acquisition time. time_unit_column: Partitioning by time unit column",
			},
			"time_partitioning_type": schema.StringAttribute{
				Optional: true,
				Validators: []validator.String{
					stringvalidator.OneOf("HOUR", "DAY", "MONTH", "YEAR"),
				},
				MarkdownDescription: "Time partitioning type. If you specify anything for partitioning_type, this parameter is required",
			},
			"time_partitioning_field": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "If partitioning_type is time_unit_column, this parameter is required",
			},
			"time_partitioning_expiration_ms": schema.Int64Attribute{
				Optional:            true,
				MarkdownDescription: "Duration of partition(milliseconds). Duration of the partition (in milliseconds). There is no minimum value. The date of the partition plus this integer value is the expiration date. The default value is unspecified (keep forever).",
			},
			"location": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("US"),
				MarkdownDescription: "Location",
			},
			"template_table": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Template table. Generate schema information for inclusion in Google BigQuery from schema information in this table",
			},
			"bigquery_connection_id": schema.Int64Attribute{
				Required: true,
				Validators: []validator.Int64{
					int64validator.AtLeast(1),
				},
				MarkdownDescription: "Id of BigQuery connection",
			},
			"bigquery_output_option_column_options": schema.ListNestedAttribute{
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
								stringvalidator.OneOf(
									"BOOLEAN",
									"INTEGER",
									"FLOAT",
									"STRING",
									"TIMESTAMP",
									"DATETIME",
									"DATE",
									"RECORD",
									"NUMERIC",
								),
							},
							MarkdownDescription: "Column type",
						},
						"mode": schema.StringAttribute{
							Required: true,
							Validators: []validator.String{
								stringvalidator.OneOf(
									"NULLABLE",
									"REQUIRED",
									"REPEATED",
								),
							},
							MarkdownDescription: "Mode",
						},
						"timestamp_format": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Timestamp format",
						},
						"timezone": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Time zone",
						},
						"description": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Description",
						},
					},
				},
			},
			"bigquery_output_option_clustering_fields": schema.SetAttribute{
				Optional:            true,
				Computed:            true,
				ElementType:         types.StringType,
				MarkdownDescription: "Clustered column. Clustering can only be set when creating a new table. A maximum of four clustered columns can be specified.",
				PlanModifiers: []planmodifier.Set{
					planmodifier2.EmptySetForNull(),
				},
			},
			"bigquery_output_option_merge_keys": schema.SetAttribute{
				Optional:            true,
				Computed:            true,
				ElementType:         types.StringType,
				MarkdownDescription: "Merge key. The column to be used as the merge key. Required when mode is 'merge'.",
				Validators: []validator.Set{
					mergeKeysValidator.MergeKeysRequiredOnlyForMergeMode(),
				},
				PlanModifiers: []planmodifier.Set{
					planmodifier2.EmptySetForNull(),
				},
			},
			"custom_variable_settings": CustomVariableSettingsSchema(),
		},
		PlanModifiers: []planmodifier.Object{
			&planmodifier2.BigqueryOutputOptionPlanModifier{},
		},
	}
}
