package input_options

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"terraform-provider-trocco/internal/client/entities/job_definitions/input_options"
	"terraform-provider-trocco/internal/provider/models"
)

type MySQLInputOption struct {
	Database                  types.String                    `tfsdk:"database"`
	Table                     types.String                    `tfsdk:"table"`
	Query                     types.String                    `tfsdk:"query"`
	IncrementalColumns        types.String                    `tfsdk:"incremental_columns"`
	LastRecord                types.String                    `tfsdk:"last_record"`
	IncrementalLoadingEnabled types.Bool                      `tfsdk:"incremental_loading_enabled"`
	FetchRows                 types.Int64                     `tfsdk:"fetch_rows"`
	ConnectTimeout            types.Int64                     `tfsdk:"connect_timeout"`
	SocketTimeout             types.Int64                     `tfsdk:"socket_timeout"`
	DefaultTimeZone           types.String                    `tfsdk:"default_time_zone"`
	UseLegacyDatetimeCode     types.Bool                      `tfsdk:"use_legacy_datetime_code"`
	MySQLConnectionID         types.Int64                     `tfsdk:"mysql_connection_id"`
	InputOptionColumns        []inputOptionColumn             `tfsdk:"input_option_columns"`
	CustomVariableSettings    *[]models.CustomVariableSetting `tfsdk:"custom_variable_settings"`
}

type inputOptionColumn struct {
	Name types.String `tfsdk:"name"`
	Type types.String `tfsdk:"type"`
}

func NewMysqlInputOption(mysqlInputOption *input_options.MySQLInputOption) *MySQLInputOption {
	if mysqlInputOption == nil {
		return nil
	}

	return &MySQLInputOption{
		Database:                  types.StringValue(mysqlInputOption.Database),
		Table:                     types.StringPointerValue(mysqlInputOption.Table),
		Query:                     types.StringValue(mysqlInputOption.Query),
		IncrementalColumns:        types.StringPointerValue(mysqlInputOption.IncrementalColumns),
		LastRecord:                types.StringPointerValue(mysqlInputOption.LastRecord),
		IncrementalLoadingEnabled: types.BoolValue(mysqlInputOption.IncrementalLoadingEnabled),
		FetchRows:                 types.Int64Value(mysqlInputOption.FetchRows),
		ConnectTimeout:            types.Int64Value(mysqlInputOption.ConnectTimeout),
		SocketTimeout:             types.Int64Value(mysqlInputOption.SocketTimeout),
		DefaultTimeZone:           types.StringPointerValue(mysqlInputOption.DefaultTimeZone),
		UseLegacyDatetimeCode:     types.BoolPointerValue(mysqlInputOption.UseLegacyDatetimeCode),
		MySQLConnectionID:         types.Int64Value(mysqlInputOption.MySQLConnectionID),
		InputOptionColumns:        newInputOptionColumns(mysqlInputOption.InputOptionColumns),
		CustomVariableSettings:    models.NewCustomVariableSettings(mysqlInputOption.CustomVariableSettings),
	}
}

func newInputOptionColumns(inputOptionColumns []input_options.InputOptionColumn) []inputOptionColumn {
	if inputOptionColumns == nil {
		return nil
	}
	columns := make([]inputOptionColumn, 0, len(inputOptionColumns))
	for _, input := range inputOptionColumns {
		column := inputOptionColumn{
			Name: types.StringValue(input.Name),
			Type: types.StringValue(input.Type),
		}
		columns = append(columns, column)
	}
	return columns
}
