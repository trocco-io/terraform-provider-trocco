package provider

import (
	"context"
	"fmt"
	"strconv"
	"terraform-provider-trocco/internal/client"
	pdp "terraform-provider-trocco/internal/client/parameters/pipeline_definition"
	pdm "terraform-provider-trocco/internal/provider/models/pipeline_definition"
	pds "terraform-provider-trocco/internal/provider/schemas/pipeline_definition"

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

type pipelineDefinitionModel struct {
	ID               types.Int64           `tfsdk:"id"`
	Name             types.String          `tfsdk:"name"`
	Description      types.String          `tfsdk:"description"`
	Labels           []types.String        `tfsdk:"labels"`
	Notifications    []pdm.Notification    `tfsdk:"notifications"`
	Schedules        []pdm.Schedule        `tfsdk:"schedules"`
	Tasks            []*pdm.Task           `tfsdk:"tasks"`
	TaskDependencies []*pdm.TaskDependency `tfsdk:"task_dependencies"`
}

func (m *pipelineDefinitionModel) ToCreateWorkflowInput() *client.CreateWorkflowInput {
	labels := []string{}
	for _, l := range m.Labels {
		labels = append(labels, l.ValueString())
	}

	notifications := []pdp.Notification{}
	for _, n := range m.Notifications {
		notifications = append(notifications, n.ToInput())
	}

	schedules := []pdp.Schedule{}
	for _, s := range m.Schedules {
		schedules = append(schedules, s.ToInput())
	}

	tasks := []pdp.Task{}
	for _, t := range m.Tasks {
		tasks = append(tasks, *t.ToInput(map[string]int64{}))
	}

	taskDependencies := []pdp.TaskDependency{}
	for _, d := range m.TaskDependencies {
		taskDependencies = append(taskDependencies, pdp.TaskDependency{
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

func (m *pipelineDefinitionModel) ToUpdateWorkflowInput(state *pipelineDefinitionModel) *client.UpdateWorkflowInput {
	labels := []string{}
	for _, l := range m.Labels {
		labels = append(labels, l.ValueString())
	}

	notifications := []pdp.Notification{}
	for _, n := range m.Notifications {
		notifications = append(notifications, n.ToInput())
	}

	schedules := []pdp.Schedule{}
	for _, s := range m.Schedules {
		schedules = append(schedules, s.ToInput())
	}

	stateTaskIdentifiers := map[string]int64{}
	for _, s := range state.Tasks {
		stateTaskIdentifiers[s.Key.ValueString()] = s.TaskIdentifier.ValueInt64()
	}

	tasks := []pdp.Task{}
	for _, t := range m.Tasks {
		tasks = append(tasks, *t.ToInput(stateTaskIdentifiers))
	}

	taskDependencies := []pdp.TaskDependency{}
	for _, d := range m.TaskDependencies {
		taskDependencies = append(taskDependencies, pdp.TaskDependency{
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
	customVariables := pds.NewCustomVariableAttribute()

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
						"trocco_transfer_config":           pds.NewTroccoTransferTaskConfigAttribute(),
						"trocco_transfer_bulk_config":      pds.NewTroccoTransferBulkTaskConfigAttribute(),
						"dbt_config":                       pds.NewDBTTaskConfigAttribute(),
						"trocco_agent_config":              pds.NewTroccoAgentTaskConfigAttribute(),
						"trocco_bigquery_datamart_config":  pds.NewBigQueryDatamartTaskConfigAttribute(),
						"trocco_redshift_datamart_config":  pds.NewRedshiftDatamartTaskConfigAttribute(),
						"trocco_snowflake_datamart_config": pds.NewSnowflakeDatamartTaskConfigAttribute(),
						"trocco_pipeline_config":           pds.NewTroccoPiplineTaskConfigAttribute(),
						"slack_notification_config":        pds.NewSlackNotificationTaskConfigAttribute(),
						"tableau_data_extraction_config":   pds.NewTableauDataExtractionTaskConfigAttribute(),
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
	plan := &pipelineDefinitionModel{}
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

	keys := map[int64]types.String{}
	for _, t := range workflow.Tasks {
		keys[t.TaskIdentifier] = types.StringValue(t.Key)
	}

	newState := pipelineDefinitionModel{
		ID:               types.Int64Value(workflow.ID),
		Name:             types.StringPointerValue(workflow.Name),
		Description:      types.StringPointerValue(workflow.Description),
		Labels:           pdm.NewLabels(workflow.Labels, plan.Labels == nil),
		Notifications:    pdm.NewNotifications(workflow.Notifications, plan.Notifications == nil),
		Schedules:        pdm.NewSchedules(workflow.Schedules, plan.Schedules == nil),
		Tasks:            pdm.NewTasks(workflow.Tasks, keys),
		TaskDependencies: pdm.NewTaskDependencies(workflow.TaskDependencies, keys),
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}

func (r *workflowResource) Update(
	ctx context.Context,
	req resource.UpdateRequest,
	resp *resource.UpdateResponse,
) {
	state := &pipelineDefinitionModel{}
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	plan := &pipelineDefinitionModel{}
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

	keys := map[int64]types.String{}
	for _, t := range workflow.Tasks {
		keys[t.TaskIdentifier] = types.StringValue(t.Key)
	}

	newState := pipelineDefinitionModel{
		ID:               types.Int64Value(workflow.ID),
		Name:             types.StringPointerValue(workflow.Name),
		Description:      types.StringPointerValue(workflow.Description),
		Labels:           pdm.NewLabels(workflow.Labels, plan.Labels == nil),
		Notifications:    pdm.NewNotifications(workflow.Notifications, plan.Notifications == nil),
		Schedules:        pdm.NewSchedules(workflow.Schedules, plan.Schedules == nil),
		Tasks:            pdm.NewTasks(workflow.Tasks, keys),
		TaskDependencies: pdm.NewTaskDependencies(workflow.TaskDependencies, keys),
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}

func (r *workflowResource) Read(
	ctx context.Context,
	req resource.ReadRequest,
	resp *resource.ReadResponse,
) {
	state := &pipelineDefinitionModel{}
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

	keys := map[int64]types.String{}
	for _, t := range state.Tasks {
		keys[t.TaskIdentifier.ValueInt64()] = t.Key
	}

	newState := pipelineDefinitionModel{
		ID:               types.Int64Value(workflow.ID),
		Name:             types.StringPointerValue(workflow.Name),
		Description:      types.StringPointerValue(workflow.Description),
		Tasks:            pdm.NewTasks(workflow.Tasks, keys),
		Labels:           pdm.NewLabels(workflow.Labels, state.Labels == nil),
		Notifications:    pdm.NewNotifications(workflow.Notifications, state.Notifications == nil),
		Schedules:        pdm.NewSchedules(workflow.Schedules, state.Schedules == nil),
		TaskDependencies: state.TaskDependencies,
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}

func (r *workflowResource) Delete(
	ctx context.Context,
	req resource.DeleteRequest,
	resp *resource.DeleteResponse,
) {
	s := &pipelineDefinitionModel{}
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
