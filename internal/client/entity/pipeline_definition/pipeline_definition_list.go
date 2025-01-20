package pipeline_definition

type PipelineDefinitionList struct {
	Items      []*PipelineDefinition `json:"items"`
	NextCursor string                `json:"next_cursor"`
}
