package pipeline_definition

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

func Tasks() schema.Attribute {
	return schema.SetNestedAttribute{
		MarkdownDescription: "The tasks of the workflow.",
		Optional:            true,
		NestedObject: schema.NestedAttributeObject{
			Attributes: map[string]schema.Attribute{
				"key": schema.StringAttribute{
					Required: true,
				},
				"task_identifier": schema.Int64Attribute{
					Optional: true,
					Computed: true,
				},
				"type": schema.StringAttribute{
					Required: true,
					Validators: []validator.String{
						stringvalidator.OneOf(
							"trocco_transfer",
							"trocco_transfer_bulk",
							"trocco_bigquery_datamart",
							"trocco_dbt",
							"trocco_redshift_datamart",
							"trocco_snowflake_datamart",
							"trocco_agent",
							"trocco_pipeline",
							"slack_notify",
							"tableau_extract",
							"bigquery_data_check",
							"snowflake_data_check",
							"redshift_data_check",
							"http_request",
						),
					},
				},
				"trocco_transfer_config":           TroccoTransferTaskConfig(),
				"trocco_transfer_bulk_config":      TroccoTransferBulkTaskConfig(),
				"trocco_dbt_config":                TroccoDBTTaskConfig(),
				"trocco_agent_config":              TroccoAgentTaskConfig(),
				"trocco_bigquery_datamart_config":  BigQueryDatamartTaskConfig(),
				"trocco_redshift_datamart_config":  RedshiftDatamartTaskConfig(),
				"trocco_snowflake_datamart_config": SnowflakeDatamartTaskConfig(),
				"trocco_pipeline_config":           TroccoPiplineTaskConfig(),
				"slack_notification_config":        SlackNotificationTaskConfig(),
				"tableau_data_extraction_config":   TableauDataExtractionTaskConfig(),
				"http_request_config":              HTTPRequestTaskConfig(),
				"bigquery_data_check_config":       BigqueryDatacheckTaskConfig(),
				"snowflake_data_check_config":      SnowflakeDatacheckTaskConfig(),
				"redshift_data_check_config":       RedshiftDatacheckTaskConfig(),
			},
		},
	}
}
