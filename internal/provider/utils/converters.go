package utils

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ConvertSetToSlice converts a types.Set to a slice using a converter function.
func ConvertSetToSlice[T any, U any](
	ctx context.Context,
	source types.Set,
	converter func(T) U,
	diags *diag.Diagnostics,
) []U {
	if source.IsNull() || source.IsUnknown() {
		return []U{}
	}

	var values []T
	diags.Append(source.ElementsAs(ctx, &values, false)...)
	if diags.HasError() {
		return []U{}
	}

	result := make([]U, 0, len(values))
	for _, v := range values {
		result = append(result, converter(v))
	}
	return result
}

// ConvertListToSlice converts a types.List to a slice using a converter function.
func ConvertListToSlice[T any, U any](
	ctx context.Context,
	source types.List,
	converter func(T) U,
	diags *diag.Diagnostics,
) []U {
	if source.IsNull() || source.IsUnknown() {
		return []U{}
	}

	var values []T
	diags.Append(source.ElementsAs(ctx, &values, false)...)
	if diags.HasError() {
		return []U{}
	}

	result := make([]U, 0, len(values))
	for _, v := range values {
		result = append(result, converter(v))
	}
	return result
}

// ConvertStringSet converts types.Set to string slice with error handling.
func ConvertStringSet(ctx context.Context, source types.Set) ([]string, bool) {
	if source.IsNull() || source.IsUnknown() {
		return []string{}, true
	}

	var values []types.String
	if diags := source.ElementsAs(ctx, &values, false); !diags.HasError() {
		result := make([]string, 0, len(values))
		for _, v := range values {
			result = append(result, v.ValueString())
		}
		return result, true
	}
	return []string{}, false
}

// ConvertStringList converts types.List to string slice.
func ConvertStringList(ctx context.Context, source types.List) []string {
	if source.IsNull() || source.IsUnknown() {
		return []string{}
	}

	var values []types.String
	if diags := source.ElementsAs(ctx, &values, false); !diags.HasError() {
		result := make([]string, 0, len(values))
		for _, v := range values {
			result = append(result, v.ValueString())
		}
		return result
	}
	return []string{}
}
