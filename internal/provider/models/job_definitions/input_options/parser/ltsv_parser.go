package parser

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"terraform-provider-trocco/internal/client/entities/job_definitions"
	job_definitions2 "terraform-provider-trocco/internal/client/parameters/job_definitions"
)

type LtsvParser struct {
	Newline types.String       `tfsdk:"newline"`
	Charset types.String       `tfsdk:"charset"`
	Columns []LtsvParserColumn `tfsdk:"columns"`
}

type LtsvParserColumn struct {
	Name        types.String `tfsdk:"name"`
	Type        types.String `tfsdk:"type"`
	Format      types.String `tfsdk:"format"`
	ColumnOrder types.Int64  `tfsdk:"column_order"`
}

func NewLtsvParser(ltsvParser *job_definitions.LtsvParser) *LtsvParser {
	if ltsvParser == nil {
		return nil
	}
	columns := make([]LtsvParserColumn, 0, len(ltsvParser.Columns))
	for _, input := range ltsvParser.Columns {
		column := LtsvParserColumn{
			Name:        types.StringValue(input.Name),
			Type:        types.StringValue(input.Type),
			Format:      types.StringPointerValue(input.Format),
			ColumnOrder: types.Int64Value(input.ColumnOrder),
		}
		columns = append(columns, column)
	}
	return &LtsvParser{
		Newline: types.StringPointerValue(ltsvParser.Newline),
		Charset: types.StringPointerValue(ltsvParser.Charset),
		Columns: columns,
	}
}

func (ltsvParser *LtsvParser) ToLtsvParserInput() *job_definitions2.LtsvParserInput {
	if ltsvParser == nil {
		return nil
	}
	columns := make([]job_definitions2.LtsvParserColumnInput, 0, len(ltsvParser.Columns))
	for _, input := range ltsvParser.Columns {
		column := job_definitions2.LtsvParserColumnInput{
			Name:   input.Name.String(),
			Type:   input.Type.String(),
			Format: input.Format.ValueStringPointer(),
		}
		columns = append(columns, column)
	}

	return &job_definitions2.LtsvParserInput{
		Newline: ltsvParser.Newline.ValueStringPointer(),
		Charset: ltsvParser.Charset.ValueStringPointer(),
		Columns: columns,
	}
}

func ToLtsvParserModel(ltsvParser *job_definitions.LtsvParser) *LtsvParser {
	if ltsvParser == nil {
		return nil
	}
	columns := make([]LtsvParserColumn, 0, len(ltsvParser.Columns))
	for _, input := range ltsvParser.Columns {
		column := LtsvParserColumn{
			Name:        types.StringValue(input.Name),
			Type:        types.StringValue(input.Type),
			Format:      types.StringPointerValue(input.Format),
			ColumnOrder: types.Int64Value(input.ColumnOrder),
		}
		columns = append(columns, column)
	}
	return &LtsvParser{
		Newline: types.StringPointerValue(ltsvParser.Newline),
		Charset: types.StringPointerValue(ltsvParser.Charset),
		Columns: columns,
	}
}