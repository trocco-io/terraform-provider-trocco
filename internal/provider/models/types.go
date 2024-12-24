package models

import (
	"terraform-provider-trocco/internal/client/parameters"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

func NewNullableBool(v types.Bool) *parameters.NullableBool {
	return &parameters.NullableBool{Valid: !v.IsNull(), Value: v.ValueBool()}
}

func NewNullableInt64(v types.Int64) *parameters.NullableInt64 {
	return &parameters.NullableInt64{Valid: !v.IsNull(), Value: v.ValueInt64()}
}
