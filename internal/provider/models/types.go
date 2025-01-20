package models

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"terraform-provider-trocco/internal/client/parameters"
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

func NewNullableString(v types.String) *parameters.NullableString {
	if v.IsUnknown() {
		return nil
	}
	return &parameters.NullableString{Valid: !v.IsNull(), Value: v.ValueString()}
}

func WrapObject[T any](v *T) *parameters.NullableObject[T] {
	return &parameters.NullableObject[T]{Valid: v != nil, Value: v}
}
