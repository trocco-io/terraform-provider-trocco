package input_option

import "terraform-provider-trocco/internal/client/entity"

type RedshiftInputOption struct {
	RedshiftConnectionID   int64                           `json:"redshift_connection_id"`
	Schema                 string                          `json:"schema"`
	Database               string                          `json:"database"`
	Query                  string                          `json:"query"`
	FetchRows              int64                           `json:"fetch_rows"`
	ConnectTimeout         int64                           `json:"connect_timeout"`
	SocketTimeout          int64                           `json:"socket_timeout"`
	CustomVariableSettings *[]entity.CustomVariableSetting `json:"custom_variable_settings"`
}
