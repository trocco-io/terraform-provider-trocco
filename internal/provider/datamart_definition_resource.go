package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"terraform-provider-trocco/internal/client"
)

func NewDatamartDefinitionResource() resource.Resource {
	return &datamartDefinitionResource{}
}

type datamartDefinitionResource struct {
	client *client.TroccoClient
}

type datamartDefinitionModel struct {
	ID                     types.Int64  `tfsdk:"id"`
	Name                   types.String `tfsdk:"name"`
	DataWarehouseType      types.String `tfsdk:"data_warehouse_type"`
	IsRunnableConcurrently types.Bool   `tfsdk:"is_runnable_concurrently"`
	// TODO: nest & optional & other values
	BiqueryConnectionID types.Int64  `tfsdk:"bigquery_connection_id"`
	QueryMode           types.String `tfsdk:"query_mode"`
	Query               types.String `tfsdk:"query"`
	DestinationDataset  types.String `tfsdk:"destination_dataset"`
	DestinationTable    types.String `tfsdk:"destination_table"`
	WriteDisposition    types.String `tfsdk:"write_disposition"`
}

func (r *datamartDefinitionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "datamart_definition"
}

func (r *datamartDefinitionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *datamartDefinitionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Required: true,
			},
			"data_warehouse_type": schema.StringAttribute{
				Required: true,
			},
			"is_runnable_concurrently": schema.BoolAttribute{
				Required: true,
			},
			"bigquery_connection_id": schema.Int64Attribute{
				Required: true,
			},
			"query_mode": schema.StringAttribute{
				Required: true,
			},
			"query": schema.StringAttribute{
				Required: true,
			},
			"destination_dataset": schema.StringAttribute{
				Required: true,
			},
			"destination_table": schema.StringAttribute{
				Required: true,
			},
			"write_disposition": schema.StringAttribute{
				Required: true,
			},
		},
	}
}

func (r *datamartDefinitionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan datamartDefinitionModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	destiantionDataset := plan.DestinationDataset.ValueString()
	destiantionTable := plan.DestinationTable.ValueString()
	writeDisposition := plan.WriteDisposition.ValueString()
	res, err := r.client.CreateDatamartDefinition(&client.CreateDatamartDefinitionsInput{
		Name:                   plan.Name.ValueString(),
		DataWarehouseType:      plan.DataWarehouseType.ValueString(),
		IsRunnableConcurrently: plan.IsRunnableConcurrently.ValueBool(),
		DatamartBigqueryOption: client.CreateDatamartBigqueryOption{
			BigqueryConnectionID: plan.BiqueryConnectionID.ValueInt64(),
			QueryMode:            plan.QueryMode.ValueString(),
			Query:                plan.Query.ValueString(),
			DestinationDataset:   &destiantionDataset,
			DestinationTable:     &destiantionTable,
			WriteDisposition:     &writeDisposition,
		},
	})
	if err != nil {
		resp.Diagnostics.AddError(
			"failed to create datamart definition",
			fmt.Sprintf("failed to create datamart definition: %v", err),
		)
		return
	}

	datamartDefinition, err := r.client.GetDatamartDefinition(res.ID)
	if err != nil {
		resp.Diagnostics.AddError(
			"failed to get datamart definition",
			fmt.Sprintf("failed to get datamart definition: %v", err),
		)
		return
	}
	plan = datamartDefinitionModel{
		ID:                     types.Int64Value(datamartDefinition.ID),
		Name:                   types.StringValue(datamartDefinition.Name),
		DataWarehouseType:      types.StringValue(datamartDefinition.DataWarehouseType),
		IsRunnableConcurrently: types.BoolValue(datamartDefinition.IsRunnableConcurrently),
		BiqueryConnectionID:    types.Int64Value(datamartDefinition.DatamartBigqueryOption.BigqueryConnectionID),
		QueryMode:              types.StringValue(datamartDefinition.DatamartBigqueryOption.QueryMode),
		Query:                  types.StringValue(datamartDefinition.DatamartBigqueryOption.Query),
		DestinationDataset:     types.StringValue(datamartDefinition.DatamartBigqueryOption.DestinationDataset),
		DestinationTable:       types.StringValue(datamartDefinition.DatamartBigqueryOption.DestinationTable),
		WriteDisposition:       types.StringValue(datamartDefinition.DatamartBigqueryOption.WriteDisposition),
	}

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *datamartDefinitionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state datamartDefinitionModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	id := state.ID.ValueInt64()
	datamartDefinition, err := r.client.GetDatamartDefinition(id)
	if err != nil {
		resp.Diagnostics.AddError(
			"failed to get datamart definition",
			fmt.Sprintf("failed to get datamart definition: %v", err),
		)
		return
	}
	state = datamartDefinitionModel{
		ID:                     types.Int64Value(datamartDefinition.ID),
		Name:                   types.StringValue(datamartDefinition.Name),
		DataWarehouseType:      types.StringValue(datamartDefinition.DataWarehouseType),
		IsRunnableConcurrently: types.BoolValue(datamartDefinition.IsRunnableConcurrently),
		BiqueryConnectionID:    types.Int64Value(datamartDefinition.DatamartBigqueryOption.BigqueryConnectionID),
		QueryMode:              types.StringValue(datamartDefinition.DatamartBigqueryOption.QueryMode),
		Query:                  types.StringValue(datamartDefinition.DatamartBigqueryOption.Query),
		DestinationDataset:     types.StringValue(datamartDefinition.DatamartBigqueryOption.DestinationDataset),
		DestinationTable:       types.StringValue(datamartDefinition.DatamartBigqueryOption.DestinationTable),
		WriteDisposition:       types.StringValue(datamartDefinition.DatamartBigqueryOption.WriteDisposition),
	}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *datamartDefinitionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state datamartDefinitionModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	ID := state.ID.ValueInt64()
	err := r.client.DeleteDatamartDefinition(ID)
	if err != nil {
		resp.Diagnostics.AddError(
			"failed to delete datamart definition",
			fmt.Sprintf("failed to delete datamart definition: %v", err),
		)
		return
	}
}

func (r *datamartDefinitionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state datamartDefinitionModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	updateName := plan.Name.ValueString()
	updateIsRunnableConcurrently := plan.IsRunnableConcurrently.ValueBool()
	updateBiqueryConnectionID := plan.BiqueryConnectionID.ValueInt64()
	updateQueryMode := plan.QueryMode.ValueString()
	updateQuery := plan.Query.ValueString()
	updateDestinationDataset := plan.DestinationDataset.ValueString()
	updateDestinationTable := plan.DestinationTable.ValueString()
	updateWriteDisposition := plan.WriteDisposition.ValueString()
	updateRequest := client.UpdateDatamartDefinitionsInput{
		Name:                   &updateName,
		IsRunnableConcurrently: &updateIsRunnableConcurrently,
		DatamartBigqueryOption: client.UpdateDatamartBigqueryOption{
			BigqueryConnectionID: &updateBiqueryConnectionID,
			QueryMode:            &updateQueryMode,
			Query:                &updateQuery,
			DestinationDataset:   &updateDestinationDataset,
			DestinationTable:     &updateDestinationTable,
			WriteDisposition:     &updateWriteDisposition,
		},
	}
	err := r.client.UpdateDatamartDefinition(state.ID.ValueInt64(), &updateRequest)
	if err != nil {
		resp.Diagnostics.AddError(
			"failed to update datamart definition",
			fmt.Sprintf("failed to update datamart definition: %v", err),
		)
		return
	}

	datamartDefinition, err := r.client.GetDatamartDefinition(state.ID.ValueInt64())
	if err != nil {
		resp.Diagnostics.AddError(
			"failed to get datamart definition",
			fmt.Sprintf("failed to get datamart definition: %v", err),
		)
		return
	}
	state = datamartDefinitionModel{
		ID:                     state.ID,
		Name:                   types.StringValue(datamartDefinition.Name),
		DataWarehouseType:      types.StringValue(datamartDefinition.DataWarehouseType),
		IsRunnableConcurrently: types.BoolValue(datamartDefinition.IsRunnableConcurrently),
		BiqueryConnectionID:    types.Int64Value(datamartDefinition.DatamartBigqueryOption.BigqueryConnectionID),
		QueryMode:              types.StringValue(datamartDefinition.DatamartBigqueryOption.QueryMode),
		Query:                  types.StringValue(datamartDefinition.DatamartBigqueryOption.Query),
		DestinationDataset:     types.StringValue(datamartDefinition.DatamartBigqueryOption.DestinationDataset),
		DestinationTable:       types.StringValue(datamartDefinition.DatamartBigqueryOption.DestinationTable),
		WriteDisposition:       types.StringValue(datamartDefinition.DatamartBigqueryOption.WriteDisposition),
	}

	diags := resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}
