package pipeline_definition

import (
	troccoSetPlanValidator "terraform-provider-trocco/internal/provider/validator/common/list"
	troccoPipelineDefinitionValidator "terraform-provider-trocco/internal/provider/validator/pipeline_definition"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

func TasksSchema() schema.Attribute {
	return schema.SetNestedAttribute{
		MarkdownDescription: "The tasks of the workflow.",
		Optional:            true,
		Validators: []validator.Set{
			troccoSetPlanValidator.UniqueObjectAttributeValue{
				AttributeName: "key",
			},
			troccoPipelineDefinitionValidator.TaskConfig{},
		},
		NestedObject: schema.NestedAttributeObject{
			Attributes: map[string]schema.Attribute{
				"key": schema.StringAttribute{
					MarkdownDescription: "The key of the task.",
					Required:            true,
					Validators: []validator.String{
						stringvalidator.LengthBetween(1, 100),
					},
				},
				"task_identifier": schema.Int64Attribute{
					MarkdownDescription: "The task identifier.",
					Computed:            true,
					PlanModifiers: []planmodifier.Int64{
						int64planmodifier.UseStateForUnknown(),
					},
				},
				"type": schema.StringAttribute{
					MarkdownDescription: "The type of the task.",
					Required:            true,
					Validators: []validator.String{
						stringvalidator.OneOf(
							"trocco_transfer",
							"trocco_transfer_bulk",
							"trocco_bigquery_datamart",
							"trocco_dbt",
							"trocco_redshift_datamart",
							"trocco_snowflake_datamart",
							"trocco_azure_synapse_analytics_datamart",
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
				"trocco_transfer_config":                         TroccoTransferTaskConfigSchema(),
				"trocco_transfer_bulk_config":                    TroccoTransferBulkTaskConfigSchema(),
				"trocco_dbt_config":                              TroccoDBTTaskConfigSchema(),
				"trocco_bigquery_datamart_config":                BigQueryDatamartTaskConfigSchema(),
				"trocco_redshift_datamart_config":                RedshiftDatamartTaskConfigSchema(),
				"trocco_snowflake_datamart_config":               SnowflakeDatamartTaskConfigSchema(),
				"trocco_azure_synapse_analytics_datamart_config": AzureSynapseAnalyticsDatamartTaskConfigSchema(),
				"trocco_pipeline_config":                         TroccoPipelineTaskConfigSchema(),
				"slack_notification_config":                      SlackNotificationTaskConfigSchema(),
				"tableau_data_extraction_config":                 TableauDataExtractionTaskConfigSchema(),
				"http_request_config":                            HTTPRequestTaskConfigSchema(),
				"bigquery_data_check_config":                     BigqueryDatacheckTaskConfigSchema(),
				"snowflake_data_check_config":                    SnowflakeDatacheckTaskConfigSchema(),
				"redshift_data_check_config":                     RedshiftDatacheckTaskConfigSchema(),
			},
		},
	}
}
