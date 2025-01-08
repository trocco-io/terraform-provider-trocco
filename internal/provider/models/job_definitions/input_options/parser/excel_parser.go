package parser

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"terraform-provider-trocco/internal/client/entities/job_definitions"
	job_definitions2 "terraform-provider-trocco/internal/client/parameters/job_definitions"
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
			FormulaHandling: types.StringValue(input.FormulaHandling),
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

func (excelParser *ExcelParser) ToExcelParserInput() *job_definitions2.ExcelParserInput {
	if excelParser == nil {
		return nil
	}
	columns := make([]job_definitions2.ExcelParserColumnInput, 0, len(excelParser.Columns))
	for _, input := range excelParser.Columns {
		column := job_definitions2.ExcelParserColumnInput{
			Name:            input.Name.ValueString(),
			Type:            input.Type.ValueString(),
			Format:          input.Format.ValueStringPointer(),
			FormulaHandling: input.FormulaHandling.ValueString(),
		}
		columns = append(columns, column)
	}

	return &job_definitions2.ExcelParserInput{
		DefaultTimeZone: excelParser.DefaultTimeZone.ValueString(),
		SheetName:       excelParser.SheetName.ValueString(),
		SkipHeaderLines: excelParser.SkipHeaderLines.ValueInt64(),
		Columns:         columns,
	}
}

func ToExcelParserModel(excelParser *job_definitions.ExcelParser) *ExcelParser {
	if excelParser == nil {
		return nil
	}
	columns := make([]ExcelParserColumn, 0, len(excelParser.Columns))
	for _, input := range excelParser.Columns {
		column := ExcelParserColumn{
			Name:            types.StringValue(input.Name),
			Type:            types.StringValue(input.Type),
			Format:          types.StringPointerValue(input.Format),
			FormulaHandling: types.StringValue(input.FormulaHandling),
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
