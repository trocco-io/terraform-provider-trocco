package workflow

import (
	we "terraform-provider-trocco/internal/client/entities/pipeline_definition"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

type Task struct {
	Key            types.String `tfsdk:"key"`
	TaskIdentifier types.Int64  `tfsdk:"task_identifier"`
	Type           types.String `tfsdk:"type"`

	BigqueryDataCheckConfig       *BigqueryDataCheckTaskConfig       `tfsdk:"bigquery_data_check_config"`
	DBTConfig                     *DBTTaskConfig                     `tfsdk:"dbt_config"`
	HTTPRequestConfig             *HTTPRequestTaskConfig             `tfsdk:"http_request_config"`
	RedshiftDataCheckConfig       *RedshiftDataCheckTaskConfig       `tfsdk:"redshift_data_check_config"`
	SlackNotificationConfig       *SlackNotificationTaskConfig       `tfsdk:"slack_notification_config"`
	SnowflakeDataCheckConfig      *SnowflakeDataCheckTaskConfig      `tfsdk:"snowflake_data_check_config"`
	TableauDataExtractionConfig   *TableauDataExtractionTaskConfig   `tfsdk:"tableau_data_extraction_config"`
	TroccoAgentConfig             *TroccoAgentTaskConfig             `tfsdk:"trocco_agent_config"`
	TroccoBigQueryDatamartConfig  *TroccoBigqueryDatamartTaskConfig  `tfsdk:"trocco_bigquery_datamart_config"`
	TroccoRedshiftDatamartConfig  *TroccoRedshiftDatamartTaskConfig  `tfsdk:"trocco_redshift_datamart_config"`
	TroccoSnowflakeDatamartConfig *TroccoSnowflakeDatamartTaskConfig `tfsdk:"trocco_snowflake_datamart_config"`
	TroccoTransferBulkConfig      *TroccoTransferBulkTaskConfig      `tfsdk:"trocco_transfer_bulk_config"`
	TroccoTransferConfig          *TroccoTransferTaskConfig          `tfsdk:"trocco_transfer_config"`
	WorkflowConfig                *WorkflowTaskConfig                `tfsdk:"workflow_config"`
}

func NewTasks(ens []*we.Task) []*Task {
	if ens == nil {
		return nil
	}

	var tasks []*Task
	for _, en := range ens {
		tasks = append(tasks, NewTask(en))
	}

	return tasks
}

func NewTask(en *we.Task) *Task {
	if en == nil {
		return nil
	}

	return &Task{
		Key:            types.StringValue(en.Key),
		TaskIdentifier: types.Int64Value(en.TaskIdentifier),
		Type:           types.StringValue(en.Type),

		TroccoTransferConfig:          NewTroccoTransferTaskConfig(en.TroccoTransferConfig),
		TroccoTransferBulkConfig:      NewTroccoTransferBulkTaskConfig(en.TroccoTransferBulkConfig),
		DBTConfig:                     NewDBTTaskConfig(en.DBTConfig),
		TroccoAgentConfig:             NewTroccoAgentTaskConfig(en.TroccoAgentConfig),
		TroccoBigQueryDatamartConfig:  NewTroccoBigQueryDatamartTaskConfig(en.TroccoBigQueryDatamartConfig),
		TroccoRedshiftDatamartConfig:  NewTroccoRedshiftDatamartTaskConfig(en.TroccoRedshiftDatamartConfig),
		TroccoSnowflakeDatamartConfig: NewTroccoSnowflakeDatamartTaskConfig(en.TroccoSnowflakeDatamartConfig),
		WorkflowConfig:                NewWorkflowTaskConfig(en.WorkflowConfig),
		SlackNotificationConfig:       NewSlackNotificationTaskConfig(en.SlackNotificationConfig),
		TableauDataExtractionConfig:   NewTableauDataExtractionTaskConfig(en.TableauDataExtractionConfig),
		BigqueryDataCheckConfig:       NewBigqueryDataCheckTaskConfig(en.BigqueryDataCheckConfig),
		SnowflakeDataCheckConfig:      NewSnowflakeDataCheckTaskConfig(en.SnowflakeDataCheckConfig),
		RedshiftDataCheckConfig:       NewRedshiftDataCheckTaskConfig(en.RedshiftDataCheckConfig),
		HTTPRequestConfig:             NewHTTPRequestTaskConfig(en.HTTPRequestConfig),
	}
}
