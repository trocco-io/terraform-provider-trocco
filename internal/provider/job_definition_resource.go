package provider

import (
	"context"
	"fmt"
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
	"strconv"
	"terraform-provider-trocco/internal/client"
	"terraform-provider-trocco/internal/client/parameters"
	job_definitions2 "terraform-provider-trocco/internal/client/parameters/job_definitions"
	filter2 "terraform-provider-trocco/internal/client/parameters/job_definitions/filter"
	"terraform-provider-trocco/internal/provider/models"
	"terraform-provider-trocco/internal/provider/models/job_definitions"
	"terraform-provider-trocco/internal/provider/models/job_definitions/filter"
	"terraform-provider-trocco/internal/provider/schema/job_definition"
	"terraform-provider-trocco/internal/provider/schema/job_definition/filters"
)

var (
	_ resource.Resource                = &jobDefinitionResource{}
	_ resource.ResourceWithConfigure   = &jobDefinitionResource{}
	_ resource.ResourceWithImportState = &jobDefinitionResource{}
)

func NewJobDefinitionResource() resource.Resource {
	return &jobDefinitionResource{}
}

type jobDefinitionResource struct {
	client *client.TroccoClient
}

func (r *jobDefinitionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	id, err := strconv.ParseInt(req.ID, 10, 64)
	if err != nil {
		resp.Diagnostics.AddError(
			"Importing job definition",
			fmt.Sprintf("Unable to parse id, got error: %s", err),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), id)...)
}

func (r *jobDefinitionResource) Metadata(
	ctx context.Context,
	req resource.MetadataRequest,
	resp *resource.MetadataResponse,
) {
	resp.TypeName = req.ProviderTypeName + "_job_definition"
}

func (r *jobDefinitionResource) Configure(
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

func (r *jobDefinitionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Provides a TROCCO job definitions.",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "The ID of the job definition",
			},
			"name": schema.StringAttribute{
				Required: true,
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtMost(255),
				},
				MarkdownDescription: "Name of the job definition. It must be less than 256 characters",
			},
			"description": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Description of the job definition. It must be at least 1 character",
			},
			"resource_group_id": schema.Int64Attribute{
				Optional: true,
				Validators: []validator.Int64{
					int64validator.AtLeast(1),
				},
				MarkdownDescription: "ID of the resource group to which the job definition belongs",
			},
			"resource_enhancement": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Resource size to be used when executing the job. If not specified, the resource size specified in the transfer settings is applied. The value that can be specified varies depending on the connector. (This parameter is available only in the Professional plan.",
				Validators: []validator.String{
					stringvalidator.OneOf("medium", "custom_spec", "large", "xlarge"),
				},
			},
			"is_runnable_concurrently": schema.BoolAttribute{
				Required:            true,
				MarkdownDescription: "Specifies whether or not to run a job if another job with the same job definition is running at the time the job is run",
			},
			"retry_limit": schema.Int64Attribute{
				Required: true,
				Validators: []validator.Int64{
					int64validator.Between(0, 10),
				},
				MarkdownDescription: "Maximum number of retries. if set 0, the job will not be retried",
			},
			"input_option_type": schema.StringAttribute{
				Required: true,
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtLeast(1),
				},
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: "Input option type.",
			},
			"input_option": job_definition.InputOptionSchema(),
			"output_option_type": schema.StringAttribute{
				Required: true,
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtLeast(1),
				},
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: "Output option type.",
			},
			"output_option":               job_definition.OutputOptionSchema(),
			"filter_columns":              filters.FilterColumnsSchema(),
			"filter_rows":                 filters.FilterRowsSchema(),
			"filter_masks":                filters.FilterMasksSchema(),
			"filter_add_time":             filters.FilterAddTimeSchema(),
			"filter_gsub":                 filters.FilterGsubSchema(),
			"filter_string_transforms":    filters.FilterStringTransformsSchema(),
			"filter_hashes":               filters.FilterHashesSchema(),
			"filter_unixtime_conversions": filters.FilterUnixTimeConversionsSchema(),
			"schedules":                   job_definition.SchedulesSchema(),
			"notifications":               job_definition.NotificationsSchema(),
			"labels":                      job_definition.LabelsSchema(),
		},
	}
}

type jobDefinitionResourceModel struct {
	ID                        types.Int64                                 `tfsdk:"id"`
	Name                      types.String                                `tfsdk:"name"`
	Description               types.String                                `tfsdk:"description"`
	ResourceGroupID           types.Int64                                 `tfsdk:"resource_group_id"`
	IsRunnableConcurrently    types.Bool                                  `tfsdk:"is_runnable_concurrently"`
	RetryLimit                types.Int64                                 `tfsdk:"retry_limit"`
	ResourceEnhancement       types.String                                `tfsdk:"resource_enhancement"`
	InputOptionType           types.String                                `tfsdk:"input_option_type"`
	InputOption               *job_definitions.InputOption                `tfsdk:"input_option"`
	OutputOptionType          types.String                                `tfsdk:"output_option_type"`
	OutputOption              *job_definitions.OutputOption               `tfsdk:"output_option"`
	FilterColumns             []filter.FilterColumn                       `tfsdk:"filter_columns"`
	FilterRows                *filter.FilterRows                          `tfsdk:"filter_rows"`
	FilterMasks               []filter.FilterMask                         `tfsdk:"filter_masks"`
	FilterAddTime             *filter.FilterAddTime                       `tfsdk:"filter_add_time"`
	FilterGsub                []filter.FilterGsub                         `tfsdk:"filter_gsub"`
	FilterStringTransforms    []filter.FilterStringTransform              `tfsdk:"filter_string_transforms"`
	FilterHashes              []filter.FilterHash                         `tfsdk:"filter_hashes"`
	FilterUnixTimeConversions []filter.FilterUnixTimeConversion           `tfsdk:"filter_unixtime_conversions"`
	Notifications             []job_definitions.JobDefinitionNotification `tfsdk:"notifications"`
	Schedules                 []models.Schedule                           `tfsdk:"schedules"`
	Labels                    []models.LabelModel                         `tfsdk:"labels"`
}

func (model *jobDefinitionResourceModel) ToCreateJobDefinitionInput() *client.CreateJobDefinitionInput {
	var labels []string
	if model.Labels != nil {
		for _, l := range model.Labels {
			labels = append(labels, l.Name.ValueString())
		}
	}
	var notifications []job_definitions2.JobDefinitionNotificationInput
	if model.Notifications != nil {
		for _, n := range model.Notifications {
			notifications = append(notifications, n.ToInput())
		}
	}
	var schedules []parameters.ScheduleInput
	if model.Schedules != nil {
		for _, s := range model.Schedules {
			schedules = append(schedules, s.ToInput())
		}
	}

	var filterColumns []filter2.FilterColumnInput
	if model.FilterColumns != nil {
		for _, f := range model.FilterColumns {
			filterColumns = append(filterColumns, f.ToInput())
		}
	}

	var filterMasks []filter2.FilterMaskInput
	for _, f := range model.FilterMasks {
		filterMasks = append(filterMasks, f.ToInput())
	}

	var filterGsub []filter2.FilterGsubInput
	for _, f := range model.FilterGsub {
		filterGsub = append(filterGsub, f.ToInput())
	}

	var filterStringTransforms []filter2.FilterStringTransformInput
	for _, f := range model.FilterStringTransforms {
		filterStringTransforms = append(filterStringTransforms, f.ToInput())
	}

	var filterHashes []filter2.FilterHashInput
	for _, f := range model.FilterHashes {
		filterHashes = append(filterHashes, f.ToInput())
	}

	var filterUnixTimeconversions []filter2.FilterUnixTimeConversionInput
	for _, f := range model.FilterUnixTimeConversions {
		filterUnixTimeconversions = append(filterUnixTimeconversions, f.ToInput())
	}

	return &client.CreateJobDefinitionInput{
		Name:                      model.Name.ValueString(),
		Description:               models.NewNullableString(model.Description),
		ResourceGroupID:           models.NewNullableInt64(model.ResourceGroupID),
		IsRunnableConcurrently:    model.IsRunnableConcurrently.ValueBool(),
		RetryLimit:                model.RetryLimit.ValueInt64(),
		ResourceEnhancement:       model.ResourceEnhancement.ValueStringPointer(),
		FilterColumns:             filterColumns,
		FilterRows:                models.WrapObject(model.FilterRows.ToInput()),
		FilterMasks:               filterMasks,
		FilterAddTime:             models.WrapObject(model.FilterAddTime.ToInput()),
		FilterGsub:                filterGsub,
		FilterStringTransforms:    filterStringTransforms,
		FilterHashes:              filterHashes,
		FilterUnixTimeConversions: filterUnixTimeconversions,
		InputOptionType:           model.InputOptionType.ValueString(),
		InputOption:               model.InputOption.ToInput(),
		OutputOptionType:          model.OutputOptionType.ValueString(),
		OutputOption:              model.OutputOption.ToInput(),
		Labels:                    labels,
		Schedules:                 schedules,
		Notifications:             notifications,
	}
}

func (r *jobDefinitionResource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	state := &jobDefinitionResourceModel{}
	response.Diagnostics.Append(request.State.Get(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}

	plan := &jobDefinitionResourceModel{}
	response.Diagnostics.Append(request.Plan.Get(ctx, &plan)...)
	if response.Diagnostics.HasError() {
		return
	}

	jobDefinition, err := r.client.UpdateJobDefinition(
		state.ID.ValueInt64(),
		plan.ToUpdateJobDefinitionInput(),
	)
	if err != nil {
		response.Diagnostics.AddError(
			"Updating job definition",
			fmt.Sprintf("Unable to update job definition, got error: %s", err),
		)
		return
	}

	newState := jobDefinitionResourceModel{
		ID:                        types.Int64Value(jobDefinition.ID),
		Name:                      types.StringValue(jobDefinition.Name),
		Description:               types.StringPointerValue(jobDefinition.Description),
		ResourceGroupID:           types.Int64PointerValue(jobDefinition.ResourceGroupID),
		IsRunnableConcurrently:    types.BoolPointerValue(jobDefinition.IsRunnableConcurrently),
		RetryLimit:                types.Int64Value(jobDefinition.RetryLimit),
		ResourceEnhancement:       types.StringPointerValue(jobDefinition.ResourceEnhancement),
		InputOptionType:           types.StringValue(jobDefinition.InputOptionType),
		InputOption:               job_definitions.NewInputOption(jobDefinition.InputOption),
		OutputOptionType:          types.StringValue(jobDefinition.OutputOptionType),
		OutputOption:              job_definitions.NewOutputOption(jobDefinition.OutputOption),
		FilterColumns:             filter.NewFilterColumns(jobDefinition.FilterColumns),
		FilterRows:                filter.NewFilterRows(jobDefinition.FilterRows),
		FilterMasks:               filter.NewFilterMasks(jobDefinition.FilterMasks),
		FilterAddTime:             filter.NewFilterAddTime(jobDefinition.FilterAddTime),
		FilterGsub:                filter.NewFilterGsub(jobDefinition.FilterGsub),
		FilterStringTransforms:    filter.NewFilterStringTransforms(jobDefinition.FilterStringTransforms),
		FilterHashes:              filter.NewFilterHashes(jobDefinition.FilterHashes),
		FilterUnixTimeConversions: filter.NewFilterUnixTimeConversions(jobDefinition.FilterUnixTimeConversions),
		Notifications:             job_definitions.NewJobDefinitionNotifications(jobDefinition.Notifications),
		Schedules:                 models.NewSchedules(jobDefinition.Schedules),
		Labels:                    models.NewLabels(jobDefinition.Labels),
	}
	response.Diagnostics.Append(response.State.Set(ctx, newState)...)
}

func (model *jobDefinitionResourceModel) ToUpdateJobDefinitionInput() *client.UpdateJobDefinitionInput {
	labels := []string{}
	if model.Labels != nil {
		for _, l := range model.Labels {
			labels = append(labels, l.Name.ValueString())
		}
	}
	notifications := []job_definitions2.JobDefinitionNotificationInput{}
	if model.Notifications != nil {
		for _, n := range model.Notifications {
			notifications = append(notifications, n.ToInput())
		}
	}

	schedules := []parameters.ScheduleInput{}
	if model.Schedules != nil {
		for _, s := range model.Schedules {
			schedules = append(schedules, s.ToInput())
		}
	}

	filterColumns := []filter2.FilterColumnInput{}
	if model.FilterColumns != nil {
		for _, f := range model.FilterColumns {
			filterColumns = append(filterColumns, f.ToInput())
		}
	}

	filterMasks := []filter2.FilterMaskInput{}
	for _, f := range model.FilterMasks {
		filterMasks = append(filterMasks, f.ToInput())
	}

	filterGsub := []filter2.FilterGsubInput{}
	for _, f := range model.FilterGsub {
		filterGsub = append(filterGsub, f.ToInput())
	}

	filterStringTransforms := []filter2.FilterStringTransformInput{}
	for _, f := range model.FilterStringTransforms {
		filterStringTransforms = append(filterStringTransforms, f.ToInput())
	}

	filterHashes := []filter2.FilterHashInput{}
	for _, f := range model.FilterHashes {
		filterHashes = append(filterHashes, f.ToInput())
	}

	filterUnixTimeconversions := []filter2.FilterUnixTimeConversionInput{}
	for _, f := range model.FilterUnixTimeConversions {
		filterUnixTimeconversions = append(filterUnixTimeconversions, f.ToInput())
	}

	return &client.UpdateJobDefinitionInput{
		Name:                      model.Name.ValueStringPointer(),
		Description:               models.NewNullableString(model.Description),
		ResourceGroupID:           models.NewNullableInt64(model.ResourceGroupID),
		IsRunnableConcurrently:    model.IsRunnableConcurrently.ValueBoolPointer(),
		RetryLimit:                model.RetryLimit.ValueInt64Pointer(),
		ResourceEnhancement:       model.ResourceEnhancement.ValueStringPointer(),
		FilterColumns:             &filterColumns,
		FilterRows:                models.WrapObject(model.FilterRows.ToInput()),
		FilterMasks:               &filterMasks,
		FilterAddTime:             models.WrapObject(model.FilterAddTime.ToInput()),
		FilterGsub:                &filterGsub,
		FilterStringTransforms:    &filterStringTransforms,
		FilterHashes:              &filterHashes,
		FilterUnixTimeConversions: &filterUnixTimeconversions,
		InputOption:               model.InputOption.ToUpdateInput(),
		OutputOption:              model.OutputOption.ToUpdateInput(),
		Labels:                    &labels,
		Schedules:                 &schedules,
		Notifications:             &notifications,
	}
}

func (r *jobDefinitionResource) Create(
	ctx context.Context,
	req resource.CreateRequest,
	resp *resource.CreateResponse,
) {
	plan := &jobDefinitionResourceModel{}
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	jobDefinition, err := r.client.CreateJobDefinition(
		plan.ToCreateJobDefinitionInput(),
	)
	if err != nil {
		resp.Diagnostics.AddError(
			"Creating job definition",
			fmt.Sprintf("Unable to create job definition, got error: %s", err),
		)
		return
	}

	newState := jobDefinitionResourceModel{
		ID:                        types.Int64Value(jobDefinition.ID),
		Name:                      types.StringValue(jobDefinition.Name),
		Description:               types.StringPointerValue(jobDefinition.Description),
		ResourceGroupID:           types.Int64PointerValue(jobDefinition.ResourceGroupID),
		IsRunnableConcurrently:    types.BoolPointerValue(jobDefinition.IsRunnableConcurrently),
		RetryLimit:                types.Int64Value(jobDefinition.RetryLimit),
		ResourceEnhancement:       types.StringPointerValue(jobDefinition.ResourceEnhancement),
		InputOptionType:           types.StringValue(jobDefinition.InputOptionType),
		InputOption:               job_definitions.NewInputOption(jobDefinition.InputOption),
		OutputOptionType:          types.StringValue(jobDefinition.OutputOptionType),
		OutputOption:              job_definitions.NewOutputOption(jobDefinition.OutputOption),
		FilterColumns:             filter.NewFilterColumns(jobDefinition.FilterColumns),
		FilterRows:                filter.NewFilterRows(jobDefinition.FilterRows),
		FilterMasks:               filter.NewFilterMasks(jobDefinition.FilterMasks),
		FilterAddTime:             filter.NewFilterAddTime(jobDefinition.FilterAddTime),
		FilterGsub:                filter.NewFilterGsub(jobDefinition.FilterGsub),
		FilterStringTransforms:    filter.NewFilterStringTransforms(jobDefinition.FilterStringTransforms),
		FilterHashes:              filter.NewFilterHashes(jobDefinition.FilterHashes),
		FilterUnixTimeConversions: filter.NewFilterUnixTimeConversions(jobDefinition.FilterUnixTimeConversions),
		Notifications:             job_definitions.NewJobDefinitionNotifications(jobDefinition.Notifications),
		Schedules:                 models.NewSchedules(jobDefinition.Schedules),
		Labels:                    models.NewLabels(jobDefinition.Labels),
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}

func (r *jobDefinitionResource) Read(
	ctx context.Context,
	req resource.ReadRequest,
	resp *resource.ReadResponse,
) {
	state := &jobDefinitionResourceModel{}
	resp.Diagnostics.Append(req.State.Get(ctx, state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	jobDefinition, err := r.client.GetJobDefinition(state.ID.ValueInt64())
	if err != nil {
		resp.Diagnostics.AddError(
			"Reading job definition",
			fmt.Sprintf("Unable to read job definition, got error: %s", err),
		)
		return
	}

	newState := jobDefinitionResourceModel{
		ID:                        types.Int64Value(jobDefinition.ID),
		Name:                      types.StringValue(jobDefinition.Name),
		Description:               types.StringPointerValue(jobDefinition.Description),
		ResourceGroupID:           types.Int64PointerValue(jobDefinition.ResourceGroupID),
		IsRunnableConcurrently:    types.BoolPointerValue(jobDefinition.IsRunnableConcurrently),
		RetryLimit:                types.Int64Value(jobDefinition.RetryLimit),
		ResourceEnhancement:       types.StringPointerValue(jobDefinition.ResourceEnhancement),
		InputOptionType:           types.StringValue(jobDefinition.InputOptionType),
		InputOption:               job_definitions.NewInputOption(jobDefinition.InputOption),
		OutputOptionType:          types.StringValue(jobDefinition.OutputOptionType),
		OutputOption:              job_definitions.NewOutputOption(jobDefinition.OutputOption),
		FilterColumns:             filter.NewFilterColumns(jobDefinition.FilterColumns),
		FilterRows:                filter.NewFilterRows(jobDefinition.FilterRows),
		FilterMasks:               filter.NewFilterMasks(jobDefinition.FilterMasks),
		FilterAddTime:             filter.NewFilterAddTime(jobDefinition.FilterAddTime),
		FilterGsub:                filter.NewFilterGsub(jobDefinition.FilterGsub),
		FilterStringTransforms:    filter.NewFilterStringTransforms(jobDefinition.FilterStringTransforms),
		FilterHashes:              filter.NewFilterHashes(jobDefinition.FilterHashes),
		FilterUnixTimeConversions: filter.NewFilterUnixTimeConversions(jobDefinition.FilterUnixTimeConversions),
		Notifications:             job_definitions.NewJobDefinitionNotifications(jobDefinition.Notifications),
		Schedules:                 models.NewSchedules(jobDefinition.Schedules),
		Labels:                    models.NewLabels(jobDefinition.Labels),
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}

func (r *jobDefinitionResource) Delete(
	ctx context.Context,
	req resource.DeleteRequest,
	resp *resource.DeleteResponse,
) {
	s := &jobDefinitionResourceModel{}
	resp.Diagnostics.Append(req.State.Get(ctx, s)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if err := r.client.DeleteJobDefinition(s.ID.ValueInt64()); err != nil {
		resp.Diagnostics.AddError(
			"Deleting job definition",
			fmt.Sprintf("Unable to delete job definition, got error: %s", err),
		)
		return
	}
}
