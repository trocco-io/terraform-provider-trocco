package parser

import (
	"terraform-provider-trocco/internal/client/entities/job_definitions"
	params "terraform-provider-trocco/internal/client/parameter/job_definitions"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ParquetParser struct {
	Columns []ParquetParserColumn `tfsdk:"columns"`
}

type ParquetParserColumn struct {
	Name   types.String `tfsdk:"name"`
	Type   types.String `tfsdk:"type"`
	Format types.String `tfsdk:"format"`
}

func NewParquetParser(parquetParser *job_definitions.ParquetParser) *ParquetParser {
	if parquetParser == nil {
		return nil
	}
	columns := make([]ParquetParserColumn, 0, len(parquetParser.Columns))
	for _, input := range parquetParser.Columns {
		column := ParquetParserColumn{
			Name:   types.StringValue(input.Name),
			Type:   types.StringValue(input.Type),
			Format: types.StringPointerValue(input.Format),
		}
		columns = append(columns, column)
	}
	return &ParquetParser{
		Columns: columns,
	}
}

func (parquetParser *ParquetParser) ToParquetParserInput() *params.ParquetParserInput {
	if parquetParser == nil {
		return nil
	}
	columns := make([]params.ParquetParserColumnInput, 0, len(parquetParser.Columns))
	for _, input := range parquetParser.Columns {
		column := params.ParquetParserColumnInput{
			Name:   input.Name.ValueString(),
			Type:   input.Type.ValueString(),
			Format: input.Format.ValueStringPointer(),
		}
		columns = append(columns, column)
	}

	return &params.ParquetParserInput{Columns: columns}
}
