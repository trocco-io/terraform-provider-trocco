package provider

import (
	"context"
	"fmt"
	"regexp"
	"strconv"
	"terraform-provider-trocco/internal/client"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ resource.Resource                = &userResource{}
	_ resource.ResourceWithConfigure   = &userResource{}
	_ resource.ResourceWithImportState = &userResource{}
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
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Validators: []validator.String{
					stringvalidator.RegexMatches(
						regexp.MustCompile(`^[^@\s]+@[^@\s]+$`),
						"invalid email address",
					),
				},
				MarkdownDescription: "The email of the user.",
			},
			"password": schema.StringAttribute{
				Optional:  true,
				Computed:  true,
				Sensitive: true,
				PlanModifiers: []planmodifier.String{
					&IgnoreChangesPlanModifier{},
					&RequiredOnCreatePlanModifier{"password"},
				},
				Validators: []validator.String{
					// see: https://documents.trocco.io/docs/password-policy
					stringvalidator.LengthBetween(8, 128),
					stringvalidator.All(
						stringvalidator.RegexMatches(regexp.MustCompile(`[a-zA-Z]`), "must contain at least one letter"),
						stringvalidator.RegexMatches(regexp.MustCompile(`[0-9]`), "must contain at least one number"),
					),
				},
				MarkdownDescription: "The password of the user. It must be at least 8 characters long and contain at least one letter and one number. It is required when creating a new user but optional during updates.",
			},
			"role": schema.StringAttribute{
				Required: true,
				Validators: []validator.String{
					stringvalidator.OneOf("super_admin", "admin", "member"),
				},
				MarkdownDescription: "The role of the user. Valid value is `super_admin`, `admin`, or `member`.",
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
		Email:                        plan.Email.ValueString(),
		Password:                     plan.Password.ValueString(),
		Role:                         plan.Role.ValueString(),
		CanUseAuditLog:               plan.CanUseAuditLog.ValueBoolPointer(),
		IsRestrictedConnectionModify: plan.IsRestrictedConnectionModify.ValueBoolPointer(),
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
		Role:                         plan.Role.ValueStringPointer(),
		CanUseAuditLog:               plan.CanUseAuditLog.ValueBoolPointer(),
		IsRestrictedConnectionModify: plan.IsRestrictedConnectionModify.ValueBoolPointer(),
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

func (r *userResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	id, err := strconv.ParseInt(req.ID, 10, 64)
	if err != nil {
		resp.Diagnostics.AddError(
			"Importing user",
			fmt.Sprintf("Unable to parse id, got error: %s", err),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), id)...)
}
