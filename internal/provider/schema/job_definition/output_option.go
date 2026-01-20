package job_definition

import (
	planModifier "terraform-provider-trocco/internal/provider/planmodifier"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
)

func OutputOptionSchema() schema.Attribute {
	return schema.SingleNestedAttribute{
		Required: true,
		Attributes: map[string]schema.Attribute{
			"bigquery_output_option":            BigqueryOutputOptionSchema(),
			"snowflake_output_option":           SnowflakeOutputOptionSchema(),
			"salesforce_output_option":          SalesforceOutputOptionSchema(),
			"google_spreadsheets_output_option": GoogleSpreadsheetsOutputOptionSchema(),
			"databricks_output_option":          DatabricksOutputOptionSchema(),
			"mysql_output_option":               MysqlOutputOptionSchema(),
		},
		PlanModifiers: []planmodifier.Object{
			&planModifier.OutputOptionPlanModifier{},
		},
	}
}
