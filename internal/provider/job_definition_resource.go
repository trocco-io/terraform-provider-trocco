package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"terraform-provider-trocco/internal/client"
	job_definitions2 "terraform-provider-trocco/internal/client/parameters/job_definitions"
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
	//var schedules []parameters.ScheduleInput
	//if model.Schedules != nil {
	//	for _, s := range *model.Schedules {
	//		schedules = append(schedules, s.ToInput())
	//	}
	//}

	return &client.CreateJobDefinitionInput{
		Name:                      model.Name.ValueString(),
		Description:               model.Description.ValueStringPointer(),
		ResourceGroupID:           newNullableFromTerraformInt64(model.ResourceGroupID),
		IsRunnableConcurrently:    model.IsRunnableConcurrently.ValueBoolPointer(),
		RetryLimit:                model.RetryLimit.ValueInt64(),
		ResourceEnhancement:       model.ResourceEnhancement.ValueStringPointer(),
		FilterColumns:             nil,
		FilterRows:                nil,
		FilterMasks:               nil,
		FilterAddTime:             nil,
		FilterGsub:                nil,
		FilterStringTransforms:    nil,
		FilterHashes:              nil,
		FilterUnixTimeConversions: nil,
		InputOptionType:           model.InputOptionType.ValueString(),
		InputOption:               model.InputOption.ToInput(),
		OutputOptionType:          model.OutputOptionType.ValueString(),
		OutputOption:              model.OutputOption.ToInput(),
		Labels:                    labels,
		Schedules:                 nil,
		Notifications:             notifications,
	}
}

//
//func (r *connectionResource) Create(
//	ctx context.Context,
//	req resource.CreateRequest,
//	resp *resource.CreateResponse,
//) {
//	plan := &jobDefinitionResourceModel{}
//	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
//	if resp.Diagnostics.HasError() {
//		return
//	}
//
//	connection, err := r.client.CreateJobDefinition(
//		plan.ToCreateJobDefinitionInput(),
//	)
//	if err != nil {
//		resp.Diagnostics.AddError(
//			"Creating job definition",
//			fmt.Sprintf("Unable to create job definition, got error: %s", err),
//		)
//		return
//	}
//
//	newState := connectionResourceModel{
//		// Common Fields
//		ConnectionType:  plan.ConnectionType,
//		ID:              types.Int64Value(connection.ID),
//		Name:            types.StringPointerValue(connection.Name),
//		Description:     types.StringPointerValue(connection.Description),
//		ResourceGroupID: types.Int64PointerValue(connection.ResourceGroupID),
//
//		// BigQuery Fields
//		ProjectID:             types.StringPointerValue(connection.ProjectID),
//		ServiceAccountJSONKey: plan.ServiceAccountJSONKey,
//
//		// Snowflake Fields
//		Host:       types.StringPointerValue(connection.Host),
//		UserName:   types.StringPointerValue(connection.UserName),
//		Role:       types.StringPointerValue(connection.Role),
//		AuthMethod: types.StringPointerValue(connection.AuthMethod),
//		Password:   plan.Password,
//		PrivateKey: plan.PrivateKey,
//	}
//	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
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
