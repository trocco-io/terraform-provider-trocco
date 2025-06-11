package pipeline_definition

import (
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"

	we "terraform-provider-trocco/internal/client/entity/pipeline_definition"
	wp "terraform-provider-trocco/internal/client/parameter/pipeline_definition"
)

type TroccoDBTTaskConfig struct {
	DefinitionID types.Int64 `tfsdk:"definition_id"`
}

func NewTroccoDBTTaskConfig(c *we.TroccoDBTTaskConfig) *TroccoDBTTaskConfig {
	if c == nil {
		return nil
	}

	return &TroccoDBTTaskConfig{
		DefinitionID: types.Int64Value(c.DefinitionID),
	}
}

func (c *TroccoDBTTaskConfig) ToInput() *wp.TroccoDBTTaskConfig {
	return &wp.TroccoDBTTaskConfig{
		DefinitionID: c.DefinitionID.ValueInt64(),
	}
}

func TroccoDBTTaskConfigAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"definition_id": types.Int64Type,
	}
}
