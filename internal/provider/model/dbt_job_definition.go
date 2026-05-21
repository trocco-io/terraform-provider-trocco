package model

import (
	"terraform-provider-trocco/internal/client/entity"
	"terraform-provider-trocco/internal/client/parameter"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

type DbtJobDefinitionModel struct {
	ID                     types.Int64                  `tfsdk:"id"`
	Name                   types.String                 `tfsdk:"name"`
	Description            types.String                 `tfsdk:"description"`
	ResourceGroupID        types.Int64                  `tfsdk:"resource_group_id"`
	AdapterType            types.String                 `tfsdk:"adapter_type"`
	DbtGitRepositoryID     types.Int64                  `tfsdk:"dbt_git_repository_id"`
	Threads                types.Int64                  `tfsdk:"threads"`
	Target                 types.String                 `tfsdk:"target"`
	BigquerySetting        *DbtBigquerySettingModel     `tfsdk:"bigquery_setting"`
	SnowflakeSetting       *DbtSnowflakeSettingModel    `tfsdk:"snowflake_setting"`
	RedshiftSetting        *DbtRedshiftSettingModel     `tfsdk:"redshift_setting"`
	Commands               []DbtCommandModel            `tfsdk:"commands"`
	CustomVariableSettings *[]CustomVariableSetting     `tfsdk:"custom_variable_settings"`
}

type DbtBigquerySettingModel struct {
	ConnectionID types.Int64  `tfsdk:"connection_id"`
	Dataset      types.String `tfsdk:"dataset"`
	Location     types.String `tfsdk:"location"`
}

type DbtSnowflakeSettingModel struct {
	ConnectionID types.Int64  `tfsdk:"connection_id"`
	Warehouse    types.String `tfsdk:"warehouse"`
	Database     types.String `tfsdk:"database"`
	Schema       types.String `tfsdk:"schema"`
	Role         types.String `tfsdk:"role"`
}

type DbtRedshiftSettingModel struct {
	ConnectionID types.Int64  `tfsdk:"connection_id"`
	Database     types.String `tfsdk:"database"`
	Schema       types.String `tfsdk:"schema"`
}

type DbtCommandModel struct {
	Command types.String            `tfsdk:"command"`
	Value   types.String            `tfsdk:"value"`
	Options []DbtCommandOptionModel `tfsdk:"options"`
}

type DbtCommandOptionModel struct {
	Key   types.String `tfsdk:"key"`
	Value types.String `tfsdk:"value"`
}

// NewDbtJobDefinitionModel hydrates the TF model from the API entity.
func NewDbtJobDefinitionModel(def *entity.DbtJobDefinition) DbtJobDefinitionModel {
	m := DbtJobDefinitionModel{
		ID:                 types.Int64Value(def.ID),
		Name:               types.StringValue(def.Name),
		Description:        types.StringPointerValue(def.Description),
		ResourceGroupID:    types.Int64PointerValue(def.ResourceGroupID),
		AdapterType:        types.StringValue(def.AdapterType),
		DbtGitRepositoryID: types.Int64Value(def.DbtGitRepositoryID),
		Threads:            types.Int64Value(def.Threads),
		Target:             types.StringValue(def.Target),
	}

	if def.BigquerySetting != nil {
		m.BigquerySetting = &DbtBigquerySettingModel{
			ConnectionID: types.Int64Value(def.BigquerySetting.ConnectionID),
			Dataset:      types.StringValue(def.BigquerySetting.Dataset),
			Location:     types.StringPointerValue(def.BigquerySetting.Location),
		}
	}
	if def.SnowflakeSetting != nil {
		m.SnowflakeSetting = &DbtSnowflakeSettingModel{
			ConnectionID: types.Int64Value(def.SnowflakeSetting.ConnectionID),
			Warehouse:    types.StringValue(def.SnowflakeSetting.Warehouse),
			Database:     types.StringValue(def.SnowflakeSetting.Database),
			Schema:       types.StringValue(def.SnowflakeSetting.Schema),
			Role:         types.StringPointerValue(def.SnowflakeSetting.Role),
		}
	}
	if def.RedshiftSetting != nil {
		m.RedshiftSetting = &DbtRedshiftSettingModel{
			ConnectionID: types.Int64Value(def.RedshiftSetting.ConnectionID),
			Database:     types.StringValue(def.RedshiftSetting.Database),
			Schema:       types.StringValue(def.RedshiftSetting.Schema),
		}
	}

	m.Commands = make([]DbtCommandModel, 0, len(def.Commands))
	for _, c := range def.Commands {
		cm := DbtCommandModel{
			Command: types.StringValue(c.Command),
			Value:   types.StringPointerValue(c.Value),
		}
		if len(c.Options) > 0 {
			cm.Options = make([]DbtCommandOptionModel, 0, len(c.Options))
			for _, opt := range c.Options {
				cm.Options = append(cm.Options, DbtCommandOptionModel{
					Key:   types.StringValue(opt.Key),
					Value: types.StringPointerValue(opt.Value),
				})
			}
		}
		m.Commands = append(m.Commands, cm)
	}

	m.CustomVariableSettings = NewCustomVariableSettings(&def.CustomVariableSettings)

	return m
}

// ToCreateInput converts a plan into a Create request payload.
func (m *DbtJobDefinitionModel) ToCreateInput() parameter.CreateDbtJobDefinitionInput {
	input := parameter.CreateDbtJobDefinitionInput{
		Name:                   m.Name.ValueString(),
		Description:            m.Description.ValueStringPointer(),
		ResourceGroupID:        m.ResourceGroupID.ValueInt64Pointer(),
		DbtGitRepositoryID:     m.DbtGitRepositoryID.ValueInt64(),
		Threads:                m.Threads.ValueInt64Pointer(),
		Target:                 m.Target.ValueStringPointer(),
		BigquerySetting:        bigquerySettingToInput(m.BigquerySetting),
		SnowflakeSetting:       snowflakeSettingToInput(m.SnowflakeSetting),
		RedshiftSetting:        redshiftSettingToInput(m.RedshiftSetting),
		Commands:               commandsToInput(m.Commands),
		CustomVariableSettings: customVariableSettingsToInput(m.CustomVariableSettings),
	}
	return input
}

// ToUpdateInput converts a plan into a Update request payload.
// Always sends every field so that Terraform state is the single source of truth
// (the API otherwise falls back to the previous revision for missing keys).
func (m *DbtJobDefinitionModel) ToUpdateInput() parameter.UpdateDbtJobDefinitionInput {
	input := parameter.UpdateDbtJobDefinitionInput{
		Name:                   m.Name.ValueStringPointer(),
		Description:            m.Description.ValueStringPointer(),
		ResourceGroupID:        m.ResourceGroupID.ValueInt64Pointer(),
		DbtGitRepositoryID:     m.DbtGitRepositoryID.ValueInt64Pointer(),
		Threads:                m.Threads.ValueInt64Pointer(),
		Target:                 m.Target.ValueStringPointer(),
		BigquerySetting:        bigquerySettingToInput(m.BigquerySetting),
		SnowflakeSetting:       snowflakeSettingToInput(m.SnowflakeSetting),
		RedshiftSetting:        redshiftSettingToInput(m.RedshiftSetting),
		Commands:               commandsToInput(m.Commands),
		CustomVariableSettings: customVariableSettingsToInput(m.CustomVariableSettings),
	}
	return input
}

func bigquerySettingToInput(s *DbtBigquerySettingModel) *parameter.DbtBigquerySettingInput {
	if s == nil {
		return nil
	}
	return &parameter.DbtBigquerySettingInput{
		ConnectionID: s.ConnectionID.ValueInt64(),
		Dataset:      s.Dataset.ValueString(),
		Location:     s.Location.ValueStringPointer(),
	}
}

func snowflakeSettingToInput(s *DbtSnowflakeSettingModel) *parameter.DbtSnowflakeSettingInput {
	if s == nil {
		return nil
	}
	return &parameter.DbtSnowflakeSettingInput{
		ConnectionID: s.ConnectionID.ValueInt64(),
		Warehouse:    s.Warehouse.ValueString(),
		Database:     s.Database.ValueString(),
		Schema:       s.Schema.ValueString(),
		Role:         s.Role.ValueStringPointer(),
	}
}

func redshiftSettingToInput(s *DbtRedshiftSettingModel) *parameter.DbtRedshiftSettingInput {
	if s == nil {
		return nil
	}
	return &parameter.DbtRedshiftSettingInput{
		ConnectionID: s.ConnectionID.ValueInt64(),
		Database:     s.Database.ValueString(),
		Schema:       s.Schema.ValueString(),
	}
}

func commandsToInput(commands []DbtCommandModel) []parameter.DbtCommandInput {
	out := make([]parameter.DbtCommandInput, 0, len(commands))
	for _, c := range commands {
		cmd := parameter.DbtCommandInput{
			Command: c.Command.ValueString(),
			Value:   c.Value.ValueStringPointer(),
		}
		if len(c.Options) > 0 {
			cmd.Options = make([]parameter.DbtCommandOptionInput, 0, len(c.Options))
			for _, opt := range c.Options {
				cmd.Options = append(cmd.Options, parameter.DbtCommandOptionInput{
					Key:   opt.Key.ValueString(),
					Value: opt.Value.ValueStringPointer(),
				})
			}
		}
		out = append(out, cmd)
	}
	return out
}

func customVariableSettingsToInput(settings *[]CustomVariableSetting) []parameter.CustomVariableSettingInput {
	if settings == nil {
		return []parameter.CustomVariableSettingInput{}
	}
	inputs := ToCustomVariableSettingInputs(settings)
	if inputs == nil {
		return []parameter.CustomVariableSettingInput{}
	}
	return *inputs
}

