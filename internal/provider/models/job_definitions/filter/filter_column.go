package filter

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"terraform-provider-trocco/internal/client/entities/job_definitions/filter"
)

type FilterColumn struct {
	Name                     types.String       `tfsdk:"name"`
	Src                      types.String       `tfsdk:"src"`
	Type                     types.String       `tfsdk:"type"`
	Default                  types.String       `tfsdk:"default"`
	HasParser                types.Bool         `tfsdk:"has_parser"`
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
		expandColumns := make([]jsonExpandColumn, 0, len(input.JSONExpandColumns))
		for _, input := range input.JSONExpandColumns {
			column := jsonExpandColumn{
				Name:     types.StringValue(input.Name),
				JSONPath: types.StringValue(input.JSONPath),
				Type:     types.StringValue(input.Type),
				Format:   types.StringPointerValue(input.Format),
				Timezone: types.StringValue(input.Type),
			}
			expandColumns = append(expandColumns, column)
		}

		filterColumn := FilterColumn{
			Name:                     types.StringValue(input.Name),
			Src:                      types.StringValue(input.Src),
			Type:                     types.StringValue(input.Type),
			Default:                  types.StringPointerValue(input.Default),
			HasParser:                types.BoolValue(input.HasParser),
			Format:                   types.StringPointerValue(input.Format),
			JSONExpandEnabled:        types.BoolValue(input.JSONExpandEnabled),
			JSONExpandKeepBaseColumn: types.BoolValue(input.JSONExpandKeepBaseColumn),
			JSONExpandColumns:        expandColumns,
		}
		outputs = append(outputs, filterColumn)
	}
	return outputs
}
