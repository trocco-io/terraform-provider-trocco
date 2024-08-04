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
		MarkdownDescription: "The datamart definition resource allows you to create, read, update, and delete a datamart definition.",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Required: true,
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtMost(255),
				},
				MarkdownDescription: "It must be less than 256 characters",
			},
			"description": schema.StringAttribute{
				Optional: true,
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtLeast(1),
				},
				MarkdownDescription: "It must be at least 1 character",
			},
			"is_runnable_concurrently": schema.BoolAttribute{
				Required:            true,
				MarkdownDescription: "Whether or not to run a job if another job with the same data mart definition is running at the time the job is run.",
			},
			"resource_group_id": schema.Int64Attribute{
				Optional: true,
				Validators: []validator.Int64{
					int64validator.AtLeast(1),
				},
				MarkdownDescription: "Resource group ID to which the datamart definition belongs",
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
							MarkdownDescription: "It must start and end with `$`",
						},
						"type": schema.StringAttribute{
							Required: true,
							Validators: []validator.String{
								stringvalidator.OneOf("string", "timestamp", "timestamp_runtime"),
							},
							MarkdownDescription: "The following types are supported: `string`, `timestamp`, `timestamp_runtime`",
						},
						"value": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Required for `string` type.",
						},
						"quantity": schema.Int32Attribute{
							Optional: true,
							Validators: []validator.Int32{
								int32validator.AtLeast(0),
							},
							MarkdownDescription: "Required for `timestamp` and `timestamp_runtime` types.",
						},
						"unit": schema.StringAttribute{
							Optional: true,
							Validators: []validator.String{
								stringvalidator.OneOf("hour", "date", "month"),
							},
							MarkdownDescription: "The following units are supported: `hour`, `date`, `month`. Required for `timestamp` and `timestamp_runtime` types.",
						},
						"direction": schema.StringAttribute{
							Optional: true,
							Validators: []validator.String{
								stringvalidator.OneOf("ago", "later"),
							},
							MarkdownDescription: "The following directions are supported: `ago`, `later`. Required for `timestamp` and `timestamp_runtime` types.",
						},
						"format": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Required for `timestamp` and `timestamp_runtime` types.",
						},
						"time_zone": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Required for `timestamp` and `timestamp_runtime` types.",
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
			},
			"query_mode": schema.StringAttribute{
				Required: true,
				Validators: []validator.String{
					stringvalidator.OneOf("insert", "query"),
				},
				MarkdownDescription: "The following query modes are supported: `insert`, `query`",
			},
			"query": schema.StringAttribute{
				Required:   true,
				CustomType: trimmedStringType{},
			},
			"destination_dataset": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Required for `insert` mode",
			},
			"destination_table": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Required for `insert` mode",
			},
			"write_disposition": schema.StringAttribute{
				Optional: true,
				Validators: []validator.String{
					stringvalidator.OneOf("append", "truncate"),
				},
				MarkdownDescription: "The following write dispositions are supported: `append`, `truncate`. Required for `insert` mode",
			},
			"before_load": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Valid for `insert` mode",
			},
			"partitioning": schema.StringAttribute{
				Optional: true,
				Validators: []validator.String{
					stringvalidator.OneOf("ingestion_time", "time_unit_column"),
				},
				MarkdownDescription: "The following partitioning types are supported: `ingestion_time`, `time_unit_column`. Valid for `insert` mode",
			},
			"partitioning_time": schema.StringAttribute{
				Optional: true,
				Validators: []validator.String{
					stringvalidator.OneOf("DAY", "HOUR", "MONTH", "YEAR"),
				},
				MarkdownDescription: "The following partitioning time units are supported: `DAY`, `HOUR`, `MONTH`, `YEAR`. Valid for `insert` mode. Required when `partitioning` is set",
			},
			"partitioning_field": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Required when `partitioning` is `time_unit_column`",
			},
			"clustering_fields": schema.ListAttribute{
				Optional:    true,
				ElementType: types.StringType,
				Validators: []validator.List{
					listvalidator.SizeAtMost(4),
				},
				MarkdownDescription: "Valid for `insert` mode. At most 4 fields can be specified.",
			},
			"location": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Valid for `query` mode",
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
							MarkdownDescription: "The following frequencies are supported: `hourly`, `daily`, `weekly`, `monthly`",
						},
						"minute": schema.Int32Attribute{
							Required: true,
							Validators: []validator.Int32{
								int32validator.Between(0, 59),
							},
						},
						"hour": schema.Int32Attribute{
							Optional: true,
							Validators: []validator.Int32{
								int32validator.Between(0, 23),
							},
							MarkdownDescription: "Required for `daily`, `weekly`, and `monthly` schedules",
						},
						"day_of_week": schema.Int32Attribute{
							Optional: true,
							Validators: []validator.Int32{
								int32validator.Between(0, 6),
							},
							MarkdownDescription: "Sunday - Saturday is represented as 0 - 6. Required for `weekly` schedule",
						},
						"day": schema.Int32Attribute{
							Optional: true,
							Validators: []validator.Int32{
								int32validator.Between(1, 31),
							},
							MarkdownDescription: "Required for `monthly` schedule",
						},
						"time_zone": schema.StringAttribute{
							Required:            true,
							MarkdownDescription: "Time zone to calculate the schedule time",
						},
					},
					PlanModifiers: []planmodifier.Object{
						&schedulePlanModifier{},
					},
				},
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
							MarkdownDescription: "The following destination types are supported: `slack`, `email`",
						},
						"slack_channel_id": schema.Int64Attribute{
							Optional: true,
							Validators: []validator.Int64{
								int64validator.AtLeast(1),
							},
							MarkdownDescription: "Required when `destination_type` is `slack`",
						},
						"email_id": schema.Int64Attribute{
							Optional: true,
							Validators: []validator.Int64{
								int64validator.AtLeast(1),
							},
							MarkdownDescription: "Required when `destination_type` is `email`",
						},
						"notification_type": schema.StringAttribute{
							Required: true,
							Validators: []validator.String{
								stringvalidator.OneOf("job", "record"),
							},
							MarkdownDescription: "The following notification types are supported: `job`, `record`",
						},
						"notify_when": schema.StringAttribute{
							Optional: true,
							Validators: []validator.String{
								stringvalidator.OneOf("finished", "failed"),
							},
							MarkdownDescription: "The following notify when types are supported: `finished`, `failed`. Required for `job` notification type",
						},
						"record_count": schema.Int64Attribute{
							Optional:            true,
							MarkdownDescription: "Required for `record` notification type",
						},
						"record_operator": schema.StringAttribute{
							Optional: true,
							Validators: []validator.String{
								stringvalidator.OneOf("above", "below"),
							},
							MarkdownDescription: "The following record operators are supported: `above`, `below`. Required for `record` notification type",
						},
						"message": schema.StringAttribute{
							Required: true,
						},
					},
					PlanModifiers: []planmodifier.Object{
						&datamartNotificationPlanModifier{},
					},
				},
			},
			"labels": schema.SetNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.Int64Attribute{
							Computed: true,
						},
						"name": schema.StringAttribute{
							Required: true,
						},
					},
				},
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
