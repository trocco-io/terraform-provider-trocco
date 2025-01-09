package provider

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"terraform-provider-trocco/internal/client"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
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

var (
	_ resource.Resource                = &connectionResource{}
	_ resource.ResourceWithConfigure   = &connectionResource{}
	_ resource.ResourceWithImportState = &connectionResource{}
)

type connectionResourceModel struct {
	// Common Fields
	ConnectionType  types.String `tfsdk:"connection_type"`
	ID              types.Int64  `tfsdk:"id"`
	Name            types.String `tfsdk:"name"`
	Description     types.String `tfsdk:"description"`
	ResourceGroupID types.Int64  `tfsdk:"resource_group_id"`

	// BigQuery Fields
	ProjectID             types.String `tfsdk:"project_id"`
	ServiceAccountJSONKey types.String `tfsdk:"service_account_json_key"`

	// Snowflake Fields
	Host       types.String `tfsdk:"host"`
	UserName   types.String `tfsdk:"user_name"`
	Role       types.String `tfsdk:"role"`
	AuthMethod types.String `tfsdk:"auth_method"`
	Password   types.String `tfsdk:"password"`
	PrivateKey types.String `tfsdk:"private_key"`

	// GCS Fields
	ApplicationName     types.String `tfsdk:"application_name"`
	ServiceAccountEmail types.String `tfsdk:"service_account_email"`
}

func (m *connectionResourceModel) ToCreateConnectionInput() *client.CreateConnectionInput {
	return &client.CreateConnectionInput{
		// Common Fields
		Name:            m.Name.ValueString(),
		Description:     m.Description.ValueStringPointer(),
		ResourceGroupID: newNullableFromTerraformInt64(m.ResourceGroupID),

		// BigQuery Fields
		ProjectID:             m.ProjectID.ValueStringPointer(),
		ServiceAccountJSONKey: m.ServiceAccountJSONKey.ValueStringPointer(),

		// Snowflake Fields
		Host:       m.Host.ValueStringPointer(),
		UserName:   m.UserName.ValueStringPointer(),
		Role:       m.Role.ValueStringPointer(),
		AuthMethod: m.AuthMethod.ValueStringPointer(),
		Password:   m.Password.ValueStringPointer(),
		PrivateKey: m.PrivateKey.ValueStringPointer(),

		// GCS Fields
		ApplicationName:     m.ApplicationName.ValueStringPointer(),
		ServiceAccountEmail: m.ServiceAccountEmail.ValueStringPointer(),
	}
}

func (m *connectionResourceModel) ToUpdateConnectionInput() *client.UpdateConnectionInput {
	return &client.UpdateConnectionInput{
		// Common Fields
		Name:            m.Name.ValueStringPointer(),
		Description:     m.Description.ValueStringPointer(),
		ResourceGroupID: newNullableFromTerraformInt64(m.ResourceGroupID),

		// BigQuery Fields
		ProjectID:             m.ProjectID.ValueStringPointer(),
		ServiceAccountJSONKey: m.ServiceAccountJSONKey.ValueStringPointer(),

		// Snowflake Fields
		Host:       m.Host.ValueStringPointer(),
		UserName:   m.UserName.ValueStringPointer(),
		Role:       m.Role.ValueStringPointer(),
		AuthMethod: m.AuthMethod.ValueStringPointer(),
		Password:   m.Password.ValueStringPointer(),
		PrivateKey: m.PrivateKey.ValueStringPointer(),

		// GCS Fields
		ApplicationName:     m.ApplicationName.ValueStringPointer(),
		ServiceAccountEmail: m.ServiceAccountEmail.ValueStringPointer(),
	}
}

type connectionResource struct {
	client *client.TroccoClient
}

func NewConnectionResource() resource.Resource {
	return &connectionResource{}
}

func (r *connectionResource) Metadata(
	ctx context.Context,
	req resource.MetadataRequest,
	resp *resource.MetadataResponse,
) {
	resp.TypeName = fmt.Sprintf("%s_connection", req.ProviderTypeName)
}

func (r *connectionResource) Configure(
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

func (r *connectionResource) Schema(
	ctx context.Context,
	req resource.SchemaRequest,
	resp *resource.SchemaResponse,
) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Provides a TROCCO connection resource.",
		Attributes: map[string]schema.Attribute{
			// Common Fields
			"connection_type": schema.StringAttribute{
				MarkdownDescription: "The type of the connection. It must be one of `bigquery`, `snowflake` or `gcs`.",
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf("bigquery", "snowflake", "gcs"),
				},
			},
			"id": schema.Int64Attribute{
				MarkdownDescription: "The ID of the connection.",
				Computed:            true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
				Validators: []validator.Int64{
					int64validator.AtLeast(1),
				},
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "The name of the connection.",
				Required:            true,
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtMost(255),
				},
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "The description of the connection.",
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtLeast(1),
				},
			},
			"resource_group_id": schema.Int64Attribute{
				MarkdownDescription: "The ID of the resource group the connection belongs to.",
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.AtLeast(1),
				},
			},

			// BigQuery Fields
			"project_id": schema.StringAttribute{
				MarkdownDescription: "BigQuery, GCS: A GCP project ID.",
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtLeast(1),
				},
			},
			"service_account_json_key": schema.StringAttribute{
				MarkdownDescription: "BigQuery: A GCP service account key.",
				Optional:            true,
				Sensitive:           true,
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtLeast(1),
				},
			},

			// Snowflake Fields
			"host": schema.StringAttribute{
				MarkdownDescription: "Snowflake: The host of a Snowflake account.",
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtLeast(1),
				},
			},
			"user_name": schema.StringAttribute{
				MarkdownDescription: "Snowflake: The name of a Snowflake user.",
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtLeast(1),
				},
			},
			"role": schema.StringAttribute{
				MarkdownDescription: "Snowflake: A role attached to the Snowflake user.",
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtLeast(1),
				},
			},
			"auth_method": schema.StringAttribute{
				MarkdownDescription: "Snowflake: The authentication method for the Snowflake user. It must be one of `key_pair` or `user_password`.",
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("key_pair", "user_password"),
				},
			},
			"password": schema.StringAttribute{
				MarkdownDescription: "Snowflake: The password for the Snowflake user.",
				Optional:            true,
				Sensitive:           true,
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtLeast(1),
				},
			},
			"private_key": schema.StringAttribute{
				MarkdownDescription: "Snowflake: A private key for the Snowflake user.",
				Optional:            true,
				Sensitive:           true,
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtLeast(1),
				},
			},

			// GCS Fields
			"application_name": schema.StringAttribute{
				MarkdownDescription: "GCS: Application name.",
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtLeast(1),
				},
			},
			"service_account_email": schema.StringAttribute{
				MarkdownDescription: "GCS: A GCP service account email.",
				Optional:            true,
				Sensitive:           true,
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtLeast(1),
				},
			},
		},
	}
}

func (r *connectionResource) Create(
	ctx context.Context,
	req resource.CreateRequest,
	resp *resource.CreateResponse,
) {
	plan := &connectionResourceModel{}
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	connection, err := r.client.CreateConnection(
		plan.ConnectionType.ValueString(),
		plan.ToCreateConnectionInput(),
	)
	if err != nil {
		resp.Diagnostics.AddError(
			"Creating connection",
			fmt.Sprintf("Unable to create connection, got error: %s", err),
		)
		return
	}

	newState := connectionResourceModel{
		// Common Fields
		ConnectionType:  plan.ConnectionType,
		ID:              types.Int64Value(connection.ID),
		Name:            types.StringPointerValue(connection.Name),
		Description:     types.StringPointerValue(connection.Description),
		ResourceGroupID: types.Int64PointerValue(connection.ResourceGroupID),

		// BigQuery Fields
		ProjectID:             types.StringPointerValue(connection.ProjectID),
		ServiceAccountJSONKey: plan.ServiceAccountJSONKey,

		// Snowflake Fields
		Host:       types.StringPointerValue(connection.Host),
		UserName:   types.StringPointerValue(connection.UserName),
		Role:       types.StringPointerValue(connection.Role),
		AuthMethod: types.StringPointerValue(connection.AuthMethod),
		Password:   plan.Password,
		PrivateKey: plan.PrivateKey,

		// GCS Fields
		ApplicationName:     types.StringPointerValue(connection.ApplicationName),
		ServiceAccountEmail: plan.ServiceAccountEmail,
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}

func (r *connectionResource) Update(
	ctx context.Context,
	req resource.UpdateRequest,
	resp *resource.UpdateResponse,
) {
	state := &connectionResourceModel{}
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	plan := &connectionResourceModel{}
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	connection, err := r.client.UpdateConnection(
		state.ConnectionType.ValueString(),
		state.ID.ValueInt64(),
		plan.ToUpdateConnectionInput(),
	)
	if err != nil {
		resp.Diagnostics.AddError(
			"Updating connection",
			fmt.Sprintf("Unable to update connection, got error: %s", err),
		)
		return
	}

	newState := connectionResourceModel{
		// Common Fields
		ConnectionType:  state.ConnectionType,
		ID:              types.Int64Value(connection.ID),
		Name:            types.StringPointerValue(connection.Name),
		Description:     types.StringPointerValue(connection.Description),
		ResourceGroupID: types.Int64PointerValue(connection.ResourceGroupID),

		// BigQuery Fields
		ProjectID:             types.StringPointerValue(connection.ProjectID),
		ServiceAccountJSONKey: plan.ServiceAccountJSONKey,

		// Snowflake Fields
		Host:       types.StringPointerValue(connection.Host),
		UserName:   types.StringPointerValue(connection.UserName),
		Role:       types.StringPointerValue(connection.Role),
		AuthMethod: types.StringPointerValue(connection.AuthMethod),
		Password:   plan.Password,
		PrivateKey: plan.PrivateKey,

		// GCS Fields
		ApplicationName:     types.StringPointerValue(connection.ApplicationName),
		ServiceAccountEmail: plan.ServiceAccountEmail,
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}

func (r *connectionResource) Read(
	ctx context.Context,
	req resource.ReadRequest,
	resp *resource.ReadResponse,
) {
	state := &connectionResourceModel{}
	resp.Diagnostics.Append(req.State.Get(ctx, state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	connection, err := r.client.GetConnection(
		state.ConnectionType.ValueString(),
		state.ID.ValueInt64(),
	)
	if err != nil {
		resp.Diagnostics.AddError(
			"Reading connection",
			fmt.Sprintf("Unable to read connection, got error: %s", err),
		)
		return
	}

	newState := connectionResourceModel{
		// Common Fields
		ConnectionType:  state.ConnectionType,
		ID:              types.Int64Value(connection.ID),
		Name:            types.StringPointerValue(connection.Name),
		Description:     types.StringPointerValue(connection.Description),
		ResourceGroupID: types.Int64PointerValue(connection.ResourceGroupID),

		// BigQuery Fields
		ProjectID:             types.StringPointerValue(connection.ProjectID),
		ServiceAccountJSONKey: state.ServiceAccountJSONKey,

		// Snowflake Fields
		Host:       types.StringPointerValue(connection.Host),
		UserName:   types.StringPointerValue(connection.UserName),
		Role:       types.StringPointerValue(connection.Role),
		AuthMethod: types.StringPointerValue(connection.AuthMethod),
		Password:   state.Password,
		PrivateKey: state.PrivateKey,

		// GCS Fields
		ApplicationName:     types.StringPointerValue(connection.ApplicationName),
		ServiceAccountEmail: state.ServiceAccountEmail,
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}

func (r *connectionResource) Delete(
	ctx context.Context,
	req resource.DeleteRequest,
	resp *resource.DeleteResponse,
) {
	s := &connectionResourceModel{}
	resp.Diagnostics.Append(req.State.Get(ctx, s)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if err := r.client.DeleteConnection(
		s.ConnectionType.ValueString(),
		s.ID.ValueInt64(),
	); err != nil {
		resp.Diagnostics.AddError(
			"Deleting connection",
			fmt.Sprintf("Unable to delete connection, got error: %s", err),
		)
		return
	}
}

func (r *connectionResource) ImportState(
	ctx context.Context,
	req resource.ImportStateRequest,
	resp *resource.ImportStateResponse,
) {
	// We must write custom logic because the Read method requires two attributes to refresh.
	// For more information, see https://developer.hashicorp.com/terraform/plugin/framework/resources/import#multiple-attributes.

	idParts := strings.Split(req.ID, ",")
	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		resp.Diagnostics.AddError(
			"Importing connection",
			fmt.Sprintf("Expected import identifier with format: connection_type,connection_id. Got: %q", req.ID),
		)
		return
	}

	connectionType := idParts[0]

	connectionID, err := strconv.ParseInt(idParts[1], 10, 64)
	if err != nil {
		resp.Diagnostics.AddError(
			"Importing connection",
			fmt.Sprintf("Failed to parse connection ID: %s", err),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("connection_type"), connectionType)...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), connectionID)...)
}
