package provider

import (
	"terraform-provider-trocco/internal/client"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

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
