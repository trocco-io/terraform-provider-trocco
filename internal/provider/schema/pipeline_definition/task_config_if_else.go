package pipeline_definition

import (
	pipelineDefinitionValidator "terraform-provider-trocco/internal/provider/validator/pipeline_definition"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func IfElseTaskConfigSchema() schema.Attribute {
	return schema.SingleNestedAttribute{
		MarkdownDescription: "The task configuration for the if-else task.",
		Optional:            true,
		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				MarkdownDescription: "The name of the task",
				Required:            true,
			},
			"condition_groups": schema.SingleNestedAttribute{
				MarkdownDescription: "The condition groups for the if-else task",
				Required:            true,
				Attributes: map[string]schema.Attribute{
					"set_type": schema.StringAttribute{
						MarkdownDescription: "The type of condition set (and, or)",
						Required:            true,
						Validators: []validator.String{
							stringvalidator.OneOf("and", "or"),
						},
					},
					"conditions": schema.ListNestedAttribute{
						MarkdownDescription: "The list of conditions",
						Required:            true,
						Validators: []validator.List{
							listvalidator.SizeAtLeast(1),
						},
						NestedObject: schema.NestedAttributeObject{
							Validators: []validator.Object{
								pipelineDefinitionValidator.ConditionTaskKey{},
							},
							Attributes: map[string]schema.Attribute{
								"variable": schema.StringAttribute{
									MarkdownDescription: "The variable to check (e.g., current_time, environment, status, response_status_code, transfer_record_count, check_result)",
									Required:            true,
									Validators: []validator.String{
										stringvalidator.OneOf(
											"current_time",
											"environment",
											"status",
											"response_status_code",
											"transfer_record_count",
											"check_result",
										),
									},
								},
								"task_key": schema.StringAttribute{
									MarkdownDescription: "The task key (required for task-scoped variables like status, response_status_code, transfer_record_count, check_result)",
									Optional:            true,
								},
								"operator": schema.StringAttribute{
									MarkdownDescription: "The operator for comparison",
									Required:            true,
									Validators: []validator.String{
										stringvalidator.OneOf(
											"equal",
											"not_equal",
											"greater",
											"greater_equal",
											"less",
											"less_equal",
										),
									},
								},
								"value": schema.StringAttribute{
									MarkdownDescription: "The value to compare against",
									Required:            true,
								},
							},
						},
					},
				},
			},
			"destinations": schema.SingleNestedAttribute{
				MarkdownDescription: "The destination tasks for the if and else branches",
				Required:            true,
				Validators: []validator.Object{
					pipelineDefinitionValidator.DestinationsNotBothEmpty{},
				},
				Attributes: map[string]schema.Attribute{
					"if": schema.ListAttribute{
						MarkdownDescription: "The list of task keys to execute when the condition is true. Specify an empty list `[]` if not needed.",
						ElementType:         types.StringType,
						Required:            true,
					},
					"else": schema.ListAttribute{
						MarkdownDescription: "The list of task keys to execute when the condition is false. Specify an empty list `[]` if not needed.",
						ElementType:         types.StringType,
						Required:            true,
					},
				},
			},
		},
	}
}
