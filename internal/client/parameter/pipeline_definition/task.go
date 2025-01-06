package pipeline_definition

type Task struct {
	Key            string `json:"key,omitempty"`
	TaskIdentifier int64  `json:"task_identifier,omitempty"`
	Type           string `json:"type,omitempty"`

	TroccoTransferConfig                      *TroccoTransferTaskConfig                      `json:"trocco_transfer_config,omitempty"`
	TroccoTransferBulkConfig                  *TroccoTransferBulkTaskConfig                  `json:"trocco_transfer_bulk_config,omitempty"`
	TroccoDBTConfig                           *TroccoDBTTaskConfig                           `json:"trocco_dbt_config,omitempty"`
	TroccoBigQueryDatamartConfig              *TroccoBigqueryDatamartTaskConfig              `json:"trocco_bigquery_datamart_config,omitempty"`
	TroccoRedshiftDatamartConfig              *TroccoRedshiftDatamartTaskConfig              `json:"trocco_redshift_datamart_config,omitempty"`
	TroccoSnowflakeDatamartConfig             *TroccoSnowflakeDatamartTaskConfig             `json:"trocco_snowflake_datamart_config,omitempty"`
	TroccoAzureSynapseAnalyticsDatamartConfig *TroccoAzureSynapseAnalyticsDatamartTaskConfig `json:"trocco_azure_synapse_analytics_datamart_config,omitempty"`
	TroccoPipelineConfig                      *TroccoPipelineTaskConfig                      `json:"trocco_pipeline_config,omitempty"`
	SlackNotificationConfig                   *SlackNotificationTaskConfig                   `json:"slack_notification_config,omitempty"`
	TableauDataExtractionConfig               *TableauDataExtractionTaskConfig               `json:"tableau_data_extraction_config,omitempty"`
	BigqueryDataCheckConfig                   *BigqueryDataCheckTaskConfigInput              `json:"bigquery_data_check_config,omitempty"`
	SnowflakeDataCheckConfig                  *SnowflakeDataCheckTaskConfigInput             `json:"snowflake_data_check_config,omitempty"`
	RedshiftDataCheckConfig                   *RedshiftDataCheckTaskConfigInput              `json:"redshift_data_check_config,omitempty"`
	HTTPRequestConfig                         *HTTPRequestTaskConfig                         `json:"http_request_config,omitempty"`
}
