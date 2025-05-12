package pipeline_definition

import (
	we "terraform-provider-trocco/internal/client/entity/pipeline_definition"
	p "terraform-provider-trocco/internal/client/parameter"
	wp "terraform-provider-trocco/internal/client/parameter/pipeline_definition"
	"terraform-provider-trocco/internal/provider/custom_type"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

type SnowflakeDataCheckTaskConfig struct {
	Name            types.String                   `tfsdk:"name"`
	ConnectionID    types.Int64                    `tfsdk:"connection_id"`
	Query           custom_type.TrimmedStringValue `tfsdk:"query"`
	Operator        types.String                   `tfsdk:"operator"`
	QueryResult     types.Int64                    `tfsdk:"query_result"`
	AcceptsNull     types.Bool                     `tfsdk:"accepts_null"`
	Warehouse       types.String                   `tfsdk:"warehouse"`
	CustomVariables []CustomVariable               `tfsdk:"custom_variables"`
}

func NewSnowflakeDataCheckTaskConfig(c *we.SnowflakeDataCheckTaskConfig) *SnowflakeDataCheckTaskConfig {
	if c == nil {
		return nil
	}

	return &SnowflakeDataCheckTaskConfig{
		Name:            types.StringValue(c.Name),
		ConnectionID:    types.Int64Value(c.ConnectionID),
		Query:           custom_type.TrimmedStringValue{types.StringValue(c.Query)},
		Operator:        types.StringValue(c.Operator),
		QueryResult:     types.Int64Value(c.QueryResult),
		AcceptsNull:     types.BoolValue(c.AcceptsNull),
		Warehouse:       types.StringValue(c.Warehouse),
		CustomVariables: NewCustomVariables(c.CustomVariables),
	}
}

func (c *SnowflakeDataCheckTaskConfig) ToInput() *wp.SnowflakeDataCheckTaskConfigInput {
	customVariables := []wp.CustomVariable{}
	for _, v := range c.CustomVariables {
		customVariables = append(customVariables, v.ToInput())
	}

	return &wp.SnowflakeDataCheckTaskConfigInput{
		Name:            c.Name.ValueString(),
		ConnectionID:    c.ConnectionID.ValueInt64(),
		Query:           c.Query.ValueString(),
		Operator:        c.Operator.ValueString(),
		QueryResult:     &p.NullableInt64{Valid: !c.QueryResult.IsNull(), Value: c.QueryResult.ValueInt64()},
		AcceptsNull:     &p.NullableBool{Valid: !c.AcceptsNull.IsNull(), Value: c.AcceptsNull.ValueBool()},
		Warehouse:       c.Warehouse.ValueString(),
		CustomVariables: customVariables,
	}
}
