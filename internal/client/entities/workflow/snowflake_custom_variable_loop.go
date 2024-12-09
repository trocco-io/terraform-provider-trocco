package workflow

type SnowflakeCustomVariableLoopConfig struct {
	ConnectionID int64    `json:"connection_id"`
	Query        string   `json:"query"`
	Warehouse    string   `json:"warehouse"`
	Variables    []string `json:"variables"`
}
