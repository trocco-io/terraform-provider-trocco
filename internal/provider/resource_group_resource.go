package provider

import (
	"context"
	"fmt"
	"strconv"
	"terraform-provider-trocco/internal/client"
	model "terraform-provider-trocco/internal/provider/model/resource_group"
	troccoValidator "terraform-provider-trocco/internal/provider/validator"

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
	_ resource.Resource                = &resourceGroupResource{}
	_ resource.ResourceWithConfigure   = &resourceGroupResource{}
	_ resource.ResourceWithImportState = &resourceGroupResource{}
)

func NewResourceGroupResource() resource.Resource {
	return &resourceGroupResource{}
}

type resourceGroupResource struct {
	client *client.TroccoClient
}

func (r *resourceGroupResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_resource_group"
}

func (r *resourceGroupResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *resourceGroupResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Provides a TROCCO resource_group resource.",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "The ID of the resource group.",
			},
			"name": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "The name of the resource group.",
			},
			"description": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "The description of the resource group.",
			},
			"teams": schema.SetNestedAttribute{
				Required:            true,
				MarkdownDescription: "The team roles of the resource group.",
				Validators: []validator.Set{
					troccoValidator.UniqueTeamValidator{},
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"team_id": schema.Int64Attribute{
							Required:            true,
							MarkdownDescription: "The team ID of the role.",
						},
						"role": schema.StringAttribute{
							Required: true,
							Validators: []validator.String{
								stringvalidator.OneOf("administrator", "editor", "operator", "viewer"),
							},
							MarkdownDescription: "The role of the team. Valid values are `administrator`, `editor`, `operator`, `viewer`.",
						},
					},
				},
			},
		},
	}
}

func (r *resourceGroupResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan model.ResourceGroupResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	input := client.CreateResourceGroupInput{
		Name:        plan.Name.ValueString(),
		Description: plan.Description.ValueStringPointer(),
		Teams:     []client.TeamRoleInput{},
	}
	for _, m := range plan.Teams {
		input.Teams = append(input.Teams, client.TeamRoleInput{
			TeamID: m.TeamID.ValueInt64(),
			Role:   m.Role.ValueString(),
		})

	}

	resourceGroup, err := r.client.CreateResourceGroup(&input)
	if err != nil {
		resp.Diagnostics.AddError(
			"Creating resource group",
			fmt.Sprintf("Unable to create resource group, got error: %s", err),
		)
		return
	}

	newState := model.ResourceGroupResourceModel{
		ID:          types.Int64Value(resourceGroup.ID),
		Name:        types.StringValue(resourceGroup.Name),
		Description: types.StringPointerValue(resourceGroup.Description),
		Teams:     []model.TeamRoleResourceModel{},
	}
	for _, m := range resourceGroup.Teams {
		newState.Teams = append(newState.Teams, model.TeamRoleResourceModel{
			TeamID: types.Int64Value(m.TeamID),
			Role:   types.StringValue(m.Role),
		})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}

func (r *resourceGroupResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state model.ResourceGroupResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resourceGroup, err := r.client.GetResourceGroup(state.ID.ValueInt64())
	if err != nil {
		resp.Diagnostics.AddError(
			"Reading resource group",
			fmt.Sprintf("Unable to read resource group, got error: %s", err),
		)
		return
	}

	newState := model.ResourceGroupResourceModel{
		ID:          types.Int64Value(resourceGroup.ID),
		Name:        types.StringValue(resourceGroup.Name),
		Description: types.StringPointerValue(resourceGroup.Description),
		Teams:     []model.TeamRoleResourceModel{},
	}
	for _, m := range resourceGroup.Teams {
		newState.Teams = append(newState.Teams, model.TeamRoleResourceModel{
			TeamID: types.Int64Value(m.TeamID),
			Role:   types.StringValue(m.Role),
		})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}
func (r *resourceGroupResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state model.ResourceGroupResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	input := client.UpdateResourceGroupInput{
		Name:        *plan.Name.ValueStringPointer(),
		Description: plan.Description.ValueStringPointer(),
		Teams:     []client.TeamRoleInput{},
	}
	for _, m := range plan.Teams {
		input.Teams = append(input.Teams, client.TeamRoleInput{
			TeamID: m.TeamID.ValueInt64(),
			Role:   m.Role.ValueString(),
		})
	}

	resourceGroup, err := r.client.UpdateResourceGroup(state.ID.ValueInt64(), &input)
	if err != nil {
		resp.Diagnostics.AddError(
			"Updating team",
			fmt.Sprintf("Unable to update team, got error: %s", err),
		)
		return
	}

	newState := model.ResourceGroupResourceModel{
		ID:          types.Int64Value(resourceGroup.ID),
		Name:        types.StringValue(resourceGroup.Name),
		Description: types.StringPointerValue(resourceGroup.Description),
		Teams:     []model.TeamRoleResourceModel{},
	}
	for _, m := range resourceGroup.Teams {
		newState.Teams = append(newState.Teams, model.TeamRoleResourceModel{
			TeamID: types.Int64Value(m.TeamID),
			Role:   types.StringValue(m.Role),
		})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}
func (r *resourceGroupResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state model.ResourceGroupResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := r.client.DeleteResourceGroup(state.ID.ValueInt64())
	if err != nil {
		resp.Diagnostics.AddError(
			"Deleting resource group",
			fmt.Sprintf("Unable to delete resource group, got error: %s", err),
		)
		return
	}
}

func (r *resourceGroupResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	id, err := strconv.ParseInt(req.ID, 10, 64)
	if err != nil {
		resp.Diagnostics.AddError(
			"Importing resource group",
			fmt.Sprintf("Unable to parse id, got error: %s", err),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), id)...)
}
