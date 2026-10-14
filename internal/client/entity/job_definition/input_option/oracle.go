package input_option

import (
	"encoding/json"
	"strings"

	"terraform-provider-trocco/internal/client/entity"
)

type OracleInputOption struct {
	OracleConnectionID        int64                            `json:"oracle_connection_id"`
	Database                  *string                          `json:"database"`
	ConnectionType            string                           `json:"connection_type"`
	NetServiceName            *string                          `json:"net_service_name"`
	Schema                    *string                          `json:"schema"`
	Query                     *string                          `json:"query"`
	IncrementalLoadingEnabled bool                             `json:"incremental_loading_enabled"`
	Table                     *string                          `json:"table"`
	IncrementalColumns        *string                          `json:"incremental_columns"`
	LastRecord                *string                          `json:"last_record"`
	SourceTimeZone            *string                          `json:"source_time_zone"`
	InputOptionColumns        *[]OracleInputOptionColumn       `json:"input_option_columns"`
	InputOptionColumnOptions  *[]OracleInputOptionColumnOption `json:"input_option_column_options"`
	CustomVariableSettings    *[]entity.CustomVariableSetting  `json:"custom_variable_settings"`
}

// UnmarshalJSON normalizes last_record, which the Public API returns as
// either a JSON string or a JSON array of strings (one value per
// incremental column) depending on how the record was last saved. Both
// forms are joined into a single comma-separated string, matching the
// format the API also accepts on write.
func (o *OracleInputOption) UnmarshalJSON(data []byte) error {
	type alias OracleInputOption
	aux := &struct {
		LastRecord json.RawMessage `json:"last_record"`
		*alias
	}{alias: (*alias)(o)}

	if err := json.Unmarshal(data, aux); err != nil {
		return err
	}

	if len(aux.LastRecord) == 0 || string(aux.LastRecord) == "null" {
		return nil
	}

	var s string
	if err := json.Unmarshal(aux.LastRecord, &s); err == nil {
		o.LastRecord = &s
		return nil
	}

	var arr []string
	if err := json.Unmarshal(aux.LastRecord, &arr); err != nil {
		return err
	}
	joined := strings.Join(arr, ",")
	o.LastRecord = &joined
	return nil
}

type OracleInputOptionColumn struct {
	Name     string  `json:"name"`
	Type     string  `json:"type"`
	Format   *string `json:"format"`
	Timezone *string `json:"timezone"`
}

type OracleInputOptionColumnOption struct {
	ColumnName      string `json:"column_name"`
	ColumnValueType string `json:"column_value_type"`
}
