package pipeline_definition

import (
	we "terraform-provider-trocco/internal/client/entity/pipeline_definition"
	p "terraform-provider-trocco/internal/client/parameter"
	wp "terraform-provider-trocco/internal/client/parameter/pipeline_definition"
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
	CustomVariables []CustomVariable               `tfsdk:"custom_variables"`
}

func NewBigqueryDataCheckTaskConfig(c *we.BigqueryDataCheckTaskConfig) *BigqueryDataCheckTaskConfig {
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
		CustomVariables: NewCustomVariables(c.CustomVariables),
	}
}

func (c *BigqueryDataCheckTaskConfig) ToInput() *wp.BigqueryDataCheckTaskConfigInput {
	customVariables := []wp.CustomVariable{}
	for _, v := range c.CustomVariables {
		customVariables = append(customVariables, v.ToInput())
	}

	return &wp.BigqueryDataCheckTaskConfigInput{
		Name:            c.Name.ValueString(),
		ConnectionID:    c.ConnectionID.ValueInt64(),
		Query:           c.Query.ValueString(),
		Operator:        c.Operator.ValueString(),
		QueryResult:     &p.NullableInt64{Valid: !c.QueryResult.IsNull(), Value: c.QueryResult.ValueInt64()},
		AcceptsNull:     &p.NullableBool{Valid: !c.AcceptsNull.IsNull(), Value: c.AcceptsNull.ValueBool()},
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
