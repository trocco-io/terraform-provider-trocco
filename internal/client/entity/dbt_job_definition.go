package entity

type DbtJobDefinition struct {
	ID                     int64                       `json:"id"`
	Name                   string                      `json:"name"`
	Description            *string                     `json:"description"`
	ResourceGroupID        *int64                      `json:"resource_group_id"`
	DbtGitRepositoryID     int64                       `json:"dbt_git_repository_id"`
	Threads                int64                       `json:"threads"`
	Target                 string                      `json:"target"`
	BigquerySetting        *DbtBigquerySetting         `json:"bigquery_setting,omitempty"`
	SnowflakeSetting       *DbtSnowflakeSetting        `json:"snowflake_setting,omitempty"`
	RedshiftSetting        *DbtRedshiftSetting         `json:"redshift_setting,omitempty"`
	Commands               []DbtCommand                `json:"commands"`
	CustomVariableSettings []CustomVariableSetting     `json:"custom_variable_settings"`
	CreatedAt              string                      `json:"created_at"`
	UpdatedAt              string                      `json:"updated_at"`
	CreatedBy              *DbtJobDefinitionCreatedBy  `json:"created_by"`
}

type DbtBigquerySetting struct {
	ConnectionID int64   `json:"connection_id"`
	Dataset      string  `json:"dataset"`
	Location     *string `json:"location"`
}

type DbtSnowflakeSetting struct {
	ConnectionID int64   `json:"connection_id"`
	Warehouse    string  `json:"warehouse"`
	Database     string  `json:"database"`
	Schema       string  `json:"schema"`
	Role         *string `json:"role"`
}

type DbtRedshiftSetting struct {
	ConnectionID int64  `json:"connection_id"`
	Database     string `json:"database"`
	Schema       string `json:"schema"`
}

type DbtCommand struct {
	Command string             `json:"command"`
	Value   *string            `json:"value"`
	Options []DbtCommandOption `json:"options,omitempty"`
}

type DbtCommandOption struct {
	Key   string  `json:"key"`
	Value *string `json:"value"`
}

type DbtJobDefinitionCreatedBy struct {
	ID    int64  `json:"id"`
	Email string `json:"email"`
}
