package pipeline_definition

import (
	"context"
	pipelineDefinitionEntities "terraform-provider-trocco/internal/client/entity/pipeline_definition"
	parameter "terraform-provider-trocco/internal/client/parameter"
	pipelineDefinitionParameters "terraform-provider-trocco/internal/client/parameter/pipeline_definition"
	"terraform-provider-trocco/internal/provider/custom_type"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type BigqueryDataCheckTaskConfig struct {
	Name            types.String                   `tfsdk:"name"`
	ConnectionID    types.Int64                    `tfsdk:"connection_id"`
	Query           custom_type.TrimmedStringValue `tfsdk:"query"`
	Operator        types.String                   `tfsdk:"operator"`
	QueryResult     types.Int64                    `tfsdk:"query_result"`
	AcceptsNull     types.Bool                     `tfsdk:"accepts_null"`
	CustomVariables types.Set                      `tfsdk:"custom_variables"`
}

func NewBigqueryDataCheckTaskConfig(ctx context.Context, c *pipelineDefinitionEntities.BigqueryDataCheckTaskConfig) *BigqueryDataCheckTaskConfig {
	if c == nil {
		return nil
	}

	return &BigqueryDataCheckTaskConfig{
		Name:            types.StringValue(c.Name),
		ConnectionID:    types.Int64Value(c.ConnectionID),
		Query:           custom_type.TrimmedStringValue{StringValue: types.StringValue(c.Query)},
		Operator:        types.StringValue(c.Operator),
		QueryResult:     types.Int64Value(c.QueryResult),
		AcceptsNull:     types.BoolValue(c.AcceptsNull),
		CustomVariables: NewCustomVariables(ctx, c.CustomVariables),
	}
}

func (c *BigqueryDataCheckTaskConfig) ToInput(ctx context.Context) *pipelineDefinitionParameters.BigqueryDataCheckTaskConfigInput {
	customVariables := []pipelineDefinitionParameters.CustomVariable{}
	if !c.CustomVariables.IsNull() && !c.CustomVariables.IsUnknown() {
		var customVariableValues []CustomVariable
		diags := c.CustomVariables.ElementsAs(ctx, &customVariableValues, false)
		if !diags.HasError() {
			for _, v := range customVariableValues {
				customVariables = append(customVariables, v.ToInput())
			}
		}
	}

	return &pipelineDefinitionParameters.BigqueryDataCheckTaskConfigInput{
		Name:            c.Name.ValueString(),
		ConnectionID:    c.ConnectionID.ValueInt64(),
		Query:           c.Query.ValueString(),
		Operator:        c.Operator.ValueString(),
		QueryResult:     &parameter.NullableInt64{Valid: !c.QueryResult.IsNull(), Value: c.QueryResult.ValueInt64()},
		AcceptsNull:     &parameter.NullableBool{Valid: !c.AcceptsNull.IsNull(), Value: c.AcceptsNull.ValueBool()},
		CustomVariables: customVariables,
	}
}

func BigqueryDataCheckTaskConfigAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"name":             types.StringType,
		"connection_id":    types.Int64Type,
		"query":            custom_type.TrimmedStringType{},
		"operator":         types.StringType,
		"query_result":     types.Int64Type,
		"accepts_null":     types.BoolType,
		"custom_variables": types.SetType{ElemType: types.ObjectType{AttrTypes: CustomVariableAttrTypes()}},
	}
}
