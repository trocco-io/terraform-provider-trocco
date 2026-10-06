package output_options

import (
	"context"
	"testing"

	outputOptionEntities "terraform-provider-trocco/internal/client/entity/job_definition/output_option"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

func namedValueList(ctx context.Context, t *testing.T, values ...CustomConnectorNamedValue) types.List {
	t.Helper()

	if values == nil {
		// A nil slice would produce a null list rather than an empty one.
		values = []CustomConnectorNamedValue{}
	}

	list, diags := types.ListValueFrom(
		ctx,
		types.ObjectType{AttrTypes: CustomConnectorNamedValueAttrTypes()},
		values,
	)
	if diags.HasError() {
		t.Fatalf("failed to build list: %v", diags)
	}
	return list
}

func nullNamedValueList() types.List {
	return types.ListNull(types.ObjectType{AttrTypes: CustomConnectorNamedValueAttrTypes()})
}

func TestCustomConnectorEndpointSettingsToInput(t *testing.T) {
	ctx := context.Background()

	t.Run("nil settings omit the whole key", func(t *testing.T) {
		var settings *CustomConnectorEndpointSettings

		input, diags := settings.toInput(ctx)
		if diags.HasError() {
			t.Fatalf("unexpected diagnostics: %v", diags)
		}
		if input != nil {
			t.Errorf("expected nil input, got %+v", input)
		}
	})

	t.Run("a null collection omits its key so the server keeps the stored values", func(t *testing.T) {
		settings := &CustomConnectorEndpointSettings{
			Headers:         nullNamedValueList(),
			PathParameters:  nullNamedValueList(),
			QueryParameters: nullNamedValueList(),
		}

		input, diags := settings.toInput(ctx)
		if diags.HasError() {
			t.Fatalf("unexpected diagnostics: %v", diags)
		}
		if input.Headers != nil || input.PathParameters != nil || input.QueryParameters != nil {
			t.Errorf("expected every collection to be omitted, got %+v", input)
		}
	})

	t.Run("an empty collection is sent so the server clears the stored values", func(t *testing.T) {
		settings := &CustomConnectorEndpointSettings{
			Headers:         namedValueList(ctx, t),
			PathParameters:  nullNamedValueList(),
			QueryParameters: nullNamedValueList(),
		}

		input, diags := settings.toInput(ctx)
		if diags.HasError() {
			t.Fatalf("unexpected diagnostics: %v", diags)
		}
		if input.Headers == nil {
			t.Fatal("expected headers to be sent")
		}
		if len(*input.Headers) != 0 {
			t.Errorf("expected headers to be empty, got %+v", *input.Headers)
		}
	})

	t.Run("values are sent as a full replacement", func(t *testing.T) {
		settings := &CustomConnectorEndpointSettings{
			Headers:        nullNamedValueList(),
			PathParameters: nullNamedValueList(),
			QueryParameters: namedValueList(ctx, t, CustomConnectorNamedValue{
				Name:  types.StringValue("dry_run"),
				Value: types.StringValue("true"),
			}),
		}

		input, diags := settings.toInput(ctx)
		if diags.HasError() {
			t.Fatalf("unexpected diagnostics: %v", diags)
		}
		if input.QueryParameters == nil || len(*input.QueryParameters) != 1 {
			t.Fatalf("expected a single query parameter, got %+v", input.QueryParameters)
		}
		if (*input.QueryParameters)[0].Name != "dry_run" || (*input.QueryParameters)[0].Value != "true" {
			t.Errorf("unexpected query parameter: %+v", (*input.QueryParameters)[0])
		}
	})
}

func TestNewCustomConnectorOutputOption(t *testing.T) {
	ctx := context.Background()

	t.Run("the server's cleared update_key becomes null", func(t *testing.T) {
		emptyUpdateKey := ""
		outputOption, diags := NewCustomConnectorOutputOption(ctx, &outputOptionEntities.CustomConnectorOutputOption{
			Mode:      "insert",
			UpdateKey: &emptyUpdateKey,
		}, nil)
		if diags.HasError() {
			t.Fatalf("unexpected diagnostics: %v", diags)
		}
		if !outputOption.UpdateKey.IsNull() {
			t.Errorf("expected update_key to be null, got %v", outputOption.UpdateKey)
		}
	})

	t.Run("the endpoint settings are carried over from the previous model", func(t *testing.T) {
		previous := &CustomConnectorOutputOption{
			CreateEndpointSettings: &CustomConnectorEndpointSettings{
				Headers:        nullNamedValueList(),
				PathParameters: nullNamedValueList(),
				QueryParameters: namedValueList(ctx, t, CustomConnectorNamedValue{
					Name:  types.StringValue("dry_run"),
					Value: types.StringValue("true"),
				}),
			},
		}

		// The API never echoes the endpoint settings back.
		outputOption, diags := NewCustomConnectorOutputOption(ctx, &outputOptionEntities.CustomConnectorOutputOption{
			Mode: "insert",
		}, previous)
		if diags.HasError() {
			t.Fatalf("unexpected diagnostics: %v", diags)
		}
		if outputOption.CreateEndpointSettings != previous.CreateEndpointSettings {
			t.Errorf("expected create_endpoint_settings to be carried over, got %+v", outputOption.CreateEndpointSettings)
		}
		if outputOption.UpdateEndpointSettings != nil {
			t.Errorf("expected update_endpoint_settings to stay null, got %+v", outputOption.UpdateEndpointSettings)
		}
	})

	t.Run("the endpoint snapshots are exposed as a list", func(t *testing.T) {
		outputOption, diags := NewCustomConnectorOutputOption(ctx, &outputOptionEntities.CustomConnectorOutputOption{
			Mode: "upsert",
			Endpoints: []outputOptionEntities.CustomConnectorOutputOptionEndpoint{
				{Operation: "create", Method: "POST", Path: "/users"},
				{
					Operation: "update",
					Method:    "PATCH",
					Path:      "/users/{{user_id}}",
					PathParameters: []outputOptionEntities.CustomConnectorNamedValue{
						{Name: "user_id", Value: "id"},
					},
				},
			},
		}, nil)
		if diags.HasError() {
			t.Fatalf("unexpected diagnostics: %v", diags)
		}

		var endpoints []CustomConnectorOutputOptionEndpoint
		if diags := outputOption.Endpoints.ElementsAs(ctx, &endpoints, false); diags.HasError() {
			t.Fatalf("unexpected diagnostics: %v", diags)
		}
		if len(endpoints) != 2 {
			t.Fatalf("expected 2 endpoints, got %d", len(endpoints))
		}
		if endpoints[1].Operation.ValueString() != "update" {
			t.Errorf("expected the second endpoint to be the update one, got %v", endpoints[1].Operation)
		}

		var pathParameters []CustomConnectorNamedValue
		if diags := endpoints[1].PathParameters.ElementsAs(ctx, &pathParameters, false); diags.HasError() {
			t.Fatalf("unexpected diagnostics: %v", diags)
		}
		if len(pathParameters) != 1 || pathParameters[0].Value.ValueString() != "id" {
			t.Errorf("unexpected path parameters: %+v", pathParameters)
		}
	})
}

func TestCustomConnectorOutputOptionToUpdateInput(t *testing.T) {
	ctx := context.Background()

	t.Run("the update-only attributes are omitted while mode is insert", func(t *testing.T) {
		outputOption := &CustomConnectorOutputOption{
			CustomConnectorConnectionID:           types.Int64Value(1),
			Mode:                                  types.StringValue("insert"),
			UpdateKey:                             types.StringNull(),
			CreateCustomConnectorOutputEndpointID: types.Int64Value(2),
			UpdateCustomConnectorOutputEndpointID: types.Int64Null(),
			CustomVariableSettings:                types.ListNull(types.StringType),
		}

		input, diags := outputOption.ToUpdateInput(ctx)
		if diags.HasError() {
			t.Fatalf("unexpected diagnostics: %v", diags)
		}
		if input.UpdateKey != nil {
			t.Errorf("expected update_key to be omitted, got %v", *input.UpdateKey)
		}
		if input.UpdateCustomConnectorOutputEndpointID != nil {
			t.Errorf("expected the update endpoint id to be omitted, got %v", *input.UpdateCustomConnectorOutputEndpointID)
		}
		if input.Mode == nil || *input.Mode != "insert" {
			t.Errorf("expected mode to be sent as insert, got %v", input.Mode)
		}
	})
}
