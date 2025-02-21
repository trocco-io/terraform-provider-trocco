package job_definition

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"

	planmodifier2 "terraform-provider-trocco/internal/provider/planmodifier"
)

func InputOptionSchema() schema.Attribute {
	return schema.SingleNestedAttribute{
		Required: true,
		Attributes: map[string]schema.Attribute{
			"mysql_input_option":               MysqlInputOptionSchema(),
			"gcs_input_option":                 GcsInputOptionSchema(),
			"snowflake_input_option":           SnowflakeInputOptionSchema(),
			"salesforce_input_option":          SalesforceInputOptionSchema(),
			"google_spreadsheets_input_option": GoogleSpreadsheetsInputOptionSchema(),
			"s3_input_option":                  S3InputOptionSchema(),
			"bigquery_input_option":            BigqueryInputOptionSchema(),
			"postgresql_input_option":          PostgresqlInputOptionSchema(),
		},
		PlanModifiers: []planmodifier.Object{
			&planmodifier2.InputOptionPlanModifier{},
		},
	}
}
