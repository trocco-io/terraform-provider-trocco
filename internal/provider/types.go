// internal/provider/types.go

package provider

import (
	"terraform-provider-trocco/internal/client"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

// NewNullableFromTerraformInt64 create a client.NullableInt64 from a types.Int64.
func newNullableFromTerraformInt64(v types.Int64) *client.NullableInt64 {
	return &client.NullableInt64{Valid: !v.IsNull(), Value: v.ValueInt64()}
}
