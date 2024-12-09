package workflow

type SnowflakeCustomVariableLoopConfig struct {
	ConnectionID int64    `json:"connection_id,omitempty"`
	Query        string   `json:"query,omitempty"`
	Warehouse    string   `json:"warehouse,omitempty"`
	Variables    []string `json:"variables,omitempty"`
}
