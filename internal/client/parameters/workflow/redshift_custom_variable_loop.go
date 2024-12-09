package workflow

type RedshiftCustomVariableLoopConfig struct {
	ConnectionID int64    `json:"connection_id,omitempty"`
	Query        string   `json:"query,omitempty"`
	Database     string   `json:"database,omitempty"`
	Variables    []string `json:"variables,omitempty"`
}
