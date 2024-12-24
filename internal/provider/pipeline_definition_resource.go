package provider

import (
	"context"
	"fmt"
	"strconv"
	"terraform-provider-trocco/internal/client"
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
)

var (
	_ resource.Resource                = &workflowResource{}
	_ resource.ResourceWithConfigure   = &workflowResource{}
	_ resource.ResourceWithImportState = &workflowResource{}
)

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
			"resource_group_id": schema.Int64Attribute{
				Optional: true,
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
			"max_task_parallelism": schema.Int64Attribute{
				Optional: true,
				Computed: true,
			},
			"execution_timeout": schema.Int64Attribute{
				Optional: true,
				Computed: true,
			},
			"max_retries": schema.Int64Attribute{
				Optional: true,
				Computed: true,
			},
			"min_retry_interval": schema.Int64Attribute{
				Optional: true,
				Computed: true,
			},
			"is_concurrent_execution_skipped": schema.BoolAttribute{
				Optional: true,
				Computed: true,
			},
			"is_stopped_on_errors": schema.BoolAttribute{
				Optional: true,
				Computed: true,
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
			"schedules":         pds.Schedule(),
			"tasks":             pds.Tasks(),
			"task_dependencies": pds.TaskDependencies(),
		},
	}
}

func (r *workflowResource) Create(
	ctx context.Context,
	req resource.CreateRequest,
	resp *resource.CreateResponse,
) {
	plan := &pdm.PipelineDefinition{}
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	workflow, err := r.client.CreateWorkflow(
		plan.ToCreateInput(),
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

	newState := pdm.PipelineDefinition{
		ID:                           types.Int64Value(workflow.ID),
		ResourceGroupID:              types.Int64PointerValue(workflow.ResourceGroupID),
		Name:                         types.StringPointerValue(workflow.Name),
		Description:                  types.StringPointerValue(workflow.Description),
		MaxTaskParallelism:           types.Int64PointerValue(workflow.MaxTaskParallelism),
		ExecutionTimeout:             types.Int64PointerValue(workflow.ExecutionTimeout),
		MaxRetries:                   types.Int64PointerValue(workflow.MaxRetries),
		MinRetryInterval:             types.Int64PointerValue(workflow.MinRetryInterval),
		IsConcurrentExecutionSkipped: types.BoolPointerValue(workflow.IsConcurrentExecutionSkipped),
		IsStoppedOnErrors:            types.BoolPointerValue(workflow.IsStoppedOnErrors),
		Labels:                       pdm.NewLabels(workflow.Labels, plan.Labels == nil),
		Notifications:                pdm.NewNotifications(workflow.Notifications, plan.Notifications == nil),
		Schedules:                    pdm.NewSchedules(workflow.Schedules, plan.Schedules == nil),
		Tasks:                        pdm.NewTasks(workflow.Tasks, keys, plan),
		TaskDependencies:             pdm.NewTaskDependencies(workflow.TaskDependencies, keys),
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}

func (r *workflowResource) Update(
	ctx context.Context,
	req resource.UpdateRequest,
	resp *resource.UpdateResponse,
) {
	state := &pdm.PipelineDefinition{}
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	plan := &pdm.PipelineDefinition{}
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

	newState := pdm.PipelineDefinition{
		ID:                           types.Int64Value(workflow.ID),
		ResourceGroupID:              types.Int64PointerValue(workflow.ResourceGroupID),
		Name:                         types.StringPointerValue(workflow.Name),
		Description:                  types.StringPointerValue(workflow.Description),
		MaxTaskParallelism:           types.Int64PointerValue(workflow.MaxTaskParallelism),
		ExecutionTimeout:             types.Int64PointerValue(workflow.ExecutionTimeout),
		MaxRetries:                   types.Int64PointerValue(workflow.MaxRetries),
		MinRetryInterval:             types.Int64PointerValue(workflow.MinRetryInterval),
		IsConcurrentExecutionSkipped: types.BoolPointerValue(workflow.IsConcurrentExecutionSkipped),
		IsStoppedOnErrors:            types.BoolPointerValue(workflow.IsStoppedOnErrors),
		Labels:                       pdm.NewLabels(workflow.Labels, plan.Labels == nil),
		Notifications:                pdm.NewNotifications(workflow.Notifications, plan.Notifications == nil),
		Schedules:                    pdm.NewSchedules(workflow.Schedules, plan.Schedules == nil),
		Tasks:                        pdm.NewTasks(workflow.Tasks, keys, plan),
		TaskDependencies:             pdm.NewTaskDependencies(workflow.TaskDependencies, keys),
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}

func (r *workflowResource) Read(
	ctx context.Context,
	req resource.ReadRequest,
	resp *resource.ReadResponse,
) {
	state := &pdm.PipelineDefinition{}
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

	newState := pdm.PipelineDefinition{
		ID:                           types.Int64Value(workflow.ID),
		ResourceGroupID:              types.Int64PointerValue(workflow.ResourceGroupID),
		Name:                         types.StringPointerValue(workflow.Name),
		Description:                  types.StringPointerValue(workflow.Description),
		MaxTaskParallelism:           types.Int64PointerValue(workflow.MaxTaskParallelism),
		ExecutionTimeout:             types.Int64PointerValue(workflow.ExecutionTimeout),
		MaxRetries:                   types.Int64PointerValue(workflow.MaxRetries),
		MinRetryInterval:             types.Int64PointerValue(workflow.MinRetryInterval),
		IsConcurrentExecutionSkipped: types.BoolPointerValue(workflow.IsConcurrentExecutionSkipped),
		IsStoppedOnErrors:            types.BoolPointerValue(workflow.IsStoppedOnErrors),
		Tasks:                        pdm.NewTasks(workflow.Tasks, keys, state),
		Labels:                       pdm.NewLabels(workflow.Labels, state.Labels == nil),
		Notifications:                pdm.NewNotifications(workflow.Notifications, state.Notifications == nil),
		Schedules:                    pdm.NewSchedules(workflow.Schedules, state.Schedules == nil),
		TaskDependencies:             state.TaskDependencies,
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}

func (r *workflowResource) Delete(
	ctx context.Context,
	req resource.DeleteRequest,
	resp *resource.DeleteResponse,
) {
	s := &pdm.PipelineDefinition{}
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
