package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"terraform-provider-trocco/internal/client"
	job_definitions3 "terraform-provider-trocco/internal/client/entities/job_definitions"
	"terraform-provider-trocco/internal/client/parameters"
	job_definitions2 "terraform-provider-trocco/internal/client/parameters/job_definitions"
	filter2 "terraform-provider-trocco/internal/client/parameters/job_definitions/filter"
	"terraform-provider-trocco/internal/provider/models"
	"terraform-provider-trocco/internal/provider/models/job_definitions"
	"terraform-provider-trocco/internal/provider/models/job_definitions/filter"
	"terraform-provider-trocco/internal/provider/models/job_definitions/input_options"
)

type jobDefinitionResourceModel struct {
	ID                        types.Int64                                  `tfsdk:"id"`
	Name                      types.String                                 `tfsdk:"name"`
	Description               types.String                                 `tfsdk:"description"`
	ResourceGroupID           types.Int64                                  `tfsdk:"resource_group_id"`
	IsRunnableConcurrently    types.Bool                                   `tfsdk:"is_runnable_concurrently"`
	RetryLimit                types.Int64                                  `tfsdk:"retry_limit"`
	ResourceEnhancement       types.String                                 `tfsdk:"resource_enhancement"`
	InputOptionType           types.String                                 `tfsdk:"input_option_type"`
	InputOption               job_definitions.InputOption                  `tfsdk:"input_option"`
	OutputOptionType          types.String                                 `tfsdk:"output_option_type"`
	OutputOption              job_definitions.OutputOption                 `tfsdk:"output_option"`
	FilterColumns             []filter.FilterColumn                        `tfsdk:"filter_columns"`
	FilterRows                *filter.FilterRows                           `tfsdk:"filter_rows"`
	FilterMasks               []filter.FilterMask                          `tfsdk:"filter_masks"`
	FilterAddTime             *filter.FilterAddTime                        `tfsdk:"filter_add_time"`
	FilterGsub                []filter.FilterGsub                          `tfsdk:"filter_gsub"`
	FilterStringTransforms    []filter.FilterStringTransform               `tfsdk:"filter_string_transforms"`
	FilterHashes              []filter.FilterHash                          `tfsdk:"filter_hashes"`
	FilterUnixTimeConversions []filter.FilterUnixTimeConversion            `tfsdk:"filter_unixtime_conversions"`
	Notifications             *[]job_definitions.JobDefinitionNotification `tfsdk:"notifications"`
	Schedules                 *[]models.Schedule                           `tfsdk:"schedules"`
	Labels                    *[]models.LabelModel                         `tfsdk:"labels"`
}

type jobDefinitionResource struct {
	client *client.TroccoClient
}

func (model *jobDefinitionResourceModel) ToCreateJobDefinitionInput() *client.CreateJobDefinitionInput {
	var labels []string
	if model.Labels != nil {
		for _, l := range *model.Labels {
			labels = append(labels, l.Name.ValueString())
		}
	}
	var notifications []job_definitions2.JobDefinitionNotificationInput
	if model.Notifications != nil {
		for _, n := range *model.Notifications {
			notifications = append(notifications, n.ToInput())
		}
	}
	var schedules []parameters.ScheduleInput
	if model.Schedules != nil {
		for _, s := range *model.Schedules {
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
		Description:               model.Description.ValueStringPointer(),
		ResourceGroupID:           newNullableFromTerraformInt64(model.ResourceGroupID),
		IsRunnableConcurrently:    model.IsRunnableConcurrently.ValueBoolPointer(),
		RetryLimit:                model.RetryLimit.ValueInt64(),
		ResourceEnhancement:       model.ResourceEnhancement.ValueStringPointer(),
		FilterColumns:             filterColumns,
		FilterRows:                model.FilterRows.ToInput(),
		FilterMasks:               filterMasks,
		FilterAddTime:             model.FilterAddTime.ToInput(),
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

	var mysqlInputOptionColumns []input_options.InputOptionColumn
	for _, m := range jobDefinition.InputOption.MySQLInputOption.InputOptionColumns {
		mysqlInputOptionColumns = append(mysqlInputOptionColumns, input_options.InputOptionColumn{
			Name: types.StringValue(m.Name),
			Type: types.StringValue(m.Type),
		})
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
		InputOption: job_definitions.InputOption{
			MySQLInputOption: input_options.ToMysqlInputOptionModel(jobDefinition.InputOption.MySQLInputOption),
			GcsInputOption:   input_options.ToGcsInputOptionModel(jobDefinition.InputOption.GcsInputOption),
		},
		OutputOptionType:          types.StringValue(jobDefinition.OutputOptionType),
		OutputOption:              job_definitions.OutputOption{},
		FilterColumns:             nil,
		FilterRows:                nil,
		FilterMasks:               nil,
		FilterAddTime:             nil,
		FilterGsub:                nil,
		FilterStringTransforms:    nil,
		FilterHashes:              nil,
		FilterUnixTimeConversions: nil,
		Notifications:             toJobDefinitionNotificationModel(jobDefinition.Notifications),
		Schedules:                 nil,
		Labels:                    nil,
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}

func toJobDefinitionNotificationModel(notifications *[]job_definitions3.JobDefinitionNotification) *[]job_definitions.JobDefinitionNotification {
	if notifications == nil {
		return nil
	}
	outputs := make([]job_definitions.JobDefinitionNotification, 0, len(*notifications))
	for _, input := range *notifications {
		outputs = append(outputs, job_definitions.JobDefinitionNotification{
			DestinationType:  types.StringValue(input.DestinationType),
			SlackChannelID:   types.Int64PointerValue(input.SlackChannelID),
			EmailID:          types.Int64PointerValue(input.EmailID),
			NotificationType: types.StringValue(input.NotificationType),
			NotifyWhen:       types.StringPointerValue(input.NotifyWhen),
			Message:          types.StringValue(input.Message),
			RecordCount:      types.Int64PointerValue(input.RecordCount),
			RecordOperator:   types.StringPointerValue(input.RecordOperator),
			RecordType:       types.StringPointerValue(input.RecordType),
			Minutes:          types.Int64PointerValue(input.Minutes),
		})
	}
	return &outputs
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
