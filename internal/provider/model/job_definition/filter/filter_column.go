package filter

import (
	"context"
	"terraform-provider-trocco/internal/client/entity/job_definition/filter"
	filterParameters "terraform-provider-trocco/internal/client/parameter/job_definition/filter"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type FilterColumn struct {
	Name                     types.String `tfsdk:"name"`
	Src                      types.String `tfsdk:"src"`
	Type                     types.String `tfsdk:"type"`
	Default                  types.String `tfsdk:"default"`
	Format                   types.String `tfsdk:"format"`
	JSONExpandEnabled        types.Bool   `tfsdk:"json_expand_enabled"`
	JSONExpandKeepBaseColumn types.Bool   `tfsdk:"json_expand_keep_base_column"`
	JSONExpandColumns        types.List   `tfsdk:"json_expand_columns"`
}

type jsonExpandColumn struct {
	Name     types.String `tfsdk:"name"`
	JSONPath types.String `tfsdk:"json_path"`
	Type     types.String `tfsdk:"type"`
	Format   types.String `tfsdk:"format"`
	Timezone types.String `tfsdk:"timezone"`
}

func NewFilterColumns(ctx context.Context, filterColumns []filter.FilterColumn) ([]FilterColumn, diag.Diagnostics) {
	outputs := make([]FilterColumn, 0, len(filterColumns))
	var diags diag.Diagnostics

	for _, input := range filterColumns {
		var expandColumns []jsonExpandColumn
		for _, jc := range input.JSONExpandColumns {
			column := jsonExpandColumn{
				Name:     types.StringValue(jc.Name),
				JSONPath: types.StringValue(jc.JSONPath),
				Type:     types.StringValue(jc.Type),
				Format:   types.StringPointerValue(jc.Format),
				Timezone: types.StringPointerValue(jc.Timezone),
			}
			expandColumns = append(expandColumns, column)
		}

		var expandColumnsValue types.List
		if len(expandColumns) > 0 {
			v, d := types.ListValueFrom(
				ctx,
				types.ObjectType{AttrTypes: jsonExpandColumn{}.AttrTypes()},
				expandColumns,
			)
			diags.Append(d...)
			expandColumnsValue = v
		} else {
			expandColumnsValue = types.ListNull(types.ObjectType{
				AttrTypes: jsonExpandColumn{}.AttrTypes(),
			})
		}

		filterColumn := FilterColumn{
			Name:                     types.StringValue(input.Name),
			Src:                      types.StringValue(input.Src),
			Type:                     types.StringValue(input.Type),
			Default:                  types.StringPointerValue(input.Default),
			Format:                   types.StringPointerValue(input.Format),
			JSONExpandEnabled:        types.BoolValue(input.JSONExpandEnabled),
			JSONExpandKeepBaseColumn: types.BoolValue(input.JSONExpandKeepBaseColumn),
			JSONExpandColumns:        expandColumnsValue,
		}
		outputs = append(outputs, filterColumn)
	}

	return outputs, diags
}

func (filterColumn FilterColumn) ToInput(ctx context.Context) filterParameters.FilterColumnInput {
	var expandColumns []jsonExpandColumn
	if !filterColumn.JSONExpandColumns.IsNull() && !filterColumn.JSONExpandColumns.IsUnknown() {
		_ = filterColumn.JSONExpandColumns.ElementsAs(ctx, &expandColumns, false)
	}

	return filterParameters.FilterColumnInput{
		Name:                     filterColumn.Name.ValueString(),
		Src:                      filterColumn.Src.ValueString(),
		Type:                     filterColumn.Type.ValueString(),
		Default:                  filterColumn.Default.ValueStringPointer(),
		Format:                   filterColumn.Format.ValueStringPointer(),
		JSONExpandEnabled:        filterColumn.JSONExpandEnabled.ValueBool(),
		JSONExpandKeepBaseColumn: filterColumn.JSONExpandKeepBaseColumn.ValueBool(),
		JSONExpandColumns:        jsonExpandColumns(expandColumns),
	}
}

func jsonExpandColumns(columns []jsonExpandColumn) []filterParameters.JSONExpandColumnInput {
	outputs := make([]filterParameters.JSONExpandColumnInput, 0, len(columns))
	for _, input := range columns {
		column := filterParameters.JSONExpandColumnInput{
			Name:     input.Name.ValueString(),
			JSONPath: input.JSONPath.ValueString(),
			Type:     input.Type.ValueString(),
			Format:   input.Format.ValueStringPointer(),
			Timezone: input.Timezone.ValueStringPointer(),
		}
		outputs = append(outputs, column)
	}
	return outputs
}

func (c FilterColumn) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"name":                         types.StringType,
		"src":                          types.StringType,
		"type":                         types.StringType,
		"default":                      types.StringType,
		"format":                       types.StringType,
		"json_expand_enabled":          types.BoolType,
		"json_expand_keep_base_column": types.BoolType,
		"json_expand_columns": types.ListType{ElemType: types.ObjectType{
			AttrTypes: jsonExpandColumn{}.AttrTypes(),
		}},
	}
}

func (c jsonExpandColumn) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"name":      types.StringType,
		"json_path": types.StringType,
		"type":      types.StringType,
		"format":    types.StringType,
		"timezone":  types.StringType,
	}
}
