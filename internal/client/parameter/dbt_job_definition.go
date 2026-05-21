package parameter

type CreateDbtJobDefinitionInput struct {
	Name                   string                       `json:"name"`
	Description            *string                      `json:"description,omitempty"`
	ResourceGroupID        *int64                       `json:"resource_group_id,omitempty"`
	DbtGitRepositoryID     int64                        `json:"dbt_git_repository_id"`
	Threads                *int64                       `json:"threads,omitempty"`
	Target                 *string                      `json:"target,omitempty"`
	BigquerySetting        *DbtBigquerySettingInput     `json:"bigquery_setting,omitempty"`
	SnowflakeSetting       *DbtSnowflakeSettingInput    `json:"snowflake_setting,omitempty"`
	RedshiftSetting        *DbtRedshiftSettingInput     `json:"redshift_setting,omitempty"`
	Commands               []DbtCommandInput            `json:"commands"`
	CustomVariableSettings []CustomVariableSettingInput `json:"custom_variable_settings"`
}

func (input *CreateDbtJobDefinitionInput) SetDescription(v string)        { input.Description = &v }
func (input *CreateDbtJobDefinitionInput) SetResourceGroupID(v int64)     { input.ResourceGroupID = &v }
func (input *CreateDbtJobDefinitionInput) SetThreads(v int64)             { input.Threads = &v }
func (input *CreateDbtJobDefinitionInput) SetTarget(v string)             { input.Target = &v }
func (input *CreateDbtJobDefinitionInput) SetBigquerySetting(v DbtBigquerySettingInput) {
	input.BigquerySetting = &v
}
func (input *CreateDbtJobDefinitionInput) SetSnowflakeSetting(v DbtSnowflakeSettingInput) {
	input.SnowflakeSetting = &v
}
func (input *CreateDbtJobDefinitionInput) SetRedshiftSetting(v DbtRedshiftSettingInput) {
	input.RedshiftSetting = &v
}
func (input *CreateDbtJobDefinitionInput) SetCommands(v []DbtCommandInput) { input.Commands = v }
func (input *CreateDbtJobDefinitionInput) SetCustomVariableSettings(v []CustomVariableSettingInput) {
	input.CustomVariableSettings = v
}

type UpdateDbtJobDefinitionInput struct {
	Name                   *string                      `json:"name,omitempty"`
	Description            *string                      `json:"description,omitempty"`
	ResourceGroupID        *int64                       `json:"resource_group_id,omitempty"`
	DbtGitRepositoryID     *int64                       `json:"dbt_git_repository_id,omitempty"`
	Threads                *int64                       `json:"threads,omitempty"`
	Target                 *string                      `json:"target,omitempty"`
	BigquerySetting        *DbtBigquerySettingInput     `json:"bigquery_setting,omitempty"`
	SnowflakeSetting       *DbtSnowflakeSettingInput    `json:"snowflake_setting,omitempty"`
	RedshiftSetting        *DbtRedshiftSettingInput     `json:"redshift_setting,omitempty"`
	Commands               []DbtCommandInput            `json:"commands"`
	CustomVariableSettings []CustomVariableSettingInput `json:"custom_variable_settings"`
}

func (input *UpdateDbtJobDefinitionInput) SetName(v string)               { input.Name = &v }
func (input *UpdateDbtJobDefinitionInput) SetDescription(v string)        { input.Description = &v }
func (input *UpdateDbtJobDefinitionInput) SetResourceGroupID(v int64)     { input.ResourceGroupID = &v }
func (input *UpdateDbtJobDefinitionInput) SetDbtGitRepositoryID(v int64)  { input.DbtGitRepositoryID = &v }
func (input *UpdateDbtJobDefinitionInput) SetThreads(v int64)             { input.Threads = &v }
func (input *UpdateDbtJobDefinitionInput) SetTarget(v string)             { input.Target = &v }
func (input *UpdateDbtJobDefinitionInput) SetBigquerySetting(v DbtBigquerySettingInput) {
	input.BigquerySetting = &v
}
func (input *UpdateDbtJobDefinitionInput) SetSnowflakeSetting(v DbtSnowflakeSettingInput) {
	input.SnowflakeSetting = &v
}
func (input *UpdateDbtJobDefinitionInput) SetRedshiftSetting(v DbtRedshiftSettingInput) {
	input.RedshiftSetting = &v
}
func (input *UpdateDbtJobDefinitionInput) SetCommands(v []DbtCommandInput) { input.Commands = v }
func (input *UpdateDbtJobDefinitionInput) SetCustomVariableSettings(v []CustomVariableSettingInput) {
	input.CustomVariableSettings = v
}

type DbtBigquerySettingInput struct {
	ConnectionID int64   `json:"connection_id"`
	Dataset      string  `json:"dataset"`
	Location     *string `json:"location,omitempty"`
}

func (s *DbtBigquerySettingInput) SetLocation(v string) { s.Location = &v }

type DbtSnowflakeSettingInput struct {
	ConnectionID int64   `json:"connection_id"`
	Warehouse    string  `json:"warehouse"`
	Database     string  `json:"database"`
	Schema       string  `json:"schema"`
	Role         *string `json:"role,omitempty"`
}

func (s *DbtSnowflakeSettingInput) SetRole(v string) { s.Role = &v }

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

func (c *DbtCommandInput) SetValue(v string)                     { c.Value = &v }
func (c *DbtCommandInput) SetOptions(v []DbtCommandOptionInput)  { c.Options = v }

type DbtCommandOptionInput struct {
	Key   string  `json:"key"`
	Value *string `json:"value,omitempty"`
}

func (o *DbtCommandOptionInput) SetValue(v string) { o.Value = &v }
