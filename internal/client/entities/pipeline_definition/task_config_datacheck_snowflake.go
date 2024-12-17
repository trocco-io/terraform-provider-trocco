package pipeline_definition

type SnowflakeDataCheckTaskConfig struct {
	Name            string           `json:"name"`
	ConnectionID    int64            `json:"connection_id"`
	Query           string           `json:"query"`
	Operator        string           `json:"operator"`
	QueryResult     int64            `json:"query_result"`
	AcceptsNull     bool             `json:"accepts_null"`
	Warehouse       string           `json:"warehouse"`
	CustomVariables []CustomVariable `json:"custom_variables"`
}
