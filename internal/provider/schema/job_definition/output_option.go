package job_definition

import (
	planmodifier2 "terraform-provider-trocco/internal/provider/planmodifier"

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
		},
		PlanModifiers: []planmodifier.Object{
			&planmodifier2.OutputOptionPlanModifier{},
		},
	}
}
