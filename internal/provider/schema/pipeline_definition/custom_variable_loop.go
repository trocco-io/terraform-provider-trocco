package pipeline_definition

import (
	troccoPipelineDefinitionValidator "terraform-provider-trocco/internal/provider/validator/pipeline_definition"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

func CustomVariableLoop() schema.Attribute {
	return schema.SingleNestedAttribute{
		MarkdownDescription: "The custom variable loop of the pipeline definition",
		Optional:            true,
		Validators: []validator.Object{
			troccoPipelineDefinitionValidator.CustomVariableLoop{},
		},
		Attributes: map[string]schema.Attribute{
			"type": schema.StringAttribute{
				MarkdownDescription: "The type of the custom variable loop. Allowed values: \"string\", \"period\", \"bigquery\", \"snowflake\", \"redshift\".",
				Required:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("string", "period", "bigquery", "snowflake", "redshift"),
				},
			},
			"is_parallel_execution_allowed": schema.BoolAttribute{
				MarkdownDescription: "Whether parallel execution is allowed",
				Optional:            true,
				Computed:            true,
			},
			"is_stopped_on_errors": schema.BoolAttribute{
				MarkdownDescription: "Whether the loop is stopped on errors",
				Optional:            true,
				Computed:            true,
			},
			"max_errors": schema.Int64Attribute{
				MarkdownDescription: "The maximum number of errors",
				Optional:            true,
				Computed:            true,
			},
			"string_config":    StringCustomVariableLoopConfig(),
			"period_config":    PeriodCustomVariableLoopConfig(),
			"bigquery_config":  BigqueryCustomVariableLoopConfig(),
			"snowflake_config": SnowflakeCustomVariableLoopConfig(),
			"redshift_config":  RedshiftCustomVariableLoopConfig(),
		},
	}
}
