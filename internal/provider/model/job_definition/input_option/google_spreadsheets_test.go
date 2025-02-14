package input_options

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestGoogleSpreadsheetsInputOption_SpreadsheetsURL(t *testing.T) {
	type fields struct {
		GoogleSpreadsheetsConnectionID *int64
		SpreadsheetsID                 *string
		WorksheetTitle                 *string
		StartRow                       *int64
		StartColumn                    *string
		DefaultTimeZone                *string
		NullString                     *string
	}
	tests := []struct {
		name   string
		fields fields
		want   *string
	}{
		{
			name: "SpreadsheetsID is nil",
			fields: fields{
				SpreadsheetsID: nil,
			},
			want: nil,
		},
		{
			name: "SpreadsheetsID is not null",
			fields: fields{
				SpreadsheetsID: types.StringValue("MY_SHEETS_ID").ValueStringPointer(),
			},
			want: types.StringValue("https://docs.google.com/spreadsheets/d/MY_SHEETS_ID/edit#gid=0").ValueStringPointer(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inputOption := &GoogleSpreadsheetsInputOption{
				GoogleSpreadsheetsConnectionID: types.Int64PointerValue(tt.fields.GoogleSpreadsheetsConnectionID),
				SpreadsheetsID:                 types.StringPointerValue(tt.fields.SpreadsheetsID),
				WorksheetTitle:                 types.StringPointerValue(tt.fields.WorksheetTitle),
				StartRow:                       types.Int64PointerValue(tt.fields.StartRow),
				StartColumn:                    types.StringPointerValue(tt.fields.StartColumn),
				DefaultTimeZone:                types.StringPointerValue(tt.fields.DefaultTimeZone),
				NullString:                     types.StringPointerValue(tt.fields.NullString),
				InputOptionColumns:             nil,
				CustomVariableSettings:         nil,
			}
			if got := inputOption.SpreadsheetsURL(); got != types.StringPointerValue(tt.want) {
				t.Errorf("GoogleSpreadsheetsInputOption.SpreadsheetsURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
