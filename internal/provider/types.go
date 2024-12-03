// internal/provider/types.go

package provider

import (
	"terraform-provider-trocco/internal/client"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

// NewNullableFromTerraformString create a client.NullableString from a types.String.
func newNullableFromTerraformString(v types.String) *client.NullableString {
	return &client.NullableString{Valid: !v.IsNull(), Value: v.ValueString()}
}

// NewNullableFromTerraformInt64 create a client.NullableInt64 from a types.Int64.
func newNullableFromTerraformInt64(v types.Int64) *client.NullableInt64 {
	return &client.NullableInt64{Valid: !v.IsNull(), Value: v.ValueInt64()}
}

// NewNullableFromTerraformBool create a client.NullableBool from a types.Bool.
func newNullableFromTerraformBool(v types.Bool) *client.NullableBool {
	return &client.NullableBool{Valid: !v.IsNull(), Value: v.ValueBool()}
}
