package provider

import (
	"context"
	"fmt"

	"terraform-provider-trocco/internal/client"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func NewDatamartDefinitionResource() resource.Resource {
	return &datamartDefinitionResource{}
}

type datamartDefinitionResource struct {
	client *client.TroccoClient
}

type datamartDefinitionModel struct {
	ID                     types.Int64             `tfsdk:"id"`
	Name                   types.String            `tfsdk:"name"`
	Description            types.String            `tfsdk:"description"`
	DataWarehouseType      types.String            `tfsdk:"data_warehouse_type"`
	IsRunnableConcurrently types.Bool              `tfsdk:"is_runnable_concurrently"`
	DatamartBigqueryOption *datamartBigqueryOption `tfsdk:"datamart_bigquery_option"`
}
type datamartBigqueryOption struct {
	BigqueryConnectionID types.Int64  `tfsdk:"bigquery_connection_id"`
	QueryMode            types.String `tfsdk:"query_mode"`
	Query                types.String `tfsdk:"query"`
	DestinationDataset   types.String `tfsdk:"destination_dataset"`
	DestinationTable     types.String `tfsdk:"destination_table"`
	WriteDisposition     types.String `tfsdk:"write_disposition"`
	Location             types.String `tfsdk:"location"`
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
			"description": schema.StringAttribute{
				Optional: true,
			},
			"data_warehouse_type": schema.StringAttribute{
				Required: true,
			},
			"is_runnable_concurrently": schema.BoolAttribute{
				Required: true,
			},
			"datamart_bigquery_option": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
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
						Optional: true,
					},
					"destination_table": schema.StringAttribute{
						Optional: true,
					},
					"write_disposition": schema.StringAttribute{
						Optional: true,
					},
					"location": schema.StringAttribute{
						Optional: true,
					},
				},
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

	var description *string
	var destinationDataset *string
	var destinationTable *string
	var writeDisposition *string
	var location *string
	if plan.DatamartBigqueryOption.QueryMode.ValueString() == "insert" {
		descriptionValue := plan.Description.ValueString()
		description = &descriptionValue
		destiantionDatasetValue := plan.DatamartBigqueryOption.DestinationDataset.ValueString()
		destinationDataset = &destiantionDatasetValue
		destiantionTableValue := plan.DatamartBigqueryOption.DestinationTable.ValueString()
		destinationTable = &destiantionTableValue
		writeDispositionValue := plan.DatamartBigqueryOption.WriteDisposition.ValueString()
		writeDisposition = &writeDispositionValue
	}
	if plan.DatamartBigqueryOption.QueryMode.ValueString() == "query" {
		locationValue := plan.DatamartBigqueryOption.Location.ValueString()
		location = &locationValue
	}
	res, err := r.client.CreateDatamartDefinition(&client.CreateDatamartDefinitionsInput{
		Name:                   plan.Name.ValueString(),
		Description:            description,
		DataWarehouseType:      plan.DataWarehouseType.ValueString(),
		IsRunnableConcurrently: plan.IsRunnableConcurrently.ValueBool(),
		DatamartBigqueryOption: client.CreateDatamartBigqueryOption{
			BigqueryConnectionID: plan.DatamartBigqueryOption.BigqueryConnectionID.ValueInt64(),
			QueryMode:            plan.DatamartBigqueryOption.QueryMode.ValueString(),
			Query:                plan.DatamartBigqueryOption.Query.ValueString(),
			DestinationDataset:   destinationDataset,
			DestinationTable:     destinationTable,
			WriteDisposition:     writeDisposition,
			Location:             location,
		},
	})
	if err != nil {
		resp.Diagnostics.AddError(
			"failed to create datamart definition",
			fmt.Sprintf("failed to create datamart definition: %v", err),
		)
		return
	}

	plan, err = r.fetchModel(res.ID)
	if err != nil {
		resp.Diagnostics.AddError(
			"failed to get datamart definition",
			fmt.Sprintf("failed to get datamart definition: %v", err),
		)
		return
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
	state, err := r.fetchModel(id)
	if err != nil {
		resp.Diagnostics.AddError(
			"failed to get datamart definition",
			fmt.Sprintf("failed to get datamart definition: %v", err),
		)
		return
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

	var name *string
	var description *string
	var isRunnableConcurrently *bool
	var biqueryConnectionID *int64
	var queryMode *string
	var query *string
	var destinationDataset *string
	var destinationTable *string
	var writeDisposition *string
	var location *string
	nameValue := plan.Name.ValueString()
	name = &nameValue
	descriptionValue := plan.Description.ValueString()
	description = &descriptionValue
	isRunnableConcurrentlyValue := plan.IsRunnableConcurrently.ValueBool()
	isRunnableConcurrently = &isRunnableConcurrentlyValue
	biqueryConnectionIDValue := plan.DatamartBigqueryOption.BigqueryConnectionID.ValueInt64()
	biqueryConnectionID = &biqueryConnectionIDValue
	queryModeValue := plan.DatamartBigqueryOption.QueryMode.ValueString()
	queryMode = &queryModeValue
	queryValue := plan.DatamartBigqueryOption.Query.ValueString()
	query = &queryValue

	nextQueryMode := state.DatamartBigqueryOption.QueryMode.ValueString()
	if !plan.DatamartBigqueryOption.QueryMode.IsNull() {
		nextQueryMode = plan.DatamartBigqueryOption.QueryMode.ValueString()
	}
	if nextQueryMode == "insert" {
		destiantionDatasetValue := plan.DatamartBigqueryOption.DestinationDataset.ValueString()
		destinationDataset = &destiantionDatasetValue
		destiantionTableValue := plan.DatamartBigqueryOption.DestinationTable.ValueString()
		destinationTable = &destiantionTableValue
		writeDispositionValue := plan.DatamartBigqueryOption.WriteDisposition.ValueString()
		writeDisposition = &writeDispositionValue
	}
	if nextQueryMode == "query" {
		locationValue := plan.DatamartBigqueryOption.Location.ValueString()
		location = &locationValue
	}
	updateRequest := client.UpdateDatamartDefinitionsInput{
		Name:                   name,
		Description:            description,
		IsRunnableConcurrently: isRunnableConcurrently,
		DatamartBigqueryOption: client.UpdateDatamartBigqueryOption{
			BigqueryConnectionID: biqueryConnectionID,
			QueryMode:            queryMode,
			Query:                query,
			DestinationDataset:   destinationDataset,
			DestinationTable:     destinationTable,
			WriteDisposition:     writeDisposition,
			Location:             location,
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

	state, err = r.fetchModel(state.ID.ValueInt64())
	if err != nil {
		resp.Diagnostics.AddError(
			"failed to get datamart definition",
			fmt.Sprintf("failed to get datamart definition: %v", err),
		)
		return
	}

	diags := resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *datamartDefinitionResource) fetchModel(id int64) (datamartDefinitionModel, error) {
	datamartDefinition, err := r.client.GetDatamartDefinition(id)
	if err != nil {
		return datamartDefinitionModel{}, err
	}
	var description basetypes.StringValue
	if datamartDefinition.Description != "" {
		description = types.StringValue(datamartDefinition.Description)
	}
	var destinationDataset basetypes.StringValue
	if datamartDefinition.DatamartBigqueryOption.DestinationDataset != "" {
		destinationDataset = types.StringValue(datamartDefinition.DatamartBigqueryOption.DestinationDataset)
	}
	var destinationTable basetypes.StringValue
	if datamartDefinition.DatamartBigqueryOption.DestinationTable != "" {
		destinationTable = types.StringValue(datamartDefinition.DatamartBigqueryOption.DestinationTable)
	}
	var writeDisposition basetypes.StringValue
	if datamartDefinition.DatamartBigqueryOption.WriteDisposition != "" {
		writeDisposition = types.StringValue(datamartDefinition.DatamartBigqueryOption.WriteDisposition)
	}
	var location basetypes.StringValue
	if datamartDefinition.DatamartBigqueryOption.Location != "" {
		location = types.StringValue(datamartDefinition.DatamartBigqueryOption.Location)
	}
	model := datamartDefinitionModel{
		ID:                     types.Int64Value(datamartDefinition.ID),
		Name:                   types.StringValue(datamartDefinition.Name),
		Description:            description,
		DataWarehouseType:      types.StringValue(datamartDefinition.DataWarehouseType),
		IsRunnableConcurrently: types.BoolValue(datamartDefinition.IsRunnableConcurrently),
		DatamartBigqueryOption: &datamartBigqueryOption{
			BigqueryConnectionID: types.Int64Value(datamartDefinition.DatamartBigqueryOption.BigqueryConnectionID),
			QueryMode:            types.StringValue(datamartDefinition.DatamartBigqueryOption.QueryMode),
			Query:                types.StringValue(datamartDefinition.DatamartBigqueryOption.Query),
			DestinationDataset:   destinationDataset,
			DestinationTable:     destinationTable,
			WriteDisposition:     writeDisposition,
			Location:             location,
		},
	}
	return model, nil
}
