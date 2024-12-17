package pipeline_definition

import (
	we "terraform-provider-trocco/internal/client/entities/pipeline_definition"
	p "terraform-provider-trocco/internal/client/parameters"
	wp "terraform-provider-trocco/internal/client/parameters/pipeline_definition"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

type BigqueryDataCheckTaskConfig struct {
	Name            types.String     `tfsdk:"name"`
	ConnectionID    types.Int64      `tfsdk:"connection_id"`
	Query           types.String     `tfsdk:"query"`
	Operator        types.String     `tfsdk:"operator"`
	QueryResult     types.Int64      `tfsdk:"query_result"`
	AcceptsNull     types.Bool       `tfsdk:"accepts_null"`
	CustomVariables []CustomVariable `tfsdk:"custom_variables"`
}

func NewBigqueryDataCheckTaskConfig(c *we.BigqueryDataCheckTaskConfig) *BigqueryDataCheckTaskConfig {
	if c == nil {
		return nil
	}

	return &BigqueryDataCheckTaskConfig{
		Name:            types.StringValue(c.Name),
		ConnectionID:    types.Int64Value(c.ConnectionID),
		Query:           types.StringValue(c.Query),
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
