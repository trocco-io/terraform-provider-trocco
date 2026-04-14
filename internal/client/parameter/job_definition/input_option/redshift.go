package input_options

import "terraform-provider-trocco/internal/client/parameter"

type RedshiftInputOptionInput struct {
	RedshiftConnectionID   int64                                   `json:"redshift_connection_id"`
	Schema                 *parameter.NullableString               `json:"schema,omitempty"`
	Database               string                                  `json:"database"`
	Query                  string                                  `json:"query"`
	FetchRows              int64                                   `json:"fetch_rows"`
	ConnectTimeout         int64                                   `json:"connect_timeout"`
	SocketTimeout          int64                                   `json:"socket_timeout"`
	CustomVariableSettings *[]parameter.CustomVariableSettingInput `json:"custom_variable_settings,omitempty"`
}

type UpdateRedshiftInputOptionInput struct {
	RedshiftConnectionID   *int64                                  `json:"redshift_connection_id,omitempty"`
	Schema                 *parameter.NullableString               `json:"schema,omitempty"`
	Database               *string                                 `json:"database,omitempty"`
	Query                  *string                                 `json:"query,omitempty"`
	FetchRows              *int64                                  `json:"fetch_rows,omitempty"`
	ConnectTimeout         *int64                                  `json:"connect_timeout,omitempty"`
	SocketTimeout          *int64                                  `json:"socket_timeout,omitempty"`
	CustomVariableSettings *[]parameter.CustomVariableSettingInput `json:"custom_variable_settings,omitempty"`
}
