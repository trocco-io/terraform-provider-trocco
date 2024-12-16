// internal/provider/types.go

package provider

import (
	"terraform-provider-trocco/internal/client/parameters"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

// NewNullableFromTerraformInt64 create a client.NullableInt64 from a types.Int64.
func newNullableFromTerraformInt64(v types.Int64) *parameters.NullableInt64 {
	return &parameters.NullableInt64{Valid: !v.IsNull(), Value: v.ValueInt64()}
}

// NewNullableFromTerraformBool create a client.NullableBool from a types.Bool.
func newNullableFromTerraformBool(v types.Bool) *parameters.NullableBool {
	return &parameters.NullableBool{Valid: !v.IsNull(), Value: v.ValueBool()}
}
