package job_definition

import (
	planmodifier2 "terraform-provider-trocco/internal/provider/planmodifier"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
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
			"google_analytics4_input_option":   GoogleAnalytics4InputOptionSchema(),
			"http_input_option":                HTTPInputOptionSchema(),
			"kintone_input_option":             KintoneInputOptionSchema(),
			"yahoo_ads_api_yss_input_option":   YahooAdsApiYssInputOptionSchema(),
		},
		PlanModifiers: []planmodifier.Object{
			&planmodifier2.InputOptionPlanModifier{},
		},
	}
}
