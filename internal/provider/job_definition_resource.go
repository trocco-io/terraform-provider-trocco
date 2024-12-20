package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"terraform-provider-trocco/internal/client"
	"terraform-provider-trocco/internal/provider/models"
	"terraform-provider-trocco/internal/provider/models/job_definitions"
	"terraform-provider-trocco/internal/provider/models/job_definitions/filter"
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

//func (model *jobDefinitionResourceModel) ToCreateJobDefinitionInput() *client.CreateJobDefinitionInput {
//	return &client.CreateJobDefinitionInput{
//		// Common Fields
//		Name:                      model.Name.ValueString(),
//		Description:               model.Description.ValueStringPointer(),
//		ResourceGroupID:           newNullableFromTerraformInt64(model.ResourceGroupID),
//		IsRunnableConcurrently:    model.IsRunnableConcurrently.ValueBoolPointer(),
//		RetryLimit:                model.RetryLimit.ValueInt64(),
//		ResourceEnhancement:       model.ResourceEnhancement.ValueStringPointer(),
//		FilterColumns:             nil,
//		FilterRows:                nil,
//		FilterMasks:               nil,
//		FilterAddTime:             nil,
//		FilterGsub:                nil,
//		FilterStringTransforms:    nil,
//		FilterHashes:              nil,
//		FilterUnixTimeConversions: nil,
//		InputOptionType:           model.InputOptionType.ValueString(),
//		InputOption:               client.InputOptionInput{},
//		OutputOptionType:          model.OutputOptionType.ValueString(),
//		OutputOption:              client.OutputOptionInput{},
//		Labels:                    nil,
//		Schedules:                 nil,
//		Notifications:             nil,
//	}
//}

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
