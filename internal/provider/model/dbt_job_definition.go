package model

import (
	"context"
	"fmt"
	"terraform-provider-trocco/internal/client/entity"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type DbtJobDefinitionModel struct {
	ID                     types.Int64               `tfsdk:"id"`
	Name                   types.String              `tfsdk:"name"`
	Description            types.String              `tfsdk:"description"`
	ResourceGroupID        types.Int64               `tfsdk:"resource_group_id"`
	DbtGitRepositoryID     types.Int64               `tfsdk:"dbt_git_repository_id"`
	Threads                types.Int64               `tfsdk:"threads"`
	Target                 types.String              `tfsdk:"target"`
	BigquerySetting        *DbtBigquerySettingModel  `tfsdk:"bigquery_setting"`
	SnowflakeSetting       *DbtSnowflakeSettingModel `tfsdk:"snowflake_setting"`
	RedshiftSetting        *DbtRedshiftSettingModel  `tfsdk:"redshift_setting"`
	Commands               types.List                `tfsdk:"commands"`
	CustomVariableSettings types.List                `tfsdk:"custom_variable_settings"`
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

func DbtCommandOptionAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"key":   types.StringType,
		"value": types.StringType,
	}
}

func DbtCommandAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"command": types.StringType,
		"value":   types.StringType,
		"options": types.ListType{ElemType: types.ObjectType{AttrTypes: DbtCommandOptionAttrTypes()}},
	}
}

// NewDbtJobDefinitionModel hydrates the TF model from the API entity.
// Pass-through of empty arrays as empty lists is intentional: the schema is
// Optional+Computed and UseStateForUnknown propagates the value, so distinguishing
// empty from null only matters when the user explicitly wrote `= []`.
func NewDbtJobDefinitionModel(ctx context.Context, def *entity.DbtJobDefinition) (DbtJobDefinitionModel, error) {
	m := DbtJobDefinitionModel{
		ID:                 types.Int64Value(def.ID),
		Name:               types.StringValue(def.Name),
		Description:        types.StringPointerValue(def.Description),
		ResourceGroupID:    types.Int64PointerValue(def.ResourceGroupID),
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

	commands := make([]DbtCommandModel, 0, len(def.Commands))
	for _, c := range def.Commands {
		cm := DbtCommandModel{
			Command: types.StringValue(c.Command),
			Value:   types.StringPointerValue(c.Value),
		}
		cm.Options = make([]DbtCommandOptionModel, 0, len(c.Options))
		for _, opt := range c.Options {
			cm.Options = append(cm.Options, DbtCommandOptionModel{
				Key:   types.StringValue(opt.Key),
				Value: types.StringPointerValue(opt.Value),
			})
		}
		commands = append(commands, cm)
	}
	commandsList, diags := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DbtCommandAttrTypes()}, commands)
	if diags.HasError() {
		return m, fmt.Errorf("failed to convert commands to ListValue: %v", diags)
	}
	m.Commands = commandsList

	cvSettings := make([]CustomVariableSetting, 0, len(def.CustomVariableSettings))
	for _, s := range def.CustomVariableSettings {
		setting := CustomVariableSetting{
			Name:      types.StringValue(s.Name),
			Type:      types.StringValue(s.Type),
			Value:     types.StringPointerValue(s.Value),
			Quantity:  types.Int64PointerValue(s.Quantity),
			Unit:      types.StringPointerValue(s.Unit),
			Direction: types.StringPointerValue(s.Direction),
			Format:    types.StringPointerValue(s.Format),
			TimeZone:  types.StringPointerValue(s.TimeZone),
		}
		cvSettings = append(cvSettings, setting)
	}
	cvList, diags := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: CustomVariableSettingAttrTypes()}, cvSettings)
	if diags.HasError() {
		return m, fmt.Errorf("failed to convert custom_variable_settings to ListValue: %v", diags)
	}
	m.CustomVariableSettings = cvList

	return m, nil
}

func CustomVariableSettingAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"name":      types.StringType,
		"type":      types.StringType,
		"value":     types.StringType,
		"quantity":  types.Int64Type,
		"unit":      types.StringType,
		"direction": types.StringType,
		"format":    types.StringType,
		"time_zone": types.StringType,
	}
}
