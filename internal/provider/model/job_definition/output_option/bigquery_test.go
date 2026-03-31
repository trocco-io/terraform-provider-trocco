package output_options

import (
	"testing"

	"terraform-provider-trocco/internal/client/entity/job_definition/output_option"
	outputOptionParameters "terraform-provider-trocco/internal/client/parameter/job_definition/output_option"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestBigqueryColumnOptionAttrTypes(t *testing.T) {
	t.Run("depth 0 has no fields attribute", func(t *testing.T) {
		attrTypes := bigqueryColumnOptionAttrTypes(0)
		if _, ok := attrTypes["fields"]; ok {
			t.Error("depth 0 should not have fields attribute")
		}
		expectedKeys := []string{"name", "type", "mode", "timestamp_format", "timezone", "description"}
		for _, key := range expectedKeys {
			if _, ok := attrTypes[key]; !ok {
				t.Errorf("missing expected attribute: %s", key)
			}
		}
	})

	t.Run("depth 1 has fields attribute", func(t *testing.T) {
		attrTypes := bigqueryColumnOptionAttrTypes(1)
		fieldsType, ok := attrTypes["fields"]
		if !ok {
			t.Fatal("depth 1 should have fields attribute")
		}
		listType, ok := fieldsType.(types.ListType)
		if !ok {
			t.Fatal("fields should be a ListType")
		}
		objType, ok := listType.ElemType.(types.ObjectType)
		if !ok {
			t.Fatal("fields element should be an ObjectType")
		}
		if _, ok := objType.AttrTypes["fields"]; ok {
			t.Error("nested fields at depth 0 should not have further fields attribute")
		}
	})

	t.Run("depth 2 has two levels of fields", func(t *testing.T) {
		attrTypes := bigqueryColumnOptionAttrTypes(2)
		fieldsType, ok := attrTypes["fields"].(types.ListType)
		if !ok {
			t.Fatal("fields should be a ListType")
		}
		level1, ok := fieldsType.ElemType.(types.ObjectType)
		if !ok {
			t.Fatal("fields element should be an ObjectType")
		}
		level1Fields, ok := level1.AttrTypes["fields"]
		if !ok {
			t.Fatal("level 1 should have fields attribute")
		}
		level1FieldsList, ok := level1Fields.(types.ListType)
		if !ok {
			t.Fatal("level 1 fields should be a ListType")
		}
		level2, ok := level1FieldsList.ElemType.(types.ObjectType)
		if !ok {
			t.Fatal("level 1 fields element should be an ObjectType")
		}
		if _, ok := level2.AttrTypes["fields"]; ok {
			t.Error("level 2 (deepest) should not have fields attribute")
		}
	})
}

func TestNewBigqueryOutputOptionColumnOptions(t *testing.T) {
	t.Run("nil input returns null list", func(t *testing.T) {
		result, err := newBigqueryOutputOptionColumnOptions(nil, 2)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !result.IsNull() {
			t.Error("nil input should return null list")
		}
	})

	t.Run("empty input returns empty list", func(t *testing.T) {
		result, err := newBigqueryOutputOptionColumnOptions([]output_option.BigQueryOutputOptionColumnOption{}, 2)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if result.IsNull() {
			t.Error("empty input should not return null list")
		}
		if len(result.Elements()) != 0 {
			t.Errorf("expected 0 elements, got %d", len(result.Elements()))
		}
	})

	t.Run("flat column options round-trip", func(t *testing.T) {
		desc := "test description"
		input := []output_option.BigQueryOutputOptionColumnOption{
			{
				Name:        "col1",
				Type:        "STRING",
				Mode:        "NULLABLE",
				Description: &desc,
			},
		}
		stateList, err := newBigqueryOutputOptionColumnOptions(input, 2)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(stateList.Elements()) != 1 {
			t.Fatalf("expected 1 element, got %d", len(stateList.Elements()))
		}
		// Verify by converting back to parameter input
		paramResult := toInputBigqueryOutputOptionColumnOptions(stateList)
		if paramResult == nil || len(*paramResult) != 1 {
			t.Fatal("round-trip should produce 1 element")
		}
		assertColumnOption(t, (*paramResult)[0], "col1", "STRING")
		if (*paramResult)[0].Description == nil || *(*paramResult)[0].Description != "test description" {
			t.Error("description should be 'test description'")
		}
	})

	t.Run("nested RECORD fields round-trip", func(t *testing.T) {
		nestedFields := []output_option.BigQueryOutputOptionColumnOption{
			{Name: "nested_col", Type: "INTEGER", Mode: "NULLABLE"},
		}
		input := []output_option.BigQueryOutputOptionColumnOption{
			{Name: "record_col", Type: "RECORD", Mode: "NULLABLE", Fields: &nestedFields},
		}
		stateList, err := newBigqueryOutputOptionColumnOptions(input, 2)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		paramResult := toInputBigqueryOutputOptionColumnOptions(stateList)
		if paramResult == nil || len(*paramResult) != 1 {
			t.Fatal("round-trip should produce 1 element")
		}
		parent := (*paramResult)[0]
		assertColumnOption(t, parent, "record_col", "RECORD")
		if parent.Fields == nil || len(*parent.Fields) != 1 {
			t.Fatal("parent should have 1 nested field")
		}
		assertColumnOption(t, (*parent.Fields)[0], "nested_col", "INTEGER")
	})

	t.Run("deeply nested RECORD fields (3 levels) round-trip", func(t *testing.T) {
		level2Fields := []output_option.BigQueryOutputOptionColumnOption{
			{Name: "deep_col", Type: "STRING", Mode: "NULLABLE"},
		}
		level1Fields := []output_option.BigQueryOutputOptionColumnOption{
			{Name: "mid_col", Type: "RECORD", Mode: "NULLABLE", Fields: &level2Fields},
		}
		input := []output_option.BigQueryOutputOptionColumnOption{
			{Name: "top_col", Type: "RECORD", Mode: "NULLABLE", Fields: &level1Fields},
		}

		stateList, err := newBigqueryOutputOptionColumnOptions(input, 2)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		paramResult := toInputBigqueryOutputOptionColumnOptions(stateList)
		if paramResult == nil {
			t.Fatal("result should not be nil")
		}

		assertColumnOption(t, (*paramResult)[0], "top_col", "RECORD")
		if (*paramResult)[0].Fields == nil {
			t.Fatal("top level should have fields")
		}
		assertColumnOption(t, (*(*paramResult)[0].Fields)[0], "mid_col", "RECORD")
		if (*(*paramResult)[0].Fields)[0].Fields == nil {
			t.Fatal("mid level should have fields")
		}
		assertColumnOption(t, (*(*(*paramResult)[0].Fields)[0].Fields)[0], "deep_col", "STRING")

		deepest := (*(*(*paramResult)[0].Fields)[0].Fields)[0]
		if deepest.Fields != nil {
			t.Error("deepest level should have nil fields")
		}
	})
}

func TestToInputBigqueryOutputOptionColumnOptions(t *testing.T) {
	t.Run("null list returns nil", func(t *testing.T) {
		nullList := types.ListNull(types.ObjectType{AttrTypes: bigqueryColumnOptionAttrTypes(2)})
		result := toInputBigqueryOutputOptionColumnOptions(nullList)
		if result != nil {
			t.Error("null list should return nil")
		}
	})

	t.Run("nested RECORD round-trip preserves all fields", func(t *testing.T) {
		ts := "yyyy-MM-dd"
		tz := "Asia/Tokyo"
		desc := "a record"
		nestedFields := []output_option.BigQueryOutputOptionColumnOption{
			{
				Name:            "child_col",
				Type:            "TIMESTAMP",
				Mode:            "REQUIRED",
				TimestampFormat: &ts,
				Timezone:        &tz,
			},
		}
		entityInput := []output_option.BigQueryOutputOptionColumnOption{
			{
				Name:        "parent_col",
				Type:        "RECORD",
				Mode:        "NULLABLE",
				Description: &desc,
				Fields:      &nestedFields,
			},
		}

		stateList, err := newBigqueryOutputOptionColumnOptions(entityInput, 2)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		paramResult := toInputBigqueryOutputOptionColumnOptions(stateList)
		if paramResult == nil || len(*paramResult) != 1 {
			t.Fatal("result should have 1 element")
		}

		parent := (*paramResult)[0]
		assertColumnOption(t, parent, "parent_col", "RECORD")
		if parent.Description == nil || *parent.Description != "a record" {
			t.Error("parent description should be preserved")
		}
		if parent.Fields == nil || len(*parent.Fields) != 1 {
			t.Fatal("parent should have 1 child")
		}

		child := (*parent.Fields)[0]
		assertColumnOption(t, child, "child_col", "TIMESTAMP")
		if child.Mode != "REQUIRED" {
			t.Errorf("expected child mode 'REQUIRED', got '%s'", child.Mode)
		}
		if child.TimestampFormat == nil || *child.TimestampFormat != "yyyy-MM-dd" {
			t.Error("child timestamp_format should be preserved")
		}
		if child.Timezone == nil || *child.Timezone != "Asia/Tokyo" {
			t.Error("child timezone should be preserved")
		}
	})
}

func assertColumnOption(t *testing.T, opt outputOptionParameters.BigQueryOutputOptionColumnOptionInput, expectedName, expectedType string) {
	t.Helper()
	if opt.Name != expectedName {
		t.Errorf("expected name '%s', got '%s'", expectedName, opt.Name)
	}
	if opt.Type != expectedType {
		t.Errorf("expected type '%s', got '%s'", expectedType, opt.Type)
	}
}
