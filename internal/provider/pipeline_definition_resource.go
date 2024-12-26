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
	_ resource.Resource                = &pipelineDefinitionResource{}
	_ resource.ResourceWithConfigure   = &pipelineDefinitionResource{}
	_ resource.ResourceWithImportState = &pipelineDefinitionResource{}
)

type pipelineDefinitionResource struct {
	client *client.TroccoClient
}

func NewPipelineDefinitionResource() resource.Resource {
	return &pipelineDefinitionResource{}
}

func (r *pipelineDefinitionResource) Metadata(
	ctx context.Context,
	req resource.MetadataRequest,
	resp *resource.MetadataResponse,
) {
	resp.TypeName = fmt.Sprintf("%s_pipeline_definition", req.ProviderTypeName)
}

func (r *pipelineDefinitionResource) Configure(
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

func (r *pipelineDefinitionResource) Schema(
	ctx context.Context,
	req resource.SchemaRequest,
	resp *resource.SchemaResponse,
) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Provides a TROCCO pipelineDefinition resource.",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				MarkdownDescription: "The ID of the pipelineDefinition",
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
				Computed: true,
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

func (r *pipelineDefinitionResource) Create(
	ctx context.Context,
	req resource.CreateRequest,
	resp *resource.CreateResponse,
) {
	plan := &pdm.PipelineDefinition{}
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	en, err := r.client.CreatePipelineDefinition(
		plan.ToCreateInput(),
	)
	if err != nil {
		resp.Diagnostics.AddError(
			"Creating pipeline definition",
			fmt.Sprintf("Unable to create pipeline definition, got error: %s", err),
		)
		return
	}

	keys := map[int64]types.String{}
	for _, t := range en.Tasks {
		keys[t.TaskIdentifier] = types.StringValue(t.Key)
	}

	newState := pdm.NewPipelineDefinition(en, keys, plan)

	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}

func (r *pipelineDefinitionResource) Update(
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

	en, err := r.client.UpdatePipelineDefinition(
		state.ID.ValueInt64(),
		plan.ToUpdateWorkflowInput(state),
	)
	if err != nil {
		resp.Diagnostics.AddError(
			"Updating pipeline definition",
			fmt.Sprintf("Unable to update pipeline definition, got error: %s", err),
		)
		return
	}

	keys := map[int64]types.String{}
	for _, t := range en.Tasks {
		keys[t.TaskIdentifier] = types.StringValue(t.Key)
	}

	newState := pdm.NewPipelineDefinition(en, keys, plan)

	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}

func (r *pipelineDefinitionResource) Read(
	ctx context.Context,
	req resource.ReadRequest,
	resp *resource.ReadResponse,
) {
	state := &pdm.PipelineDefinition{}
	resp.Diagnostics.Append(req.State.Get(ctx, state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	en, err := r.client.GetPipelineDefinition(
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

	newState := pdm.NewPipelineDefinition(en, keys, state)

	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}

func (r *pipelineDefinitionResource) Delete(
	ctx context.Context,
	req resource.DeleteRequest,
	resp *resource.DeleteResponse,
) {
	s := &pdm.PipelineDefinition{}
	resp.Diagnostics.Append(req.State.Get(ctx, s)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if err := r.client.DeletePipelineDefinition(
		s.ID.ValueInt64(),
	); err != nil {
		resp.Diagnostics.AddError(
			"Deleting pipeline definition",
			fmt.Sprintf("Unable to delete pipeline definition, got error: %s", err),
		)
		return
	}
}

func (r *pipelineDefinitionResource) ImportState(
	ctx context.Context,
	req resource.ImportStateRequest,
	resp *resource.ImportStateResponse,
) {
	id, err := strconv.ParseInt(req.ID, 10, 64)
	if err != nil {
		resp.Diagnostics.AddError(
			"Importing pipeline definition",
			fmt.Sprintf("Unable to parse id, got error: %s", err),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), id)...)
}
