package provider

import (
	"context"
	"fmt"
	"strconv"
	"terraform-provider-trocco/internal/client"

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
	_ resource.Resource                = &teamResource{}
	_ resource.ResourceWithConfigure   = &teamResource{}
	_ resource.ResourceWithImportState = &teamResource{}
)

func NewTeamResource() resource.Resource {
	return &teamResource{}
}

type teamResource struct {
	client *client.TroccoClient
}

type teamResourceModel struct {
	ID          types.Int64               `tfsdk:"id"`
	Name        types.String              `tfsdk:"name"`
	Description types.String              `tfsdk:"description"`
	Members     []teamMemberResourceModel `tfsdk:"members"`
}

type teamMemberResourceModel struct {
	UserID types.Int64  `tfsdk:"user_id"`
	Role   types.String `tfsdk:"role"`
}

func (r *teamResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_team"
}

func (r *teamResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *teamResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Provides a Trocco team resource.",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "The ID of the team.",
			},
			"name": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "The name of the team.",
			},
			"description": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "The description of the team.",
			},
			"members": schema.ListNestedAttribute{
				Required: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"user_id": schema.Int64Attribute{
							Required:            true,
							MarkdownDescription: "The user ID of the team member.",
						},
						"role": schema.StringAttribute{
							Required: true,
							Validators: []validator.String{
								stringvalidator.OneOf("team_admin", "team_member"),
							},
							MarkdownDescription: "The role of the team member. Valid values are `team_admin` or `team_member`.",
						},
					},
				},
			},
		},
	}
}

func (r *teamResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan teamResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	input := client.CreateTeamInput{
		Name:        plan.Name.ValueString(),
		Description: plan.Description.ValueStringPointer(),
		Members:     []client.MemberInput{},
	}
	for _, m := range plan.Members {
		input.Members = append(input.Members, client.MemberInput{
			UserID: m.UserID.ValueInt64(),
			Role:   m.Role.ValueString(),
		})

	}

	team, err := r.client.CreateTeam(&input)
	if err != nil {
		resp.Diagnostics.AddError(
			"Creating team",
			fmt.Sprintf("Unable to create team, got error: %s", err),
		)
		return
	}

	newState := teamResourceModel{
		ID:          types.Int64Value(team.ID),
		Name:        types.StringValue(team.Name),
		Description: types.StringPointerValue(team.Description),
		Members:     []teamMemberResourceModel{},
	}
	for _, m := range team.Members {
		newState.Members = append(newState.Members, teamMemberResourceModel{
			UserID: types.Int64Value(m.UserID),
			Role:   types.StringValue(m.Role),
		})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}

func (r *teamResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state teamResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	team, err := r.client.GetTeam(state.ID.ValueInt64())
	if err != nil {
		resp.Diagnostics.AddError(
			"Reading team",
			fmt.Sprintf("Unable to read team, got error: %s", err),
		)
		return
	}

	newState := teamResourceModel{
		ID:          types.Int64Value(team.ID),
		Name:        types.StringValue(team.Name),
		Description: types.StringPointerValue(team.Description),
		Members:     []teamMemberResourceModel{},
	}
	for _, m := range team.Members {
		newState.Members = append(newState.Members, teamMemberResourceModel{
			UserID: types.Int64Value(m.UserID),
			Role:   types.StringValue(m.Role),
		})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}
func (r *teamResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state teamResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	input := client.UpdateTeamInput{
		Name:        plan.Name.ValueStringPointer(),
		Description: plan.Description.ValueStringPointer(),
		Members:     []client.MemberInput{},
	}
	for _, m := range plan.Members {
		input.Members = append(input.Members, client.MemberInput{
			UserID: m.UserID.ValueInt64(),
			Role:   m.Role.ValueString(),
		})
	}

	team, err := r.client.UpdateTeam(state.ID.ValueInt64(), &input)
	if err != nil {
		resp.Diagnostics.AddError(
			"Updating team",
			fmt.Sprintf("Unable to update team, got error: %s", err),
		)
		return
	}

	newState := teamResourceModel{
		ID:          types.Int64Value(team.ID),
		Name:        types.StringValue(team.Name),
		Description: types.StringPointerValue(team.Description),
		Members:     []teamMemberResourceModel{},
	}
	for _, m := range team.Members {
		newState.Members = append(newState.Members, teamMemberResourceModel{
			UserID: types.Int64Value(m.UserID),
			Role:   types.StringValue(m.Role),
		})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}
func (r *teamResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state teamResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := r.client.DeleteTeam(state.ID.ValueInt64())
	if err != nil {
		resp.Diagnostics.AddError(
			"Deleting team",
			fmt.Sprintf("Unable to delete team, got error: %s", err),
		)
		return
	}
}

func (r *teamResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	id, err := strconv.ParseInt(req.ID, 10, 64)
	if err != nil {
		resp.Diagnostics.AddError(
			"Importing team",
			fmt.Sprintf("Unable to parse id, got error: %s", err),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), id)...)
}
