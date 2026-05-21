package parameter

type CreateDbtJobDefinitionInput struct {
	Name                   string                            `json:"name"`
	Description            *string                           `json:"description,omitempty"`
	ResourceGroupID        *int64                            `json:"resource_group_id,omitempty"`
	DbtGitRepositoryID     int64                             `json:"dbt_git_repository_id"`
	Threads                *int64                            `json:"threads,omitempty"`
	Target                 *string                           `json:"target,omitempty"`
	BigquerySetting        *DbtBigquerySettingInput          `json:"bigquery_setting,omitempty"`
	SnowflakeSetting       *DbtSnowflakeSettingInput         `json:"snowflake_setting,omitempty"`
	RedshiftSetting        *DbtRedshiftSettingInput          `json:"redshift_setting,omitempty"`
	Commands               []DbtCommandInput                 `json:"commands"`
	CustomVariableSettings []CustomVariableSettingInput      `json:"custom_variable_settings"`
}

type UpdateDbtJobDefinitionInput struct {
	Name                   *string                       `json:"name,omitempty"`
	Description            *string                       `json:"description,omitempty"`
	ResourceGroupID        *int64                        `json:"resource_group_id,omitempty"`
	DbtGitRepositoryID     *int64                        `json:"dbt_git_repository_id,omitempty"`
	Threads                *int64                        `json:"threads,omitempty"`
	Target                 *string                       `json:"target,omitempty"`
	BigquerySetting        *DbtBigquerySettingInput      `json:"bigquery_setting,omitempty"`
	SnowflakeSetting       *DbtSnowflakeSettingInput     `json:"snowflake_setting,omitempty"`
	RedshiftSetting        *DbtRedshiftSettingInput      `json:"redshift_setting,omitempty"`
	Commands               []DbtCommandInput             `json:"commands"`
	CustomVariableSettings []CustomVariableSettingInput  `json:"custom_variable_settings"`
}

type DbtBigquerySettingInput struct {
	ConnectionID int64   `json:"connection_id"`
	Dataset      string  `json:"dataset"`
	Location     *string `json:"location,omitempty"`
}

type DbtSnowflakeSettingInput struct {
	ConnectionID int64   `json:"connection_id"`
	Warehouse    string  `json:"warehouse"`
	Database     string  `json:"database"`
	Schema       string  `json:"schema"`
	Role         *string `json:"role,omitempty"`
}

type DbtRedshiftSettingInput struct {
	ConnectionID int64  `json:"connection_id"`
	Database     string `json:"database"`
	Schema       string `json:"schema"`
}

type DbtCommandInput struct {
	Command string                  `json:"command"`
	Value   *string                 `json:"value,omitempty"`
	Options []DbtCommandOptionInput `json:"options,omitempty"`
}

type DbtCommandOptionInput struct {
	Key   string  `json:"key"`
	Value *string `json:"value,omitempty"`
}
