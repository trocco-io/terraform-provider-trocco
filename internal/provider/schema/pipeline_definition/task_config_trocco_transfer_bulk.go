package pipeline_definition

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
)

func TroccoTransferBulkTaskConfig() schema.Attribute {
	return schema.SingleNestedAttribute{
		MarkdownDescription: "The task configuration for the trocco transfer bulk task.",
		Optional:            true,
		Attributes: map[string]schema.Attribute{
			"definition_id": schema.Int64Attribute{
				MarkdownDescription: "The definition id to use for the trocco transfer bulk task",
				Required:            true,
			},
			"is_parallel_execution_allowed": schema.BoolAttribute{
				MarkdownDescription: "Whether the task is allowed to run in parallel",
				Optional:            true,
				Computed:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"is_stopped_on_errors": schema.BoolAttribute{
				MarkdownDescription: "Whether the task should stop on errors",
				Optional:            true,
				Computed:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"max_errors": schema.Int64Attribute{
				MarkdownDescription: "The maximum number of errors allowed before the task stops",
				Optional:            true,
				Computed:            true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}
