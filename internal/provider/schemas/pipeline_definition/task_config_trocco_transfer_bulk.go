package pipeline_definition

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

func NewTroccoTransferBulkTaskConfigAttribute() schema.Attribute {
	return schema.SingleNestedAttribute{
		Optional: true,
		Attributes: map[string]schema.Attribute{
			"definition_id": schema.Int64Attribute{
				Required: true,
			},
			"is_parallel_execution_allowed": schema.BoolAttribute{
				Optional: true,
				Computed: true,
			},
			"is_stopped_on_errors": schema.BoolAttribute{
				Optional: true,
				Computed: true,
			},
			"max_errors": schema.Int64Attribute{
				Optional: true,
				Computed: true,
			},
		},
	}
}
