package parser

import (
	"terraform-provider-trocco/internal/client/entities/job_definitions"
	params "terraform-provider-trocco/internal/client/parameter/job_definitions"
	"terraform-provider-trocco/internal/provider/models"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

type LtsvParser struct {
	Newline types.String       `tfsdk:"newline"`
	Charset types.String       `tfsdk:"charset"`
	Columns []LtsvParserColumn `tfsdk:"columns"`
}

type LtsvParserColumn struct {
	Name   types.String `tfsdk:"name"`
	Type   types.String `tfsdk:"type"`
	Format types.String `tfsdk:"format"`
}

func NewLtsvParser(ltsvParser *job_definitions.LtsvParser) *LtsvParser {
	if ltsvParser == nil {
		return nil
	}
	columns := make([]LtsvParserColumn, 0, len(ltsvParser.Columns))
	for _, input := range ltsvParser.Columns {
		column := LtsvParserColumn{
			Name:   types.StringValue(input.Name),
			Type:   types.StringValue(input.Type),
			Format: types.StringPointerValue(input.Format),
		}
		columns = append(columns, column)
	}
	return &LtsvParser{
		Newline: types.StringPointerValue(ltsvParser.Newline),
		Charset: types.StringPointerValue(ltsvParser.Charset),
		Columns: columns,
	}
}

func (ltsvParser *LtsvParser) ToLtsvParserInput() *params.LtsvParserInput {
	if ltsvParser == nil {
		return nil
	}
	columns := make([]params.LtsvParserColumnInput, 0, len(ltsvParser.Columns))
	for _, input := range ltsvParser.Columns {
		column := params.LtsvParserColumnInput{
			Name:   input.Name.ValueString(),
			Type:   input.Type.ValueString(),
			Format: input.Format.ValueStringPointer(),
		}
		columns = append(columns, column)
	}

	return &params.LtsvParserInput{
		Newline: models.NewNullableString(ltsvParser.Newline),
		Charset: models.NewNullableString(ltsvParser.Charset),
		Columns: columns,
	}
}
