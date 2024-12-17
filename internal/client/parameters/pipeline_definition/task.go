package workflow

type Task struct {
	Key            string `json:"key,omitempty"`
	TaskIdentifier int64  `json:"task_identifier,omitempty"`
	Type           string `json:"type,omitempty"`

	TroccoTransferConfig          *TroccoTransferTaskConfig          `json:"trocco_transfer_config,omitempty"`
	TroccoTransferBulkConfig      *TroccoTransferBulkTaskConfig      `json:"trocco_transfer_bulk_config,omitempty"`
	DBTConfig                     *DBTTaskConfig                     `json:"dbt_config,omitempty"`
	TroccoAgentConfig             *TroccoAgentTaskConfig             `json:"trocco_agent_config,omitempty"`
	TroccoBigQueryDatamartConfig  *TroccoBigQueryDatamartTaskConfig  `json:"trocco_bigquery_datamart_config,omitempty"`
	TroccoRedshiftDatamartConfig  *TroccoRedshiftDatamartTaskConfig  `json:"trocco_redshift_datamart_config,omitempty"`
	TroccoSnowflakeDatamartConfig *TroccoSnowflakeDatamartTaskConfig `json:"trocco_snowflake_datamart_config,omitempty"`
	WorkflowConfig                *TroccoPipelineTaskConfig          `json:"workflow_config,omitempty"`
	SlackNotificationConfig       *SlackNotificationTaskConfig       `json:"slack_notification_config,omitempty"`
	TableauDataExtractionConfig   *TableauDataExtractionTaskConfig   `json:"tableau_data_extraction_config,omitempty"`
	BigqueryDataCheckConfig       *BigqueryDataCheckTaskConfigInput  `json:"bigquery_data_check_config,omitempty"`
	SnowflakeDataCheckConfig      *SnowflakeDataCheckTaskConfigInput `json:"snowflake_data_check_config,omitempty"`
	RedshiftDataCheckConfig       *RedshiftDataCheckTaskConfigInput  `json:"redshift_data_check_config,omitempty"`
	HTTPRequestConfig             *HTTPRequestTaskConfig             `json:"http_request_config,omitempty"`
}
