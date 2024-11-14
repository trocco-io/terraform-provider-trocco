package provider

import (
	"terraform-provider-trocco/internal/client"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

// NewNullableFromTerraformString create a client.NullableString from a types.String.
func newNullableFromTerraformString(v types.String) *client.NullableString {
	return &client.NullableString{Valid: !v.IsNull(), Value: v.ValueString()}
}

// NewNullableFromTerraformBool create a client.NullableBool from a types.Bool.
func newNullableFromTerraformBool(v types.Bool) *client.NullableBool {
	return &client.NullableBool{Valid: !v.IsNull(), Value: v.ValueBool()}
}
