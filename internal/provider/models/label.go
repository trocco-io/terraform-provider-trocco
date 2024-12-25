package models

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"terraform-provider-trocco/internal/client/entities"
)

type LabelModel struct {
	ID   types.Int64  `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

func NewLabels(labels *[]entities.Label) *[]LabelModel {
	if labels == nil {
		return nil
	}

	outputs := make([]LabelModel, 0, len(*labels))
	for _, input := range *labels {
		label := LabelModel{
			ID:   types.Int64Value(input.ID),
			Name: types.StringValue(input.Name),
		}
		outputs = append(outputs, label)
	}
	return &outputs
}
