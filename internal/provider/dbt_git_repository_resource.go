package provider

import (
	"context"
	"fmt"
	"regexp"
	"strconv"
	"terraform-provider-trocco/internal/client"
	"terraform-provider-trocco/internal/provider/model"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	supportedDbtAdapterTypes = []string{"bigquery", "snowflake", "redshift"}
	supportedDbtRefTypes     = []string{"branch", "tag", "commit_hash"}
)

const defaultDbtRefType = "branch"

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
			"ref_type": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Default:  stringdefault.StaticString(defaultDbtRefType),
				Validators: []validator.String{
					stringvalidator.OneOf(supportedDbtRefTypes...),
				},
				MarkdownDescription: "The Git reference type. Must be one of `branch`, `tag`, `commit_hash`. Defaults to `branch`. Exactly the matching attribute (`branch` / `tag` / `commit_hash`) must be set; the others must be left unset.",
			},
			"branch": schema.StringAttribute{
				Optional: true,
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtLeast(1),
				},
				MarkdownDescription: "The branch of the Git repository to sync from. Required when `ref_type` is `branch`.",
			},
			"tag": schema.StringAttribute{
				Optional: true,
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtLeast(1),
				},
				MarkdownDescription: "The tag of the Git repository to sync from. Required when `ref_type` is `tag`.",
			},
			"commit_hash": schema.StringAttribute{
				Optional: true,
				Validators: []validator.String{
					stringvalidator.RegexMatches(
						regexp.MustCompile(`^[0-9a-f]{40}$`),
						"must be a 40-character lowercase hexadecimal Git commit hash",
					),
				},
				MarkdownDescription: "The commit hash of the Git repository to sync from (40-character lowercase hex). Required when `ref_type` is `commit_hash`.",
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

	validateRefConsistency(&plan, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	refType := plan.RefType.ValueString()
	input := client.CreateDbtGitRepositoryInput{
		Name:            plan.Name.ValueString(),
		Description:     plan.Description.ValueStringPointer(),
		AdapterType:     plan.AdapterType.ValueString(),
		DbtVersion:      plan.DbtVersion.ValueString(),
		URL:             plan.URL.ValueString(),
		RefType:         &refType,
		Branch:          plan.Branch.ValueStringPointer(),
		Tag:             plan.Tag.ValueStringPointer(),
		CommitHash:      plan.CommitHash.ValueStringPointer(),
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

	validateRefConsistency(&plan, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	input := client.UpdateDbtGitRepositoryInput{
		Name:            plan.Name.ValueStringPointer(),
		Description:     plan.Description.ValueStringPointer(),
		DbtVersion:      plan.DbtVersion.ValueStringPointer(),
		URL:             plan.URL.ValueStringPointer(),
		RefType:         plan.RefType.ValueStringPointer(),
		Branch:          plan.Branch.ValueStringPointer(),
		Tag:             plan.Tag.ValueStringPointer(),
		CommitHash:      plan.CommitHash.ValueStringPointer(),
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
		RefType:         types.StringValue(repo.RefType),
		Branch:          types.StringPointerValue(repo.Branch),
		Tag:             types.StringPointerValue(repo.Tag),
		CommitHash:      types.StringPointerValue(repo.CommitHash),
		Subdirectory:    types.StringPointerValue(repo.Subdirectory),
		ResourceGroupID: types.Int64PointerValue(repo.ResourceGroupID),
	}
}

// validateRefConsistency ensures the user sets exactly the ref attribute matching ref_type.
// The server silently nulls non-matching ref values, which would otherwise cause a drift loop.
func validateRefConsistency(plan *model.DbtGitRepositoryModel, diags *diag.Diagnostics) {
	refType := plan.RefType.ValueString()
	branchSet := !plan.Branch.IsNull() && !plan.Branch.IsUnknown()
	tagSet := !plan.Tag.IsNull() && !plan.Tag.IsUnknown()
	commitSet := !plan.CommitHash.IsNull() && !plan.CommitHash.IsUnknown()

	switch refType {
	case "branch":
		if !branchSet {
			diags.AddAttributeError(path.Root("branch"), "Missing branch", "`branch` is required when `ref_type` is `branch`.")
		}
		if tagSet {
			diags.AddAttributeError(path.Root("tag"), "Unexpected tag", "`tag` must not be set when `ref_type` is `branch`.")
		}
		if commitSet {
			diags.AddAttributeError(path.Root("commit_hash"), "Unexpected commit_hash", "`commit_hash` must not be set when `ref_type` is `branch`.")
		}
	case "tag":
		if !tagSet {
			diags.AddAttributeError(path.Root("tag"), "Missing tag", "`tag` is required when `ref_type` is `tag`.")
		}
		if branchSet {
			diags.AddAttributeError(path.Root("branch"), "Unexpected branch", "`branch` must not be set when `ref_type` is `tag`.")
		}
		if commitSet {
			diags.AddAttributeError(path.Root("commit_hash"), "Unexpected commit_hash", "`commit_hash` must not be set when `ref_type` is `tag`.")
		}
	case "commit_hash":
		if !commitSet {
			diags.AddAttributeError(path.Root("commit_hash"), "Missing commit_hash", "`commit_hash` is required when `ref_type` is `commit_hash`.")
		}
		if branchSet {
			diags.AddAttributeError(path.Root("branch"), "Unexpected branch", "`branch` must not be set when `ref_type` is `commit_hash`.")
		}
		if tagSet {
			diags.AddAttributeError(path.Root("tag"), "Unexpected tag", "`tag` must not be set when `ref_type` is `commit_hash`.")
		}
	}
}
