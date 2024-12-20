package parser

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"terraform-provider-trocco/internal/client/entities/job_definitions"
)

type ExcelParser struct {
	DefaultTimeZone types.String        `tfsdk:"default_time_zone"`
	SheetName       types.String        `tfsdk:"sheet_name"`
	SkipHeaderLines types.Int64         `tfsdk:"skip_header_lines"`
	Columns         []ExcelParserColumn `tfsdk:"columns"`
}

type ExcelParserColumn struct {
	Name            types.String `tfsdk:"name"`
	Type            types.String `tfsdk:"type"`
	Format          types.String `tfsdk:"format"`
	FormulaHandling types.String `tfsdk:"formula_handling"`
	ColumnOrder     types.Int64  `tfsdk:"column_order"`
}

func NewExcelParser(excelParser *job_definitions.ExcelParser) *ExcelParser {
	if excelParser == nil {
		return nil
	}
	columns := make([]ExcelParserColumn, 0, len(excelParser.Columns))
	for _, input := range excelParser.Columns {
		column := ExcelParserColumn{
			Name:            types.StringValue(input.Name),
			Type:            types.StringValue(input.Type),
			Format:          types.StringPointerValue(input.Format),
			FormulaHandling: types.StringPointerValue(input.FormulaHandling),
			ColumnOrder:     types.Int64Value(input.ColumnOrder),
		}
		columns = append(columns, column)
	}
	return &ExcelParser{
		DefaultTimeZone: types.StringValue(excelParser.DefaultTimeZone),
		SheetName:       types.StringValue(excelParser.SheetName),
		SkipHeaderLines: types.Int64Value(excelParser.SkipHeaderLines),
		Columns:         columns,
	}
}
