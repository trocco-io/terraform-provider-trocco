package parser

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"terraform-provider-trocco/internal/client/entities/job_definitions"
	job_definitions2 "terraform-provider-trocco/internal/client/parameters/job_definitions"
)

type ParquetParser struct {
	Columns []ParquetParserColumn `tfsdk:"columns"`
}

type ParquetParserColumn struct {
	Name        types.String `tfsdk:"name"`
	Type        types.String `tfsdk:"type"`
	Format      types.String `tfsdk:"format"`
	ColumnOrder types.Int64  `tfsdk:"column_order"`
}

func NewParquetParser(parquetParser *job_definitions.ParquetParser) *ParquetParser {
	if parquetParser == nil {
		return nil
	}
	columns := make([]ParquetParserColumn, 0, len(parquetParser.Columns))
	for _, input := range parquetParser.Columns {
		column := ParquetParserColumn{
			Name:        types.StringValue(input.Name),
			Type:        types.StringValue(input.Type),
			Format:      types.StringPointerValue(input.Format),
			ColumnOrder: types.Int64Value(input.ColumnOrder),
		}
		columns = append(columns, column)
	}
	return &ParquetParser{
		Columns: columns,
	}
}

func (parquetParser *ParquetParser) ToParquetParserInput() *job_definitions2.ParquetParserInput {
	if parquetParser == nil {
		return nil
	}
	columns := make([]job_definitions2.ParquetParserColumnInput, 0, len(parquetParser.Columns))
	for _, input := range parquetParser.Columns {
		column := job_definitions2.ParquetParserColumnInput{
			Name:   input.Name.String(),
			Type:   input.Type.String(),
			Format: input.Format.ValueStringPointer(),
		}
		columns = append(columns, column)
	}

	return &job_definitions2.ParquetParserInput{Columns: columns}
}

func ToParquetParserModel(parquetParser *job_definitions.ParquetParser) *ParquetParser {
	if parquetParser == nil {
		return nil
	}
	columns := make([]ParquetParserColumn, 0, len(parquetParser.Columns))
	for _, input := range parquetParser.Columns {
		column := ParquetParserColumn{
			Name:        types.StringValue(input.Name),
			Type:        types.StringValue(input.Type),
			Format:      types.StringPointerValue(input.Format),
			ColumnOrder: types.Int64Value(input.ColumnOrder),
		}
		columns = append(columns, column)
	}

	return &ParquetParser{Columns: columns}
}
