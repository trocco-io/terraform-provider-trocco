package provider

import (
	"terraform-provider-trocco/internal/client"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

// NewNullableFromTerraformString create a client.NullableString from a types.String.
func newNullableFromTerraformString(v types.String) *client.NullableString {
	return &client.NullableString{Valid: !v.IsNull(), Value: v.ValueString()}
}

func ExampleNewNullableFromTerraformInt64() {
	newNullableFromTerraformInt64(types.Int64Null())
	// Output: i1.Valid = false

	newNullableFromTerraformInt64(types.Int64Value(42))
	// Output: i2.Valid = true, i2.Value = 42
}

// NewNullableFromTerraformInt64 create a client.NullableInt64 from a types.Int64.
func newNullableFromTerraformInt64(v types.Int64) *client.NullableInt64 {
	return &client.NullableInt64{Valid: !v.IsNull(), Value: v.ValueInt64()}
}

func ExampleNewNullableFromTerraformString() {
	newNullableFromTerraformString(types.StringNull())
	// Output: s1.Valid = false

	newNullableFromTerraformString(types.StringValue("example"))
	// Output: s2.Valid = true, s2.Value = "example"
}

// NewNullableFromTerraformBool create a client.NullableBool from a types.Bool.
func newNullableFromTerraformBool(v types.Bool) *client.NullableBool {
	return &client.NullableBool{Valid: !v.IsNull(), Value: v.ValueBool()}
}

func ExampleNewNullableFromTerraformBool() {
	newNullableFromTerraformBool(types.BoolNull())
	// Output: b1.Valid = false

	newNullableFromTerraformBool(types.BoolValue(true))
	// Output: b2.Valid = true, b2.Value = true
}
