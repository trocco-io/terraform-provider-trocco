package workflow

type BigqueryCustomVariableLoopConfig struct {
	ConnectionID int64    `json:"connection_id,omitempty"`
	Query        string   `json:"query,omitempty"`
	Variables    []string `json:"variables,omitempty"`
}
