package models

import (
	"terraform-provider-trocco/internal/client/parameters"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

func NewNullableBool(v types.Bool) *parameters.NullableBool {
	if v.IsUnknown() {
		return nil
	}

	return &parameters.NullableBool{Valid: !v.IsNull(), Value: v.ValueBool()}
}

func NewNullableInt64(v types.Int64) *parameters.NullableInt64 {
	if v.IsUnknown() {
		return nil
	}

	return &parameters.NullableInt64{Valid: !v.IsNull(), Value: v.ValueInt64()}
}
