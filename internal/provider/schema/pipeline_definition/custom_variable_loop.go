package pipeline_definition

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
)

func CustomVariableLoop() schema.Attribute {
	return schema.SingleNestedAttribute{
		MarkdownDescription: "The custom variable loop of the pipeline definition",
		Optional:            true,
		Attributes: map[string]schema.Attribute{
			"type": schema.StringAttribute{
				MarkdownDescription: "The type of the custom variable loop",
				Required:            true,
			},
			"is_parallel_execution_allowed": schema.BoolAttribute{
				MarkdownDescription: "Whether parallel execution is allowed",
				Optional:            true,
				Computed:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"is_stopped_on_errors": schema.BoolAttribute{
				MarkdownDescription: "Whether the loop is stopped on errors",
				Optional:            true,
				Computed:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"max_errors": schema.Int64Attribute{
				MarkdownDescription: "The maximum number of errors",
				Optional:            true,
				Computed:            true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"string_config":    StringCustomVariableLoopConfig(),
			"period_config":    PeriodCustomVariableLoopConfig(),
			"bigquery_config":  BigqueryCustomVariableLoopConfig(),
			"snowflake_config": SnowflakeCustomVariableLoopConfig(),
			"redshift_config":  RedshiftCustomVariableLoopConfig(),
		},
	}
}