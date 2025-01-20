package pipeline_definition

type TableauDataExtractionTaskConfig struct {
	Name         string `json:"name"`
	ConnectionID int64  `json:"connection_id"`
	TaskID       string `json:"task_id"`
}
