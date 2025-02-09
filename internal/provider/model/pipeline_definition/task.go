package pipeline_definition

import (
	we "terraform-provider-trocco/internal/client/entity/pipeline_definition"
	wp "terraform-provider-trocco/internal/client/parameter/pipeline_definition"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/samber/lo"
)

type Task struct {
	Key            types.String `tfsdk:"key"`
	TaskIdentifier types.Int64  `tfsdk:"task_identifier"`
	Type           types.String `tfsdk:"type"`

	BigqueryDataCheckConfig                   *BigqueryDataCheckTaskConfig                   `tfsdk:"bigquery_data_check_config"`
	HTTPRequestConfig                         *HTTPRequestTaskConfig                         `tfsdk:"http_request_config"`
	RedshiftDataCheckConfig                   *RedshiftDataCheckTaskConfig                   `tfsdk:"redshift_data_check_config"`
	SlackNotificationConfig                   *SlackNotificationTaskConfig                   `tfsdk:"slack_notification_config"`
	SnowflakeDataCheckConfig                  *SnowflakeDataCheckTaskConfig                  `tfsdk:"snowflake_data_check_config"`
	TableauDataExtractionConfig               *TableauDataExtractionTaskConfig               `tfsdk:"tableau_data_extraction_config"`
	TroccoBigQueryDatamartConfig              *TroccoBigqueryDatamartTaskConfig              `tfsdk:"trocco_bigquery_datamart_config"`
	TroccoDBTConfig                           *TroccoDBTTaskConfig                           `tfsdk:"trocco_dbt_config"`
	TroccoPipelineConfig                      *TroccoPipelineTaskConfig                      `tfsdk:"trocco_pipeline_config"`
	TroccoRedshiftDatamartConfig              *TroccoRedshiftDatamartTaskConfig              `tfsdk:"trocco_redshift_datamart_config"`
	TroccoSnowflakeDatamartConfig             *TroccoSnowflakeDatamartTaskConfig             `tfsdk:"trocco_snowflake_datamart_config"`
	TroccoAzureSynapseAnalyticsDatamartConfig *TroccoAzureSynapseAnalyticsDatamartTaskConfig `tfsdk:"trocco_azure_synapse_analytics_datamart_config"`
	TroccoTransferBulkConfig                  *TroccoTransferBulkTaskConfig                  `tfsdk:"trocco_transfer_bulk_config"`
	TroccoTransferConfig                      *TroccoTransferTaskConfig                      `tfsdk:"trocco_transfer_config"`
}

func NewTasks(ens []*we.Task, keys map[int64]types.String, previous *PipelineDefinition) []*Task {
	if ens == nil {
		return nil
	}

	// If the attribute in the plan (or state) is nil, the provider should sets nil to the state.
	if len(ens) == 0 && previous.Tasks == nil {
		return nil
	}

	tasks := []*Task{}
	for i, en := range ens {
		var previousTask *Task
		if len(previous.Tasks) > i {
			previousTask = previous.Tasks[i]
		}

		tasks = append(tasks, NewTask(en, keys, previousTask))
	}

	return tasks
}

func NewTask(en *we.Task, keys map[int64]types.String, previous *Task) *Task {
	if en == nil {
		return nil
	}

	// This function accepts keys as an argument.
	//
	// Keys are client-only data, so the APIs:
	//
	// - Cannot return them on `READ`
	// - Returns provided keys as is on `CREATE` and `UPDATE`
	//
	// Consequently, the client cannot set keys to an entity on `READ`.
	//
	// However, even if an entity lacks keys, the provider must set them to the
	// state. Therefore, on `READ`, the provider searches keys in the state by
	// identifiers and sets them to the state.
	//
	// Moreover, to simplify the code, on `CREATE` and `UPDATE`, the provider
	// creates a map of keys and identifiers from entity and searches keys
	// from the map.
	//
	// To archive the above behavior, this function accepts keys as an argument.

	var previousHTTPRequestConfig *HTTPRequestTaskConfig
	if previous != nil {
		previousHTTPRequestConfig = previous.HTTPRequestConfig
	}

	return &Task{
		Key:            keys[en.TaskIdentifier],
		TaskIdentifier: types.Int64Value(en.TaskIdentifier),
		Type:           types.StringValue(en.Type),

		TroccoTransferConfig:                      NewTroccoTransferTaskConfig(en.TroccoTransferConfig),
		TroccoTransferBulkConfig:                  NewTroccoTransferBulkTaskConfig(en.TroccoTransferBulkConfig),
		TroccoDBTConfig:                           NewTroccoDBTTaskConfig(en.TroccoDBTConfig),
		TroccoBigQueryDatamartConfig:              NewTroccoBigqueryDatamartTaskConfig(en.TroccoBigQueryDatamartConfig),
		TroccoRedshiftDatamartConfig:              NewTroccoRedshiftDatamartTaskConfig(en.TroccoRedshiftDatamartConfig),
		TroccoSnowflakeDatamartConfig:             NewTroccoSnowflakeDatamartTaskConfig(en.TroccoSnowflakeDatamartConfig),
		TroccoAzureSynapseAnalyticsDatamartConfig: NewTroccoAzureSynapseAnalyticsDatamartTaskConfig(en.TroccoAzureSynapseAnalyticsDatamartConfig),
		TroccoPipelineConfig:                      NewTroccoPipelineTaskConfig(en.TroccoPipelineTaskConfig),
		SlackNotificationConfig:                   NewSlackNotificationTaskConfig(en.SlackNotificationConfig),
		TableauDataExtractionConfig:               NewTableauDataExtractionTaskConfig(en.TableauDataExtractionConfig),
		BigqueryDataCheckConfig:                   NewBigqueryDataCheckTaskConfig(en.BigqueryDataCheckConfig),
		SnowflakeDataCheckConfig:                  NewSnowflakeDataCheckTaskConfig(en.SnowflakeDataCheckConfig),
		RedshiftDataCheckConfig:                   NewRedshiftDataCheckTaskConfig(en.RedshiftDataCheckConfig),
		HTTPRequestConfig:                         NewHTTPRequestTaskConfig(en.HTTPRequestConfig, previousHTTPRequestConfig),
	}
}

func (t *Task) ToInput(identifiers map[string]int64) *wp.Task {
	in := &wp.Task{
		Key:            t.Key.ValueString(),
		TaskIdentifier: lo.ValueOr(identifiers, t.Key.ValueString(), t.TaskIdentifier.ValueInt64()),
		Type:           t.Type.ValueString(),
	}

	if t.TroccoTransferConfig != nil {
		in.TroccoTransferConfig = t.TroccoTransferConfig.ToInput()
	}
	if t.TroccoTransferBulkConfig != nil {
		in.TroccoTransferBulkConfig = t.TroccoTransferBulkConfig.ToInput()
	}
	if t.TroccoDBTConfig != nil {
		in.TroccoDBTConfig = t.TroccoDBTConfig.ToInput()
	}
	if t.TroccoBigQueryDatamartConfig != nil {
		in.TroccoBigQueryDatamartConfig = t.TroccoBigQueryDatamartConfig.ToInput()
	}
	if t.TroccoRedshiftDatamartConfig != nil {
		in.TroccoRedshiftDatamartConfig = t.TroccoRedshiftDatamartConfig.ToInput()
	}
	if t.TroccoSnowflakeDatamartConfig != nil {
		in.TroccoSnowflakeDatamartConfig = t.TroccoSnowflakeDatamartConfig.ToInput()
	}
	if t.TroccoAzureSynapseAnalyticsDatamartConfig != nil {
		in.TroccoAzureSynapseAnalyticsDatamartConfig = t.TroccoAzureSynapseAnalyticsDatamartConfig.ToInput()
	}
	if t.TroccoPipelineConfig != nil {
		in.TroccoPipelineConfig = t.TroccoPipelineConfig.ToInput()
	}
	if t.SlackNotificationConfig != nil {
		in.SlackNotificationConfig = t.SlackNotificationConfig.ToInput()
	}
	if t.TableauDataExtractionConfig != nil {
		in.TableauDataExtractionConfig = t.TableauDataExtractionConfig.ToInput()
	}
	if t.BigqueryDataCheckConfig != nil {
		in.BigqueryDataCheckConfig = t.BigqueryDataCheckConfig.ToInput()
	}
	if t.SnowflakeDataCheckConfig != nil {
		in.SnowflakeDataCheckConfig = t.SnowflakeDataCheckConfig.ToInput()
	}
	if t.RedshiftDataCheckConfig != nil {
		in.RedshiftDataCheckConfig = t.RedshiftDataCheckConfig.ToInput()
	}
	if t.HTTPRequestConfig != nil {
		in.HTTPRequestConfig = t.HTTPRequestConfig.ToInput()
	}

	return in
}
