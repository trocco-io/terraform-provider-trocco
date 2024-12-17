package workflow

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

func NewTroccoAgentTaskConfigAttribute() schema.Attribute {
	return schema.SingleNestedAttribute{
		Optional: true,
		Attributes: map[string]schema.Attribute{
			"definition_id": schema.Int64Attribute{
				Required: true,
			},
		},
	}
}
