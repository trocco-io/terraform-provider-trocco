package pipeline_definition

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// CustomVariableLoop validates that the appropriate config is provided based on the type.
type CustomVariableLoop struct{}

// Description returns a plain text description of the validator's behavior.
func (v CustomVariableLoop) Description(ctx context.Context) string {
	return "Validates that the appropriate config is provided based on the type"
}

// MarkdownDescription returns a markdown formatted description of the validator's behavior.
func (v CustomVariableLoop) MarkdownDescription(ctx context.Context) string {
	return "Validates that the appropriate config is provided based on the type. When type is 'string', string_config is required. When type is 'period', period_config is required. When type is 'bigquery', bigquery_config is required. When type is 'snowflake', snowflake_config is required. When type is 'redshift', redshift_config is required."
}

// ValidateObject implements the validator.Object interface.
func (v CustomVariableLoop) ValidateObject(ctx context.Context, req validator.ObjectRequest, resp *validator.ObjectResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}

	attrs := req.ConfigValue.Attributes()

	typeAttr, ok := attrs["type"].(basetypes.StringValue)
	if !ok || typeAttr.IsNull() || typeAttr.IsUnknown() {
		return
	}

	typeValue := typeAttr.ValueString()

	typeToConfig := map[string]string{
		"string":    "string_config",
		"period":    "period_config",
		"bigquery":  "bigquery_config",
		"snowflake": "snowflake_config",
		"redshift":  "redshift_config",
	}

	configName, ok := typeToConfig[typeValue]
	if !ok {
		return
	}

	configAttr, exists := attrs[configName]
	if !exists || configAttr == nil {
		resp.Diagnostics.AddAttributeError(
			path.Root(configName),
			"Missing Required Configuration",
			fmt.Sprintf("When type is '%s', %s is required.", typeValue, configName),
		)
		return
	}

	if objectVal, ok := configAttr.(basetypes.ObjectValue); !ok || objectVal.IsNull() {
		resp.Diagnostics.AddAttributeError(
			path.Root(configName),
			"Missing Required Configuration",
			fmt.Sprintf("When type is '%s', %s must be a valid object.", typeValue, configName),
		)
	}
}
