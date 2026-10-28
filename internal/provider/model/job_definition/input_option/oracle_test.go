package input_options

import (
	"context"
	"encoding/json"
	"strings"
	"testing"

	"terraform-provider-trocco/internal/provider/model/job_definition/common"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Regression test for the "removing an attribute from config doesn't clear it
// server-side" bug class: a field sent as a plain *string with `omitempty` is
// dropped from the request entirely when unset, so the API never learns the
// value should be cleared. Fields that can legitimately be unset after being
// set must use *parameter.NullableString (or the Nullable* equivalent) so the
// JSON request body carries an explicit `null`, not a missing key.
//
// Unlike acceptance tests, this requires no TROCCO account or network access,
// and runs as a normal `go test` so it is exercised on every pull request.
func TestOracleInputOptionToUpdateInput_ClearsNullableStringFieldsOnRemoval(t *testing.T) {
	ctx := context.Background()

	base := &OracleInputOption{
		OracleConnectionID:        types.Int64Value(1),
		Database:                  types.StringNull(),
		ConnectionType:            types.StringValue("sid"),
		NetServiceName:            types.StringNull(),
		Schema:                    types.StringNull(),
		Query:                     types.StringNull(),
		IncrementalLoadingEnabled: types.BoolValue(true),
		Table:                     types.StringNull(),
		IncrementalColumns:        types.StringNull(),
		LastRecord:                types.StringNull(),
		SourceTimeZone:            types.StringNull(),
		InputOptionColumns:        types.ListNull(types.ObjectType{AttrTypes: oracleInputOptionColumnAttrTypes}),
		InputOptionColumnOptions:  types.ListNull(types.ObjectType{AttrTypes: oracleInputOptionColumnOptionAttrTypes}),
		CustomVariableSettings:    types.ListNull(types.ObjectType{AttrTypes: common.GetCustomVariableSettingAttrTypes()}),
	}

	// Every field below is nullable in the API and is NOT backed by a
	// server-side default (unlike connection_type/driver), so clearing it in
	// config must always produce an explicit `null` key in the update request.
	nullableFields := []string{"source_time_zone", "query", "table"}

	input := base.ToUpdateInput(ctx)
	b, err := json.Marshal(input)
	if err != nil {
		t.Fatalf("failed to marshal UpdateOracleInputOptionInput: %s", err)
	}
	body := string(b)

	for _, field := range nullableFields {
		t.Run(field+" is sent as an explicit null, not omitted", func(t *testing.T) {
			key := `"` + field + `":null`
			if !strings.Contains(body, key) {
				t.Errorf("expected request body to contain %s (key present with explicit null) when the field is unset, got: %s", key, body)
			}
		})
	}
}

func TestOracleInputOptionToUpdateInput_SendsNullableStringFieldValues(t *testing.T) {
	ctx := context.Background()

	base := &OracleInputOption{
		OracleConnectionID:        types.Int64Value(1),
		Database:                  types.StringNull(),
		ConnectionType:            types.StringValue("sid"),
		NetServiceName:            types.StringNull(),
		Schema:                    types.StringNull(),
		Query:                     types.StringNull(),
		IncrementalLoadingEnabled: types.BoolValue(false),
		Table:                     types.StringValue("test_table"),
		IncrementalColumns:        types.StringNull(),
		LastRecord:                types.StringNull(),
		SourceTimeZone:            types.StringValue("Asia/Tokyo"),
		InputOptionColumns:        types.ListNull(types.ObjectType{AttrTypes: oracleInputOptionColumnAttrTypes}),
		InputOptionColumnOptions:  types.ListNull(types.ObjectType{AttrTypes: oracleInputOptionColumnOptionAttrTypes}),
		CustomVariableSettings:    types.ListNull(types.ObjectType{AttrTypes: common.GetCustomVariableSettingAttrTypes()}),
	}

	input := base.ToUpdateInput(ctx)
	b, err := json.Marshal(input)
	if err != nil {
		t.Fatalf("failed to marshal UpdateOracleInputOptionInput: %s", err)
	}
	body := string(b)

	for _, want := range []string{`"source_time_zone":"Asia/Tokyo"`, `"table":"test_table"`} {
		if !strings.Contains(body, want) {
			t.Errorf("expected request body to contain %s, got: %s", want, body)
		}
	}
}

// Regression test for the "removing input_option_column_options from config
// doesn't clear it server-side" bug: a field sent as *[]T with `omitempty` is
// dropped from the request entirely when the list is unset, so the backend's
// `unless ... .nil?` guard never fires and stale column options remain.
// InputOptionColumnOptions must marshal through NullableObjectList so the key
// is always present, with `[]` as an explicit "clear" signal when unset.
func TestOracleInputOptionToUpdateInput_ClearsInputOptionColumnOptionsOnRemoval(t *testing.T) {
	ctx := context.Background()

	base := &OracleInputOption{
		OracleConnectionID:        types.Int64Value(1),
		Database:                  types.StringNull(),
		ConnectionType:            types.StringValue("sid"),
		NetServiceName:            types.StringNull(),
		Schema:                    types.StringNull(),
		Query:                     types.StringNull(),
		IncrementalLoadingEnabled: types.BoolValue(true),
		Table:                     types.StringNull(),
		IncrementalColumns:        types.StringNull(),
		LastRecord:                types.StringNull(),
		SourceTimeZone:            types.StringNull(),
		InputOptionColumns:        types.ListNull(types.ObjectType{AttrTypes: oracleInputOptionColumnAttrTypes}),
		InputOptionColumnOptions:  types.ListNull(types.ObjectType{AttrTypes: oracleInputOptionColumnOptionAttrTypes}),
		CustomVariableSettings:    types.ListNull(types.ObjectType{AttrTypes: common.GetCustomVariableSettingAttrTypes()}),
	}

	input := base.ToUpdateInput(ctx)
	b, err := json.Marshal(input)
	if err != nil {
		t.Fatalf("failed to marshal UpdateOracleInputOptionInput: %s", err)
	}
	body := string(b)

	key := `"input_option_column_options":[]`
	if !strings.Contains(body, key) {
		t.Errorf("expected request body to contain %s (key present with an explicit empty array to clear existing options) when the field is unset, got: %s", key, body)
	}
}

func TestOracleInputOptionToUpdateInput_SendsInputOptionColumnOptionsValues(t *testing.T) {
	ctx := context.Background()

	columnOptions, diags := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: oracleInputOptionColumnOptionAttrTypes}, []OracleInputOptionColumnOption{
		{ColumnName: types.StringValue("amount"), ColumnValueType: types.StringValue("string")},
	})
	if diags.HasError() {
		t.Fatalf("failed to build input_option_column_options list: %v", diags)
	}

	base := &OracleInputOption{
		OracleConnectionID:        types.Int64Value(1),
		Database:                  types.StringNull(),
		ConnectionType:            types.StringValue("sid"),
		NetServiceName:            types.StringNull(),
		Schema:                    types.StringNull(),
		Query:                     types.StringNull(),
		IncrementalLoadingEnabled: types.BoolValue(true),
		Table:                     types.StringNull(),
		IncrementalColumns:        types.StringNull(),
		LastRecord:                types.StringNull(),
		SourceTimeZone:            types.StringNull(),
		InputOptionColumns:        types.ListNull(types.ObjectType{AttrTypes: oracleInputOptionColumnAttrTypes}),
		InputOptionColumnOptions:  columnOptions,
		CustomVariableSettings:    types.ListNull(types.ObjectType{AttrTypes: common.GetCustomVariableSettingAttrTypes()}),
	}

	input := base.ToUpdateInput(ctx)
	b, err := json.Marshal(input)
	if err != nil {
		t.Fatalf("failed to marshal UpdateOracleInputOptionInput: %s", err)
	}
	body := string(b)

	for _, want := range []string{`"column_name":"amount"`, `"column_value_type":"string"`} {
		if !strings.Contains(body, want) {
			t.Errorf("expected request body to contain %s, got: %s", want, body)
		}
	}
}
