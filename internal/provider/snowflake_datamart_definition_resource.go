package provider

import (
	"context"
	"fmt"
	"strconv"

	"terraform-provider-trocco/internal/client"
	"terraform-provider-trocco/internal/provider/custom_type"
	troccoPlanModifier "terraform-provider-trocco/internal/provider/planmodifier"
	troccoValidator "terraform-provider-trocco/internal/provider/validator"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.Resource = &snowflakeDatamartDefinitionResource{}
var _ resource.ResourceWithImportState = &snowflakeDatamartDefinitionResource{}

func NewSnowflakeDatamartDefinitionResource() resource.Resource {
	return &snowflakeDatamartDefinitionResource{}
}

type snowflakeDatamartDefinitionResource struct {
	client *client.TroccoClient
}

type snowflakeDatamartDefinitionModel struct {
	ID                     types.Int64                    `tfsdk:"id"`
	Name                   types.String                   `tfsdk:"name"`
	Description            types.String                   `tfsdk:"description"`
	IsRunnableConcurrently types.Bool                     `tfsdk:"is_runnable_concurrently"`
	ResourceGroupID        types.Int64                    `tfsdk:"resource_group_id"`
	CustomVariableSettings types.List                     `tfsdk:"custom_variable_settings"`
	SnowflakeConnectionID  types.Int64                    `tfsdk:"snowflake_connection_id"`
	QueryMode              types.String                   `tfsdk:"query_mode"`
	Query                  custom_type.TrimmedStringValue `tfsdk:"query"`
	Warehouse              types.String                   `tfsdk:"warehouse"`
	StatementTimeout       types.Int64                    `tfsdk:"statement_timeout"`
	DestinationDatabase    types.String                   `tfsdk:"destination_database"`
	DestinationSchema      types.String                   `tfsdk:"destination_schema"`
	DestinationTable       types.String                   `tfsdk:"destination_table"`
	WriteDisposition       types.String                   `tfsdk:"write_disposition"`
	Notifications          types.Set                      `tfsdk:"notifications"`
	Schedules              types.Set                      `tfsdk:"schedules"`
	Labels                 types.Set                      `tfsdk:"labels"`
}

func (r *snowflakeDatamartDefinitionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_snowflake_datamart_definition"
}

func (r *snowflakeDatamartDefinitionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *snowflakeDatamartDefinitionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Provides a TROCCO datamart definitions for Snowflake resource.",
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
			"snowflake_connection_id": schema.Int64Attribute{
				Required: true,
				Validators: []validator.Int64{
					int64validator.AtLeast(1),
				},
				MarkdownDescription: "ID of the Snowflake connection which is used to communicate with Snowflake",
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
			"warehouse": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "Virtual warehouse to be used for query execution",
			},
			"statement_timeout": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				Validators: []validator.Int64{
					int64validator.AtLeast(0),
				},
				MarkdownDescription: "Query timeout in seconds. If 0 is specified, Snowflake's STATEMENT_TIMEOUT_IN_SECONDS is used",
			},
			"destination_database": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Destination database where the query result will be inserted. Required in `insert` mode",
			},
			"destination_schema": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Destination schema where the query result will be inserted. Required in `insert` mode",
			},
			"destination_table": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Destination table where the query result will be inserted. Required in `insert` mode",
			},
			"write_disposition": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Validators: []validator.String{
					stringvalidator.OneOf("append", "truncate", "replace"),
				},
				MarkdownDescription: "The following write dispositions are supported: `append`, `truncate`, `replace`. In the case of `append`, the result of the query execution is appended after the records of the existing table. In the case of `truncate`, records in the existing table are deleted and replaced with the results of the query execution. In the case of `replace`, the entire table is replaced. Required in `insert` mode",
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

func (r *snowflakeDatamartDefinitionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan snowflakeDatamartDefinitionModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	input := client.CreateDatamartDefinitionInput{
		Name:                   plan.Name.ValueString(),
		DatawarehouseType:      "snowflake",
		IsRunnableConcurrently: plan.IsRunnableConcurrently.ValueBool(),
	}
	if !plan.Description.IsNull() {
		input.SetDescription(plan.Description.ValueString())
	}
	if !plan.ResourceGroupID.IsNull() {
		input.SetResourceGroupID(plan.ResourceGroupID.ValueInt64())
	}
	if customVariableSettingInputs := convertSnowflakeCustomVariableSettingsForCreate(ctx, plan.CustomVariableSettings, resp); customVariableSettingInputs != nil && !resp.Diagnostics.HasError() {
		input.SetCustomVariableSettings(customVariableSettingInputs)
	}
	if resp.Diagnostics.HasError() {
		return
	}

	if plan.QueryMode.ValueString() == "insert" {
		optionInput := client.NewInsertModeCreateDatamartSnowflakeSettingInput(
			plan.SnowflakeConnectionID.ValueInt64(),
			plan.Query.ValueString(),
			plan.Warehouse.ValueString(),
			plan.DestinationDatabase.ValueString(),
			plan.DestinationSchema.ValueString(),
			plan.DestinationTable.ValueString(),
			plan.WriteDisposition.ValueString(),
		)
		if !plan.StatementTimeout.IsNull() {
			optionInput.SetStatementTimeout(plan.StatementTimeout.ValueInt64())
		}
		input.SetDatamartSnowflakeSetting(optionInput)
	} else {
		optionInput := client.NewQueryModeCreateDatamartSnowflakeSettingInput(
			plan.SnowflakeConnectionID.ValueInt64(),
			plan.Query.ValueString(),
			plan.Warehouse.ValueString(),
		)
		if !plan.StatementTimeout.IsNull() {
			optionInput.SetStatementTimeout(plan.StatementTimeout.ValueInt64())
		}
		input.SetDatamartSnowflakeSetting(optionInput)
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
				scheduleInputs[i] = client.NewHourlyScheduleInput(
					int(v.Minute.ValueInt64()),
					v.TimeZone.ValueString(),
				)
			case "daily":
				scheduleInputs[i] = client.NewDailyScheduleInput(
					int(v.Hour.ValueInt64()),
					int(v.Minute.ValueInt64()),
					v.TimeZone.ValueString(),
				)
			case "weekly":
				scheduleInputs[i] = client.NewWeeklyScheduleInput(
					int(v.DayOfWeek.ValueInt64()),
					int(v.Hour.ValueInt64()),
					int(v.Minute.ValueInt64()),
					v.TimeZone.ValueString(),
				)
			case "monthly":
				scheduleInputs[i] = client.NewMonthlyScheduleInput(
					int(v.Day.ValueInt64()),
					int(v.Hour.ValueInt64()),
					int(v.Minute.ValueInt64()),
					v.TimeZone.ValueString(),
				)
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

	if labelInputs := convertSnowflakeLabelsForCreate(ctx, plan.Labels, resp); labelInputs != nil && !resp.Diagnostics.HasError() {
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

	data, err := parseToSnowflakeDatamartDefinitionModel(ctx, res.DatamartDefinition)
	if err != nil {
		resp.Diagnostics.AddError(
			"Reading datamart_definition",
			fmt.Sprintf("Unable to read datamart_definition (id: %d), got error: %s", res.ID, err),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, data)...)
}

func (r *snowflakeDatamartDefinitionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state snowflakeDatamartDefinitionModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	id := state.ID.ValueInt64()
	data, err := r.fetchSnowflakeModel(ctx, id)
	if err != nil {
		resp.Diagnostics.AddError(
			"Reading datamart_definition",
			fmt.Sprintf("Unable to read datamart_definition (id: %d), got error: %s", id, err),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, data)...)
}

func (r *snowflakeDatamartDefinitionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state snowflakeDatamartDefinitionModel
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

	optionInput := client.UpdateDatamartSnowflakeSettingInput{}
	optionInput.SetSnowflakeConnectionID(plan.SnowflakeConnectionID.ValueInt64())
	optionInput.SetQueryMode(plan.QueryMode.ValueString())
	optionInput.SetQuery(plan.Query.ValueString())
	optionInput.SetWarehouse(plan.Warehouse.ValueString())
	if !plan.StatementTimeout.IsNull() {
		optionInput.SetStatementTimeout(plan.StatementTimeout.ValueInt64())
	} else {
		optionInput.SetStatementTimeoutEmpty()
	}
	if !plan.DestinationDatabase.IsNull() {
		optionInput.SetDestinationDatabase(plan.DestinationDatabase.ValueString())
	}
	if !plan.DestinationSchema.IsNull() {
		optionInput.SetDestinationSchema(plan.DestinationSchema.ValueString())
	}
	if !plan.DestinationTable.IsNull() {
		optionInput.SetDestinationTable(plan.DestinationTable.ValueString())
	}
	if !plan.WriteDisposition.IsNull() {
		optionInput.SetWriteDisposition(plan.WriteDisposition.ValueString())
	}
	input.SetDatamartSnowflakeSetting(optionInput)

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
				scheduleInputs[i] = client.NewHourlyScheduleInput(
					int(v.Minute.ValueInt64()),
					v.TimeZone.ValueString(),
				)
			case "daily":
				scheduleInputs[i] = client.NewDailyScheduleInput(
					int(v.Hour.ValueInt64()),
					int(v.Minute.ValueInt64()),
					v.TimeZone.ValueString(),
				)
			case "weekly":
				scheduleInputs[i] = client.NewWeeklyScheduleInput(
					int(v.DayOfWeek.ValueInt64()),
					int(v.Hour.ValueInt64()),
					int(v.Minute.ValueInt64()),
					v.TimeZone.ValueString(),
				)
			case "monthly":
				scheduleInputs[i] = client.NewMonthlyScheduleInput(
					int(v.Day.ValueInt64()),
					int(v.Hour.ValueInt64()),
					int(v.Minute.ValueInt64()),
					v.TimeZone.ValueString(),
				)
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

	model, err := parseToSnowflakeDatamartDefinitionModel(ctx, data.DatamartDefinition)
	if err != nil {
		resp.Diagnostics.AddError(
			"Parsing datamart definition",
			fmt.Sprintf("Unable to parse datamart definition (id: %d), got error: %s", state.ID.ValueInt64(), err),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, model)...)
}

func (r *snowflakeDatamartDefinitionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state snowflakeDatamartDefinitionModel
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

func (r *snowflakeDatamartDefinitionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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

func (r snowflakeDatamartDefinitionResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data snowflakeDatamartDefinitionModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if data.QueryMode.ValueString() == "insert" {
		if data.DestinationDatabase.IsNull() {
			resp.Diagnostics.AddAttributeError(
				path.Root("destination_database"),
				"Missing Destination Database",
				"destination_database is required for insert query mode",
			)
		}
		if data.DestinationSchema.IsNull() {
			resp.Diagnostics.AddAttributeError(
				path.Root("destination_schema"),
				"Missing Destination Schema",
				"destination_schema is required for insert query mode",
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
	}
}

func parseToSnowflakeDatamartDefinitionModel(ctx context.Context, response client.DatamartDefinition) (*snowflakeDatamartDefinitionModel, error) {
	model := snowflakeDatamartDefinitionModel{
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

	if response.DatamartSnowflakeSetting != nil {
		model.SnowflakeConnectionID = types.Int64Value(response.DatamartSnowflakeSetting.SnowflakeConnectionID)
		model.QueryMode = types.StringValue(response.DatamartSnowflakeSetting.QueryMode)
		model.Query = custom_type.TrimmedStringValue{StringValue: types.StringValue(response.DatamartSnowflakeSetting.Query)}
		model.Warehouse = types.StringValue(response.DatamartSnowflakeSetting.Warehouse)
		if response.DatamartSnowflakeSetting.StatementTimeout != nil {
			model.StatementTimeout = types.Int64Value(*response.DatamartSnowflakeSetting.StatementTimeout)
		}
		if response.DatamartSnowflakeSetting.DestinationDatabase != nil {
			model.DestinationDatabase = types.StringValue(*response.DatamartSnowflakeSetting.DestinationDatabase)
		}
		if response.DatamartSnowflakeSetting.DestinationSchema != nil {
			model.DestinationSchema = types.StringValue(*response.DatamartSnowflakeSetting.DestinationSchema)
		}
		if response.DatamartSnowflakeSetting.DestinationTable != nil {
			model.DestinationTable = types.StringValue(*response.DatamartSnowflakeSetting.DestinationTable)
		}
		if response.DatamartSnowflakeSetting.WriteDisposition != nil {
			model.WriteDisposition = types.StringValue(*response.DatamartSnowflakeSetting.WriteDisposition)
		}
	} else {
		return nil, fmt.Errorf("datamartSnowflakeSetting is nil")
	}

	if response.Notifications != nil {
		notifications := make([]datamartNotificationModel, len(response.Notifications))
		for i, v := range response.Notifications {
			notifications[i] = datamartNotificationModel{
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

		objectType := types.ObjectType{
			AttrTypes: datamartNotificationModel{}.attrTypes(),
		}

		setValue, diags := types.SetValueFrom(ctx, objectType, notifications)
		if diags.HasError() {
			return nil, fmt.Errorf("failed to convert notifications to SetValue")
		}
		model.Notifications = setValue
	} else {
		model.Notifications = types.SetNull(types.ObjectType{
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

func (r *snowflakeDatamartDefinitionResource) fetchSnowflakeModel(ctx context.Context, id int64) (*snowflakeDatamartDefinitionModel, error) {
	datamartDefinition, err := r.client.GetDatamartDefinition(id)
	if err != nil {
		return nil, err
	}
	model, _ := parseToSnowflakeDatamartDefinitionModel(ctx, datamartDefinition.DatamartDefinition)
	return model, nil
}

func convertSnowflakeCustomVariableSettingsForCreate(ctx context.Context, source types.List, diags *resource.CreateResponse) []client.CustomVariableSettingInput {
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

func convertSnowflakeLabelsForCreate(ctx context.Context, source types.Set, diags *resource.CreateResponse) []string {
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

