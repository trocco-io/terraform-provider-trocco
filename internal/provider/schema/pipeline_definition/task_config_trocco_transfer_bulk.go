package pipeline_definition

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
)

func TroccoTransferBulkTaskConfigSchema() schema.Attribute {
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
				Default:             booldefault.StaticBool(false),
			},
			"is_stopped_on_errors": schema.BoolAttribute{
				MarkdownDescription: "Whether the task should stop on errors",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"max_errors": schema.Int64Attribute{
				MarkdownDescription: "The maximum number of errors allowed before the task stops",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(0),
			},
		},
	}
}
