package workflow

type TableauDataExtractionTaskConfig struct {
	Name         string `json:"name,omitempty"`
	ConnectionID int64  `json:"connection_id,omitempty"`
	TaskID       string `json:"task_id,omitempty"`
}
