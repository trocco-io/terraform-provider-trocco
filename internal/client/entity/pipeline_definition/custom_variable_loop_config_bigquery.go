package pipeline_definition

type BigqueryCustomVariableLoopConfig struct {
	ConnectionID int64    `json:"connection_id"`
	Query        string   `json:"query"`
	Variables    []string `json:"variables"`
}
