package provider

import (
	"context"
	"fmt"
	"terraform-provider-trocco/internal/client"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ resource.Resource              = &userResource{}
	_ resource.ResourceWithConfigure = &userResource{}
)

func NewUserResource() resource.Resource {
	return &userResource{}
}

type userResource struct {
	client *client.TroccoClient
}

type userResourceModel struct {
	ID                           types.Int64  `tfsdk:"id"`
	Email                        types.String `tfsdk:"email"`
	Password                     types.String `tfsdk:"password"`
	Role                         types.String `tfsdk:"role"`
	CanUseAuditLog               types.Bool   `tfsdk:"can_use_audit_log"`
	IsRestrictedConnectionModify types.Bool   `tfsdk:"is_restricted_connection_modify"`
}

func (r *userResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_user"
}

func (r *userResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *userResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Provides a TROCCO user resource.",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "The ID of the user.",
			},
			"email": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "The email of the user.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"password": schema.StringAttribute{
				Optional:  true,
				Computed:  true,
				Sensitive: true,
				PlanModifiers: []planmodifier.String{
					&IgnoreChangesPlanModifier{},
					&RequiredOnCreatePlanModifier{"password"},
				},
				MarkdownDescription: "The password of the user.",
			},
			"role": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "The role of the user.",
			},
			"can_use_audit_log": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Whether the user can use the audit log.",
			},
			"is_restricted_connection_modify": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Whether the user is restricted to modify connections.",
			},
		},
	}
}

func (r *userResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan userResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	input := client.CreateUserInput{
		Email:    plan.Email.ValueString(),
		Password: plan.Password.ValueString(),
		Role:     plan.Role.ValueString(),
	}
	if !plan.CanUseAuditLog.IsNull() {
		input.CanUseAuditLog = plan.CanUseAuditLog.ValueBool()
	}
	if !plan.IsRestrictedConnectionModify.IsNull() {
		input.IsRestrictedConnectionModify = plan.IsRestrictedConnectionModify.ValueBool()
	}

	user, err := r.client.CreateUser(&input)
	if err != nil {
		resp.Diagnostics.AddError(
			"Creating user",
			fmt.Sprintf("Unable to create user, got error: %s", err),
		)
		return
	}

	data := userResourceModel{
		ID:                           types.Int64Value(user.ID),
		Password:                     types.StringValue(plan.Password.ValueString()),
		Email:                        types.StringValue(user.Email),
		Role:                         types.StringValue(user.Role),
		CanUseAuditLog:               types.BoolValue(user.CanUseAuditLog),
		IsRestrictedConnectionModify: types.BoolValue(user.IsRestrictedConnectionModify),
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, data)...)
}

func (r *userResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state userResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	user, err := r.client.GetUser(state.ID.ValueInt64())
	if err != nil {
		resp.Diagnostics.AddError(
			"Reading user",
			fmt.Sprintf("Unable to read user, got error: %s", err),
		)
		return
	}

	data := userResourceModel{
		ID:                           types.Int64Value(user.ID),
		Password:                     types.StringValue(state.Password.ValueString()),
		Email:                        types.StringValue(user.Email),
		Role:                         types.StringValue(user.Role),
		CanUseAuditLog:               types.BoolValue(user.CanUseAuditLog),
		IsRestrictedConnectionModify: types.BoolValue(user.IsRestrictedConnectionModify),
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, data)...)
}

func (r *userResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state userResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	input := client.UpdateUserInput{
		Role:                         plan.Role.ValueString(),
		CanUseAuditLog:               plan.CanUseAuditLog.ValueBool(),
		IsRestrictedConnectionModify: plan.IsRestrictedConnectionModify.ValueBool(),
	}

	user, err := r.client.UpdateUser(state.ID.ValueInt64(), &input)
	if err != nil {
		resp.Diagnostics.AddError(
			"Updating user",
			fmt.Sprintf("Unable to update user, got error: %s", err),
		)
		return
	}

	data := userResourceModel{
		ID:                           types.Int64Value(user.ID),
		Password:                     types.StringValue(state.Password.ValueString()),
		Email:                        types.StringValue(user.Email),
		Role:                         types.StringValue(user.Role),
		CanUseAuditLog:               types.BoolValue(user.CanUseAuditLog),
		IsRestrictedConnectionModify: types.BoolValue(user.IsRestrictedConnectionModify),
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, data)...)
}

func (r *userResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state userResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := r.client.DeleteUser(state.ID.ValueInt64())
	if err != nil {
		resp.Diagnostics.AddError(
			"Deleting user",
			fmt.Sprintf("Unable to delete user, got error: %s", err),
		)
		return
	}
}
