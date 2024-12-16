package workflow

import (
	"github.com/hashicorp/terraform-plugin-framework/types"

	we "terraform-provider-trocco/internal/client/entities/pipeline_definition"
	wp "terraform-provider-trocco/internal/client/parameters/pipeline_definition"
)

type DBTTaskConfig struct {
	DefinitionID types.Int64 `tfsdk:"definition_id"`
}

func NewDBTTaskConfig(c *we.DBTTaskConfig) *DBTTaskConfig {
	if c == nil {
		return nil
	}

	return &DBTTaskConfig{
		DefinitionID: types.Int64Value(c.DefinitionID),
	}
}

func (c *DBTTaskConfig) ToInput() *wp.DBTTaskConfig {
	return &wp.DBTTaskConfig{
		DefinitionID: c.DefinitionID.ValueInt64(),
	}
}
