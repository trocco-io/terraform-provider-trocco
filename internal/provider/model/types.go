package model

import (
	"terraform-provider-trocco/internal/client/parameter"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

func NewNullableBool(v types.Bool) *parameter.NullableBool {
	if v.IsUnknown() {
		return nil
	}

	return &parameter.NullableBool{Valid: !v.IsNull(), Value: v.ValueBool()}
}

func NewNullableInt64(v types.Int64) *parameter.NullableInt64 {
	if v.IsUnknown() {
		return nil
	}

	return &parameter.NullableInt64{Valid: !v.IsNull(), Value: v.ValueInt64()}
}

func NewNullableString(v types.String) *parameter.NullableString {
	if v.IsUnknown() {
		return nil
	}

	return &parameter.NullableString{Valid: !v.IsNull(), Value: v.ValueString()}
}

func WrapObject[T any](v *T) *parameter.NullableObject[T] {
	return &parameter.NullableObject[T]{Valid: v != nil, Value: v}
}

func WrapObjectList[T parameter.SliceConstraint[E], E any](v *T) *parameter.NullableObjectList[T, E] {
	return &parameter.NullableObjectList[T, E]{Valid: v != nil, Value: v}
}
