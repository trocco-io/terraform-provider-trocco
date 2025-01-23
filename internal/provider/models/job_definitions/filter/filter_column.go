package filter

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"terraform-provider-trocco/internal/client/entities/job_definitions/filter"
	filter2 "terraform-provider-trocco/internal/client/parameters/job_definitions/filter"
)

type FilterColumn struct {
	Name                     types.String       `tfsdk:"name"`
	Src                      types.String       `tfsdk:"src"`
	Type                     types.String       `tfsdk:"type"`
	Default                  types.String       `tfsdk:"default"`
	Format                   types.String       `tfsdk:"format"`
	JSONExpandEnabled        types.Bool         `tfsdk:"json_expand_enabled"`
	JSONExpandKeepBaseColumn types.Bool         `tfsdk:"json_expand_keep_base_column"`
	JSONExpandColumns        []jsonExpandColumn `tfsdk:"json_expand_columns"`
}

type jsonExpandColumn struct {
	Name     types.String `tfsdk:"name"`
	JSONPath types.String `tfsdk:"json_path"`
	Type     types.String `tfsdk:"type"`
	Format   types.String `tfsdk:"format"`
	Timezone types.String `tfsdk:"timezone"`
}

func NewFilterColumns(filterColumns []filter.FilterColumn) []FilterColumn {
	outputs := make([]FilterColumn, 0, len(filterColumns))
	for _, input := range filterColumns {
		var expandColumns []jsonExpandColumn
		for _, input := range input.JSONExpandColumns {
			column := jsonExpandColumn{
				Name:     types.StringValue(input.Name),
				JSONPath: types.StringValue(input.JSONPath),
				Type:     types.StringValue(input.Type),
				Format:   types.StringPointerValue(input.Format),
				Timezone: types.StringPointerValue(input.Timezone),
			}
			expandColumns = append(expandColumns, column)
		}

		filterColumn := FilterColumn{
			Name:                     types.StringValue(input.Name),
			Src:                      types.StringValue(input.Src),
			Type:                     types.StringValue(input.Type),
			Default:                  types.StringPointerValue(input.Default),
			Format:                   types.StringPointerValue(input.Format),
			JSONExpandEnabled:        types.BoolValue(input.JSONExpandEnabled),
			JSONExpandKeepBaseColumn: types.BoolValue(input.JSONExpandKeepBaseColumn),
			JSONExpandColumns:        expandColumns,
		}
		outputs = append(outputs, filterColumn)
	}
	return outputs
}

func (filterColumn FilterColumn) ToInput() filter2.FilterColumnInput {
	return filter2.FilterColumnInput{
		Name:                     filterColumn.Name.ValueString(),
		Src:                      filterColumn.Src.ValueString(),
		Type:                     filterColumn.Type.ValueString(),
		Default:                  filterColumn.Default.ValueStringPointer(),
		Format:                   filterColumn.Format.ValueStringPointer(),
		JSONExpandEnabled:        filterColumn.JSONExpandEnabled.ValueBool(),
		JSONExpandKeepBaseColumn: filterColumn.JSONExpandKeepBaseColumn.ValueBool(),
		JSONExpandColumns:        jsonExpandColumns(filterColumn.JSONExpandColumns),
	}
}

func jsonExpandColumns(columns []jsonExpandColumn) []filter2.JSONExpandColumnInput {
	outputs := make([]filter2.JSONExpandColumnInput, 0, len(columns))
	for _, input := range columns {
		column := filter2.JSONExpandColumnInput{
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
