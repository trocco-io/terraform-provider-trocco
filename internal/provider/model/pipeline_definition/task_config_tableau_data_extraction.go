package pipeline_definition

import (
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"

	pipelineDefinitionEntities "terraform-provider-trocco/internal/client/entity/pipeline_definition"
	pipelineDefinitionParameters "terraform-provider-trocco/internal/client/parameter/pipeline_definition"
)

type TableauDataExtractionTaskConfig struct {
	Name         types.String `tfsdk:"name"`
	ConnectionID types.Int64  `tfsdk:"connection_id"`
	TaskID       types.String `tfsdk:"task_id"`
}

func NewTableauDataExtractionTaskConfig(c *pipelineDefinitionEntities.TableauDataExtractionTaskConfig) *TableauDataExtractionTaskConfig {
	if c == nil {
		return nil
	}

	return &TableauDataExtractionTaskConfig{
		Name:         types.StringValue(c.Name),
		ConnectionID: types.Int64Value(c.ConnectionID),
		TaskID:       types.StringValue(c.TaskID),
	}
}

func (c *TableauDataExtractionTaskConfig) ToInput() *pipelineDefinitionParameters.TableauDataExtractionTaskConfig {
	return &pipelineDefinitionParameters.TableauDataExtractionTaskConfig{
		Name:         c.Name.ValueString(),
		ConnectionID: c.ConnectionID.ValueInt64(),
		TaskID:       c.TaskID.ValueString(),
	}
}

func TableauDataExtractionTaskConfigAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"name":          types.StringType,
		"connection_id": types.Int64Type,
		"task_id":       types.StringType,
	}
}
