package provider

import (
	"context"
	"fmt"
	"strconv"

	"terraform-provider-trocco/internal/client"

	"github.com/hashicorp/terraform-plugin-framework-validators/int32validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.Resource = &bigqueryDatamartDefinitionResource{}
var _ resource.ResourceWithImportState = &bigqueryDatamartDefinitionResource{}

func newBigqueryDatamartDefinitionResource() resource.Resource {
	return &bigqueryDatamartDefinitionResource{}
}

type bigqueryDatamartDefinitionResource struct {
	client *client.TroccoClient
}

type bigqueryDatamartDefinitionModel struct {
	ID                     types.Int64                  `tfsdk:"id"`
	Name                   types.String                 `tfsdk:"name"`
	Description            types.String                 `tfsdk:"description"`
	IsRunnableConcurrently types.Bool                   `tfsdk:"is_runnable_concurrently"`
	ResourceGroupID        types.Int64                  `tfsdk:"resource_group_id"`
	CustomVariableSettings []customVariableSettingModel `tfsdk:"custom_variable_settings"`
	BigqueryConnectionID   types.Int64                  `tfsdk:"bigquery_connection_id"`
	QueryMode              types.String                 `tfsdk:"query_mode"`
	Query                  trimmedStringValue           `tfsdk:"query"`
	DestinationDataset     types.String                 `tfsdk:"destination_dataset"`
	DestinationTable       types.String                 `tfsdk:"destination_table"`
	WriteDisposition       types.String                 `tfsdk:"write_disposition"`
	BeforeLoad             types.String                 `tfsdk:"before_load"`
	Partitioning           types.String                 `tfsdk:"partitioning"`
	PartitioningTime       types.String                 `tfsdk:"partitioning_time"`
	PartitioningField      types.String                 `tfsdk:"partitioning_field"`
	ClusteringFields       []types.String               `tfsdk:"clustering_fields"`
	Location               types.String                 `tfsdk:"location"`
	Notifications          []datamartNotificationModel  `tfsdk:"notifications"`
	Schedules              []scheduleModel              `tfsdk:"schedules"`
	Labels                 []labelModel                 `tfsdk:"labels"`
}

type customVariableSettingModel struct {
	Name      types.String `tfsdk:"name"`
	Type      types.String `tfsdk:"type"`
	Value     types.String `tfsdk:"value"`
	Quantity  types.Int32  `tfsdk:"quantity"`
	Unit      types.String `tfsdk:"unit"`
	Direction types.String `tfsdk:"direction"`
	Format    types.String `tfsdk:"format"`
	TimeZone  types.String `tfsdk:"time_zone"`
}

type datamartNotificationModel struct {
	DestinationType  types.String `tfsdk:"destination_type"`
	SlackChannelID   types.Int64  `tfsdk:"slack_channel_id"`
	EmailID          types.Int64  `tfsdk:"email_id"`
	NotificationType types.String `tfsdk:"notification_type"`
	NotifyWhen       types.String `tfsdk:"notify_when"`
	RecordCount      types.Int64  `tfsdk:"record_count"`
	RecordOperator   types.String `tfsdk:"record_operator"`
	Message          types.String `tfsdk:"message"`
}

type scheduleModel struct {
	Frequency types.String `tfsdk:"frequency"`
	Minute    types.Int32  `tfsdk:"minute"`
	Hour      types.Int32  `tfsdk:"hour"`
	Day       types.Int32  `tfsdk:"day"`
	DayOfWeek types.Int32  `tfsdk:"day_of_week"`
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
								wrappingDollarValidator{},
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
						"quantity": schema.Int32Attribute{
							Optional: true,
							Validators: []validator.Int32{
								int32validator.AtLeast(0),
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
						&customVariableSettingPlanModifier{},
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
				CustomType:          trimmedStringType{},
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
					stringvalidator.OneOf("append", "truncate"),
				},
				MarkdownDescription: "The following write dispositions are supported: `append`, `truncate`. In the case of `append`, the result of the query execution is appended after the records of the existing table. In the case of `truncate`, records in the existing table are deleted and replaced with the results of the query execution. Required in `insert` mode",
			},
			"before_load": schema.StringAttribute{
				Optional:            true,
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
						"minute": schema.Int32Attribute{
							Required: true,
							Validators: []validator.Int32{
								int32validator.Between(0, 59),
							},
							MarkdownDescription: "Value of minute. Required for all schedules",
						},
						"hour": schema.Int32Attribute{
							Optional: true,
							Validators: []validator.Int32{
								int32validator.Between(0, 23),
							},
							MarkdownDescription: "Value of hour. Required in `daily`, `weekly`, and `monthly` schedules",
						},
						"day_of_week": schema.Int32Attribute{
							Optional: true,
							Validators: []validator.Int32{
								int32validator.Between(0, 6),
							},
							MarkdownDescription: "Value of day of week. Sunday - Saturday is represented as 0 - 6. Required in `weekly` schedule",
						},
						"day": schema.Int32Attribute{
							Optional: true,
							Validators: []validator.Int32{
								int32validator.Between(1, 31),
							},
							MarkdownDescription: "Value of day. Required in `monthly` schedule",
						},
						"time_zone": schema.StringAttribute{
							Required:            true,
							MarkdownDescription: "Time zone to be used for calculation",
						},
					},
					PlanModifiers: []planmodifier.Object{
						&schedulePlanModifier{},
					},
				},
				MarkdownDescription: "Schedules to be attached to the datamart definition",
			},
			"notifications": schema.SetNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
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
							MarkdownDescription: "The message to be sent with the notification",
						},
					},
					PlanModifiers: []planmodifier.Object{
						&datamartNotificationPlanModifier{},
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
	if plan.CustomVariableSettings != nil {
		customVariableSettingInputs := make([]client.CustomVariableSettingInput, len(plan.CustomVariableSettings))
		for i, v := range plan.CustomVariableSettings {
			if v.Type.ValueString() == "string" {
				customVariableSettingInputs[i] = client.NewStringTypeCustomVariableSettingInput(
					v.Name.ValueString(),
					v.Value.ValueString(),
				)
			} else {
				customVariableSettingInputs[i] = client.NewTimestampTypeCustomVariableSettingInput(
					v.Name.ValueString(),
					v.Type.ValueString(),
					int(v.Quantity.ValueInt32()),
					v.Unit.ValueString(),
					v.Direction.ValueString(),
					v.Format.ValueString(),
					v.TimeZone.ValueString(),
				)
			}
		}
		input.SetCustomVariableSettings(customVariableSettingInputs)
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
		if plan.ClusteringFields != nil {
			clusteringFields := make([]string, len(plan.ClusteringFields))
			for i, v := range plan.ClusteringFields {
				clusteringFields[i] = v.ValueString()
			}
			optionInput.SetClusteringFields(clusteringFields)
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
	res, err := r.client.CreateDatamartDefinition(&input)
	if err != nil {
		resp.Diagnostics.AddError(
			"Creating datamart_definition",
			fmt.Sprintf("Unable to create datamart_definition, got error: %s", err),
		)
		return
	}

	updateInput := client.UpdateDatamartDefinitionInput{}
	needUpdate := false
	if plan.Schedules != nil {
		needUpdate = true
		scheduleInputs := make([]client.ScheduleInput, len(plan.Schedules))
		for i, v := range plan.Schedules {
			switch v.Frequency.ValueString() {
			case "hourly":
				{
					scheduleInputs[i] = client.NewHourlyScheduleInput(
						int(v.Minute.ValueInt32()),
						v.TimeZone.ValueString(),
					)
				}
			case "daily":
				{
					scheduleInputs[i] = client.NewDailyScheduleInput(
						int(v.Hour.ValueInt32()),
						int(v.Minute.ValueInt32()),
						v.TimeZone.ValueString(),
					)
				}
			case "weekly":
				{
					scheduleInputs[i] = client.NewWeeklyScheduleInput(
						int(v.DayOfWeek.ValueInt32()),
						int(v.Hour.ValueInt32()),
						int(v.Minute.ValueInt32()),
						v.TimeZone.ValueString(),
					)
				}
			case "monthly":
				{
					scheduleInputs[i] = client.NewMonthlyScheduleInput(
						int(v.Day.ValueInt32()),
						int(v.Hour.ValueInt32()),
						int(v.Minute.ValueInt32()),
						v.TimeZone.ValueString(),
					)
				}
			}
		}
		updateInput.SetSchedules(scheduleInputs)
	}
	if plan.Notifications != nil {
		needUpdate = true
		notificationInputs := make([]client.DatamartNotificationInput, len(plan.Notifications))
		for i, v := range plan.Notifications {
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
		updateInput.SetNotifications(notificationInputs)
	}
	if plan.Labels != nil {
		needUpdate = true
		labelInputs := make([]string, len(plan.Labels))
		for i, v := range plan.Labels {
			labelInputs[i] = v.Name.ValueString()
		}
		updateInput.SetLabels(labelInputs)
	}
	if needUpdate {
		err = r.client.UpdateDatamartDefinition(res.ID, &updateInput)
		if err != nil {
			resp.Diagnostics.AddError(
				"Creating datamart_definition",
				fmt.Sprintf("Unable to attach schedules/notifications/labels to datamart_definition (id: %d), got error: %s", res.ID, err),
			)
			return
		}
	}

	data, err := r.fetchModel(res.ID)
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
	data, err := r.fetchModel(id)
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
	if plan.CustomVariableSettings != nil {
		customVariableSettingInputs := make([]client.CustomVariableSettingInput, len(plan.CustomVariableSettings))
		for i, v := range plan.CustomVariableSettings {
			if v.Type.ValueString() == "string" {
				customVariableSettingInputs[i] = client.NewStringTypeCustomVariableSettingInput(
					v.Name.ValueString(),
					v.Value.ValueString(),
				)
			} else {
				customVariableSettingInputs[i] = client.NewTimestampTypeCustomVariableSettingInput(
					v.Name.ValueString(),
					v.Type.ValueString(),
					int(v.Quantity.ValueInt32()),
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
	if plan.ClusteringFields != nil {
		clusteringFields := make([]string, len(plan.ClusteringFields))
		for i, v := range plan.ClusteringFields {
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
	input.SetDatamartBigqueryOption(optionInput)
	if plan.Schedules != nil {
		scheduleInputs := make([]client.ScheduleInput, len(plan.Schedules))
		for i, v := range plan.Schedules {
			switch v.Frequency.ValueString() {
			case "hourly":
				{
					scheduleInputs[i] = client.NewHourlyScheduleInput(
						int(v.Minute.ValueInt32()),
						v.TimeZone.ValueString(),
					)
				}
			case "daily":
				{
					scheduleInputs[i] = client.NewDailyScheduleInput(
						int(v.Hour.ValueInt32()),
						int(v.Minute.ValueInt32()),
						v.TimeZone.ValueString(),
					)
				}
			case "weekly":
				{
					scheduleInputs[i] = client.NewWeeklyScheduleInput(
						int(v.DayOfWeek.ValueInt32()),
						int(v.Hour.ValueInt32()),
						int(v.Minute.ValueInt32()),
						v.TimeZone.ValueString(),
					)
				}
			case "monthly":
				{
					scheduleInputs[i] = client.NewMonthlyScheduleInput(
						int(v.Day.ValueInt32()),
						int(v.Hour.ValueInt32()),
						int(v.Minute.ValueInt32()),
						v.TimeZone.ValueString(),
					)
				}
			}
		}
		input.SetSchedules(scheduleInputs)
	} else {
		input.SetSchedules([]client.ScheduleInput{})
	}
	if plan.Notifications != nil {
		notificationInputs := make([]client.DatamartNotificationInput, len(plan.Notifications))
		for i, v := range plan.Notifications {
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
	if plan.Labels != nil {
		labelInputs := make([]string, len(plan.Labels))
		for i, v := range plan.Labels {
			labelInputs[i] = v.Name.ValueString()
		}
		input.SetLabels(labelInputs)
	} else {
		input.SetLabels([]string{})
	}

	err := r.client.UpdateDatamartDefinition(state.ID.ValueInt64(), &input)
	if err != nil {
		resp.Diagnostics.AddError(
			"Updating datamart definition",
			fmt.Sprintf("Unable to update datamart definition (id: %d), got error: %s", state.ID.ValueInt64(), err),
		)
		return
	}

	data, err := r.fetchModel(state.ID.ValueInt64())
	if err != nil {
		resp.Diagnostics.AddError(
			"Reading datamart_definition",
			fmt.Sprintf("failed to get datamart definition: %v", err),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, data)...)
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
}

func (r *bigqueryDatamartDefinitionResource) fetchModel(id int64) (*bigqueryDatamartDefinitionModel, error) {
	datamartDefinition, err := r.client.GetDatamartDefinition(id)
	if err != nil {
		return nil, err
	}
	model := bigqueryDatamartDefinitionModel{
		ID:                     types.Int64Value(datamartDefinition.ID),
		Name:                   types.StringValue(datamartDefinition.Name),
		IsRunnableConcurrently: types.BoolValue(datamartDefinition.IsRunnableConcurrently),
	}
	if datamartDefinition.Description != nil {
		model.Description = types.StringValue(*datamartDefinition.Description)
	}
	if datamartDefinition.ResourceGroup != nil {
		model.ResourceGroupID = types.Int64Value(datamartDefinition.ResourceGroup.ID)
	}
	if datamartDefinition.CustomVariableSettings != nil {
		customVariableSettings := make([]customVariableSettingModel, len(datamartDefinition.CustomVariableSettings))
		for i, v := range datamartDefinition.CustomVariableSettings {
			customVariableSettings[i] = customVariableSettingModel{
				Name: types.StringValue(v.Name),
				Type: types.StringValue(v.Type),
			}
			if v.Value != nil {
				customVariableSettings[i].Value = types.StringValue(*v.Value)
			}
			if v.Quantity != nil {
				customVariableSettings[i].Quantity = types.Int32Value(int32(*v.Quantity))
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
		model.CustomVariableSettings = customVariableSettings
	}
	if datamartDefinition.DatamartBigqueryOption != nil {
		model.BigqueryConnectionID = types.Int64Value(datamartDefinition.DatamartBigqueryOption.BigqueryConnectionID)
		model.QueryMode = types.StringValue(datamartDefinition.DatamartBigqueryOption.QueryMode)
		model.Query = trimmedStringValue{types.StringValue(datamartDefinition.DatamartBigqueryOption.Query)}
		if datamartDefinition.DatamartBigqueryOption.DestinationDataset != nil {
			model.DestinationDataset = types.StringValue(*datamartDefinition.DatamartBigqueryOption.DestinationDataset)
		}
		if datamartDefinition.DatamartBigqueryOption.DestinationTable != nil {
			model.DestinationTable = types.StringValue(*datamartDefinition.DatamartBigqueryOption.DestinationTable)
		}
		if datamartDefinition.DatamartBigqueryOption.WriteDisposition != nil {
			model.WriteDisposition = types.StringValue(*datamartDefinition.DatamartBigqueryOption.WriteDisposition)
		}
		if datamartDefinition.DatamartBigqueryOption.BeforeLoad != nil {
			model.BeforeLoad = types.StringValue(*datamartDefinition.DatamartBigqueryOption.BeforeLoad)
		}
		if datamartDefinition.DatamartBigqueryOption.Partitioning != nil {
			model.Partitioning = types.StringValue(*datamartDefinition.DatamartBigqueryOption.Partitioning)
		}
		if datamartDefinition.DatamartBigqueryOption.PartitioningTime != nil {
			model.PartitioningTime = types.StringValue(*datamartDefinition.DatamartBigqueryOption.PartitioningTime)
		}
		if datamartDefinition.DatamartBigqueryOption.PartitioningField != nil {
			model.PartitioningField = types.StringValue(*datamartDefinition.DatamartBigqueryOption.PartitioningField)
		}
		if datamartDefinition.DatamartBigqueryOption.ClusteringFields != nil {
			clusteringFields := make([]types.String, len(datamartDefinition.DatamartBigqueryOption.ClusteringFields))
			for i, v := range datamartDefinition.DatamartBigqueryOption.ClusteringFields {
				clusteringFields[i] = types.StringValue(v)
			}
			model.ClusteringFields = clusteringFields
		}
		if datamartDefinition.DatamartBigqueryOption.Location != nil {
			model.Location = types.StringValue(*datamartDefinition.DatamartBigqueryOption.Location)
		}
	} else {
		return nil, fmt.Errorf("datamartBigqueryOption is nil")
	}
	if datamartDefinition.Notifications != nil {
		notifications := make([]datamartNotificationModel, len(datamartDefinition.Notifications))
		for i, v := range datamartDefinition.Notifications {
			notifications[i] = datamartNotificationModel{
				DestinationType:  types.StringValue(v.DestinationType),
				NotificationType: types.StringValue(v.NotificationType),
				Message:          types.StringValue(v.Message),
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
		model.Notifications = notifications
	}
	if datamartDefinition.Schedules != nil {
		schedules := make([]scheduleModel, len(datamartDefinition.Schedules))
		for i, v := range datamartDefinition.Schedules {
			schedules[i] = scheduleModel{
				Frequency: types.StringValue(v.Frequency),
				Minute:    types.Int32Value(int32(v.Minute)),
				TimeZone:  types.StringValue(v.TimeZone),
			}
			if v.Hour != nil {
				schedules[i].Hour = types.Int32Value(int32(*v.Hour))
			}
			if v.DayOfWeek != nil {
				schedules[i].DayOfWeek = types.Int32Value(int32(*v.DayOfWeek))
			}
			if v.Day != nil {
				schedules[i].Day = types.Int32Value(int32(*v.Day))
			}
		}
		model.Schedules = schedules
	}
	if datamartDefinition.Labels != nil {
		labels := make([]labelModel, len(datamartDefinition.Labels))
		for i, v := range datamartDefinition.Labels {
			labels[i] = labelModel{
				ID:   types.Int64Value(v.ID),
				Name: types.StringValue(v.Name),
			}
		}
		model.Labels = labels
	}

	return &model, nil
}