package input_option

import (
	"terraform-provider-trocco/internal/client/entity"
	"testing"
)

func TestGoogleSpreadsheetsInputOption_SpreadsheetsID(t *testing.T) {
	type fields struct {
		SpreadsheetsURL                string
		WorksheetTitle                 string
		StartRow                       int64
		StartColumn                    string
		DefaultTimeZone                string
		NullString                     string
		GoogleSpreadsheetsConnectionID int64
		InputOptionColumns             []GoogleSpreadsheetsInputOptionColumn
		CustomVariableSettings         *[]entity.CustomVariableSetting
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "SpreadsheetsURL is empty",
			fields: fields{
				SpreadsheetsURL: "",
			},
			want: "",
		},
		{
			name: "SpreadsheetsURL is not empty",
			fields: fields{
				SpreadsheetsURL: "https://docs.google.com/spreadsheets/d/MY_SHEETS_ID/edit#gid=0",
			},
			want: "MY_SHEETS_ID",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inputOption := &GoogleSpreadsheetsInputOption{
				SpreadsheetsURL:                tt.fields.SpreadsheetsURL,
				WorksheetTitle:                 tt.fields.WorksheetTitle,
				StartRow:                       tt.fields.StartRow,
				StartColumn:                    tt.fields.StartColumn,
				DefaultTimeZone:                tt.fields.DefaultTimeZone,
				NullString:                     tt.fields.NullString,
				GoogleSpreadsheetsConnectionID: tt.fields.GoogleSpreadsheetsConnectionID,
				InputOptionColumns:             tt.fields.InputOptionColumns,
				CustomVariableSettings:         tt.fields.CustomVariableSettings,
			}
			if got := inputOption.SpreadsheetsID(); got != tt.want {
				t.Errorf("GoogleSpreadsheetsInputOption.SpreadsheetsID() = %v, want %v", got, tt.want)
			}
		})
	}
}
