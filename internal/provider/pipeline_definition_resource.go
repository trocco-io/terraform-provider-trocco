package provider

import (
	"context"
	"fmt"
	"strconv"
	"terraform-provider-trocco/internal/client"
	wp "terraform-provider-trocco/internal/client/parameters/pipeline_definition"
	wm "terraform-provider-trocco/internal/provider/models/pipeline_definition"
	ws "terraform-provider-trocco/internal/provider/schemas/pipeline_definition"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/samber/lo"
)

var (
	_ resource.Resource                = &workflowResource{}
	_ resource.ResourceWithConfigure   = &workflowResource{}
	_ resource.ResourceWithImportState = &workflowResource{}
)

// -----------------------------------------------------------------------------
// Provider-side Data Types
// -----------------------------------------------------------------------------

type pipelineDefinitionResourceModel struct {
	ID               types.Int64         `tfsdk:"id"`
	Name             types.String        `tfsdk:"name"`
	Description      types.String        `tfsdk:"description"`
	Labels           []types.String      `tfsdk:"labels"`
	Notifications    []wm.Notification   `tfsdk:"notifications"`
	Schedules        []wm.Schedule       `tfsdk:"schedules"`
	Tasks            []wm.Task           `tfsdk:"tasks"`
	TaskDependencies []wm.TaskDependency `tfsdk:"task_dependencies"`
}

func (m *pipelineDefinitionResourceModel) ToCreateWorkflowInput() *client.CreateWorkflowInput {
	labels := []string{}
	for _, l := range m.Labels {
		labels = append(labels, l.ValueString())
	}

	notifications := []wp.Notification{}
	for _, n := range m.Notifications {
		notifications = append(notifications, n.ToInput())
	}

	schedules := []wp.Schedule{}
	for _, s := range m.Schedules {
		schedules = append(schedules, s.ToInput())
	}

	tasks := []client.WorkflowTaskInput{}
	for _, r := range m.Tasks {
		i := client.WorkflowTaskInput{
			Key:            r.Key.ValueString(),
			TaskIdentifier: r.TaskIdentifier.ValueInt64(),
			Type:           r.Type.ValueString(),
		}

		if r.TroccoTransferConfig != nil {
			i.TroccoTransferConfig = r.TroccoTransferConfig.ToInput()
		}
		if r.TroccoTransferBulkConfig != nil {
			i.TroccoTransferBulkConfig = r.TroccoTransferBulkConfig.ToInput()
		}
		if r.DBTConfig != nil {
			i.DBTConfig = r.DBTConfig.ToInput()
		}
		if r.TroccoAgentConfig != nil {
			i.TroccoAgentConfig = r.TroccoAgentConfig.ToInput()
		}
		if r.TroccoBigQueryDatamartConfig != nil {
			i.TroccoBigQueryDatamartConfig = r.TroccoBigQueryDatamartConfig.ToInput()
		}
		if r.TroccoRedshiftDatamartConfig != nil {
			i.TroccoRedshiftDatamartConfig = r.TroccoRedshiftDatamartConfig.ToInput()
		}
		if r.TroccoSnowflakeDatamartConfig != nil {
			i.TroccoSnowflakeDatamartConfig = r.TroccoSnowflakeDatamartConfig.ToInput()
		}
		if r.WorkflowConfig != nil {
			i.WorkflowConfig = r.WorkflowConfig.ToInput()
		}
		if r.SlackNotificationConfig != nil {
			i.SlackNotificationConfig = r.SlackNotificationConfig.ToInput()
		}
		if r.TableauDataExtractionConfig != nil {
			i.TableauDataExtractionConfig = r.TableauDataExtractionConfig.ToInput()
		}
		if r.BigqueryDataCheckConfig != nil {
			i.BigqueryDataCheckConfig = r.BigqueryDataCheckConfig.ToInput()
		}
		if r.SnowflakeDataCheckConfig != nil {
			i.SnowflakeDataCheckConfig = r.SnowflakeDataCheckConfig.ToInput()
		}
		if r.RedshiftDataCheckConfig != nil {
			i.RedshiftDataCheckConfig = r.RedshiftDataCheckConfig.ToInput()
		}
		if r.HTTPRequestConfig != nil {
			i.HTTPRequestConfig = r.HTTPRequestConfig.ToInput()
		}

		tasks = append(tasks, i)
	}

	taskDependencies := []wp.TaskDependency{}
	for _, d := range m.TaskDependencies {
		taskDependencies = append(taskDependencies, wp.TaskDependency{
			Source:      d.Source.ValueString(),
			Destination: d.Destination.ValueString(),
		})
	}

	return &client.CreateWorkflowInput{
		Name:             m.Name.ValueString(),
		Description:      m.Description.ValueStringPointer(),
		Labels:           lo.ToPtr(labels),
		Notifications:    lo.ToPtr(notifications),
		Schedules:        lo.ToPtr(schedules),
		Tasks:            tasks,
		TaskDependencies: taskDependencies,
	}
}

func (m *pipelineDefinitionResourceModel) ToUpdateWorkflowInput(state *pipelineDefinitionResourceModel) *client.UpdateWorkflowInput {
	labels := []string{}
	for _, l := range m.Labels {
		labels = append(labels, l.ValueString())
	}

	notifications := []wp.Notification{}
	for _, n := range m.Notifications {
		notifications = append(notifications, n.ToInput())
	}

	schedules := []wp.Schedule{}
	for _, s := range m.Schedules {
		schedules = append(schedules, s.ToInput())
	}

	stateTaskIdentifiers := map[string]int64{}
	for _, s := range state.Tasks {
		stateTaskIdentifiers[s.Key.ValueString()] = s.TaskIdentifier.ValueInt64()
	}

	tasks := []client.WorkflowTaskInput{}
	for _, t := range m.Tasks {
		identifier := stateTaskIdentifiers[t.Key.ValueString()]

		i := client.WorkflowTaskInput{
			Key:            t.Key.ValueString(),
			TaskIdentifier: identifier,
			Type:           t.Type.ValueString(),
		}

		if t.TroccoTransferConfig != nil {
			i.TroccoTransferConfig = t.TroccoTransferConfig.ToInput()
		}
		if t.TroccoTransferBulkConfig != nil {
			i.TroccoTransferBulkConfig = t.TroccoTransferBulkConfig.ToInput()
		}
		if t.DBTConfig != nil {
			i.DBTConfig = t.DBTConfig.ToInput()
		}
		if t.TroccoAgentConfig != nil {
			i.TroccoAgentConfig = t.TroccoAgentConfig.ToInput()
		}
		if t.TroccoBigQueryDatamartConfig != nil {
			i.TroccoBigQueryDatamartConfig = t.TroccoBigQueryDatamartConfig.ToInput()
		}
		if t.TroccoRedshiftDatamartConfig != nil {
			i.TroccoRedshiftDatamartConfig = t.TroccoRedshiftDatamartConfig.ToInput()
		}
		if t.TroccoSnowflakeDatamartConfig != nil {
			i.TroccoSnowflakeDatamartConfig = t.TroccoSnowflakeDatamartConfig.ToInput()
		}
		if t.WorkflowConfig != nil {
			i.WorkflowConfig = t.WorkflowConfig.ToInput()
		}
		if t.SlackNotificationConfig != nil {
			i.SlackNotificationConfig = t.SlackNotificationConfig.ToInput()
		}
		if t.TableauDataExtractionConfig != nil {
			i.TableauDataExtractionConfig = t.TableauDataExtractionConfig.ToInput()
		}
		if t.BigqueryDataCheckConfig != nil {
			i.BigqueryDataCheckConfig = t.BigqueryDataCheckConfig.ToInput()
		}
		if t.SnowflakeDataCheckConfig != nil {
			i.SnowflakeDataCheckConfig = t.SnowflakeDataCheckConfig.ToInput()
		}
		if t.RedshiftDataCheckConfig != nil {
			i.RedshiftDataCheckConfig = t.RedshiftDataCheckConfig.ToInput()
		}
		if t.HTTPRequestConfig != nil {
			i.HTTPRequestConfig = t.HTTPRequestConfig.ToInput()
		}

		tasks = append(tasks, i)
	}

	taskDependencies := []wp.TaskDependency{}
	for _, d := range m.TaskDependencies {
		taskDependencies = append(taskDependencies, wp.TaskDependency{
			Source:      d.Source.ValueString(),
			Destination: d.Destination.ValueString(),
		})
	}

	return &client.UpdateWorkflowInput{
		Name:             m.Name.ValueStringPointer(),
		Description:      m.Description.ValueStringPointer(),
		Labels:           lo.ToPtr(labels),
		Notifications:    lo.ToPtr(notifications),
		Schedules:        lo.ToPtr(schedules),
		Tasks:            tasks,
		TaskDependencies: taskDependencies,
	}
}

// -----------------------------------------------------------------------------
// Resource
// -----------------------------------------------------------------------------

type workflowResource struct {
	client *client.TroccoClient
}

func NewWorkflowResource() resource.Resource {
	return &workflowResource{}
}

func (r *workflowResource) Metadata(
	ctx context.Context,
	req resource.MetadataRequest,
	resp *resource.MetadataResponse,
) {
	resp.TypeName = fmt.Sprintf("%s_workflow", req.ProviderTypeName)
}

func (r *workflowResource) Configure(
	ctx context.Context,
	req resource.ConfigureRequest,
	resp *resource.ConfigureResponse,
) {
	if req.ProviderData == nil {
		return
	}

	c, ok := req.ProviderData.(*client.TroccoClient)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *client.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}

	r.client = c
}

func (r *workflowResource) Schema(
	ctx context.Context,
	req resource.SchemaRequest,
	resp *resource.SchemaResponse,
) {
	customVariables := schema.SetNestedAttribute{
		Optional: true,
		NestedObject: schema.NestedAttributeObject{
			Attributes: map[string]schema.Attribute{
				"name": schema.StringAttribute{
					Required: true,
				},
				"type": schema.StringAttribute{
					Required: true,
				},
				"value": schema.StringAttribute{
					Optional: true,
				},
				"quantity": schema.Int64Attribute{
					Optional: true,
				},
				"unit": schema.StringAttribute{
					Optional: true,
				},
				"direction": schema.StringAttribute{
					Optional: true,
				},
				"format": schema.StringAttribute{
					Optional: true,
				},
				"time_zone": schema.StringAttribute{
					Optional: true,
				},
			},
		},
	}

	resp.Schema = schema.Schema{
		MarkdownDescription: "Provides a TROCCO workflow resource.",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				MarkdownDescription: "The ID of the workflow",
				Computed:            true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
				Validators: []validator.Int64{
					int64validator.AtLeast(1),
				},
			},
			"name": schema.StringAttribute{
				Required: true,
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtMost(255),
				},
			},
			"description": schema.StringAttribute{
				Optional: true,
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtLeast(1),
				},
			},
			"labels": schema.SetAttribute{
				Optional:    true,
				ElementType: types.StringType,
			},
			"notifications": schema.ListNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							Optional: true,
						},
						"email_config": schema.SingleNestedAttribute{
							Optional: true,
							Attributes: map[string]schema.Attribute{
								"notification_id": schema.Int64Attribute{
									Required: true,
								},
								"notify_when": schema.StringAttribute{
									Required: true,
								},
								"message": schema.StringAttribute{
									Required: true,
								},
							},
						},
						"slack_config": schema.SingleNestedAttribute{
							Optional: true,
							Attributes: map[string]schema.Attribute{
								"notification_id": schema.Int64Attribute{
									Required: true,
								},
								"notify_when": schema.StringAttribute{
									Required: true,
								},
								"message": schema.StringAttribute{
									Required: true,
								},
							},
						},
					},
				},
			},
			"schedules": schema.SetNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							Optional: true,
						},
						"daily_config": schema.SingleNestedAttribute{
							Optional: true,
							Attributes: map[string]schema.Attribute{
								"time_zone": schema.StringAttribute{
									Required: true,
								},
								"hour": schema.Int64Attribute{
									Required: true,
								},
								"minute": schema.Int64Attribute{
									Required: true,
								},
							},
						},
						"hourly_config": schema.SingleNestedAttribute{
							Optional: true,
							Attributes: map[string]schema.Attribute{
								"time_zone": schema.StringAttribute{
									Required: true,
								},
								"minute": schema.Int64Attribute{
									Required: true,
								},
							},
						},
						"monthly_config": schema.SingleNestedAttribute{
							Optional: true,
							Attributes: map[string]schema.Attribute{
								"time_zone": schema.StringAttribute{
									Required: true,
								},
								"day": schema.Int64Attribute{
									Required: true,
								},
								"hour": schema.Int64Attribute{
									Required: true,
								},
								"minute": schema.Int64Attribute{
									Required: true,
								},
							},
						},
						"weekly_config": schema.SingleNestedAttribute{
							Optional: true,
							Attributes: map[string]schema.Attribute{
								"time_zone": schema.StringAttribute{
									Required: true,
								},
								"day_of_week": schema.Int64Attribute{
									Required: true,
								},
								"hour": schema.Int64Attribute{
									Required: true,
								},
								"minute": schema.Int64Attribute{
									Required: true,
								},
							},
						},
					},
				},
			},
			"tasks": schema.SetNestedAttribute{
				MarkdownDescription: "The tasks of the workflow.",
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"key": schema.StringAttribute{
							Required: true,
						},
						"task_identifier": schema.Int64Attribute{
							Optional: true,
							Computed: true,
						},
						"type": schema.StringAttribute{
							Required: true,
							Validators: []validator.String{
								stringvalidator.OneOf(
									"trocco_transfer",
									"trocco_transfer_bulk",
									"trocco_bigquery_datamart",
									"trocco_redshift_datamart",
									"trocco_snowflake_datamart",
									"dbt",
									"trocco_agent",
									"trocco_pipeline",
									"slack_notify",
									"tableau_extract",
									"bigquery_data_check",
									"snowflake_data_check",
									"redshift_data_check",
									"http_request",
								),
							},
						},
						"trocco_transfer_config":           ws.NewTroccoTransferTaskConfigAttribute(),
						"trocco_transfer_bulk_config":      ws.NewTroccoTransferBulkTaskConfigAttribute(),
						"dbt_config":                       ws.NewDBTTaskConfigAttribute(),
						"trocco_agent_config":              ws.NewTroccoAgentTaskConfigAttribute(),
						"trocco_bigquery_datamart_config":  ws.NewBigQueryDatamartTaskConfigAttribute(),
						"trocco_redshift_datamart_config":  ws.NewRedshiftDatamartTaskConfigAttribute(),
						"trocco_snowflake_datamart_config": ws.NewSnowflakeDatamartTaskConfigAttribute(),
						"workflow_config": schema.SingleNestedAttribute{
							Optional: true,
							Attributes: map[string]schema.Attribute{
								"definition_id": schema.Int64Attribute{
									Required: true,
								},
								"custom_variable_loop": ws.NewCustomVariableLoopAttribute(),
							},
						},
						"slack_notification_config": schema.SingleNestedAttribute{
							Optional: true,
							Attributes: map[string]schema.Attribute{
								"name": schema.StringAttribute{
									Required: true,
								},
								"connection_id": schema.Int64Attribute{
									Required: true,
								},
								"message": schema.StringAttribute{
									Required: true,
								},
							},
						},
						"tableau_data_extraction_config": schema.SingleNestedAttribute{
							Optional: true,
							Attributes: map[string]schema.Attribute{
								"name": schema.StringAttribute{
									Required: true,
								},
								"connection_id": schema.Int64Attribute{
									Required: true,
								},
								"task_id": schema.StringAttribute{
									Required: true,
								},
							},
						},
						"http_request_config": schema.SingleNestedAttribute{
							Optional: true,
							Attributes: map[string]schema.Attribute{
								"name": schema.StringAttribute{
									Required: true,
								},
								"connection_id": schema.Int64Attribute{
									Optional: true,
								},
								"http_method": schema.StringAttribute{
									Required: true,
								},
								"url": schema.StringAttribute{
									Required: true,
								},
								"request_body": schema.StringAttribute{
									Optional: true,
								},
								"request_headers": schema.SetNestedAttribute{
									Optional: true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"key": schema.StringAttribute{
												Required: true,
											},
											"value": schema.StringAttribute{
												Required: true,
											},
											"masking": schema.BoolAttribute{
												Optional: true,
											},
										},
									},
								},
								"request_parameters": schema.SetNestedAttribute{
									Optional: true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"key": schema.StringAttribute{
												Required: true,
											},
											"value": schema.StringAttribute{
												Required: true,
											},
											"masking": schema.BoolAttribute{
												Optional: true,
											},
										},
									},
								},
								"custom_variables": customVariables,
							},
						},
						"bigquery_data_check_config": schema.SingleNestedAttribute{
							Optional: true,
							Attributes: map[string]schema.Attribute{
								"name": schema.StringAttribute{
									Required: true,
								},
								"connection_id": schema.Int64Attribute{
									Required: true,
								},
								"query": schema.StringAttribute{
									Optional: true,
								},
								"operator": schema.StringAttribute{
									Optional: true,
								},
								"query_result": schema.Int64Attribute{
									Optional: true,
								},
								"accepts_null": schema.BoolAttribute{
									Optional: true,
								},
								"custom_variables": customVariables,
							},
						},
						"snowflake_data_check_config": schema.SingleNestedAttribute{
							Optional: true,
							Attributes: map[string]schema.Attribute{
								"name": schema.StringAttribute{
									Required: true,
								},
								"connection_id": schema.Int64Attribute{
									Required: true,
								},
								"query": schema.StringAttribute{
									Optional: true,
								},
								"operator": schema.StringAttribute{
									Optional: true,
								},
								"query_result": schema.Int64Attribute{
									Optional: true,
								},
								"accepts_null": schema.BoolAttribute{
									Optional: true,
								},
								"warehouse": schema.StringAttribute{
									Optional: true,
								},
								"custom_variables": customVariables,
							},
						},
						"redshift_data_check_config": schema.SingleNestedAttribute{
							Optional: true,
							Attributes: map[string]schema.Attribute{
								"name": schema.StringAttribute{
									Required: true,
								},
								"connection_id": schema.Int64Attribute{
									Required: true,
								},
								"query": schema.StringAttribute{
									Optional: true,
								},
								"operator": schema.StringAttribute{
									Optional: true,
								},
								"query_result": schema.Int64Attribute{
									Optional: true,
								},
								"accepts_null": schema.BoolAttribute{
									Optional: true,
								},
								"database": schema.StringAttribute{
									Optional: true,
								},
								"custom_variables": customVariables,
							},
						},
					},
				},
			},
			"task_dependencies": schema.SetNestedAttribute{
				MarkdownDescription: "The task dependencies of the workflow.",
				Required:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"source": schema.StringAttribute{
							Required: true,
						},
						"destination": schema.StringAttribute{
							Required: true,
						},
					},
				},
			},
		},
	}
}

func (r *workflowResource) Create(
	ctx context.Context,
	req resource.CreateRequest,
	resp *resource.CreateResponse,
) {
	plan := &pipelineDefinitionResourceModel{}
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	workflow, err := r.client.CreateWorkflow(
		plan.ToCreateWorkflowInput(),
	)
	if err != nil {
		resp.Diagnostics.AddError(
			"Creating workflow",
			fmt.Sprintf("Unable to create workflow, got error: %s", err),
		)
		return
	}

	// tasks := wm.NewTasks(workflow.Tasks)

	tasks := []wm.Task{}
	for _, t := range workflow.Tasks {
		tasks = append(tasks, wm.Task{
			Key:            types.StringValue(t.Key),
			TaskIdentifier: types.Int64Value(t.TaskIdentifier),
			Type:           types.StringValue(t.Type),

			TroccoTransferConfig:          wm.NewTroccoTransferTaskConfig(t.TroccoTransferConfig),
			TroccoTransferBulkConfig:      wm.NewTroccoTransferBulkTaskConfig(t.TroccoTransferBulkConfig),
			DBTConfig:                     wm.NewDBTTaskConfig(t.DBTConfig),
			TroccoAgentConfig:             wm.NewTroccoAgentTaskConfig(t.TroccoAgentConfig),
			TroccoBigQueryDatamartConfig:  wm.NewTroccoBigQueryDatamartTaskConfig(t.TroccoBigQueryDatamartConfig),
			TroccoRedshiftDatamartConfig:  wm.NewTroccoRedshiftDatamartTaskConfig(t.TroccoRedshiftDatamartConfig),
			TroccoSnowflakeDatamartConfig: wm.NewTroccoSnowflakeDatamartTaskConfig(t.TroccoSnowflakeDatamartConfig),
			WorkflowConfig:                wm.NewWorkflowTaskConfig(t.WorkflowConfig),
			SlackNotificationConfig:       wm.NewSlackNotificationTaskConfig(t.SlackNotificationConfig),
			TableauDataExtractionConfig:   wm.NewTableauDataExtractionTaskConfig(t.TableauDataExtractionConfig),
			BigqueryDataCheckConfig:       wm.NewBigqueryDataCheckTaskConfig(t.BigqueryDataCheckConfig),
			SnowflakeDataCheckConfig:      wm.NewSnowflakeDataCheckTaskConfig(t.SnowflakeDataCheckConfig),
			RedshiftDataCheckConfig:       wm.NewRedshiftDataCheckTaskConfig(t.RedshiftDataCheckConfig),
			HTTPRequestConfig:             wm.NewHTTPRequestTaskConfig(t.HTTPRequestConfig),
		})
	}

	keys := map[int64]types.String{}
	for _, t := range workflow.Tasks {
		keys[t.TaskIdentifier] = types.StringValue(t.Key)
	}

	taskDependencies := []wm.TaskDependency{}
	for _, d := range workflow.TaskDependencies {
		taskDependencies = append(taskDependencies, wm.TaskDependency{
			Source:      keys[d.Source],
			Destination: keys[d.Destination],
		})
	}

	newState := pipelineDefinitionResourceModel{
		ID:               types.Int64Value(workflow.ID),
		Name:             types.StringPointerValue(workflow.Name),
		Description:      types.StringPointerValue(workflow.Description),
		Labels:           wm.NewLabels(workflow.Labels, plan.Labels == nil),
		Notifications:    wm.NewNotifications(workflow.Notifications, plan.Notifications == nil),
		Schedules:        wm.NewSchedules(workflow.Schedules, plan.Schedules == nil),
		Tasks:            tasks,
		TaskDependencies: taskDependencies,
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}

func (r *workflowResource) Update(
	ctx context.Context,
	req resource.UpdateRequest,
	resp *resource.UpdateResponse,
) {
	state := &pipelineDefinitionResourceModel{}
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	plan := &pipelineDefinitionResourceModel{}
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	workflow, err := r.client.UpdateWorkflow(
		state.ID.ValueInt64(),
		plan.ToUpdateWorkflowInput(state),
	)
	if err != nil {
		resp.Diagnostics.AddError(
			"Updating workflow",
			fmt.Sprintf("Unable to update workflow, got error: %s", err),
		)
		return
	}

	tasks := []wm.Task{}
	for _, t := range workflow.Tasks {
		task := wm.Task{
			Key:            types.StringValue(t.Key),
			TaskIdentifier: types.Int64Value(t.TaskIdentifier),
			Type:           types.StringValue(t.Type),

			TroccoTransferConfig:          wm.NewTroccoTransferTaskConfig(t.TroccoTransferConfig),
			TroccoTransferBulkConfig:      wm.NewTroccoTransferBulkTaskConfig(t.TroccoTransferBulkConfig),
			DBTConfig:                     wm.NewDBTTaskConfig(t.DBTConfig),
			TroccoAgentConfig:             wm.NewTroccoAgentTaskConfig(t.TroccoAgentConfig),
			TroccoBigQueryDatamartConfig:  wm.NewTroccoBigQueryDatamartTaskConfig(t.TroccoBigQueryDatamartConfig),
			TroccoRedshiftDatamartConfig:  wm.NewTroccoRedshiftDatamartTaskConfig(t.TroccoRedshiftDatamartConfig),
			TroccoSnowflakeDatamartConfig: wm.NewTroccoSnowflakeDatamartTaskConfig(t.TroccoSnowflakeDatamartConfig),
			WorkflowConfig:                wm.NewWorkflowTaskConfig(t.WorkflowConfig),
			SlackNotificationConfig:       wm.NewSlackNotificationTaskConfig(t.SlackNotificationConfig),
			TableauDataExtractionConfig:   wm.NewTableauDataExtractionTaskConfig(t.TableauDataExtractionConfig),
			BigqueryDataCheckConfig:       wm.NewBigqueryDataCheckTaskConfig(t.BigqueryDataCheckConfig),
			SnowflakeDataCheckConfig:      wm.NewSnowflakeDataCheckTaskConfig(t.SnowflakeDataCheckConfig),
			RedshiftDataCheckConfig:       wm.NewRedshiftDataCheckTaskConfig(t.RedshiftDataCheckConfig),
			HTTPRequestConfig:             wm.NewHTTPRequestTaskConfig(t.HTTPRequestConfig),
		}

		tasks = append(tasks, task)
	}

	keys := map[int64]types.String{}
	for _, t := range workflow.Tasks {
		keys[t.TaskIdentifier] = types.StringValue(t.Key)
	}

	taskDependencies := []wm.TaskDependency{}
	for _, d := range workflow.TaskDependencies {
		taskDependencies = append(taskDependencies, wm.TaskDependency{
			Source:      keys[d.Source],
			Destination: keys[d.Destination],
		})
	}

	newState := pipelineDefinitionResourceModel{
		ID:               types.Int64Value(workflow.ID),
		Name:             types.StringPointerValue(workflow.Name),
		Description:      types.StringPointerValue(workflow.Description),
		Labels:           wm.NewLabels(workflow.Labels, plan.Labels == nil),
		Notifications:    wm.NewNotifications(workflow.Notifications, plan.Notifications == nil),
		Schedules:        wm.NewSchedules(workflow.Schedules, plan.Schedules == nil),
		Tasks:            tasks,
		TaskDependencies: taskDependencies,
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}

func (r *workflowResource) Read(
	ctx context.Context,
	req resource.ReadRequest,
	resp *resource.ReadResponse,
) {
	state := &pipelineDefinitionResourceModel{}
	resp.Diagnostics.Append(req.State.Get(ctx, state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	workflow, err := r.client.GetWorkflow(
		state.ID.ValueInt64(),
	)
	if err != nil {
		resp.Diagnostics.AddError(
			"Reading pipeline definition",
			fmt.Sprintf("Unable to read pipeline definition, got error: %s", err),
		)
		return
	}

	stateKeys := map[int64]string{}
	for _, s := range state.Tasks {
		stateKeys[s.TaskIdentifier.ValueInt64()] = s.Key.ValueString()
	}

	tasks := []wm.Task{}
	for _, t := range workflow.Tasks {
		key := stateKeys[t.TaskIdentifier]

		task := wm.Task{
			Key:            types.StringValue(key),
			TaskIdentifier: types.Int64Value(t.TaskIdentifier),
			Type:           types.StringValue(t.Type),

			BigqueryDataCheckConfig:       wm.NewBigqueryDataCheckTaskConfig(t.BigqueryDataCheckConfig),
			DBTConfig:                     wm.NewDBTTaskConfig(t.DBTConfig),
			HTTPRequestConfig:             wm.NewHTTPRequestTaskConfig(t.HTTPRequestConfig),
			RedshiftDataCheckConfig:       wm.NewRedshiftDataCheckTaskConfig(t.RedshiftDataCheckConfig),
			SlackNotificationConfig:       wm.NewSlackNotificationTaskConfig(t.SlackNotificationConfig),
			SnowflakeDataCheckConfig:      wm.NewSnowflakeDataCheckTaskConfig(t.SnowflakeDataCheckConfig),
			TableauDataExtractionConfig:   wm.NewTableauDataExtractionTaskConfig(t.TableauDataExtractionConfig),
			TroccoAgentConfig:             wm.NewTroccoAgentTaskConfig(t.TroccoAgentConfig),
			TroccoBigQueryDatamartConfig:  wm.NewTroccoBigQueryDatamartTaskConfig(t.TroccoBigQueryDatamartConfig),
			TroccoRedshiftDatamartConfig:  wm.NewTroccoRedshiftDatamartTaskConfig(t.TroccoRedshiftDatamartConfig),
			TroccoSnowflakeDatamartConfig: wm.NewTroccoSnowflakeDatamartTaskConfig(t.TroccoSnowflakeDatamartConfig),
			TroccoTransferBulkConfig:      wm.NewTroccoTransferBulkTaskConfig(t.TroccoTransferBulkConfig),
			TroccoTransferConfig:          wm.NewTroccoTransferTaskConfig(t.TroccoTransferConfig),
			WorkflowConfig:                wm.NewWorkflowTaskConfig(t.WorkflowConfig),
		}

		tasks = append(tasks, task)
	}

	newState := pipelineDefinitionResourceModel{
		ID:               types.Int64Value(workflow.ID),
		Name:             types.StringPointerValue(workflow.Name),
		Description:      types.StringPointerValue(workflow.Description),
		Tasks:            tasks,
		Labels:           wm.NewLabels(workflow.Labels, state.Labels == nil),
		Notifications:    wm.NewNotifications(workflow.Notifications, state.Notifications == nil),
		Schedules:        wm.NewSchedules(workflow.Schedules, state.Schedules == nil),
		TaskDependencies: state.TaskDependencies,
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}

func (r *workflowResource) Delete(
	ctx context.Context,
	req resource.DeleteRequest,
	resp *resource.DeleteResponse,
) {
	s := &pipelineDefinitionResourceModel{}
	resp.Diagnostics.Append(req.State.Get(ctx, s)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if err := r.client.DeleteWorkflow(
		s.ID.ValueInt64(),
	); err != nil {
		resp.Diagnostics.AddError(
			"Deleting workflow",
			fmt.Sprintf("Unable to delete workflow, got error: %s", err),
		)
		return
	}
}

func (r *workflowResource) ImportState(
	ctx context.Context,
	req resource.ImportStateRequest,
	resp *resource.ImportStateResponse,
) {
	id, err := strconv.ParseInt(req.ID, 10, 64)
	if err != nil {
		resp.Diagnostics.AddError(
			"Importing workflow",
			fmt.Sprintf("Unable to parse id, got error: %s", err),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), id)...)
}
