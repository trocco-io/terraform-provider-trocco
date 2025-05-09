package pipeline_definition

type Task struct {
	Key            string `json:"key"`
	TaskIdentifier int64  `json:"task_identifier"`
	Type           string `json:"type"`

	TroccoTransferConfig                      *TroccoTransferTaskConfig                      `json:"trocco_transfer_config"`
	TroccoTransferBulkConfig                  *TroccoTransferBulkTaskConfig                  `json:"trocco_transfer_bulk_config"`
	TroccoDBTConfig                           *TroccoDBTTaskConfig                           `json:"trocco_dbt_config"`
	TroccoPipelineTaskConfig                  *TroccoPipelineTaskConfig                      `json:"trocco_pipeline_config"`
	TroccoBigQueryDatamartConfig              *TroccoBigqueryDatamartTaskConfig              `json:"trocco_bigquery_datamart_config"`
	TroccoRedshiftDatamartConfig              *TroccoRedshiftDatamartTaskConfig              `json:"trocco_redshift_datamart_config"`
	TroccoSnowflakeDatamartConfig             *TroccoSnowflakeDatamartTaskConfig             `json:"trocco_snowflake_datamart_config"`
	TroccoAzureSynapseAnalyticsDatamartConfig *TroccoAzureSynapseAnalyticsDatamartTaskConfig `json:"trocco_azure_synapse_analytics_datamart_config"`
	SlackNotificationConfig                   *SlackNotificationTaskConfig                   `json:"slack_notification_config"`
	TableauDataExtractionConfig               *TableauDataExtractionTaskConfig               `json:"tableau_data_extraction_config"`
	BigqueryDataCheckConfig                   *BigqueryDataCheckTaskConfig                   `json:"bigquery_data_check_config"`
	SnowflakeDataCheckConfig                  *SnowflakeDataCheckTaskConfig                  `json:"snowflake_data_check_config"`
	RedshiftDataCheckConfig                   *RedshiftDataCheckTaskConfig                   `json:"redshift_data_check_config"`
	HTTPRequestConfig                         *HTTPRequestTaskConfig                         `json:"http_request_config"`
}
