package provider

import (
	"context"
	"fmt"
	"strconv"
	"terraform-provider-trocco/internal/client"
	"terraform-provider-trocco/internal/client/parameter"
	params "terraform-provider-trocco/internal/client/parameter/job_definition"
	filterParameters "terraform-provider-trocco/internal/client/parameter/job_definition/filter"
	"terraform-provider-trocco/internal/provider/model"
	job_definitions "terraform-provider-trocco/internal/provider/model/job_definition"
	"terraform-provider-trocco/internal/provider/model/job_definition/filter"
	input_options "terraform-provider-trocco/internal/provider/model/job_definition/input_option"
	"terraform-provider-trocco/internal/provider/schema/job_definition"
	"terraform-provider-trocco/internal/provider/schema/job_definition/filters"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
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
				MarkdownDescription: "Description of the job definition.",
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
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
				MarkdownDescription: "Specifies whether or not to run a job if another job with the same job definition is running at the time the job is run",
			},
			"retry_limit": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				Default:  int64default.StaticInt64(0),
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
	ID                        types.Int64                   `tfsdk:"id"`
	Name                      types.String                  `tfsdk:"name"`
	Description               types.String                  `tfsdk:"description"`
	ResourceGroupID           types.Int64                   `tfsdk:"resource_group_id"`
	IsRunnableConcurrently    types.Bool                    `tfsdk:"is_runnable_concurrently"`
	RetryLimit                types.Int64                   `tfsdk:"retry_limit"`
	ResourceEnhancement       types.String                  `tfsdk:"resource_enhancement"`
	InputOptionType           types.String                  `tfsdk:"input_option_type"`
	InputOption               *job_definitions.InputOption  `tfsdk:"input_option"`
	OutputOptionType          types.String                  `tfsdk:"output_option_type"`
	OutputOption              *job_definitions.OutputOption `tfsdk:"output_option"`
	FilterColumns             types.List                    `tfsdk:"filter_columns"`
	FilterRows                *filter.FilterRows            `tfsdk:"filter_rows"`
	FilterMasks               types.List                    `tfsdk:"filter_masks"`
	FilterAddTime             *filter.FilterAddTime         `tfsdk:"filter_add_time"`
	FilterGsub                types.List                    `tfsdk:"filter_gsub"`
	FilterStringTransforms    types.List                    `tfsdk:"filter_string_transforms"`
	FilterHashes              types.List                    `tfsdk:"filter_hashes"`
	FilterUnixTimeConversions types.List                    `tfsdk:"filter_unixtime_conversions"`
	Notifications             types.Set                     `tfsdk:"notifications"`
	Schedules                 types.Set                     `tfsdk:"schedules"`
	Labels                    types.Set                     `tfsdk:"labels"`
}

func (m *jobDefinitionResourceModel) ToCreateJobDefinitionInput(ctx context.Context, resp *resource.CreateResponse) (*client.CreateJobDefinitionInput, diag.Diagnostics) {
	var labels []string
	if !m.Labels.IsNull() && !m.Labels.IsUnknown() {
		var labelValues []job_definitions.Label
		diags := m.Labels.ElementsAs(ctx, &labelValues, false)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return nil, resp.Diagnostics
		}

		for _, l := range labelValues {
			labels = append(labels, l.Name.ValueString())
		}
	}

	var notifications []params.JobDefinitionNotificationInput
	if !m.Notifications.IsNull() && !m.Notifications.IsUnknown() {
		var notificationValues []job_definitions.JobDefinitionNotification
		diags := m.Notifications.ElementsAs(ctx, &notificationValues, false)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return nil, resp.Diagnostics
		}
		for _, n := range notificationValues {
			notifications = append(notifications, n.ToInput())
		}
	}

	var schedules []parameter.ScheduleInput
	if !m.Schedules.IsNull() && !m.Schedules.IsUnknown() {
		var scheduleValues []model.Schedule
		diags := m.Schedules.ElementsAs(ctx, &scheduleValues, false)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return nil, resp.Diagnostics
		}
		for _, s := range scheduleValues {
			schedules = append(schedules, s.ToInput())
		}
	}

	var filterColumns []filterParameters.FilterColumnInput
	if !m.FilterColumns.IsNull() && !m.FilterColumns.IsUnknown() {
		var filterColumnValues []filter.FilterColumn
		diags := m.FilterColumns.ElementsAs(ctx, &filterColumnValues, false)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return nil, resp.Diagnostics
		}
		for _, f := range filterColumnValues {
			filterColumns = append(filterColumns, f.ToInput(ctx))
		}
	}

	var filterMasks []filterParameters.FilterMaskInput
	if !m.FilterMasks.IsNull() && !m.FilterMasks.IsUnknown() {
		var filterMaskValues []filter.FilterMask
		diags := m.FilterMasks.ElementsAs(ctx, &filterMaskValues, false)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return nil, resp.Diagnostics
		}
		for _, f := range filterMaskValues {
			filterMasks = append(filterMasks, f.ToInput())
		}
	}

	var filterGsub []filterParameters.FilterGsubInput
	if !m.FilterGsub.IsNull() && !m.FilterGsub.IsUnknown() {
		var filterGsubValues []filter.FilterGsub
		diags := m.FilterGsub.ElementsAs(ctx, &filterGsubValues, false)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return nil, resp.Diagnostics
		}
		for _, f := range filterGsubValues {
			filterGsub = append(filterGsub, f.ToInput())
		}
	}

	var filterStringTransforms []filterParameters.FilterStringTransformInput
	if !m.FilterStringTransforms.IsNull() && !m.FilterStringTransforms.IsUnknown() {
		var filterStringTransformValues []filter.FilterStringTransform
		diags := m.FilterStringTransforms.ElementsAs(ctx, &filterStringTransformValues, false)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return nil, resp.Diagnostics
		}
		for _, f := range filterStringTransformValues {
			filterStringTransforms = append(filterStringTransforms, f.ToInput())
		}
	}

	var filterHashes []filterParameters.FilterHashInput
	if !m.FilterHashes.IsNull() && !m.FilterHashes.IsUnknown() {
		var filterHashValues []filter.FilterHash
		diags := m.FilterHashes.ElementsAs(ctx, &filterHashValues, false)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return nil, resp.Diagnostics
		}
		for _, f := range filterHashValues {
			filterHashes = append(filterHashes, f.ToInput())
		}
	}

	var filterUnixTimeconversions []filterParameters.FilterUnixTimeConversionInput
	if !m.FilterUnixTimeConversions.IsNull() && !m.FilterUnixTimeConversions.IsUnknown() {
		var filterUnixTimeConversionValues []filter.FilterUnixTimeConversion
		diags := m.FilterUnixTimeConversions.ElementsAs(ctx, &filterUnixTimeConversionValues, false)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return nil, resp.Diagnostics
		}
		for _, f := range filterUnixTimeConversionValues {
			filterUnixTimeconversions = append(filterUnixTimeconversions, f.ToInput())
		}
	}

	var diags diag.Diagnostics
	inputOption, d := m.InputOption.ToInput(ctx)
	diags.Append(d...)
	return &client.CreateJobDefinitionInput{
		Name:                      m.Name.ValueString(),
		Description:               model.NewNullableString(m.Description),
		ResourceGroupID:           model.NewNullableInt64(m.ResourceGroupID),
		IsRunnableConcurrently:    m.IsRunnableConcurrently.ValueBool(),
		RetryLimit:                m.RetryLimit.ValueInt64(),
		ResourceEnhancement:       m.ResourceEnhancement.ValueStringPointer(),
		FilterColumns:             filterColumns,
		FilterRows:                model.WrapObject(m.FilterRows.ToInput(ctx)),
		FilterMasks:               filterMasks,
		FilterAddTime:             model.WrapObject(m.FilterAddTime.ToInput()),
		FilterGsub:                filterGsub,
		FilterStringTransforms:    filterStringTransforms,
		FilterHashes:              filterHashes,
		FilterUnixTimeConversions: filterUnixTimeconversions,
		InputOptionType:           m.InputOptionType.ValueString(),
		InputOption:               inputOption,
		OutputOptionType:          m.OutputOptionType.ValueString(),
		OutputOption:              m.OutputOption.ToInput(ctx),
		Labels:                    labels,
		Schedules:                 schedules,
		Notifications:             notifications,
	}, diags
}

func (r *jobDefinitionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	state := &jobDefinitionResourceModel{}
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	plan := &jobDefinitionResourceModel{}
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	jobDefinitionInput, diags := plan.ToUpdateJobDefinitionInput(ctx, resp)
	if diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}

	jobDefinition, err := r.client.UpdateJobDefinition(
		state.ID.ValueInt64(),
		jobDefinitionInput,
	)
	if err != nil {
		resp.Diagnostics.AddError(
			"Updating job definition",
			fmt.Sprintf("Unable to update job definition, got error: %s", err),
		)
		return
	}

	inputOption, diags := job_definitions.NewInputOption(ctx, jobDefinition.InputOption, plan.InputOption)
	if diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}

	newState := jobDefinitionResourceModel{
		ID:                     types.Int64Value(jobDefinition.ID),
		Name:                   types.StringValue(jobDefinition.Name),
		Description:            types.StringPointerValue(jobDefinition.Description),
		ResourceGroupID:        types.Int64PointerValue(jobDefinition.ResourceGroupID),
		IsRunnableConcurrently: types.BoolPointerValue(jobDefinition.IsRunnableConcurrently),
		RetryLimit:             types.Int64Value(jobDefinition.RetryLimit),
		ResourceEnhancement:    types.StringPointerValue(jobDefinition.ResourceEnhancement),
		InputOptionType:        types.StringValue(jobDefinition.InputOptionType),
		InputOption:            inputOption,
		OutputOptionType:       types.StringValue(jobDefinition.OutputOptionType),
		OutputOption:           job_definitions.NewOutputOption(ctx, jobDefinition.OutputOption),
		FilterRows:             filter.NewFilterRows(ctx, jobDefinition.FilterRows),
		FilterAddTime:          filter.NewFilterAddTime(jobDefinition.FilterAddTime),
	}

	filterColumns, diags := filter.NewFilterColumns(ctx, jobDefinition.FilterColumns)
	resp.Diagnostics.Append(diags...)

	filterColumnsValue, diags := types.ListValueFrom(ctx, types.ObjectType{
		AttrTypes: filter.FilterColumn{}.AttrTypes(),
	}, filterColumns)
	resp.Diagnostics.Append(diags...)
	newState.FilterColumns = filterColumnsValue

	if jobDefinition.FilterMasks != nil {
		filterMasks := filter.NewFilterMasks(jobDefinition.FilterMasks)
		filterMasksValue, diags := types.ListValueFrom(ctx, types.ObjectType{
			AttrTypes: filter.FilterMask{}.AttrTypes(),
		}, filterMasks)
		resp.Diagnostics.Append(diags...)
		newState.FilterMasks = filterMasksValue
	} else {
		newState.FilterMasks = types.ListNull(types.ObjectType{
			AttrTypes: filter.FilterMask{}.AttrTypes(),
		})
	}

	if jobDefinition.FilterGsub != nil {
		filterGsub := filter.NewFilterGsub(jobDefinition.FilterGsub)
		filterGsubValue, diags := types.ListValueFrom(ctx, types.ObjectType{
			AttrTypes: filter.FilterGsub{}.AttrTypes(),
		}, filterGsub)
		resp.Diagnostics.Append(diags...)
		newState.FilterGsub = filterGsubValue
	} else {
		newState.FilterGsub = types.ListNull(types.ObjectType{
			AttrTypes: filter.FilterGsub{}.AttrTypes(),
		})
	}

	if jobDefinition.FilterStringTransforms != nil {
		filterStringTransforms := filter.NewFilterStringTransforms(jobDefinition.FilterStringTransforms)
		filterStringTransformsValue, diags := types.ListValueFrom(ctx, types.ObjectType{
			AttrTypes: filter.FilterStringTransform{}.AttrTypes(),
		}, filterStringTransforms)
		resp.Diagnostics.Append(diags...)
		newState.FilterStringTransforms = filterStringTransformsValue
	} else {
		newState.FilterStringTransforms = types.ListNull(types.ObjectType{
			AttrTypes: filter.FilterStringTransform{}.AttrTypes(),
		})
	}

	if jobDefinition.FilterHashes != nil {
		filterHashes := filter.NewFilterHashes(jobDefinition.FilterHashes)
		filterHashesValue, diags := types.ListValueFrom(ctx, types.ObjectType{
			AttrTypes: filter.FilterHash{}.AttrTypes(),
		}, filterHashes)
		resp.Diagnostics.Append(diags...)
		newState.FilterHashes = filterHashesValue
	} else {
		newState.FilterHashes = types.ListNull(types.ObjectType{
			AttrTypes: filter.FilterHash{}.AttrTypes(),
		})
	}

	if jobDefinition.FilterUnixTimeConversions != nil {
		filterUTC := filter.NewFilterUnixTimeConversions(jobDefinition.FilterUnixTimeConversions)
		filterUTCValue, diags := types.ListValueFrom(ctx, types.ObjectType{
			AttrTypes: filter.FilterUnixTimeConversion{}.AttrTypes(),
		}, filterUTC)
		resp.Diagnostics.Append(diags...)
		newState.FilterUnixTimeConversions = filterUTCValue
	} else {
		newState.FilterUnixTimeConversions = types.ListNull(types.ObjectType{
			AttrTypes: filter.FilterUnixTimeConversion{}.AttrTypes(),
		})
	}

	if jobDefinition.Notifications != nil {
		notifications := job_definitions.NewJobDefinitionNotifications(jobDefinition.Notifications)
		notificationsValue, diags := types.SetValueFrom(ctx, types.ObjectType{
			AttrTypes: job_definitions.JobDefinitionNotification{}.AttrTypes(),
		}, notifications)
		resp.Diagnostics.Append(diags...)
		newState.Notifications = notificationsValue
	} else {
		newState.Notifications = types.SetNull(types.ObjectType{
			AttrTypes: job_definitions.JobDefinitionNotification{}.AttrTypes(),
		})
	}

	if jobDefinition.Schedules != nil {
		schedules := model.NewSchedules(jobDefinition.Schedules)
		schedulesValue, diags := types.SetValueFrom(ctx, types.ObjectType{
			AttrTypes: model.Schedule{}.AttrTypes(),
		}, schedules)
		resp.Diagnostics.Append(diags...)
		newState.Schedules = schedulesValue
	} else {
		newState.Schedules = types.SetNull(types.ObjectType{
			AttrTypes: model.Schedule{}.AttrTypes(),
		})
	}

	if jobDefinition.Labels != nil {
		labels := job_definitions.NewLabels(jobDefinition.Labels)
		labelsValue, diags := types.SetValueFrom(ctx, types.ObjectType{
			AttrTypes: job_definitions.Label{}.AttrTypes(),
		}, labels)
		resp.Diagnostics.Append(diags...)
		newState.Labels = labelsValue
	} else {
		newState.Labels = types.SetNull(types.ObjectType{
			AttrTypes: job_definitions.Label{}.AttrTypes(),
		})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}

func (m *jobDefinitionResourceModel) ToUpdateJobDefinitionInput(ctx context.Context, resp *resource.UpdateResponse) (*client.UpdateJobDefinitionInput, diag.Diagnostics) {
	labels := []string{}
	if !m.Labels.IsNull() && !m.Labels.IsUnknown() {
		var labelValues []job_definitions.Label
		diags := m.Labels.ElementsAs(ctx, &labelValues, false)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return nil, resp.Diagnostics
		}

		for _, l := range labelValues {
			labels = append(labels, l.Name.ValueString())
		}
	}

	notifications := []params.JobDefinitionNotificationInput{}
	if !m.Notifications.IsNull() && !m.Notifications.IsUnknown() {
		var notificationValues []job_definitions.JobDefinitionNotification
		diags := m.Notifications.ElementsAs(ctx, &notificationValues, false)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return nil, resp.Diagnostics
		}

		for _, n := range notificationValues {
			notifications = append(notifications, n.ToInput())
		}
	}

	schedules := []parameter.ScheduleInput{}
	if !m.Schedules.IsNull() && !m.Schedules.IsUnknown() {
		var scheduleValues []model.Schedule
		diags := m.Schedules.ElementsAs(ctx, &scheduleValues, false)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return nil, resp.Diagnostics
		}

		for _, s := range scheduleValues {
			schedules = append(schedules, s.ToInput())
		}
	}

	filterColumns := []filterParameters.FilterColumnInput{}
	if !m.FilterColumns.IsNull() && !m.FilterColumns.IsUnknown() {
		var filterColumnValues []filter.FilterColumn
		diags := m.FilterColumns.ElementsAs(ctx, &filterColumnValues, false)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return nil, resp.Diagnostics
		}

		for _, f := range filterColumnValues {
			filterColumns = append(filterColumns, f.ToInput(ctx))
		}
	}

	filterMasks := []filterParameters.FilterMaskInput{}
	if !m.FilterMasks.IsNull() && !m.FilterMasks.IsUnknown() {
		var filterMaskValues []filter.FilterMask
		diags := m.FilterMasks.ElementsAs(ctx, &filterMaskValues, false)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return nil, resp.Diagnostics
		}

		for _, f := range filterMaskValues {
			filterMasks = append(filterMasks, f.ToInput())
		}
	}

	filterGsub := []filterParameters.FilterGsubInput{}
	if !m.FilterGsub.IsNull() && !m.FilterGsub.IsUnknown() {
		var filterGsubValues []filter.FilterGsub
		diags := m.FilterGsub.ElementsAs(ctx, &filterGsubValues, false)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return nil, resp.Diagnostics
		}

		for _, f := range filterGsubValues {
			filterGsub = append(filterGsub, f.ToInput())
		}
	}

	filterStringTransforms := []filterParameters.FilterStringTransformInput{}
	if !m.FilterStringTransforms.IsNull() && !m.FilterStringTransforms.IsUnknown() {
		var filterStringTransformValues []filter.FilterStringTransform
		diags := m.FilterStringTransforms.ElementsAs(ctx, &filterStringTransformValues, false)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return nil, resp.Diagnostics
		}

		for _, f := range filterStringTransformValues {
			filterStringTransforms = append(filterStringTransforms, f.ToInput())
		}
	}

	filterHashes := []filterParameters.FilterHashInput{}
	if !m.FilterHashes.IsNull() && !m.FilterHashes.IsUnknown() {
		var filterHashValues []filter.FilterHash
		diags := m.FilterHashes.ElementsAs(ctx, &filterHashValues, false)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return nil, resp.Diagnostics
		}

		for _, f := range filterHashValues {
			filterHashes = append(filterHashes, f.ToInput())
		}
	}

	filterUnixTimeconversions := []filterParameters.FilterUnixTimeConversionInput{}
	if !m.FilterUnixTimeConversions.IsNull() && !m.FilterUnixTimeConversions.IsUnknown() {
		var filterUnixTimeConversionValues []filter.FilterUnixTimeConversion
		diags := m.FilterUnixTimeConversions.ElementsAs(ctx, &filterUnixTimeConversionValues, false)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return nil, resp.Diagnostics
		}

		for _, f := range filterUnixTimeConversionValues {
			filterUnixTimeconversions = append(filterUnixTimeconversions, f.ToInput())
		}
	}

	var diags diag.Diagnostics
	inputOption, d := m.InputOption.ToUpdateInput(ctx)
	diags.Append(d...)
	return &client.UpdateJobDefinitionInput{
		Name:                      m.Name.ValueStringPointer(),
		Description:               model.NewNullableString(m.Description),
		ResourceGroupID:           model.NewNullableInt64(m.ResourceGroupID),
		IsRunnableConcurrently:    m.IsRunnableConcurrently.ValueBoolPointer(),
		RetryLimit:                m.RetryLimit.ValueInt64Pointer(),
		ResourceEnhancement:       m.ResourceEnhancement.ValueStringPointer(),
		FilterColumns:             &filterColumns,
		FilterRows:                model.WrapObject(m.FilterRows.ToInput(ctx)),
		FilterMasks:               &filterMasks,
		FilterAddTime:             model.WrapObject(m.FilterAddTime.ToInput()),
		FilterGsub:                &filterGsub,
		FilterStringTransforms:    &filterStringTransforms,
		FilterHashes:              &filterHashes,
		FilterUnixTimeConversions: &filterUnixTimeconversions,
		InputOption:               inputOption,
		OutputOption:              m.OutputOption.ToUpdateInput(ctx),
		Labels:                    &labels,
		Schedules:                 &schedules,
		Notifications:             &notifications,
	}, diags
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

	jobDefinitionInput, diags := plan.ToCreateJobDefinitionInput(ctx, resp)
	if diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	jobDefinition, err := r.client.CreateJobDefinition(jobDefinitionInput)
	if err != nil {
		resp.Diagnostics.AddError(
			"Creating job definition",
			fmt.Sprintf("Unable to create job definition, got error: %s", err),
		)
		return
	}

	inputOption, diags := job_definitions.NewInputOption(ctx, jobDefinition.InputOption, plan.InputOption)
	if diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}

	newState := jobDefinitionResourceModel{
		ID:                     types.Int64Value(jobDefinition.ID),
		Name:                   types.StringValue(jobDefinition.Name),
		Description:            types.StringPointerValue(jobDefinition.Description),
		ResourceGroupID:        types.Int64PointerValue(jobDefinition.ResourceGroupID),
		IsRunnableConcurrently: types.BoolPointerValue(jobDefinition.IsRunnableConcurrently),
		RetryLimit:             types.Int64Value(jobDefinition.RetryLimit),
		ResourceEnhancement:    types.StringPointerValue(jobDefinition.ResourceEnhancement),
		InputOptionType:        types.StringValue(jobDefinition.InputOptionType),
		InputOption:            inputOption,
		OutputOptionType:       types.StringValue(jobDefinition.OutputOptionType),
		OutputOption:           job_definitions.NewOutputOption(ctx, jobDefinition.OutputOption),
		FilterRows:             filter.NewFilterRows(ctx, jobDefinition.FilterRows),
		FilterAddTime:          filter.NewFilterAddTime(jobDefinition.FilterAddTime),
	}

	filterColumns, diags := filter.NewFilterColumns(ctx, jobDefinition.FilterColumns)
	resp.Diagnostics.Append(diags...)

	filterColumnsValue, diags := types.ListValueFrom(ctx, types.ObjectType{
		AttrTypes: filter.FilterColumn{}.AttrTypes(),
	}, filterColumns)
	resp.Diagnostics.Append(diags...)
	newState.FilterColumns = filterColumnsValue

	if jobDefinition.FilterMasks != nil {
		filterMasks := filter.NewFilterMasks(jobDefinition.FilterMasks)
		filterMasksValue, diags := types.ListValueFrom(ctx, types.ObjectType{
			AttrTypes: filter.FilterMask{}.AttrTypes(),
		}, filterMasks)
		resp.Diagnostics.Append(diags...)
		newState.FilterMasks = filterMasksValue
	} else {
		newState.FilterMasks = types.ListNull(types.ObjectType{
			AttrTypes: filter.FilterMask{}.AttrTypes(),
		})
	}

	if jobDefinition.FilterGsub != nil {
		filterGsub := filter.NewFilterGsub(jobDefinition.FilterGsub)
		filterGsubValue, diags := types.ListValueFrom(ctx, types.ObjectType{
			AttrTypes: filter.FilterGsub{}.AttrTypes(),
		}, filterGsub)
		resp.Diagnostics.Append(diags...)
		newState.FilterGsub = filterGsubValue
	} else {
		newState.FilterGsub = types.ListNull(types.ObjectType{
			AttrTypes: filter.FilterGsub{}.AttrTypes(),
		})
	}

	if jobDefinition.FilterStringTransforms != nil {
		filterStringTransforms := filter.NewFilterStringTransforms(jobDefinition.FilterStringTransforms)
		filterStringTransformsValue, diags := types.ListValueFrom(ctx, types.ObjectType{
			AttrTypes: filter.FilterStringTransform{}.AttrTypes(),
		}, filterStringTransforms)
		resp.Diagnostics.Append(diags...)
		newState.FilterStringTransforms = filterStringTransformsValue
	} else {
		newState.FilterStringTransforms = types.ListNull(types.ObjectType{
			AttrTypes: filter.FilterStringTransform{}.AttrTypes(),
		})
	}

	if jobDefinition.FilterHashes != nil {
		filterHashes := filter.NewFilterHashes(jobDefinition.FilterHashes)
		filterHashesValue, diags := types.ListValueFrom(ctx, types.ObjectType{
			AttrTypes: filter.FilterHash{}.AttrTypes(),
		}, filterHashes)
		resp.Diagnostics.Append(diags...)
		newState.FilterHashes = filterHashesValue
	} else {
		newState.FilterHashes = types.ListNull(types.ObjectType{
			AttrTypes: filter.FilterHash{}.AttrTypes(),
		})
	}

	if jobDefinition.FilterUnixTimeConversions != nil {
		filterUTC := filter.NewFilterUnixTimeConversions(jobDefinition.FilterUnixTimeConversions)
		filterUTCValue, diags := types.ListValueFrom(ctx, types.ObjectType{
			AttrTypes: filter.FilterUnixTimeConversion{}.AttrTypes(),
		}, filterUTC)
		resp.Diagnostics.Append(diags...)
		newState.FilterUnixTimeConversions = filterUTCValue
	} else {
		newState.FilterUnixTimeConversions = types.ListNull(types.ObjectType{
			AttrTypes: filter.FilterUnixTimeConversion{}.AttrTypes(),
		})
	}

	if jobDefinition.Notifications != nil {
		notifications := job_definitions.NewJobDefinitionNotifications(jobDefinition.Notifications)
		notificationsValue, diags := types.SetValueFrom(ctx, types.ObjectType{
			AttrTypes: job_definitions.JobDefinitionNotification{}.AttrTypes(),
		}, notifications)
		resp.Diagnostics.Append(diags...)
		newState.Notifications = notificationsValue
	} else {
		newState.Notifications = types.SetNull(types.ObjectType{
			AttrTypes: job_definitions.JobDefinitionNotification{}.AttrTypes(),
		})
	}

	if jobDefinition.Schedules != nil {
		schedules := model.NewSchedules(jobDefinition.Schedules)
		schedulesValue, diags := types.SetValueFrom(ctx, types.ObjectType{
			AttrTypes: model.Schedule{}.AttrTypes(),
		}, schedules)
		resp.Diagnostics.Append(diags...)
		newState.Schedules = schedulesValue
	} else {
		newState.Schedules = types.SetNull(types.ObjectType{
			AttrTypes: model.Schedule{}.AttrTypes(),
		})
	}

	if jobDefinition.Labels != nil {
		labels := job_definitions.NewLabels(jobDefinition.Labels)
		labelsValue, diags := types.SetValueFrom(ctx, types.ObjectType{
			AttrTypes: job_definitions.Label{}.AttrTypes(),
		}, labels)
		resp.Diagnostics.Append(diags...)
		newState.Labels = labelsValue
	} else {
		newState.Labels = types.SetNull(types.ObjectType{
			AttrTypes: job_definitions.Label{}.AttrTypes(),
		})
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
	inputOption, diags := job_definitions.NewInputOption(ctx, jobDefinition.InputOption, state.InputOption)
	if diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}

	newState := jobDefinitionResourceModel{
		ID:                     types.Int64Value(jobDefinition.ID),
		Name:                   types.StringValue(jobDefinition.Name),
		Description:            types.StringPointerValue(jobDefinition.Description),
		ResourceGroupID:        types.Int64PointerValue(jobDefinition.ResourceGroupID),
		IsRunnableConcurrently: types.BoolPointerValue(jobDefinition.IsRunnableConcurrently),
		RetryLimit:             types.Int64Value(jobDefinition.RetryLimit),
		ResourceEnhancement:    types.StringPointerValue(jobDefinition.ResourceEnhancement),
		InputOptionType:        types.StringValue(jobDefinition.InputOptionType),
		InputOption:            inputOption,
		OutputOptionType:       types.StringValue(jobDefinition.OutputOptionType),
		OutputOption:           job_definitions.NewOutputOption(ctx, jobDefinition.OutputOption),
		FilterRows:             filter.NewFilterRows(ctx, jobDefinition.FilterRows),
		FilterAddTime:          filter.NewFilterAddTime(jobDefinition.FilterAddTime),
	}

	filterColumns, diags := filter.NewFilterColumns(ctx, jobDefinition.FilterColumns)
	resp.Diagnostics.Append(diags...)

	filterColumnsValue, diags := types.ListValueFrom(ctx, types.ObjectType{
		AttrTypes: filter.FilterColumn{}.AttrTypes(),
	}, filterColumns)
	resp.Diagnostics.Append(diags...)
	newState.FilterColumns = filterColumnsValue

	if jobDefinition.FilterMasks != nil {
		filterMasks := filter.NewFilterMasks(jobDefinition.FilterMasks)
		filterMasksValue, diags := types.ListValueFrom(ctx, types.ObjectType{
			AttrTypes: filter.FilterMask{}.AttrTypes(),
		}, filterMasks)
		resp.Diagnostics.Append(diags...)
		newState.FilterMasks = filterMasksValue
	} else {
		newState.FilterMasks = types.ListNull(types.ObjectType{
			AttrTypes: filter.FilterMask{}.AttrTypes(),
		})
	}

	if jobDefinition.FilterGsub != nil {
		filterGsub := filter.NewFilterGsub(jobDefinition.FilterGsub)
		filterGsubValue, diags := types.ListValueFrom(ctx, types.ObjectType{
			AttrTypes: filter.FilterGsub{}.AttrTypes(),
		}, filterGsub)
		resp.Diagnostics.Append(diags...)
		newState.FilterGsub = filterGsubValue
	} else {
		newState.FilterGsub = types.ListNull(types.ObjectType{
			AttrTypes: filter.FilterGsub{}.AttrTypes(),
		})
	}

	if jobDefinition.FilterStringTransforms != nil {
		filterStringTransforms := filter.NewFilterStringTransforms(jobDefinition.FilterStringTransforms)
		filterStringTransformsValue, diags := types.ListValueFrom(ctx, types.ObjectType{
			AttrTypes: filter.FilterStringTransform{}.AttrTypes(),
		}, filterStringTransforms)
		resp.Diagnostics.Append(diags...)
		newState.FilterStringTransforms = filterStringTransformsValue
	} else {
		newState.FilterStringTransforms = types.ListNull(types.ObjectType{
			AttrTypes: filter.FilterStringTransform{}.AttrTypes(),
		})
	}

	if jobDefinition.FilterHashes != nil {
		filterHashes := filter.NewFilterHashes(jobDefinition.FilterHashes)
		filterHashesValue, diags := types.ListValueFrom(ctx, types.ObjectType{
			AttrTypes: filter.FilterHash{}.AttrTypes(),
		}, filterHashes)
		resp.Diagnostics.Append(diags...)
		newState.FilterHashes = filterHashesValue
	} else {
		newState.FilterHashes = types.ListNull(types.ObjectType{
			AttrTypes: filter.FilterHash{}.AttrTypes(),
		})
	}

	if jobDefinition.FilterUnixTimeConversions != nil {
		filterUTC := filter.NewFilterUnixTimeConversions(jobDefinition.FilterUnixTimeConversions)
		filterUTCValue, diags := types.ListValueFrom(ctx, types.ObjectType{
			AttrTypes: filter.FilterUnixTimeConversion{}.AttrTypes(),
		}, filterUTC)
		resp.Diagnostics.Append(diags...)
		newState.FilterUnixTimeConversions = filterUTCValue
	} else {
		newState.FilterUnixTimeConversions = types.ListNull(types.ObjectType{
			AttrTypes: filter.FilterUnixTimeConversion{}.AttrTypes(),
		})
	}

	if jobDefinition.Notifications != nil {
		notifications := job_definitions.NewJobDefinitionNotifications(jobDefinition.Notifications)
		notificationsValue, diags := types.SetValueFrom(ctx, types.ObjectType{
			AttrTypes: job_definitions.JobDefinitionNotification{}.AttrTypes(),
		}, notifications)
		resp.Diagnostics.Append(diags...)
		newState.Notifications = notificationsValue
	} else {
		newState.Notifications = types.SetNull(types.ObjectType{
			AttrTypes: job_definitions.JobDefinitionNotification{}.AttrTypes(),
		})
	}

	if jobDefinition.Schedules != nil {
		schedules := model.NewSchedules(jobDefinition.Schedules)
		schedulesValue, diags := types.SetValueFrom(ctx, types.ObjectType{
			AttrTypes: model.Schedule{}.AttrTypes(),
		}, schedules)
		resp.Diagnostics.Append(diags...)
		newState.Schedules = schedulesValue
	} else {
		newState.Schedules = types.SetNull(types.ObjectType{
			AttrTypes: model.Schedule{}.AttrTypes(),
		})
	}

	if jobDefinition.Labels != nil {
		labels := job_definitions.NewLabels(jobDefinition.Labels)
		labelsValue, diags := types.SetValueFrom(ctx, types.ObjectType{
			AttrTypes: job_definitions.Label{}.AttrTypes(),
		}, labels)
		resp.Diagnostics.Append(diags...)
		newState.Labels = labelsValue
	} else {
		newState.Labels = types.SetNull(types.ObjectType{
			AttrTypes: job_definitions.Label{}.AttrTypes(),
		})
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

func (r *jobDefinitionResource) ValidateConfig(
	ctx context.Context,
	req resource.ValidateConfigRequest,
	resp *resource.ValidateConfigResponse,
) {
	data := &jobDefinitionResourceModel{}
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if data.InputOptionType.ValueString() == "http" {
		if data.InputOption.HttpInputOption == nil {
			return
		}
		httpInputOption := data.InputOption.HttpInputOption
		validateHttpInputOption(httpInputOption, resp)
	}
}

func validateHttpInputOption(httpInputOption *input_options.HttpInputOption, resp *resource.ValidateConfigResponse) {
	// validate that request_body and request_params are not set at the same time
	bodySet := !httpInputOption.RequestBody.IsNull() && !httpInputOption.RequestBody.IsUnknown()
	paramsSet := !httpInputOption.RequestParams.IsNull() && len(httpInputOption.RequestParams.Elements()) > 0

	if bodySet && httpInputOption.Method.ValueString() != "POST" {
		resp.Diagnostics.AddAttributeError(
			path.Root("request_body"),
			"request_body is only allowed when method == \"POST\"",
			fmt.Sprintf("method is %q, so request_body must be removed or method changed to \"POST\".",
				httpInputOption.Method.ValueString()),
		)
	}

	if bodySet && paramsSet {
		resp.Diagnostics.AddAttributeError(
			path.Root("request_body"),
			"request_body conflicts with request_params",
			"When request_body is set, request_params must be omitted.",
		)
	}

	// validate pagination settings
	switch httpInputOption.PagerType.ValueString() {
	case "offset":
		if httpInputOption.PagerFromParam.IsNull() ||
			httpInputOption.PagerFromParam.IsUnknown() {
			resp.Diagnostics.AddAttributeError(
				path.Root("pager_from_param"),
				"pager_from_param is required when pager_type is offset",
				"pager_from_param must be set to the name of the parameter that specifies the starting offset.",
			)
		}
	case "cursor":
		if httpInputOption.CursorRequestParameterCursorName.IsNull() ||
			httpInputOption.CursorRequestParameterCursorName.IsUnknown() {
			resp.Diagnostics.AddAttributeError(
				path.Root("cursor_request_parameter_cursor_name"),
				"cursor_request_parameter_cursor_name is required when pager_type is cursor",
				"cursor_request_parameter_cursor_name must be set to the name of the parameter that specifies the cursor.",
			)
		}
		if httpInputOption.CursorResponseParameterCursorJsonPath.IsNull() ||
			httpInputOption.CursorResponseParameterCursorJsonPath.IsUnknown() {
			resp.Diagnostics.AddAttributeError(
				path.Root("cursor_response_parameter_cursor_json_path"),
				"cursor_response_parameter_cursor_json_path is required when pager_type is cursor",
				"cursor_response_parameter_cursor_json_path must be set to the JSONPath that extracts the cursor value from the response.",
			)
		}
	}
}
