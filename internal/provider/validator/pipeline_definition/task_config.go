package pipeline_definition

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/samber/lo"
)

var _ validator.Set = TaskConfig{}

var taskConfigKeys = map[string]string{
	"trocco_transfer":                         "trocco_transfer_config",
	"trocco_transfer_bulk":                    "trocco_transfer_bulk_config",
	"trocco_bigquery_datamart":                "trocco_bigquery_datamart_config",
	"trocco_dbt":                              "trocco_dbt_config",
	"trocco_redshift_datamart":                "trocco_redshift_datamart_config",
	"trocco_snowflake_datamart":               "trocco_snowflake_datamart_config",
	"trocco_azure_synapse_analytics_datamart": "trocco_azure_synapse_analytics_datamart_config",
	"trocco_pipeline":                         "trocco_pipeline_config",
	"slack_notify":                            "slack_notification_config",
	"tableau_extract":                         "tableau_data_extraction_config",
	"bigquery_data_check":                     "bigquery_data_check_config",
	"snowflake_data_check":                    "snowflake_data_check_config",
	"redshift_data_check":                     "redshift_data_check_config",
	"http_request":                            "http_request_config",
}

type TaskConfig struct {
}

func (v TaskConfig) Description(ctx context.Context) string {
	return "Ensures the task configuration is present for the specified task type."
}

func (v TaskConfig) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v TaskConfig) ValidateSet(
	ctx context.Context,
	req validator.SetRequest,
	resp *validator.SetResponse,
) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}

	var objects []types.Object
	if diags := req.ConfigValue.ElementsAs(ctx, &objects, false); diags.HasError() {
		resp.Diagnostics.Append(diags...)

		return
	}

	for _, object := range objects {
		currentTaskTypeAttribute, ok := object.Attributes()["type"].(types.String)
		if !ok {
			resp.Diagnostics.AddError(
				"Invalid Task Type",
				"Task type is invalid",
			)
		}

		currentTaskTypeValue := currentTaskTypeAttribute.ValueString()

		if !lo.Contains(lo.Keys(taskConfigKeys), currentTaskTypeValue) {
			resp.Diagnostics.AddError(
				"Invalid Task Type",
				fmt.Sprintf("Task type %s is invalid", currentTaskTypeValue),
			)

			return
		}

		for taskType, taskConfigKey := range taskConfigKeys {
			if taskType == currentTaskTypeValue {
				taskConfig := object.Attributes()[taskConfigKey]
				if taskConfig.IsUnknown() || taskConfig.IsNull() {
					resp.Diagnostics.AddError(
						"Missing Task Configuration",
						fmt.Sprintf("Task configuration %s is missing", taskConfigKey),
					)
				}
			} else {
				taskConfig := object.Attributes()[taskConfigKey]
				if !taskConfig.IsUnknown() && !taskConfig.IsNull() {
					resp.Diagnostics.AddError(
						"Unexpected Task Configuration",
						fmt.Sprintf("Task configuration %s is unexpected", taskConfigKey),
					)
				}
			}
		}
	}
}
