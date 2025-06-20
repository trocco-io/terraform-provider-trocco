package job_definitions

import (
	"terraform-provider-trocco/internal/client/entity"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type Label struct {
	ID   types.Int64  `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

func NewLabels(labels []entity.Label) []Label {
	if labels == nil {
		return nil
	}

	outputs := make([]Label, 0, len(labels))
	for _, input := range labels {
		label := Label{
			ID:   types.Int64Value(input.ID),
			Name: types.StringValue(input.Name),
		}
		outputs = append(outputs, label)
	}
	return outputs
}

func (l Label) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"id":   types.Int64Type,
		"name": types.StringType,
	}
}
