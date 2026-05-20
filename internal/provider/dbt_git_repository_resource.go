package provider

import (
	"context"
	"fmt"
	"strconv"
	"terraform-provider-trocco/internal/client"
	"terraform-provider-trocco/internal/provider/model"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var supportedDbtAdapterTypes = []string{"bigquery", "snowflake", "redshift"}

var (
	_ resource.Resource                = &dbtGitRepositoryResource{}
	_ resource.ResourceWithConfigure   = &dbtGitRepositoryResource{}
	_ resource.ResourceWithImportState = &dbtGitRepositoryResource{}
)

func NewDbtGitRepositoryResource() resource.Resource {
	return &dbtGitRepositoryResource{}
}

type dbtGitRepositoryResource struct {
	client *client.TroccoClient
}

func (r *dbtGitRepositoryResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_dbt_git_repository"
}

func (r *dbtGitRepositoryResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	c, ok := req.ProviderData.(*client.TroccoClient)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *client.TroccoClient, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}

	r.client = c
}

func (r *dbtGitRepositoryResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Provides a TROCCO dbt Git repository resource.",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "The ID of the dbt Git repository.",
			},
			"name": schema.StringAttribute{
				Required: true,
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtLeast(1),
				},
				MarkdownDescription: "The name of the dbt Git repository.",
			},
			"description": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "The description of the dbt Git repository.",
			},
			"adapter_type": schema.StringAttribute{
				Required: true,
				Validators: []validator.String{
					stringvalidator.OneOf(supportedDbtAdapterTypes...),
				},
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: "The adapter type. Must be one of `bigquery`, `snowflake`, `redshift`. Cannot be changed after creation.",
			},
			"dbt_version": schema.StringAttribute{
				Required: true,
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtLeast(1),
				},
				MarkdownDescription: "The dbt version (e.g., `1.11`). Supported values follow what TROCCO currently allows; refer to the TROCCO documentation for the latest list.",
			},
			"url": schema.StringAttribute{
				Required: true,
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtLeast(1),
				},
				MarkdownDescription: "The URL of the Git repository.",
			},
			"branch": schema.StringAttribute{
				Required: true,
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtLeast(1),
				},
				MarkdownDescription: "The branch of the Git repository to sync from.",
			},
			"subdirectory": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "The subdirectory where the dbt project is located in the Git repository.",
			},
			"resource_group_id": schema.Int64Attribute{
				Optional:            true,
				MarkdownDescription: "The ID of the resource group that the dbt Git repository belongs to.",
			},
		},
	}
}

func (r *dbtGitRepositoryResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan model.DbtGitRepositoryModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	input := client.CreateDbtGitRepositoryInput{
		Name:            plan.Name.ValueString(),
		Description:     plan.Description.ValueStringPointer(),
		AdapterType:     plan.AdapterType.ValueString(),
		DbtVersion:      plan.DbtVersion.ValueString(),
		URL:             plan.URL.ValueString(),
		Branch:          plan.Branch.ValueString(),
		Subdirectory:    plan.Subdirectory.ValueStringPointer(),
		ResourceGroupID: plan.ResourceGroupID.ValueInt64Pointer(),
	}

	repo, err := r.client.CreateDbtGitRepository(&input)
	if err != nil {
		resp.Diagnostics.AddError(
			"Creating dbt Git repository",
			fmt.Sprintf("Unable to create dbt Git repository, got error: %s", err),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, newDbtGitRepositoryModel(repo))...)
}

func (r *dbtGitRepositoryResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state model.DbtGitRepositoryModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	repo, err := r.client.GetDbtGitRepository(state.ID.ValueInt64())
	if err != nil {
		resp.Diagnostics.AddError(
			"Reading dbt Git repository",
			fmt.Sprintf("Unable to read dbt Git repository, got error: %s", err),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, newDbtGitRepositoryModel(repo))...)
}

func (r *dbtGitRepositoryResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state model.DbtGitRepositoryModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	input := client.UpdateDbtGitRepositoryInput{
		Name:            plan.Name.ValueStringPointer(),
		Description:     plan.Description.ValueStringPointer(),
		DbtVersion:      plan.DbtVersion.ValueStringPointer(),
		URL:             plan.URL.ValueStringPointer(),
		Branch:          plan.Branch.ValueStringPointer(),
		Subdirectory:    plan.Subdirectory.ValueStringPointer(),
		ResourceGroupID: plan.ResourceGroupID.ValueInt64Pointer(),
	}

	repo, err := r.client.UpdateDbtGitRepository(state.ID.ValueInt64(), &input)
	if err != nil {
		resp.Diagnostics.AddError(
			"Updating dbt Git repository",
			fmt.Sprintf("Unable to update dbt Git repository, got error: %s", err),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, newDbtGitRepositoryModel(repo))...)
}

func (r *dbtGitRepositoryResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state model.DbtGitRepositoryModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if err := r.client.DeleteDbtGitRepository(state.ID.ValueInt64()); err != nil {
		resp.Diagnostics.AddError(
			"Deleting dbt Git repository",
			fmt.Sprintf("Unable to delete dbt Git repository, got error: %s", err),
		)
		return
	}
}

func (r *dbtGitRepositoryResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	id, err := strconv.ParseInt(req.ID, 10, 64)
	if err != nil {
		resp.Diagnostics.AddError(
			"Importing dbt Git repository",
			fmt.Sprintf("Unable to parse id, got error: %s", err),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), id)...)
}

func newDbtGitRepositoryModel(repo *client.DbtGitRepository) model.DbtGitRepositoryModel {
	return model.DbtGitRepositoryModel{
		ID:              types.Int64Value(repo.ID),
		Name:            types.StringValue(repo.Name),
		Description:     types.StringPointerValue(repo.Description),
		AdapterType:     types.StringValue(repo.AdapterType),
		DbtVersion:      types.StringValue(repo.DbtVersion),
		URL:             types.StringValue(repo.URL),
		Branch:          types.StringValue(repo.Branch),
		Subdirectory:    types.StringPointerValue(repo.Subdirectory),
		ResourceGroupID: types.Int64PointerValue(repo.ResourceGroupID),
	}
}
