package provider

import (
	"context"
	"fmt"
	"strconv"

	"terraform-provider-trocco/internal/client"
	"terraform-provider-trocco/internal/provider/custom_type"
	troccoPlanModifier "terraform-provider-trocco/internal/provider/planmodifier"
	"terraform-provider-trocco/internal/provider/utils"
	troccoValidator "terraform-provider-trocco/internal/provider/validator"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.Resource = &bigqueryDatamartDefinitionResource{}
var _ resource.ResourceWithImportState = &bigqueryDatamartDefinitionResource{}

func NewBigqueryDatamartDefinitionResource() resource.Resource {
	return &bigqueryDatamartDefinitionResource{}
}

type bigqueryDatamartDefinitionResource struct {
	client *client.TroccoClient
}

type bigqueryDatamartDefinitionModel struct {
	ID                       types.Int64                    `tfsdk:"id"`
	Name                     types.String                   `tfsdk:"name"`
	Description              types.String                   `tfsdk:"description"`
	IsRunnableConcurrently   types.Bool                     `tfsdk:"is_runnable_concurrently"`
	ResourceGroupID          types.Int64                    `tfsdk:"resource_group_id"`
	CustomVariableSettings   types.List                     `tfsdk:"custom_variable_settings"`
	BigqueryConnectionID     types.Int64                    `tfsdk:"bigquery_connection_id"`
	QueryMode                types.String                   `tfsdk:"query_mode"`
	Query                    custom_type.TrimmedStringValue `tfsdk:"query"`
	DestinationDataset       types.String                   `tfsdk:"destination_dataset"`
	DestinationTable         types.String                   `tfsdk:"destination_table"`
	WriteDisposition         types.String                   `tfsdk:"write_disposition"`
	BeforeLoad               custom_type.TrimmedStringValue `tfsdk:"before_load"`
	Partitioning             types.String                   `tfsdk:"partitioning"`
	PartitioningTime         types.String                   `tfsdk:"partitioning_time"`
	PartitioningField        types.String                   `tfsdk:"partitioning_field"`
	ClusteringFields         types.List                     `tfsdk:"clustering_fields"`
	Location                 types.String                   `tfsdk:"location"`
	MergeKeys                types.List                     `tfsdk:"merge_keys"`
	OnMatchedAction          types.String                   `tfsdk:"on_matched_action"`
	IncrementalColumn        types.String                   `tfsdk:"incremental_column"`
	ValidFromColumn          types.String                   `tfsdk:"valid_from_column"`
	ValidToColumn            types.String                   `tfsdk:"valid_to_column"`
	IsCurrentColumn          types.String                   `tfsdk:"is_current_column"`
	SchemaEvolutionMode      types.String                   `tfsdk:"schema_evolution_mode"`
	LookbackPeriodColumn     types.String                   `tfsdk:"lookback_period_column"`
	LookbackPeriodColumnType types.String                   `tfsdk:"lookback_period_column_type"`
	LookbackPeriodTimezone   types.String                   `tfsdk:"lookback_period_timezone"`
	LookbackPeriodFrom       types.Int64                    `tfsdk:"lookback_period_from"`
	LookbackPeriodTo         types.Int64                    `tfsdk:"lookback_period_to"`
	LookbackPeriodUnit       types.String                   `tfsdk:"lookback_period_unit"`
	Notifications            types.List                     `tfsdk:"notifications"`
	Schedules                types.Set                      `tfsdk:"schedules"`
	Labels                   types.Set                      `tfsdk:"labels"`
}

type customVariableSettingModel struct {
	Name      types.String `tfsdk:"name"`
	Type      types.String `tfsdk:"type"`
	Value     types.String `tfsdk:"value"`
	Quantity  types.Int64  `tfsdk:"quantity"`
	Unit      types.String `tfsdk:"unit"`
	Direction types.String `tfsdk:"direction"`
	Format    types.String `tfsdk:"format"`
	TimeZone  types.String `tfsdk:"time_zone"`
}

type datamartNotificationModel struct {
	ID               types.Int64                    `tfsdk:"id"`
	DestinationType  types.String                   `tfsdk:"destination_type"`
	SlackChannelID   types.Int64                    `tfsdk:"slack_channel_id"`
	EmailID          types.Int64                    `tfsdk:"email_id"`
	NotificationType types.String                   `tfsdk:"notification_type"`
	NotifyWhen       types.String                   `tfsdk:"notify_when"`
	RecordCount      types.Int64                    `tfsdk:"record_count"`
	RecordOperator   types.String                   `tfsdk:"record_operator"`
	Message          custom_type.TrimmedStringValue `tfsdk:"message"`
}

type scheduleModel struct {
	Frequency types.String `tfsdk:"frequency"`
	Minute    types.Int64  `tfsdk:"minute"`
	Hour      types.Int64  `tfsdk:"hour"`
	Day       types.Int64  `tfsdk:"day"`
	DayOfWeek types.Int64  `tfsdk:"day_of_week"`
	TimeZone  types.String `tfsdk:"time_zone"`
}

type labelModel struct {
	ID   types.Int64  `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

func (r *bigqueryDatamartDefinitionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_bigquery_datamart_definition"
}

func (r *bigqueryDatamartDefinitionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*client.TroccoClient)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *client.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}

	r.client = client
}

func (r *bigqueryDatamartDefinitionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Provides a TROCCO datamart definitions for Google BigQuery resource.",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "The ID of the datamart definition",
			},
			"name": schema.StringAttribute{
				Required: true,
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtMost(255),
				},
				MarkdownDescription: "Name of the datamart definition. It must be less than 256 characters",
			},
			"description": schema.StringAttribute{
				Optional: true,
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtLeast(1),
				},
				MarkdownDescription: "Description of the datamart definition. It must be at least 1 character",
			},
			"is_runnable_concurrently": schema.BoolAttribute{
				Required:            true,
				MarkdownDescription: "Specifies whether or not to run a job if another job with the same datamart definition is running at the time the job is run",
			},
			"resource_group_id": schema.Int64Attribute{
				Optional: true,
				Validators: []validator.Int64{
					int64validator.AtLeast(1),
				},
				MarkdownDescription: "ID of the resource group to which the datamart definition belongs",
			},
			"custom_variable_settings": schema.ListNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Required: true,
							Validators: []validator.String{
								troccoValidator.WrappingDollarValidator{},
							},
							MarkdownDescription: "Custom variable name. It must start and end with `$`",
						},
						"type": schema.StringAttribute{
							Required: true,
							Validators: []validator.String{
								stringvalidator.OneOf("string", "timestamp", "timestamp_runtime"),
							},
							MarkdownDescription: "Custom variable type. The following types are supported: `string`, `timestamp`, `timestamp_runtime`",
						},
						"value": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Fixed string which will replace variables at runtime. Required in `string` type",
						},
						"quantity": schema.Int64Attribute{
							Optional: true,
							Validators: []validator.Int64{
								int64validator.AtLeast(0),
							},
							MarkdownDescription: "Quantity used to calculate diff from context_time. Required in `timestamp` and `timestamp_runtime` types",
						},
						"unit": schema.StringAttribute{
							Optional: true,
							Validators: []validator.String{
								stringvalidator.OneOf("hour", "date", "month"),
							},
							MarkdownDescription: "Time unit used to calculate diff from context_time. The following units are supported: `hour`, `date`, `month`. Required in `timestamp` and `timestamp_runtime` types",
						},
						"direction": schema.StringAttribute{
							Optional: true,
							Validators: []validator.String{
								stringvalidator.OneOf("ago", "later"),
							},
							MarkdownDescription: "Direction of the diff from context_time. The following directions are supported: `ago`, `later`. Required in `timestamp` and `timestamp_runtime` types",
						},
						"format": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Format used to replace variables. Required in `timestamp` and `timestamp_runtime` types",
						},
						"time_zone": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Time zone used to format the timestamp. Required in `timestamp` and `timestamp_runtime` types",
						},
					},
					PlanModifiers: []planmodifier.Object{
						&troccoPlanModifier.CustomVariableSettingPlanModifier{},
					},
				},
			},
			"bigquery_connection_id": schema.Int64Attribute{
				Required: true,
				Validators: []validator.Int64{
					int64validator.AtLeast(1),
				},
				MarkdownDescription: "ID of the BigQuery connection which is used to communicate with Google BigQuery",
			},
			"query_mode": schema.StringAttribute{
				Required: true,
				Validators: []validator.String{
					stringvalidator.OneOf("insert", "query"),
				},
				MarkdownDescription: "The following query modes are supported: `insert`, `query`. You can simply specify the query and the destination table in insert mode. In query mode, you can write and execute any DML/DDL statement",
			},
			"query": schema.StringAttribute{
				Required:            true,
				CustomType:          custom_type.TrimmedStringType{},
				MarkdownDescription: "Query to be executed.",
			},
			"destination_dataset": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Destination dataset where the query result will be inserted. Required in `insert` mode",
			},
			"destination_table": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Destination table where the query result will be inserted. Required in `insert` mode",
			},
			"write_disposition": schema.StringAttribute{
				Optional: true,
				Validators: []validator.String{
					stringvalidator.OneOf("append", "truncate", "incremental", "scd_type_2"),
				},
				MarkdownDescription: "The following write dispositions are supported: `append`, `truncate`, `incremental`, `scd_type_2`. Required in `insert` mode",
			},
			"before_load": schema.StringAttribute{
				Optional:            true,
				CustomType:          custom_type.TrimmedStringType{},
				MarkdownDescription: "The query to be executed before loading the data into the destination table. Available only in `insert` mode",
			},
			"partitioning": schema.StringAttribute{
				Optional: true,
				Validators: []validator.String{
					stringvalidator.OneOf("ingestion_time", "time_unit_column"),
				},
				MarkdownDescription: "The following partitioning types are supported: `ingestion_time`, `time_unit_column`. In the case of `ingestion_time`, partitions are cut based on TROCCO's job execution time. In the case of `time_unit_column`, partitioning is done based on the reference column. Available only in `insert` mode",
			},
			"partitioning_time": schema.StringAttribute{
				Optional: true,
				Validators: []validator.String{
					stringvalidator.OneOf("DAY", "HOUR", "MONTH", "YEAR"),
				},
				MarkdownDescription: "The granularity of table partitioning. The following units are supported: `DAY`, `HOUR`, `MONTH`, `YEAR`. Required when `partitioning` is set",
			},
			"partitioning_field": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Column name to be used for partitioning. Required when `partitioning` is `time_unit_column`",
			},
			"clustering_fields": schema.ListAttribute{
				Optional:    true,
				ElementType: types.StringType,
				Validators: []validator.List{
					listvalidator.SizeAtMost(4),
				},
				MarkdownDescription: "Column names to be used for clustering. At most 4 fields can be specified. Available only in `insert` mode",
			},
			"location": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "The location where the query will be executed. If not specified, the location is automatically determined by Google BigQuery. Available only in `query` mode",
			},
			"merge_keys": schema.ListAttribute{
				Optional:    true,
				ElementType: types.StringType,
				Validators: []validator.List{
					listvalidator.SizeAtLeast(1),
				},
				MarkdownDescription: "Key columns to uniquely identify records. Required when `write_disposition` is `incremental` or `scd_type_2`",
			},
			"on_matched_action": schema.StringAttribute{
				Optional: true,
				Validators: []validator.String{
					stringvalidator.OneOf("upsert", "skip"),
				},
				MarkdownDescription: "Behavior when a record with a matching key exists. The following actions are supported: `upsert`, `skip`. Required when `write_disposition` is `incremental`",
			},
			"incremental_column": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Incremental reference column. Required when `write_disposition` is `scd_type_2`",
			},
			"valid_from_column": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "SCD Type 2 valid-from column name. Fixed value: `trocco_valid_from`",
			},
			"valid_to_column": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "SCD Type 2 valid-to column name. Fixed value: `trocco_valid_to`",
			},
			"is_current_column": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "SCD Type 2 is-current flag column name. Fixed value: `trocco_is_current`",
			},
			"schema_evolution_mode": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Validators: []validator.String{
					stringvalidator.OneOf("detect_only", "auto_add_column"),
				},
				MarkdownDescription: "Schema evolution mode. The following modes are supported: `detect_only`, `auto_add_column`. Available when `write_disposition` is `incremental` or `scd_type_2`",
			},
			"lookback_period_column": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Column name for the lookback period. Available when `write_disposition` is `incremental` or `scd_type_2`",
			},
			"lookback_period_column_type": schema.StringAttribute{
				Optional: true,
				Validators: []validator.String{
					stringvalidator.OneOf("TIMESTAMP", "DATETIME", "DATE"),
				},
				MarkdownDescription: "Data type of the lookback period column. The following types are supported: `TIMESTAMP`, `DATETIME`, `DATE`",
			},
			"lookback_period_timezone": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Timezone for the lookback period",
			},
			"lookback_period_from": schema.Int64Attribute{
				Optional: true,
				Validators: []validator.Int64{
					int64validator.AtLeast(0),
				},
				MarkdownDescription: "Start value of the lookback period",
			},
			"lookback_period_to": schema.Int64Attribute{
				Optional: true,
				Validators: []validator.Int64{
					int64validator.AtLeast(0),
				},
				MarkdownDescription: "End value of the lookback period",
			},
			"lookback_period_unit": schema.StringAttribute{
				Optional: true,
				Validators: []validator.String{
					stringvalidator.OneOf("days", "hours"),
				},
				MarkdownDescription: "Unit of the lookback period. The following units are supported: `days`, `hours`",
			},
			"schedules": schema.SetNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"frequency": schema.StringAttribute{
							Required: true,
							Validators: []validator.String{
								stringvalidator.OneOf("hourly", "daily", "weekly", "monthly"),
							},
							MarkdownDescription: "Frequency of automatic execution. The following frequencies are supported: `hourly`, `daily`, `weekly`, `monthly`",
						},
						"minute": schema.Int64Attribute{
							Required: true,
							Validators: []validator.Int64{
								int64validator.Between(0, 59),
							},
							MarkdownDescription: "Value of minute. Required for all schedules",
						},
						"hour": schema.Int64Attribute{
							Optional: true,
							Validators: []validator.Int64{
								int64validator.Between(0, 23),
							},
							MarkdownDescription: "Value of hour. Required in `daily`, `weekly`, and `monthly` schedules",
						},
						"day_of_week": schema.Int64Attribute{
							Optional: true,
							Validators: []validator.Int64{
								int64validator.Between(0, 6),
							},
							MarkdownDescription: "Value of day of week. Sunday - Saturday is represented as 0 - 6. Required in `weekly` schedule",
						},
						"day": schema.Int64Attribute{
							Optional: true,
							Validators: []validator.Int64{
								int64validator.Between(1, 31),
							},
							MarkdownDescription: "Value of day. Required in `monthly` schedule",
						},
						"time_zone": schema.StringAttribute{
							Required:            true,
							MarkdownDescription: "Time zone to be used for calculation",
						},
					},
					PlanModifiers: []planmodifier.Object{
						&troccoPlanModifier.SchedulePlanModifier{},
					},
				},
				MarkdownDescription: "Schedules to be attached to the datamart definition",
			},
			"notifications": schema.ListNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.Int64Attribute{
							Computed:            true,
							MarkdownDescription: "Server-assigned ID of the notification. Unique within `(notification_type, destination_type)` for matching across API responses.",
						},
						"destination_type": schema.StringAttribute{
							Required: true,
							Validators: []validator.String{
								stringvalidator.OneOf("slack", "email"),
							},
							MarkdownDescription: "Destination service where the notification will be sent. The following types are supported: `slack`, `email`",
						},
						"slack_channel_id": schema.Int64Attribute{
							Optional: true,
							Validators: []validator.Int64{
								int64validator.AtLeast(1),
							},
							MarkdownDescription: "ID of the slack channel used to send notifications. Required when `destination_type` is `slack`",
						},
						"email_id": schema.Int64Attribute{
							Optional: true,
							Validators: []validator.Int64{
								int64validator.AtLeast(1),
							},
							MarkdownDescription: "ID of the email used to send notifications. Required when `destination_type` is `email`",
						},
						"notification_type": schema.StringAttribute{
							Required: true,
							Validators: []validator.String{
								stringvalidator.OneOf("job", "record"),
							},
							MarkdownDescription: "Category of condition. The following types are supported: `job`, `record`",
						},
						"notify_when": schema.StringAttribute{
							Optional: true,
							Validators: []validator.String{
								stringvalidator.OneOf("finished", "failed"),
							},
							MarkdownDescription: "Specifies the job status that trigger a notification. The following types are supported: `finished`, `failed`. Required when `notification_type` is `job`",
						},
						"record_count": schema.Int64Attribute{
							Optional:            true,
							MarkdownDescription: "The number of records to be used for condition. Required when `notification_type` is `record`",
						},
						"record_operator": schema.StringAttribute{
							Optional: true,
							Validators: []validator.String{
								stringvalidator.OneOf("above", "below"),
							},
							MarkdownDescription: "Operator to be used for condition. The following operators are supported: `above`, `below`. Required when `notification_type` is `record`",
						},
						"message": schema.StringAttribute{
							Required:            true,
							CustomType:          custom_type.TrimmedStringType{},
							MarkdownDescription: "The message to be sent with the notification",
						},
					},
					PlanModifiers: []planmodifier.Object{
						&troccoPlanModifier.DatamartNotificationPlanModifier{},
					},
				},
				MarkdownDescription: "Notifications to be attached to the datamart definition",
			},
			"labels": schema.SetNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.Int64Attribute{
							Computed:            true,
							MarkdownDescription: "The ID of the label",
						},
						"name": schema.StringAttribute{
							Required:            true,
							MarkdownDescription: "The name of the label",
						},
					},
				},
				MarkdownDescription: "Labels to be attached to the datamart definition",
			},
		},
	}
}

func (r *bigqueryDatamartDefinitionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan bigqueryDatamartDefinitionModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	input := client.CreateDatamartDefinitionInput{
		Name:                   plan.Name.ValueString(),
		DatawarehouseType:      "bigquery",
		IsRunnableConcurrently: plan.IsRunnableConcurrently.ValueBool(),
	}
	if !plan.Description.IsNull() {
		input.SetDescription(plan.Description.ValueString())
	}
	if !plan.ResourceGroupID.IsNull() {
		input.SetResourceGroupID(plan.ResourceGroupID.ValueInt64())
	}
	if customVariableSettingInputs := convertCustomVariableSettingsForCreate(ctx, plan.CustomVariableSettings, resp); customVariableSettingInputs != nil && !resp.Diagnostics.HasError() {
		input.SetCustomVariableSettings(customVariableSettingInputs)
	}
	if resp.Diagnostics.HasError() {
		return
	}
	if plan.QueryMode.ValueString() == "insert" {
		optionInput := client.NewInsertModeCreateDatamartBigqueryOptionInput(
			plan.BigqueryConnectionID.ValueInt64(),
			plan.Query.ValueString(),
			plan.DestinationDataset.ValueString(),
			plan.DestinationTable.ValueString(),
			plan.WriteDisposition.ValueString(),
		)
		if !plan.BeforeLoad.IsNull() {
			optionInput.SetBeforeLoad(plan.BeforeLoad.ValueString())
		}
		if !plan.Partitioning.IsNull() {
			optionInput.SetPartitioning(plan.Partitioning.ValueString())
		}
		if !plan.PartitioningTime.IsNull() {
			optionInput.SetPartitioningTime(plan.PartitioningTime.ValueString())
		}
		if !plan.PartitioningField.IsNull() {
			optionInput.SetPartitioningField(plan.PartitioningField.ValueString())
		}
		if clusteringFields := utils.ConvertStringList(ctx, plan.ClusteringFields); len(clusteringFields) > 0 {
			optionInput.SetClusteringFields(clusteringFields)
		}
		if mergeKeys := utils.ConvertStringList(ctx, plan.MergeKeys); len(mergeKeys) > 0 {
			optionInput.SetMergeKeys(mergeKeys)
		}
		if !plan.OnMatchedAction.IsNull() {
			optionInput.SetOnMatchedAction(plan.OnMatchedAction.ValueString())
		}
		if !plan.IncrementalColumn.IsNull() {
			optionInput.SetIncrementalColumn(plan.IncrementalColumn.ValueString())
		}
		if !plan.SchemaEvolutionMode.IsNull() && !plan.SchemaEvolutionMode.IsUnknown() {
			optionInput.SetSchemaEvolutionMode(plan.SchemaEvolutionMode.ValueString())
		}
		if !plan.LookbackPeriodColumn.IsNull() {
			optionInput.SetLookbackPeriodColumn(plan.LookbackPeriodColumn.ValueString())
		}
		if !plan.LookbackPeriodColumnType.IsNull() {
			optionInput.SetLookbackPeriodColumnType(plan.LookbackPeriodColumnType.ValueString())
		}
		if !plan.LookbackPeriodTimezone.IsNull() {
			optionInput.SetLookbackPeriodTimezone(plan.LookbackPeriodTimezone.ValueString())
		}
		if !plan.LookbackPeriodFrom.IsNull() {
			optionInput.SetLookbackPeriodFrom(plan.LookbackPeriodFrom.ValueInt64())
		}
		if !plan.LookbackPeriodTo.IsNull() {
			optionInput.SetLookbackPeriodTo(plan.LookbackPeriodTo.ValueInt64())
		}
		if !plan.LookbackPeriodUnit.IsNull() {
			optionInput.SetLookbackPeriodUnit(plan.LookbackPeriodUnit.ValueString())
		}
		if resp.Diagnostics.HasError() {
			return
		}
		input.SetDatamartBigqueryOption(optionInput)
	} else {
		optionInput := client.NewQueryModeCreateDatamartBigqueryOptionInput(
			plan.BigqueryConnectionID.ValueInt64(),
			plan.Query.ValueString(),
		)
		if !plan.Location.IsNull() {
			optionInput.SetLocation(plan.Location.ValueString())
		}
		input.SetDatamartBigqueryOption(optionInput)
	}
	if !plan.Schedules.IsNull() && !plan.Schedules.IsUnknown() {
		var scheduleValues []scheduleModel
		diags := plan.Schedules.ElementsAs(ctx, &scheduleValues, false)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

		scheduleInputs := make([]client.ScheduleInput, len(scheduleValues))
		for i, v := range scheduleValues {
			switch v.Frequency.ValueString() {
			case "hourly":
				{
					scheduleInputs[i] = client.NewHourlyScheduleInput(
						int(v.Minute.ValueInt64()),
						v.TimeZone.ValueString(),
					)
				}
			case "daily":
				{
					scheduleInputs[i] = client.NewDailyScheduleInput(
						int(v.Hour.ValueInt64()),
						int(v.Minute.ValueInt64()),
						v.TimeZone.ValueString(),
					)
				}
			case "weekly":
				{
					scheduleInputs[i] = client.NewWeeklyScheduleInput(
						int(v.DayOfWeek.ValueInt64()),
						int(v.Hour.ValueInt64()),
						int(v.Minute.ValueInt64()),
						v.TimeZone.ValueString(),
					)
				}
			case "monthly":
				{
					scheduleInputs[i] = client.NewMonthlyScheduleInput(
						int(v.Day.ValueInt64()),
						int(v.Hour.ValueInt64()),
						int(v.Minute.ValueInt64()),
						v.TimeZone.ValueString(),
					)
				}
			}
		}
		input.SetSchedules(scheduleInputs)
	}
	if !plan.Notifications.IsNull() && !plan.Notifications.IsUnknown() {
		var notificationValues []datamartNotificationModel
		diags := plan.Notifications.ElementsAs(ctx, &notificationValues, false)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

		notificationInputs := make([]client.DatamartNotificationInput, len(notificationValues))
		for i, v := range notificationValues {
			if v.DestinationType.ValueString() == "slack" {
				if v.NotificationType.ValueString() == "job" {
					notificationInputs[i] = client.NewSlackJobDatamartNotificationInput(
						v.SlackChannelID.ValueInt64(),
						v.NotifyWhen.ValueString(),
						v.Message.ValueString(),
					)
				} else {
					notificationInputs[i] = client.NewSlackRecordDatamartNotificationInput(
						v.SlackChannelID.ValueInt64(),
						v.RecordCount.ValueInt64(),
						v.RecordOperator.ValueString(),
						v.Message.ValueString(),
					)
				}
			} else {
				if v.NotificationType.ValueString() == "job" {
					notificationInputs[i] = client.NewEmailJobDatamartNotificationInput(
						v.EmailID.ValueInt64(),
						v.NotifyWhen.ValueString(),
						v.Message.ValueString(),
					)
				} else {
					notificationInputs[i] = client.NewEmailRecordDatamartNotificationInput(
						v.EmailID.ValueInt64(),
						v.RecordCount.ValueInt64(),
						v.RecordOperator.ValueString(),
						v.Message.ValueString(),
					)
				}
			}
		}
		input.SetNotifications(notificationInputs)
	}
	if labelInputs := convertLabelsForCreate(ctx, plan.Labels, resp); labelInputs != nil && !resp.Diagnostics.HasError() {
		input.SetLabels(labelInputs)
	}
	if resp.Diagnostics.HasError() {
		return
	}
	res, err := r.client.CreateDatamartDefinition(&input)
	if err != nil {
		resp.Diagnostics.AddError(
			"Creating datamart_definition",
			fmt.Sprintf("Unable to create datamart_definition, got error: %s", err),
		)
		return
	}

	var planNotifs []datamartNotificationModel
	if !plan.Notifications.IsNull() && !plan.Notifications.IsUnknown() {
		if refDiags := plan.Notifications.ElementsAs(ctx, &planNotifs, false); refDiags.HasError() {
			planNotifs = nil
		}
	}
	data, err := parseToBigqueryDatamartDefinitionModel(ctx, res.DatamartDefinition, planNotifs)
	if err != nil {
		resp.Diagnostics.AddError(
			"Reading datamart_definition",
			fmt.Sprintf("Unable to read datamart_definition (id: %d), got error: %s", res.ID, err),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, data)...)
}

func (r *bigqueryDatamartDefinitionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state bigqueryDatamartDefinitionModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	id := state.ID.ValueInt64()
	var stateNotifs []datamartNotificationModel
	if !state.Notifications.IsNull() && !state.Notifications.IsUnknown() {
		if refDiags := state.Notifications.ElementsAs(ctx, &stateNotifs, false); refDiags.HasError() {
			stateNotifs = nil
		}
	}
	data, err := r.fetchModel(ctx, id, stateNotifs)
	if err != nil {
		resp.Diagnostics.AddError(
			"Reading datamart_definition",
			fmt.Sprintf("Unable to read datamart_definition (id: %d), got error: %s", id, err),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, data)...)
}

func (r *bigqueryDatamartDefinitionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state bigqueryDatamartDefinitionModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	input := client.UpdateDatamartDefinitionInput{}
	input.SetName(plan.Name.ValueString())
	if !plan.Description.IsNull() {
		input.SetDescription(plan.Description.ValueString())
	} else {
		input.SetDescriptionEmpty()
	}
	input.SetIsRunnableConcurrently(plan.IsRunnableConcurrently.ValueBool())
	if !plan.ResourceGroupID.IsNull() {
		input.SetResourceGroupID(plan.ResourceGroupID.ValueInt64())
	} else {
		input.SetResourceGroupIDEmpty()
	}
	if !plan.CustomVariableSettings.IsNull() && !plan.CustomVariableSettings.IsUnknown() {
		var settings []customVariableSettingModel
		diags := plan.CustomVariableSettings.ElementsAs(ctx, &settings, false)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

		customVariableSettingInputs := make([]client.CustomVariableSettingInput, len(settings))
		for i, v := range settings {
			if v.Type.ValueString() == "string" {
				customVariableSettingInputs[i] = client.NewStringTypeCustomVariableSettingInput(
					v.Name.ValueString(),
					v.Value.ValueString(),
				)
			} else {
				customVariableSettingInputs[i] = client.NewTimestampTypeCustomVariableSettingInput(
					v.Name.ValueString(),
					v.Type.ValueString(),
					int(v.Quantity.ValueInt64()),
					v.Unit.ValueString(),
					v.Direction.ValueString(),
					v.Format.ValueString(),
					v.TimeZone.ValueString(),
				)
			}
		}
		input.SetCustomVariableSettings(customVariableSettingInputs)
	} else {
		input.SetCustomVariableSettings([]client.CustomVariableSettingInput{})
	}
	optionInput := client.UpdateDatamartBigqueryOptionInput{}
	optionInput.SetBigqueryConnectionID(plan.BigqueryConnectionID.ValueInt64())
	optionInput.SetQueryMode(plan.QueryMode.ValueString())
	optionInput.SetQuery(plan.Query.ValueString())
	if !plan.DestinationDataset.IsNull() {
		optionInput.SetDestinationDataset(plan.DestinationDataset.ValueString())
	}
	if !plan.DestinationTable.IsNull() {
		optionInput.SetDestinationTable(plan.DestinationTable.ValueString())
	}
	if !plan.WriteDisposition.IsNull() {
		optionInput.SetWriteDisposition(plan.WriteDisposition.ValueString())
	}
	if !plan.BeforeLoad.IsNull() {
		optionInput.SetBeforeLoad(plan.BeforeLoad.ValueString())
	} else {
		optionInput.SetBeforeLoadEmpty()
	}
	if !plan.Partitioning.IsNull() {
		optionInput.SetPartitioning(plan.Partitioning.ValueString())
	} else {
		optionInput.SetPartitioningEmpty()
	}
	if !plan.PartitioningTime.IsNull() {
		optionInput.SetPartitioningTime(plan.PartitioningTime.ValueString())
	}
	if !plan.PartitioningField.IsNull() {
		optionInput.SetPartitioningField(plan.PartitioningField.ValueString())
	}
	if !plan.ClusteringFields.IsNull() {
		var clusteringFieldsValues []types.String
		diags := plan.ClusteringFields.ElementsAs(ctx, &clusteringFieldsValues, false)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

		clusteringFields := make([]string, len(clusteringFieldsValues))
		for i, v := range clusteringFieldsValues {
			clusteringFields[i] = v.ValueString()
		}
		optionInput.SetClusteringFields(clusteringFields)
	} else {
		optionInput.SetClusteringFields([]string{})
	}
	if !plan.Location.IsNull() {
		optionInput.SetLocation(plan.Location.ValueString())
	} else {
		optionInput.SetLocationEmpty()
	}
	if !plan.MergeKeys.IsNull() {
		var mergeKeysValues []types.String
		diags := plan.MergeKeys.ElementsAs(ctx, &mergeKeysValues, false)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}
		mergeKeys := make([]string, len(mergeKeysValues))
		for i, v := range mergeKeysValues {
			mergeKeys[i] = v.ValueString()
		}
		optionInput.SetMergeKeys(mergeKeys)
	} else {
		optionInput.SetMergeKeys([]string{})
	}
	if !plan.OnMatchedAction.IsNull() {
		optionInput.SetOnMatchedAction(plan.OnMatchedAction.ValueString())
	} else {
		optionInput.SetOnMatchedActionEmpty()
	}
	if !plan.IncrementalColumn.IsNull() {
		optionInput.SetIncrementalColumn(plan.IncrementalColumn.ValueString())
	} else {
		optionInput.SetIncrementalColumnEmpty()
	}
	if !plan.SchemaEvolutionMode.IsNull() && !plan.SchemaEvolutionMode.IsUnknown() {
		optionInput.SetSchemaEvolutionMode(plan.SchemaEvolutionMode.ValueString())
	} else if plan.SchemaEvolutionMode.IsNull() {
		optionInput.SetSchemaEvolutionModeEmpty()
	}
	if !plan.LookbackPeriodColumn.IsNull() {
		optionInput.SetLookbackPeriodColumn(plan.LookbackPeriodColumn.ValueString())
	} else {
		optionInput.SetLookbackPeriodColumnEmpty()
	}
	if !plan.LookbackPeriodColumnType.IsNull() {
		optionInput.SetLookbackPeriodColumnType(plan.LookbackPeriodColumnType.ValueString())
	} else {
		optionInput.SetLookbackPeriodColumnTypeEmpty()
	}
	if !plan.LookbackPeriodTimezone.IsNull() {
		optionInput.SetLookbackPeriodTimezone(plan.LookbackPeriodTimezone.ValueString())
	} else {
		optionInput.SetLookbackPeriodTimezoneEmpty()
	}
	if !plan.LookbackPeriodFrom.IsNull() {
		optionInput.SetLookbackPeriodFrom(plan.LookbackPeriodFrom.ValueInt64())
	} else {
		optionInput.SetLookbackPeriodFromEmpty()
	}
	if !plan.LookbackPeriodTo.IsNull() {
		optionInput.SetLookbackPeriodTo(plan.LookbackPeriodTo.ValueInt64())
	} else {
		optionInput.SetLookbackPeriodToEmpty()
	}
	if !plan.LookbackPeriodUnit.IsNull() {
		optionInput.SetLookbackPeriodUnit(plan.LookbackPeriodUnit.ValueString())
	} else {
		optionInput.SetLookbackPeriodUnitEmpty()
	}
	input.SetDatamartBigqueryOption(optionInput)
	if !plan.Schedules.IsNull() && !plan.Schedules.IsUnknown() {
		var scheduleValues []scheduleModel
		diags := plan.Schedules.ElementsAs(ctx, &scheduleValues, false)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

		scheduleInputs := make([]client.ScheduleInput, len(scheduleValues))
		for i, v := range scheduleValues {
			switch v.Frequency.ValueString() {
			case "hourly":
				{
					scheduleInputs[i] = client.NewHourlyScheduleInput(
						int(v.Minute.ValueInt64()),
						v.TimeZone.ValueString(),
					)
				}
			case "daily":
				{
					scheduleInputs[i] = client.NewDailyScheduleInput(
						int(v.Hour.ValueInt64()),
						int(v.Minute.ValueInt64()),
						v.TimeZone.ValueString(),
					)
				}
			case "weekly":
				{
					scheduleInputs[i] = client.NewWeeklyScheduleInput(
						int(v.DayOfWeek.ValueInt64()),
						int(v.Hour.ValueInt64()),
						int(v.Minute.ValueInt64()),
						v.TimeZone.ValueString(),
					)
				}
			case "monthly":
				{
					scheduleInputs[i] = client.NewMonthlyScheduleInput(
						int(v.Day.ValueInt64()),
						int(v.Hour.ValueInt64()),
						int(v.Minute.ValueInt64()),
						v.TimeZone.ValueString(),
					)
				}
			}
		}
		input.SetSchedules(scheduleInputs)
	} else {
		input.SetSchedules([]client.ScheduleInput{})
	}
	if !plan.Notifications.IsNull() && !plan.Notifications.IsUnknown() {
		var notificationValues []datamartNotificationModel
		diags := plan.Notifications.ElementsAs(ctx, &notificationValues, false)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

		notificationInputs := make([]client.DatamartNotificationInput, len(notificationValues))
		for i, v := range notificationValues {
			if v.DestinationType.ValueString() == "slack" {
				if v.NotificationType.ValueString() == "job" {
					notificationInputs[i] = client.NewSlackJobDatamartNotificationInput(
						v.SlackChannelID.ValueInt64(),
						v.NotifyWhen.ValueString(),
						v.Message.ValueString(),
					)
				} else {
					notificationInputs[i] = client.NewSlackRecordDatamartNotificationInput(
						v.SlackChannelID.ValueInt64(),
						v.RecordCount.ValueInt64(),
						v.RecordOperator.ValueString(),
						v.Message.ValueString(),
					)
				}
			} else {
				if v.NotificationType.ValueString() == "job" {
					notificationInputs[i] = client.NewEmailJobDatamartNotificationInput(
						v.EmailID.ValueInt64(),
						v.NotifyWhen.ValueString(),
						v.Message.ValueString(),
					)
				} else {
					notificationInputs[i] = client.NewEmailRecordDatamartNotificationInput(
						v.EmailID.ValueInt64(),
						v.RecordCount.ValueInt64(),
						v.RecordOperator.ValueString(),
						v.Message.ValueString(),
					)
				}
			}
		}
		input.SetNotifications(notificationInputs)
	} else {
		input.SetNotifications([]client.DatamartNotificationInput{})
	}
	if !plan.Labels.IsNull() {
		var labelValues []labelModel
		diags := plan.Labels.ElementsAs(ctx, &labelValues, false)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

		labelInputs := make([]string, len(labelValues))
		for i, v := range labelValues {
			labelInputs[i] = v.Name.ValueString()
		}
		input.SetLabels(labelInputs)
	} else {
		input.SetLabels([]string{})
	}

	data, err := r.client.UpdateDatamartDefinition(state.ID.ValueInt64(), &input)

	if err != nil {
		resp.Diagnostics.AddError(
			"Updating datamart definition",
			fmt.Sprintf("Unable to update datamart definition (id: %d), got error: %s", state.ID.ValueInt64(), err),
		)
		return
	}
	var planNotifs []datamartNotificationModel
	if !plan.Notifications.IsNull() && !plan.Notifications.IsUnknown() {
		if refDiags := plan.Notifications.ElementsAs(ctx, &planNotifs, false); refDiags.HasError() {
			planNotifs = nil
		}
	}
	model, err := parseToBigqueryDatamartDefinitionModel(ctx, data.DatamartDefinition, planNotifs)
	if err != nil {
		resp.Diagnostics.AddError(
			"Parsing datamart definition",
			fmt.Sprintf("Unable to parse datamart definition (id: %d), got error: %s", state.ID.ValueInt64(), err),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, model)...)
}

func (r *bigqueryDatamartDefinitionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state bigqueryDatamartDefinitionModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	id := state.ID.ValueInt64()
	err := r.client.DeleteDatamartDefinition(id)
	if err != nil {
		resp.Diagnostics.AddError(
			"Deleting datamart_definition",
			fmt.Sprintf("Unable to delete datamart_definition (id: %d), got error: %s", id, err),
		)
		return
	}
}

func (r *bigqueryDatamartDefinitionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	id, err := strconv.ParseInt(req.ID, 10, 64)
	if err != nil {
		resp.Diagnostics.AddError(
			"Importing datamart_definition",
			fmt.Sprintf("Unable to parse id, got error: %s", err),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), id)...)
}

func (r bigqueryDatamartDefinitionResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data bigqueryDatamartDefinitionModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if !data.BeforeLoad.IsNull() {
		if data.WriteDisposition.ValueString() != "append" {
			resp.Diagnostics.AddAttributeError(
				path.Root("before_load"),
				"Invalid Before Load Query",
				"before_load is only available in insert query mode and write_disposition is append",
			)
		}
	}

	if data.QueryMode.ValueString() == "insert" {
		if data.DestinationDataset.IsNull() {
			resp.Diagnostics.AddAttributeError(
				path.Root("destination_dataset"),
				"Missing Destination Dataset",
				"destination_dataset is required for insert query mode",
			)
		}
		if data.DestinationTable.IsNull() {
			resp.Diagnostics.AddAttributeError(
				path.Root("destination_table"),
				"Missing Destination Table",
				"destination_table is required for insert query mode",
			)
		}
		if data.WriteDisposition.IsNull() {
			resp.Diagnostics.AddAttributeError(
				path.Root("write_disposition"),
				"Missing Write Disposition",
				"write_disposition is required for insert query mode",
			)
		}
		if !data.Partitioning.IsNull() {
			if data.PartitioningTime.IsNull() {
				resp.Diagnostics.AddAttributeError(
					path.Root("partitioning_time"),
					"Missing Partitioning Time",
					"partitioning_time is required when partitioning is set",
				)
			}
			if data.Partitioning.ValueString() == "time_unit_column" && data.PartitioningField.IsNull() {
				resp.Diagnostics.AddAttributeError(
					path.Root("partitioning_field"),
					"Missing Partitioning Field",
					"partitioning_field is required when partitioning is time_unit_column",
				)
			}
		}
	}

	writeDisposition := data.WriteDisposition.ValueString()

	// Validate incremental write disposition
	if writeDisposition == "incremental" {
		if data.MergeKeys.IsNull() {
			resp.Diagnostics.AddAttributeError(
				path.Root("merge_keys"),
				"Missing Merge Keys",
				"merge_keys is required when write_disposition is incremental",
			)
		}
		if data.OnMatchedAction.IsNull() {
			resp.Diagnostics.AddAttributeError(
				path.Root("on_matched_action"),
				"Missing On Matched Action",
				"on_matched_action is required when write_disposition is incremental",
			)
		}
	}

	// Validate scd_type_2 write disposition
	if writeDisposition == "scd_type_2" {
		if data.MergeKeys.IsNull() {
			resp.Diagnostics.AddAttributeError(
				path.Root("merge_keys"),
				"Missing Merge Keys",
				"merge_keys is required when write_disposition is scd_type_2",
			)
		}
		if data.IncrementalColumn.IsNull() {
			resp.Diagnostics.AddAttributeError(
				path.Root("incremental_column"),
				"Missing Incremental Column",
				"incremental_column is required when write_disposition is scd_type_2",
			)
		}
	}

	// Validate is_runnable_concurrently for incremental/scd_type_2
	if (writeDisposition == "incremental" || writeDisposition == "scd_type_2") &&
		!data.IsRunnableConcurrently.IsNull() && data.IsRunnableConcurrently.ValueBool() {
		resp.Diagnostics.AddAttributeError(
			path.Root("is_runnable_concurrently"),
			"Invalid Concurrent Execution Setting",
			"is_runnable_concurrently must be false when write_disposition is incremental or scd_type_2",
		)
	}

	// Validate lookback_period consistency for incremental/scd_type_2
	if writeDisposition == "incremental" || writeDisposition == "scd_type_2" {
		columnSet := !data.LookbackPeriodColumn.IsNull()
		fromSet := !data.LookbackPeriodFrom.IsNull()
		toSet := !data.LookbackPeriodTo.IsNull()
		fromOrToSet := fromSet || toSet

		if !columnSet && fromOrToSet {
			resp.Diagnostics.AddAttributeError(
				path.Root("lookback_period_column"),
				"Missing Lookback Period Column",
				"lookback_period_column is required when lookback_period_from or lookback_period_to is set",
			)
		}

		if columnSet {
			if data.LookbackPeriodColumnType.IsNull() {
				resp.Diagnostics.AddAttributeError(
					path.Root("lookback_period_column_type"),
					"Missing Lookback Period Column Type",
					"lookback_period_column_type is required when lookback_period_column is set",
				)
			}
			if !fromOrToSet {
				resp.Diagnostics.AddAttributeError(
					path.Root("lookback_period_from"),
					"Missing Lookback Period Range",
					"lookback_period_from or lookback_period_to is required when lookback_period_column is set",
				)
			}
		}

		if fromOrToSet {
			if data.LookbackPeriodUnit.IsNull() {
				resp.Diagnostics.AddAttributeError(
					path.Root("lookback_period_unit"),
					"Missing Lookback Period Unit",
					"lookback_period_unit is required when lookback_period_from or lookback_period_to is set",
				)
			}
			if data.LookbackPeriodTimezone.IsNull() {
				resp.Diagnostics.AddAttributeError(
					path.Root("lookback_period_timezone"),
					"Missing Lookback Period Timezone",
					"lookback_period_timezone is required when lookback_period_from or lookback_period_to is set",
				)
			}
		}

		// Validate lookback_period_from >= lookback_period_to
		if fromSet && toSet {
			if data.LookbackPeriodFrom.ValueInt64() < data.LookbackPeriodTo.ValueInt64() {
				resp.Diagnostics.AddAttributeError(
					path.Root("lookback_period_from"),
					"Invalid Lookback Period Range",
					"lookback_period_from must be greater than or equal to lookback_period_to",
				)
			}
		}
	}

}

func parseToBigqueryDatamartDefinitionModel(ctx context.Context, response client.DatamartDefinition, refNotifs []datamartNotificationModel) (*bigqueryDatamartDefinitionModel, error) {
	model := bigqueryDatamartDefinitionModel{
		ID:                     types.Int64Value(response.ID),
		Name:                   types.StringValue(response.Name),
		IsRunnableConcurrently: types.BoolValue(response.IsRunnableConcurrently),
	}
	if response.Description != nil {
		model.Description = types.StringValue(*response.Description)
	}
	if response.ResourceGroup != nil {
		model.ResourceGroupID = types.Int64Value(response.ResourceGroup.ID)
	}
	if response.CustomVariableSettings != nil {
		customVariableSettings := make([]customVariableSettingModel, len(response.CustomVariableSettings))
		for i, v := range response.CustomVariableSettings {
			customVariableSettings[i] = customVariableSettingModel{
				Name: types.StringValue(v.Name),
				Type: types.StringValue(v.Type),
			}
			if v.Value != nil {
				customVariableSettings[i].Value = types.StringValue(*v.Value)
			}
			if v.Quantity != nil {
				customVariableSettings[i].Quantity = types.Int64Value(int64(*v.Quantity))
			}
			if v.Unit != nil {
				customVariableSettings[i].Unit = types.StringValue(*v.Unit)
			}
			if v.Direction != nil {
				customVariableSettings[i].Direction = types.StringValue(*v.Direction)
			}
			if v.Format != nil {
				customVariableSettings[i].Format = types.StringValue(*v.Format)
			}
			if v.TimeZone != nil {
				customVariableSettings[i].TimeZone = types.StringValue(*v.TimeZone)
			}
		}

		objectType := types.ObjectType{
			AttrTypes: customVariableSettingModel{}.attrTypes(),
		}

		listValue, diags := types.ListValueFrom(ctx, objectType, customVariableSettings)
		if diags.HasError() {
			return nil, fmt.Errorf("failed to convert to ListValue")
		}
		model.CustomVariableSettings = listValue
	} else {
		model.CustomVariableSettings = types.ListNull(types.ObjectType{
			AttrTypes: customVariableSettingModel{}.attrTypes(),
		})
	}
	if response.DatamartBigqueryOption != nil {
		model.BigqueryConnectionID = types.Int64Value(response.DatamartBigqueryOption.BigqueryConnectionID)
		model.QueryMode = types.StringValue(response.DatamartBigqueryOption.QueryMode)
		model.Query = custom_type.TrimmedStringValue{StringValue: types.StringValue(response.DatamartBigqueryOption.Query)}
		if response.DatamartBigqueryOption.DestinationDataset != nil {
			model.DestinationDataset = types.StringValue(*response.DatamartBigqueryOption.DestinationDataset)
		}
		if response.DatamartBigqueryOption.DestinationTable != nil {
			model.DestinationTable = types.StringValue(*response.DatamartBigqueryOption.DestinationTable)
		}
		if response.DatamartBigqueryOption.WriteDisposition != nil {
			model.WriteDisposition = types.StringValue(*response.DatamartBigqueryOption.WriteDisposition)
		}
		if response.DatamartBigqueryOption.BeforeLoad != nil {
			model.BeforeLoad = custom_type.TrimmedStringValue{StringValue: types.StringValue(*response.DatamartBigqueryOption.BeforeLoad)}
		}
		if response.DatamartBigqueryOption.Partitioning != nil {
			model.Partitioning = types.StringValue(*response.DatamartBigqueryOption.Partitioning)
		}
		if response.DatamartBigqueryOption.PartitioningTime != nil {
			model.PartitioningTime = types.StringValue(*response.DatamartBigqueryOption.PartitioningTime)
		}
		if response.DatamartBigqueryOption.PartitioningField != nil {
			model.PartitioningField = types.StringValue(*response.DatamartBigqueryOption.PartitioningField)
		}
		if response.DatamartBigqueryOption.ClusteringFields != nil {
			clusteringFields := make([]types.String, len(response.DatamartBigqueryOption.ClusteringFields))
			for i, v := range response.DatamartBigqueryOption.ClusteringFields {
				clusteringFields[i] = types.StringValue(v)
			}

			listValue, diags := types.ListValueFrom(ctx, types.StringType, clusteringFields)
			if diags.HasError() {
				return nil, fmt.Errorf("failed to convert clustering fields to ListValue")
			}
			model.ClusteringFields = listValue
		} else {
			model.ClusteringFields = types.ListNull(types.StringType)
		}
		if response.DatamartBigqueryOption.Location != nil {
			model.Location = types.StringValue(*response.DatamartBigqueryOption.Location)
		}
		if response.DatamartBigqueryOption.MergeKeys != nil {
			mergeKeys := make([]types.String, len(response.DatamartBigqueryOption.MergeKeys))
			for i, v := range response.DatamartBigqueryOption.MergeKeys {
				mergeKeys[i] = types.StringValue(v)
			}
			listValue, diags := types.ListValueFrom(ctx, types.StringType, mergeKeys)
			if diags.HasError() {
				return nil, fmt.Errorf("failed to convert merge_keys to ListValue")
			}
			model.MergeKeys = listValue
		} else {
			model.MergeKeys = types.ListNull(types.StringType)
		}
		if response.DatamartBigqueryOption.OnMatchedAction != nil {
			model.OnMatchedAction = types.StringValue(*response.DatamartBigqueryOption.OnMatchedAction)
		}
		if response.DatamartBigqueryOption.IncrementalColumn != nil {
			model.IncrementalColumn = types.StringValue(*response.DatamartBigqueryOption.IncrementalColumn)
		}
		if response.DatamartBigqueryOption.ValidFromColumn != nil {
			model.ValidFromColumn = types.StringValue(*response.DatamartBigqueryOption.ValidFromColumn)
		}
		if response.DatamartBigqueryOption.ValidToColumn != nil {
			model.ValidToColumn = types.StringValue(*response.DatamartBigqueryOption.ValidToColumn)
		}
		if response.DatamartBigqueryOption.IsCurrentColumn != nil {
			model.IsCurrentColumn = types.StringValue(*response.DatamartBigqueryOption.IsCurrentColumn)
		}
		if response.DatamartBigqueryOption.SchemaEvolutionMode != nil {
			model.SchemaEvolutionMode = types.StringValue(*response.DatamartBigqueryOption.SchemaEvolutionMode)
		}
		if response.DatamartBigqueryOption.LookbackPeriodColumn != nil {
			model.LookbackPeriodColumn = types.StringValue(*response.DatamartBigqueryOption.LookbackPeriodColumn)
		}
		if response.DatamartBigqueryOption.LookbackPeriodColumnType != nil {
			model.LookbackPeriodColumnType = types.StringValue(*response.DatamartBigqueryOption.LookbackPeriodColumnType)
		}
		if response.DatamartBigqueryOption.LookbackPeriodTimezone != nil {
			model.LookbackPeriodTimezone = types.StringValue(*response.DatamartBigqueryOption.LookbackPeriodTimezone)
		}
		if response.DatamartBigqueryOption.LookbackPeriodFrom != nil {
			model.LookbackPeriodFrom = types.Int64Value(*response.DatamartBigqueryOption.LookbackPeriodFrom)
		}
		if response.DatamartBigqueryOption.LookbackPeriodTo != nil {
			model.LookbackPeriodTo = types.Int64Value(*response.DatamartBigqueryOption.LookbackPeriodTo)
		}
		if response.DatamartBigqueryOption.LookbackPeriodUnit != nil {
			model.LookbackPeriodUnit = types.StringValue(*response.DatamartBigqueryOption.LookbackPeriodUnit)
		}
	} else {
		return nil, fmt.Errorf("datamartBigqueryOption is nil")
	}
	if response.Notifications != nil {
		notifications := make([]datamartNotificationModel, len(response.Notifications))
		for i, v := range response.Notifications {
			notifications[i] = datamartNotificationModel{
				ID:               types.Int64Value(v.ID),
				DestinationType:  types.StringValue(v.DestinationType),
				NotificationType: types.StringValue(v.NotificationType),
				Message:          custom_type.TrimmedStringValue{StringValue: types.StringValue(v.Message)},
			}
			if v.SlackChannelID != nil {
				notifications[i].SlackChannelID = types.Int64Value(*v.SlackChannelID)
			}
			if v.EmailID != nil {
				notifications[i].EmailID = types.Int64Value(*v.EmailID)
			}
			if v.NotifyWhen != nil {
				notifications[i].NotifyWhen = types.StringValue(*v.NotifyWhen)
			}
			if v.RecordCount != nil {
				notifications[i].RecordCount = types.Int64Value(*v.RecordCount)
			}
			if v.RecordOperator != nil {
				notifications[i].RecordOperator = types.StringValue(*v.RecordOperator)
			}
		}

		notifications = utils.MatchByKey(notifications, refNotifs, datamartNotificationKey, datamartNotificationFallbackKey)

		objectType := types.ObjectType{
			AttrTypes: datamartNotificationModel{}.attrTypes(),
		}

		listValue, diags := types.ListValueFrom(ctx, objectType, notifications)
		if diags.HasError() {
			return nil, fmt.Errorf("failed to convert notifications to ListValue")
		}
		model.Notifications = listValue
	} else {
		model.Notifications = types.ListNull(types.ObjectType{
			AttrTypes: datamartNotificationModel{}.attrTypes(),
		})
	}
	if response.Schedules != nil {
		schedules := make([]scheduleModel, len(response.Schedules))
		for i, v := range response.Schedules {
			schedules[i] = scheduleModel{
				Frequency: types.StringValue(v.Frequency),
				Minute:    types.Int64Value(int64(v.Minute)),
				TimeZone:  types.StringValue(v.TimeZone),
			}
			if v.Hour != nil {
				schedules[i].Hour = types.Int64Value(int64(*v.Hour))
			}
			if v.DayOfWeek != nil {
				schedules[i].DayOfWeek = types.Int64Value(int64(*v.DayOfWeek))
			}
			if v.Day != nil {
				schedules[i].Day = types.Int64Value(int64(*v.Day))
			}
		}

		objectType := types.ObjectType{
			AttrTypes: scheduleModel{}.attrTypes(),
		}

		setValue, diags := types.SetValueFrom(ctx, objectType, schedules)
		if diags.HasError() {
			return nil, fmt.Errorf("failed to convert schedules to SetValue")
		}
		model.Schedules = setValue
	} else {
		model.Schedules = types.SetNull(types.ObjectType{
			AttrTypes: scheduleModel{}.attrTypes(),
		})
	}
	if response.Labels != nil {
		labels := make([]labelModel, len(response.Labels))
		for i, v := range response.Labels {
			labels[i] = labelModel{
				ID:   types.Int64Value(v.ID),
				Name: types.StringValue(v.Name),
			}
		}

		objectType := types.ObjectType{
			AttrTypes: labelModel{}.attrTypes(),
		}

		setValue, diags := types.SetValueFrom(ctx, objectType, labels)
		if diags.HasError() {
			return nil, fmt.Errorf("failed to convert labels to SetValue")
		}
		model.Labels = setValue
	} else {
		model.Labels = types.SetNull(types.ObjectType{
			AttrTypes: labelModel{}.attrTypes(),
		})
	}

	return &model, nil
}

func (r *bigqueryDatamartDefinitionResource) fetchModel(ctx context.Context, id int64, refNotifs []datamartNotificationModel) (*bigqueryDatamartDefinitionModel, error) {
	datamartDefinition, err := r.client.GetDatamartDefinition(id)
	if err != nil {
		return nil, err
	}
	model, _ := parseToBigqueryDatamartDefinitionModel(ctx, datamartDefinition.DatamartDefinition, refNotifs)
	return model, nil
}

func (c customVariableSettingModel) attrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"name":      types.StringType,
		"type":      types.StringType,
		"value":     types.StringType,
		"quantity":  types.Int64Type,
		"unit":      types.StringType,
		"direction": types.StringType,
		"format":    types.StringType,
		"time_zone": types.StringType,
	}
}

func (n datamartNotificationModel) attrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"id":                types.Int64Type,
		"destination_type":  types.StringType,
		"slack_channel_id":  types.Int64Type,
		"email_id":          types.Int64Type,
		"notification_type": types.StringType,
		"notify_when":       types.StringType,
		"record_count":      types.Int64Type,
		"record_operator":   types.StringType,
		"message":           types.StringType,
	}
}

func (s scheduleModel) attrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"frequency":   types.StringType,
		"minute":      types.Int32Type,
		"hour":        types.Int32Type,
		"day":         types.Int32Type,
		"day_of_week": types.Int32Type,
		"time_zone":   types.StringType,
	}
}

func (l labelModel) attrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"id":   types.Int64Type,
		"name": types.StringType,
	}
}

// Helper functions for BigQuery datamart definition to reduce code duplication

// convertCustomVariableSettingsForCreate converts custom variable settings for create operation.
func convertCustomVariableSettingsForCreate(ctx context.Context, source types.List, diags *resource.CreateResponse) []client.CustomVariableSettingInput {
	if source.IsNull() || source.IsUnknown() {
		return []client.CustomVariableSettingInput{}
	}

	var settings []customVariableSettingModel
	elemDiags := source.ElementsAs(ctx, &settings, false)
	diags.Diagnostics.Append(elemDiags...)
	if diags.Diagnostics.HasError() {
		return nil
	}

	result := make([]client.CustomVariableSettingInput, 0, len(settings))
	for _, v := range settings {
		if v.Type.ValueString() == "string" {
			result = append(result, client.NewStringTypeCustomVariableSettingInput(
				v.Name.ValueString(),
				v.Value.ValueString(),
			))
		} else {
			result = append(result, client.NewTimestampTypeCustomVariableSettingInput(
				v.Name.ValueString(),
				v.Type.ValueString(),
				int(v.Quantity.ValueInt64()),
				v.Unit.ValueString(),
				v.Direction.ValueString(),
				v.Format.ValueString(),
				v.TimeZone.ValueString(),
			))
		}
	}
	return result
}

// convertLabelsForCreate converts labels for create operation.
func convertLabelsForCreate(ctx context.Context, source types.Set, diags *resource.CreateResponse) []string {
	if source.IsNull() {
		return []string{}
	}

	var labelValues []labelModel
	elemDiags := source.ElementsAs(ctx, &labelValues, false)
	diags.Diagnostics.Append(elemDiags...)
	if diags.Diagnostics.HasError() {
		return nil
	}

	result := make([]string, 0, len(labelValues))
	for _, v := range labelValues {
		result = append(result, v.Name.ValueString())
	}
	return result
}

func datamartNotificationKey(n datamartNotificationModel) string {
	return fmt.Sprintf("%s|%s|%d",
		n.NotificationType.ValueString(),
		n.DestinationType.ValueString(),
		n.ID.ValueInt64(),
	)
}

func datamartNotificationFallbackKey(n datamartNotificationModel) string {
	return fmt.Sprintf("%s|%s",
		n.NotificationType.ValueString(),
		n.DestinationType.ValueString(),
	)
}
