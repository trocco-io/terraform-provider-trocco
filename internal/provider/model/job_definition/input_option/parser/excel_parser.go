package parser

import (
	"context"
	job_definitions "terraform-provider-trocco/internal/client/entity/job_definition"
	params "terraform-provider-trocco/internal/client/parameter/job_definition"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ExcelParser struct {
	DefaultTimeZone types.String `tfsdk:"default_time_zone"`
	SheetName       types.String `tfsdk:"sheet_name"`
	SkipHeaderLines types.Int64  `tfsdk:"skip_header_lines"`
	Columns         types.List   `tfsdk:"columns"`
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

	columnElements := make([]ExcelParserColumn, 0, len(excelParser.Columns))
	for _, input := range excelParser.Columns {
		column := ExcelParserColumn{
			Name:            types.StringValue(input.Name),
			Type:            types.StringValue(input.Type),
			Format:          types.StringPointerValue(input.Format),
			FormulaHandling: types.StringValue(input.FormulaHandling),
		}
		columnElements = append(columnElements, column)
	}

	columns, _ := types.ListValueFrom(
		context.Background(),
		types.ObjectType{
			AttrTypes: map[string]attr.Type{
				"name":             types.StringType,
				"type":             types.StringType,
				"format":           types.StringType,
				"formula_handling": types.StringType,
			},
		},
		columnElements,
	)
	return &ExcelParser{
		DefaultTimeZone: types.StringValue(excelParser.DefaultTimeZone),
		SheetName:       types.StringValue(excelParser.SheetName),
		SkipHeaderLines: types.Int64Value(excelParser.SkipHeaderLines),
		Columns:         columns,
	}
}

func (excelParser *ExcelParser) ToExcelParserInput() *params.ExcelParserInput {
	if excelParser == nil {
		return nil
	}

	var columnElements []ExcelParserColumn
	excelParser.Columns.ElementsAs(context.Background(), &columnElements, false)

	columns := make([]params.ExcelParserColumnInput, 0, len(columnElements))
	for _, input := range columnElements {
		column := params.ExcelParserColumnInput{
			Name:            input.Name.ValueString(),
			Type:            input.Type.ValueString(),
			Format:          input.Format.ValueStringPointer(),
			FormulaHandling: input.FormulaHandling.ValueString(),
		}
		columns = append(columns, column)
	}

	return &params.ExcelParserInput{
		DefaultTimeZone: excelParser.DefaultTimeZone.ValueString(),
		SheetName:       excelParser.SheetName.ValueString(),
		SkipHeaderLines: excelParser.SkipHeaderLines.ValueInt64(),
		Columns:         columns,
	}
}
