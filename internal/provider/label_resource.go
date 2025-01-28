package provider

import (
	"context"
	"fmt"
	"regexp"
	"strconv"
	"terraform-provider-trocco/internal/client"
	"terraform-provider-trocco/internal/provider/model"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ resource.Resource                = &labelResource{}
	_ resource.ResourceWithConfigure   = &labelResource{}
	_ resource.ResourceWithImportState = &labelResource{}
)

func NewLabelResource() resource.Resource {
	return &labelResource{}
}

type labelResource struct {
	client *client.TroccoClient
}

func (r *labelResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_label"
}

func (r *labelResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *labelResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Provides a TROCCO label resource.",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "The ID of the label.",
			},
			"name": schema.StringAttribute{
				Required: true,
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtLeast(1),
				},
				MarkdownDescription: "The name of the label. It must be at least 1 character.",
			},
			"description": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				MarkdownDescription: "The description of the label.",
			},
			"color": schema.StringAttribute{
				Required: true,
				Validators: []validator.String{
					stringvalidator.RegexMatches(
						regexp.MustCompile(`^#[0-9a-fA-F]{3}([0-9a-fA-F]{3})?$`),
						"must be in format #RRGGBB or #RGB",
					),
				},
				MarkdownDescription: "The color of the label. It must be in format #RRGGBB or #RGB.",
			},
		},
	}
}

func (r *labelResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan model.LabelModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	input := client.CreateLabelInput{
		Name:        plan.Name.ValueString(),
		Description: plan.Description.ValueStringPointer(),
		Color:       plan.Color.ValueString(),
	}

	label, err := r.client.CreateLabel(&input)
	if err != nil {
		resp.Diagnostics.AddError(
			"Creating label",
			fmt.Sprintf("Unable to create label, got error: %s", err),
		)
		return
	}

	data := model.LabelModel{
		ID:          types.Int64Value(label.ID),
		Name:        types.StringValue(label.Name),
		Description: types.StringPointerValue(label.Description),
		Color:       types.StringValue(label.Color),
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, data)...)
}

func (r *labelResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state model.LabelModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	label, err := r.client.GetLabel(state.ID.ValueInt64())
	if err != nil {
		resp.Diagnostics.AddError(
			"Reading label",
			fmt.Sprintf("Unable to read label, got error: %s", err),
		)
		return
	}

	data := model.LabelModel{
		ID:          types.Int64Value(label.ID),
		Name:        types.StringValue(label.Name),
		Description: types.StringPointerValue(label.Description),
		Color:       types.StringValue(label.Color),
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, data)...)
}

func (r *labelResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state model.LabelModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	input := client.UpdateLabelInput{
		Name:        plan.Name.ValueStringPointer(),
		Description: plan.Description.ValueStringPointer(),
		Color:       plan.Color.ValueStringPointer(),
	}

	label, err := r.client.UpdateLabel(state.ID.ValueInt64(), &input)
	if err != nil {
		resp.Diagnostics.AddError(
			"Updating label",
			fmt.Sprintf("Unable to update label, got error: %s", err),
		)
		return
	}

	data := model.LabelModel{
		ID:          types.Int64Value(label.ID),
		Name:        types.StringValue(label.Name),
		Description: types.StringPointerValue(label.Description),
		Color:       types.StringValue(label.Color),
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, data)...)
}

func (r *labelResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state model.LabelModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := r.client.DeleteLabel(state.ID.ValueInt64())
	if err != nil {
		resp.Diagnostics.AddError(
			"Deleting label",
			fmt.Sprintf("Unable to delete label, got error: %s", err),
		)
		return
	}
}

func (r *labelResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	id, err := strconv.ParseInt(req.ID, 10, 64)
	if err != nil {
		resp.Diagnostics.AddError(
			"Importing label",
			fmt.Sprintf("Unable to parse id, got error: %s", err),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), id)...)
}
