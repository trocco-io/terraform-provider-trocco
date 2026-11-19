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
// value should be cleared. source_time_zone has no server-side default (unlike
// the deprecated default_time_zone it replaces), so it must use
// *parameter.NullableString so the JSON request body carries an explicit
// `null`, not a missing key, when cleared.
//
// Unlike acceptance tests, this requires no TROCCO account or network access,
// and runs as a normal `go test` so it is exercised on every pull request.
func newBaseSQLServerInputOption() *SQLServerInputOption {
	return &SQLServerInputOption{
		SQLServerConnectionID:     types.Int64Value(1),
		Database:                  types.StringValue("test_database"),
		Schema:                    types.StringNull(),
		Query:                     types.StringNull(),
		IncrementalLoadingEnabled: types.BoolValue(true),
		Table:                     types.StringValue("test_table"),
		IncrementalColumns:        types.StringNull(),
		LastRecord:                types.StringNull(),
		FetchRows:                 types.Int64Value(10_000),
		ConnectTimeout:            types.Int64Value(300),
		SocketTimeout:             types.Int64Value(1_800),
		DefaultTimeZone:           types.StringNull(),
		SourceTimeZone:            types.StringNull(),
		InputOptionColumns:        types.ListNull(types.ObjectType{AttrTypes: SQLServerInputOptionColumn{}.attrTypes()}),
		InputOptionColumnOptions:  types.ListNull(types.ObjectType{AttrTypes: InputOptionColumnOptions{}.attrTypes()}),
		CustomVariableSettings:    types.ListNull(types.ObjectType{AttrTypes: common.GetCustomVariableSettingAttrTypes()}),
	}
}

func TestSQLServerInputOptionToUpdateInput_ClearsSourceTimeZoneOnRemoval(t *testing.T) {
	ctx := context.Background()

	base := newBaseSQLServerInputOption()

	input := base.ToUpdateInput(ctx)
	b, err := json.Marshal(input)
	if err != nil {
		t.Fatalf("failed to marshal UpdateSQLServerInputOptionInput: %s", err)
	}
	body := string(b)

	key := `"source_time_zone":null`
	if !strings.Contains(body, key) {
		t.Errorf("expected request body to contain %s (key present with explicit null) when the field is unset, got: %s", key, body)
	}
}

func TestSQLServerInputOptionToUpdateInput_SendsSourceTimeZoneValue(t *testing.T) {
	ctx := context.Background()

	base := newBaseSQLServerInputOption()
	base.SourceTimeZone = types.StringValue("Asia/Tokyo")

	input := base.ToUpdateInput(ctx)
	b, err := json.Marshal(input)
	if err != nil {
		t.Fatalf("failed to marshal UpdateSQLServerInputOptionInput: %s", err)
	}
	body := string(b)

	want := `"source_time_zone":"Asia/Tokyo"`
	if !strings.Contains(body, want) {
		t.Errorf("expected request body to contain %s, got: %s", want, body)
	}
}

// default_time_zone is deprecated and intentionally never sent to the API
// (see NewSQLServerInputOption), so it must not appear in the update request
// body at all, regardless of what value is set on the model.
func TestSQLServerInputOptionToUpdateInput_NeverSendsDefaultTimeZone(t *testing.T) {
	ctx := context.Background()

	base := newBaseSQLServerInputOption()
	base.DefaultTimeZone = types.StringValue("Asia/Tokyo")

	input := base.ToUpdateInput(ctx)
	b, err := json.Marshal(input)
	if err != nil {
		t.Fatalf("failed to marshal UpdateSQLServerInputOptionInput: %s", err)
	}
	body := string(b)

	if strings.Contains(body, "default_time_zone") {
		t.Errorf("expected request body to never contain default_time_zone, got: %s", body)
	}
}
