package output_options

import (
	"context"
	"fmt"
	"terraform-provider-trocco/internal/client/entity/job_definition/output_option"
	outputOptionParameters "terraform-provider-trocco/internal/client/parameter/job_definition/output_option"
	"terraform-provider-trocco/internal/provider/model"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type HubspotOutputOption struct {
	HubspotConnectionID types.Int64  `tfsdk:"hubspot_connection_id"`
	ObjectType          types.String `tfsdk:"object_type"`
	Mode                types.String `tfsdk:"mode"`
	UpsertKey           types.String `tfsdk:"upsert_key"`
	NumberOfParallels   types.Int64  `tfsdk:"number_of_parallels"`
	Associations        types.List   `tfsdk:"associations"`
}

type hubspotOutputAssociation struct {
	ToObjectType  types.String `tfsdk:"to_object_type"`
	FromObjectKey types.String `tfsdk:"from_object_key"`
	ToObjectKey   types.String `tfsdk:"to_object_key"`
}

func (h hubspotOutputAssociation) attrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"to_object_type":  types.StringType,
		"from_object_key": types.StringType,
		"to_object_key":   types.StringType,
	}
}

func NewHubspotOutputOption(ctx context.Context, hubspotOutputOption *output_option.HubspotOutputOption) *HubspotOutputOption {
	if hubspotOutputOption == nil {
		return nil
	}

	result := &HubspotOutputOption{
		HubspotConnectionID: types.Int64Value(hubspotOutputOption.HubspotConnectionID),
		ObjectType:          types.StringValue(hubspotOutputOption.ObjectType),
		Mode:                types.StringValue(hubspotOutputOption.Mode),
		UpsertKey:           types.StringPointerValue(hubspotOutputOption.UpsertKey),
		NumberOfParallels:   types.Int64Value(hubspotOutputOption.NumberOfParallels),
	}

	hubspotOutputAssociations, err := newHubspotOutputAssociations(ctx, hubspotOutputOption.Associations)
	if err != nil {
		return nil
	}
	result.Associations = hubspotOutputAssociations

	return result
}

func newHubspotOutputAssociations(ctx context.Context, inputAssociations *[]output_option.HubspotOutputAssociation) (types.List, error) {
	objectType := types.ObjectType{
		AttrTypes: hubspotOutputAssociation{}.attrTypes(),
	}

	if inputAssociations == nil || len(*inputAssociations) == 0 {
		return types.ListNull(objectType), nil
	}

	associations := make([]hubspotOutputAssociation, 0, len(*inputAssociations))
	for _, input := range *inputAssociations {
		association := hubspotOutputAssociation{
			ToObjectType:  types.StringValue(input.ToObjectType),
			FromObjectKey: types.StringValue(input.FromObjectKey),
			ToObjectKey:   types.StringValue(input.ToObjectKey),
		}
		associations = append(associations, association)
	}

	listValue, diags := types.ListValueFrom(ctx, objectType, associations)
	if diags.HasError() {
		return types.ListNull(objectType), fmt.Errorf("failed to convert to ListValue: %v", diags)
	}
	return listValue, nil
}

func (hubspotOutputOption *HubspotOutputOption) ToInput(ctx context.Context) *outputOptionParameters.HubspotOutputOptionInput {
	if hubspotOutputOption == nil {
		return nil
	}

	var associations *[]outputOptionParameters.HubspotOutputAssociationInput
	if !hubspotOutputOption.Associations.IsNull() && !hubspotOutputOption.Associations.IsUnknown() {
		var associationValues []hubspotOutputAssociation
		diags := hubspotOutputOption.Associations.ElementsAs(ctx, &associationValues, false)
		if diags.HasError() {
			return nil
		}

		outputs := make([]outputOptionParameters.HubspotOutputAssociationInput, 0, len(associationValues))
		for _, assoc := range associationValues {
			outputs = append(outputs, outputOptionParameters.HubspotOutputAssociationInput{
				ToObjectType:  assoc.ToObjectType.ValueString(),
				FromObjectKey: assoc.FromObjectKey.ValueString(),
				ToObjectKey:   assoc.ToObjectKey.ValueString(),
			})
		}
		associations = &outputs
	}

	return &outputOptionParameters.HubspotOutputOptionInput{
		HubspotConnectionID: hubspotOutputOption.HubspotConnectionID.ValueInt64(),
		ObjectType:          hubspotOutputOption.ObjectType.ValueString(),
		Mode:                hubspotOutputOption.Mode.ValueString(),
		UpsertKey:           model.NewNullableString(hubspotOutputOption.UpsertKey),
		NumberOfParallels:   hubspotOutputOption.NumberOfParallels.ValueInt64(),
		Associations:        model.WrapObjectList(associations),
	}
}

func (hubspotOutputOption *HubspotOutputOption) ToUpdateInput(ctx context.Context) *outputOptionParameters.UpdateHubspotOutputOptionInput {
	if hubspotOutputOption == nil {
		return nil
	}

	var associations *[]outputOptionParameters.HubspotOutputAssociationInput
	if !hubspotOutputOption.Associations.IsNull() && !hubspotOutputOption.Associations.IsUnknown() {
		var associationValues []hubspotOutputAssociation
		diags := hubspotOutputOption.Associations.ElementsAs(ctx, &associationValues, false)
		if diags.HasError() {
			return nil
		}

		outputs := make([]outputOptionParameters.HubspotOutputAssociationInput, 0, len(associationValues))
		for _, assoc := range associationValues {
			outputs = append(outputs, outputOptionParameters.HubspotOutputAssociationInput{
				ToObjectType:  assoc.ToObjectType.ValueString(),
				FromObjectKey: assoc.FromObjectKey.ValueString(),
				ToObjectKey:   assoc.ToObjectKey.ValueString(),
			})
		}
		associations = &outputs
	}

	return &outputOptionParameters.UpdateHubspotOutputOptionInput{
		HubspotConnectionID: hubspotOutputOption.HubspotConnectionID.ValueInt64Pointer(),
		ObjectType:          hubspotOutputOption.ObjectType.ValueStringPointer(),
		Mode:                hubspotOutputOption.Mode.ValueStringPointer(),
		UpsertKey:           model.NewNullableString(hubspotOutputOption.UpsertKey),
		NumberOfParallels:   hubspotOutputOption.NumberOfParallels.ValueInt64Pointer(),
		Associations:        model.WrapObjectList(associations),
	}
}
