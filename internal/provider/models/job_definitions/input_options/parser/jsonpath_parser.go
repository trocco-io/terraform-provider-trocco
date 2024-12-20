package parser

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"terraform-provider-trocco/internal/client/entities/job_definitions"
)

type JsonpathParser struct {
	Root            types.String           `tfsdk:"root"`
	DefaultTimeZone types.String           `tfsdk:"default_time_zone"`
	Columns         []JsonpathParserColumn `tfsdk:"columns"`
}

type JsonpathParserColumn struct {
	Name        types.String `tfsdk:"name"`
	Type        types.String `tfsdk:"type"`
	TimeZone    types.String `tfsdk:"time_zone"`
	Format      types.String `tfsdk:"format"`
	ColumnOrder types.Int64  `tfsdk:"column_order"`
}

func NewJsonPathParser(jsonpathParser *job_definitions.JsonpathParser) *JsonpathParser {
	if jsonpathParser == nil {
		return nil
	}
	columns := make([]JsonpathParserColumn, 0, len(jsonpathParser.Columns))
	for _, input := range jsonpathParser.Columns {
		column := JsonpathParserColumn{
			Name:        types.StringValue(input.Name),
			Type:        types.StringValue(input.Type),
			TimeZone:    types.StringPointerValue(input.TimeZone),
			Format:      types.StringPointerValue(input.Format),
			ColumnOrder: types.Int64Value(input.ColumnOrder),
		}
		columns = append(columns, column)
	}
	return &JsonpathParser{
		Root:            types.StringValue(jsonpathParser.Root),
		DefaultTimeZone: types.StringValue(jsonpathParser.DefaultTimeZone),
		Columns:         columns,
	}
}
