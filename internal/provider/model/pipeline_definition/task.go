package pipeline_definition

import (
	"context"
	pipelineDefinitionEntities "terraform-provider-trocco/internal/client/entity/pipeline_definition"
	pipelineDefinitionParameters "terraform-provider-trocco/internal/client/parameter/pipeline_definition"

	"github.com/hashicorp/terraform-plugin-framework/attr"
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
	IfElseConfig                              *IfElseTaskConfig                              `tfsdk:"if_else_config"`
}

func NewTasks(ctx context.Context, ens []*pipelineDefinitionEntities.Task, keys map[int64]types.String, previous *PipelineDefinition) types.Set {
	var TaskObjectType = types.ObjectType{
		AttrTypes: TaskObjectAttrTypes(),
	}

	if ens == nil {
		return types.SetNull(TaskObjectType)
	}

	var previousTasks []*Task
	if previous != nil && !previous.Tasks.IsNull() && !previous.Tasks.IsUnknown() {
		if diags := previous.Tasks.ElementsAs(ctx, &previousTasks, false); diags.HasError() {
			return types.SetNull(TaskObjectType)
		}
	}

	if len(ens) == 0 && previousTasks == nil {
		return types.SetNull(TaskObjectType)
	}

	tasks := []*Task{}
	for i, en := range ens {
		var previousTask *Task
		if len(previousTasks) > i {
			previousTask = previousTasks[i]
		}
		tasks = append(tasks, NewTask(ctx, en, keys, previousTask))
	}

	set, diags := types.SetValueFrom(ctx, TaskObjectType, tasks)
	if diags.HasError() {
		return types.SetNull(TaskObjectType)
	}
	return set
}

func NewTask(ctx context.Context, en *pipelineDefinitionEntities.Task, keys map[int64]types.String, previous *Task) *Task {
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

		TroccoTransferConfig:                      NewTroccoTransferTaskConfig(ctx, en.TroccoTransferConfig),
		TroccoTransferBulkConfig:                  NewTroccoTransferBulkTaskConfig(en.TroccoTransferBulkConfig),
		TroccoDBTConfig:                           NewTroccoDBTTaskConfig(en.TroccoDBTConfig),
		TroccoBigQueryDatamartConfig:              NewTroccoBigqueryDatamartTaskConfig(ctx, en.TroccoBigQueryDatamartConfig),
		TroccoRedshiftDatamartConfig:              NewTroccoRedshiftDatamartTaskConfig(ctx, en.TroccoRedshiftDatamartConfig),
		TroccoSnowflakeDatamartConfig:             NewTroccoSnowflakeDatamartTaskConfig(ctx, en.TroccoSnowflakeDatamartConfig),
		TroccoAzureSynapseAnalyticsDatamartConfig: NewTroccoAzureSynapseAnalyticsDatamartTaskConfig(ctx, en.TroccoAzureSynapseAnalyticsDatamartConfig),
		TroccoPipelineConfig:                      NewTroccoPipelineTaskConfig(ctx, en.TroccoPipelineTaskConfig),
		SlackNotificationConfig:                   NewSlackNotificationTaskConfig(en.SlackNotificationConfig),
		TableauDataExtractionConfig:               NewTableauDataExtractionTaskConfig(en.TableauDataExtractionConfig),
		BigqueryDataCheckConfig:                   NewBigqueryDataCheckTaskConfig(ctx, en.BigqueryDataCheckConfig),
		SnowflakeDataCheckConfig:                  NewSnowflakeDataCheckTaskConfig(ctx, en.SnowflakeDataCheckConfig),
		RedshiftDataCheckConfig:                   NewRedshiftDataCheckTaskConfig(ctx, en.RedshiftDataCheckConfig),
		HTTPRequestConfig:                         NewHTTPRequestTaskConfig(ctx, en.HTTPRequestConfig, previousHTTPRequestConfig),
		IfElseConfig:                              NewIfElseTaskConfig(ctx, en.IfElseConfig),
	}
}

func (t *Task) ToInput(ctx context.Context, identifiers map[string]int64) *pipelineDefinitionParameters.Task {
	in := &pipelineDefinitionParameters.Task{
		Key:            t.Key.ValueString(),
		TaskIdentifier: lo.ValueOr(identifiers, t.Key.ValueString(), t.TaskIdentifier.ValueInt64()),
		Type:           t.Type.ValueString(),
	}

	if t.TroccoTransferConfig != nil {
		in.TroccoTransferConfig = t.TroccoTransferConfig.ToInput(ctx)
	}
	if t.TroccoTransferBulkConfig != nil {
		in.TroccoTransferBulkConfig = t.TroccoTransferBulkConfig.ToInput()
	}
	if t.TroccoDBTConfig != nil {
		in.TroccoDBTConfig = t.TroccoDBTConfig.ToInput()
	}
	if t.TroccoBigQueryDatamartConfig != nil {
		in.TroccoBigQueryDatamartConfig = t.TroccoBigQueryDatamartConfig.ToInput(ctx)
	}
	if t.TroccoRedshiftDatamartConfig != nil {
		in.TroccoRedshiftDatamartConfig = t.TroccoRedshiftDatamartConfig.ToInput(ctx)
	}
	if t.TroccoSnowflakeDatamartConfig != nil {
		in.TroccoSnowflakeDatamartConfig = t.TroccoSnowflakeDatamartConfig.ToInput(ctx)
	}
	if t.TroccoAzureSynapseAnalyticsDatamartConfig != nil {
		in.TroccoAzureSynapseAnalyticsDatamartConfig = t.TroccoAzureSynapseAnalyticsDatamartConfig.ToInput(ctx)
	}
	if t.TroccoPipelineConfig != nil {
		in.TroccoPipelineConfig = t.TroccoPipelineConfig.ToInput(ctx)
	}
	if t.SlackNotificationConfig != nil {
		in.SlackNotificationConfig = t.SlackNotificationConfig.ToInput()
	}
	if t.TableauDataExtractionConfig != nil {
		in.TableauDataExtractionConfig = t.TableauDataExtractionConfig.ToInput()
	}
	if t.BigqueryDataCheckConfig != nil {
		in.BigqueryDataCheckConfig = t.BigqueryDataCheckConfig.ToInput(ctx)
	}
	if t.SnowflakeDataCheckConfig != nil {
		in.SnowflakeDataCheckConfig = t.SnowflakeDataCheckConfig.ToInput(ctx)
	}
	if t.RedshiftDataCheckConfig != nil {
		in.RedshiftDataCheckConfig = t.RedshiftDataCheckConfig.ToInput(ctx)
	}
	if t.HTTPRequestConfig != nil {
		in.HTTPRequestConfig = t.HTTPRequestConfig.ToInput(ctx)
	}
	if t.IfElseConfig != nil {
		in.IfElseConfig = t.IfElseConfig.ToInput(ctx)
	}

	return in
}

func TaskObjectAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"key":             types.StringType,
		"task_identifier": types.Int64Type,
		"type":            types.StringType,

		"bigquery_data_check_config": types.ObjectType{
			AttrTypes: BigqueryDataCheckTaskConfigAttrTypes(),
		},
		"http_request_config": types.ObjectType{
			AttrTypes: HTTPRequestTaskConfigAttrTypes(),
		},
		"redshift_data_check_config": types.ObjectType{
			AttrTypes: RedshiftDataCheckTaskConfigAttrTypes(),
		},
		"slack_notification_config": types.ObjectType{
			AttrTypes: SlackNotificationTaskConfigAttrTypes(),
		},
		"snowflake_data_check_config": types.ObjectType{
			AttrTypes: SnowflakeDataCheckTaskConfigAttrTypes(),
		},
		"tableau_data_extraction_config": types.ObjectType{
			AttrTypes: TableauDataExtractionTaskConfigAttrTypes(),
		},
		"trocco_bigquery_datamart_config": types.ObjectType{
			AttrTypes: TroccoBigqueryDatamartTaskConfigAttrTypes(),
		},
		"trocco_dbt_config": types.ObjectType{
			AttrTypes: TroccoDBTTaskConfigAttrTypes(),
		},
		"trocco_pipeline_config": types.ObjectType{
			AttrTypes: TroccoPipelineTaskConfigAttrTypes(),
		},
		"trocco_redshift_datamart_config": types.ObjectType{
			AttrTypes: TroccoRedshiftDatamartTaskConfigAttrTypes(),
		},
		"trocco_snowflake_datamart_config": types.ObjectType{
			AttrTypes: TroccoSnowflakeDatamartTaskConfigAttrTypes(),
		},
		"trocco_azure_synapse_analytics_datamart_config": types.ObjectType{
			AttrTypes: TroccoAzureSynapseAnalyticsDatamartTaskConfigAttrTypes(),
		},
		"trocco_transfer_bulk_config": types.ObjectType{
			AttrTypes: TroccoTransferBulkTaskConfigAttrTypes(),
		},
		"trocco_transfer_config": types.ObjectType{
			AttrTypes: TroccoTransferTaskConfigAttrTypes(),
		},
		"if_else_config": types.ObjectType{
			AttrTypes: IfElseTaskConfigAttrTypes(),
		},
	}
}
