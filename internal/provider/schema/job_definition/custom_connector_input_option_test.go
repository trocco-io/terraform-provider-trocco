package job_definition

import (
	"context"
	"maps"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

func customConnectorInputOptionObjectType(t *testing.T) types.ObjectType {
	t.Helper()

	objectType, ok := CustomConnectorInputOptionSchema().GetType().(types.ObjectType)
	if !ok {
		t.Fatalf("custom_connector_input_option is not an object type")
	}
	return objectType
}

func nullValue(t *testing.T, attrType attr.Type) attr.Value {
	t.Helper()

	value, err := attrType.ValueFromTerraform(
		context.Background(),
		tftypes.NewValue(attrType.TerraformType(context.Background()), nil),
	)
	if err != nil {
		t.Fatalf("building a null value: %s", err)
	}
	return value
}

// customConnectorInputOptionConfig builds a configuration object in which every
// attribute is null, then applies the given overrides. A null attribute stands
// for one that is absent from the configuration.
func customConnectorInputOptionConfig(t *testing.T, overrides map[string]attr.Value) types.Object {
	t.Helper()

	objectType := customConnectorInputOptionObjectType(t)

	attributes := make(map[string]attr.Value, len(objectType.AttrTypes))
	for name, attrType := range objectType.AttrTypes {
		attributes[name] = nullValue(t, attrType)
	}
	maps.Copy(attributes, overrides)

	config, diags := types.ObjectValue(objectType.AttrTypes, attributes)
	if diags.HasError() {
		t.Fatalf("building the configuration object: %s", diags)
	}
	return config
}

// configuredJsonpathParser is a jsonpath_parser block as it appears in a
// configuration: root is never set (the API derives it from the endpoint),
// while columns is.
func configuredJsonpathParser(t *testing.T) types.Object {
	t.Helper()

	objectType, ok := customConnectorJsonpathParserSchema().GetType().(types.ObjectType)
	if !ok {
		t.Fatalf("jsonpath_parser is not an object type")
	}

	value, diags := types.ObjectValue(objectType.AttrTypes, map[string]attr.Value{
		"root":              types.StringNull(),
		"default_time_zone": types.StringNull(),
		"columns":           emptyList(t, objectType, "columns"),
	})
	if diags.HasError() {
		t.Fatalf("building jsonpath_parser: %s", diags)
	}
	return value
}

// emptyList builds an empty list for the named list attribute of objectType,
// standing for a list that is present in the configuration.
func emptyList(t *testing.T, objectType types.ObjectType, name string) types.List {
	t.Helper()

	listType, ok := objectType.AttrTypes[name].(types.ListType)
	if !ok {
		t.Fatalf("%s is not a list type", name)
	}

	value, diags := types.ListValue(listType.ElemType, []attr.Value{})
	if diags.HasError() {
		t.Fatalf("building %s: %s", name, diags)
	}
	return value
}

func snapshotPaths(attributes []CustomConnectorSnapshotAttribute) []string {
	paths := make([]string, 0, len(attributes))
	for _, attribute := range attributes {
		paths = append(paths, attribute.Path.String())
	}
	return paths
}

func TestCustomConnectorSnapshotAttributes(t *testing.T) {
	base := path.Root("input_option").AtName("custom_connector_input_option")

	t.Run("named value lists are server-derived while omitted", func(t *testing.T) {
		config := customConnectorInputOptionConfig(t, map[string]attr.Value{
			"custom_connector_endpoint_id":   types.Int64Value(1),
			"custom_connector_connection_id": types.Int64Value(2),
			"jsonpath_parser":                configuredJsonpathParser(t),
		})

		want := []string{
			"input_option.custom_connector_input_option.auth_header_name",
			"input_option.custom_connector_input_option.auth_header_scheme",
			"input_option.custom_connector_input_option.auth_type",
			"input_option.custom_connector_input_option.endpoint_method",
			"input_option.custom_connector_input_option.endpoint_path",
			"input_option.custom_connector_input_option.endpoint_request_body",
			"input_option.custom_connector_input_option.headers",
			"input_option.custom_connector_input_option.jsonpath_parser.root",
			"input_option.custom_connector_input_option.not_retryable_codes",
			"input_option.custom_connector_input_option.paginator",
			"input_option.custom_connector_input_option.path_parameters",
			"input_option.custom_connector_input_option.query_parameters",
			"input_option.custom_connector_input_option.request_body_parameters",
			"input_option.custom_connector_input_option.request_timeout_sec",
			"input_option.custom_connector_input_option.success_codes",
			"input_option.custom_connector_input_option.url",
		}

		got := snapshotPaths(CustomConnectorSnapshotAttributes(base, config))
		if len(got) != len(want) {
			t.Fatalf("got %d attributes %v, want %d %v", len(got), got, len(want), want)
		}
		for i := range want {
			if got[i] != want[i] {
				t.Errorf("attribute %d: got %q, want %q", i, got[i], want[i])
			}
		}
	})

	t.Run("a configured named value list is not server-derived", func(t *testing.T) {
		config := customConnectorInputOptionConfig(t, map[string]attr.Value{
			"jsonpath_parser":  configuredJsonpathParser(t),
			"query_parameters": emptyList(t, customConnectorInputOptionObjectType(t), "query_parameters"),
		})

		for _, p := range snapshotPaths(CustomConnectorSnapshotAttributes(base, config)) {
			if p == "input_option.custom_connector_input_option.query_parameters" {
				t.Errorf("query_parameters is configured, so it must not be planned as unknown")
			}
		}
	})
}
