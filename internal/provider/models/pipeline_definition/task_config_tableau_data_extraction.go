package pipeline_definition

import (
	"github.com/hashicorp/terraform-plugin-framework/types"

	we "terraform-provider-trocco/internal/client/entities/pipeline_definition"
	wp "terraform-provider-trocco/internal/client/parameters/pipeline_definition"
)

type TableauDataExtractionTaskConfig struct {
	Name         types.String `tfsdk:"name"`
	ConnectionID types.Int64  `tfsdk:"connection_id"`
	TaskID       types.String `tfsdk:"task_id"`
}

func NewTableauDataExtractionTaskConfig(c *we.TableauDataExtractionTaskConfig) *TableauDataExtractionTaskConfig {
	if c == nil {
		return nil
	}

	return &TableauDataExtractionTaskConfig{
		Name:         types.StringValue(c.Name),
		ConnectionID: types.Int64Value(c.ConnectionID),
		TaskID:       types.StringValue(c.TaskID),
	}
}

func (c *TableauDataExtractionTaskConfig) ToInput() *wp.TableauDataExtractionTaskConfig {
	return &wp.TableauDataExtractionTaskConfig{
		Name:         c.Name.ValueString(),
		ConnectionID: c.ConnectionID.ValueInt64(),
		TaskID:       c.TaskID.ValueString(),
	}
}