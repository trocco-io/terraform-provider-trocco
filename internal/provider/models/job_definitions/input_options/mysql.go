package input_options

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"terraform-provider-trocco/internal/provider/models"
)

type MySQLInputOption struct {
	Database                  types.String                   `tfsdk:"database"`
	Table                     types.String                   `tfsdk:"table"`
	Query                     types.String                   `tfsdk:"query"`
	IncrementalColumns        types.String                   `tfsdk:"incremental_columns"`
	LastRecord                types.String                   `tfsdk:"last_record"`
	IncrementalLoadingEnabled types.Bool                     `tfsdk:"incremental_loading_enabled"`
	FetchRows                 types.Int64                    `tfsdk:"fetch_rows"`
	ConnectTimeout            types.Int64                    `tfsdk:"connect_timeout"`
	SocketTimeout             types.Int64                    `tfsdk:"socket_timeout"`
	DefaultTimeZone           types.String                   `tfsdk:"default_time_zone"`
	UseLegacyDatetimeCode     types.Bool                     `tfsdk:"use_legacy_datetime_code"`
	MySQLConnectionID         types.Int64                    `tfsdk:"mysql_connection_id"`
	InputOptionColumns        []inputOptionColumn            `tfsdk:"input_option_columns"`
	CustomVariableSettings    []models.CustomVariableSetting `tfsdk:"custom_variable_settings"`
}

type inputOptionColumn struct {
	Name types.String `tfsdk:"name"`
	Type types.String `tfsdk:"type"`
}
