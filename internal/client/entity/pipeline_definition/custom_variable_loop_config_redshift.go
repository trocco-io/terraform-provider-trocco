package pipeline_definition

type RedshiftCustomVariableLoopConfig struct {
	ConnectionID int64    `json:"connection_id"`
	Query        string   `json:"query"`
	Database     string   `json:"database"`
	Variables    []string `json:"variables"`
}
