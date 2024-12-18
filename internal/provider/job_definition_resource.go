package provider

import (
	//"context"
	//"fmt"
	//"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"terraform-provider-trocco/internal/client"
	filterEntities "terraform-provider-trocco/internal/client/entities/job_definitions/filter"
	"terraform-provider-trocco/internal/provider/models"
	"terraform-provider-trocco/internal/provider/models/job_definitions"
	"terraform-provider-trocco/internal/provider/models/job_definitions/filter"
	"terraform-provider-trocco/internal/provider/models/job_definitions/input_options"
	"terraform-provider-trocco/internal/provider/models/job_definitions/output_options"
)

type jobDefinitionResourceModel struct {
	ID                        types.Int64                                 `tfsdk:"id"`
	Name                      types.String                                `tfsdk:"name"`
	Description               types.String                                `tfsdk:"description"`
	ResourceGroupID           types.Int64                                 `tfsdk:"resource_group_id"`
	IsRunnableConcurrently    types.Bool                                  `tfsdk:"is_runnable_concurrently"`
	RetryLimit                types.Int64                                 `tfsdk:"retry_limit"`
	ResourceEnhancement       types.String                                `tfsdk:"resource_enhancement"`
	InputOptionType           types.String                                `tfsdk:"input_option_type"`
	InputOption               inputOption                                 `tfsdk:"input_option"`
	OutputOptionType          types.String                                `tfsdk:"output_option_type"`
	OutputOption              outputOption                                `tfsdk:"output_option"`
	FilterColumns             []filter.FilterColumn                       `tfsdk:"filter_columns"`
	FilterRows                filter.FilterRows                           `tfsdk:"filter_rows"`
	FilterMasks               []filter.FilterMask                         `tfsdk:"filter_masks"`
	FilterAddTime             filter.FilterAddTime                        `tfsdk:"filter_add_time"`
	FilterGsub                []filter.FilterGsub                         `tfsdk:"filter_gsub"`
	FilterStringTransforms    []filter.FilterStringTransform              `tfsdk:"filter_string_transforms"`
	FilterHashes              []filter.FilterHash                         `tfsdk:"filter_hashes"`
	FilterUnixTimeConversions []filterEntities.FilterUnixTimeConversion   `tfsdk:"filter_unixtime_conversions"`
	Notifications             []job_definitions.JobDefinitionNotification `tfsdk:"notifications"`
	Schedules                 []models.Schedule                           `tfsdk:"schedules"`
	Labels                    []models.LabelModel                         `tfsdk:"labels"`
}

type inputOption struct {
	MySQLInputOption input_options.MySQLInputOption `tfsdk:"mysql_input_option"`
	GcsInputOption   input_options.GcsInputOption   `tfsdk:"gcs_input_option"`
}

type outputOption struct {
	BigQueryOutputOption *output_options.BigQueryOutputOption `tfsdk:"bigquery_output_option"`
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

//
//func (r *jobDefinitionResource) Read(
//	ctx context.Context,
//	req resource.ReadRequest,
//	resp *resource.ReadResponse,
//) {
//	state := &jobDefinitionResourceModel{}
//	resp.Diagnostics.Append(req.State.Get(ctx, state)...)
//	if resp.Diagnostics.HasError() {
//		return
//	}
//
//	connection, err := r.client.GetConnection(
//		state.ConnectionType.ValueString(),
//		state.ID.ValueInt64(),
//	)
//	if err != nil {
//		resp.Diagnostics.AddError(
//			"Reading connection",
//			fmt.Sprintf("Unable to read connection, got error: %s", err),
//		)
//		return
//	}
//
//	newState := connectionResourceModel{
//		// Common Fields
//		ConnectionType:  state.ConnectionType,
//		ID:              types.Int64Value(connection.ID),
//		Name:            types.StringPointerValue(connection.Name),
//		Description:     types.StringPointerValue(connection.Description),
//		ResourceGroupID: types.Int64PointerValue(connection.ResourceGroupID),
//
//		// BigQuery Fields
//		ProjectID:             types.StringPointerValue(connection.ProjectID),
//		ServiceAccountJSONKey: state.ServiceAccountJSONKey,
//
//		// Snowflake Fields
//		Host:       types.StringPointerValue(connection.Host),
//		UserName:   types.StringPointerValue(connection.UserName),
//		Role:       types.StringPointerValue(connection.Role),
//		AuthMethod: types.StringPointerValue(connection.AuthMethod),
//		Password:   state.Password,
//		PrivateKey: state.PrivateKey,
//	}
//	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
//}
