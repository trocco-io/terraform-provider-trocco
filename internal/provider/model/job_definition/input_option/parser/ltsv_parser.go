package parser

import (
	"context"
	job_definitions "terraform-provider-trocco/internal/client/entity/job_definition"
	params "terraform-provider-trocco/internal/client/parameter/job_definition"
	"terraform-provider-trocco/internal/provider/model"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type LtsvParser struct {
	Newline types.String `tfsdk:"newline"`
	Charset types.String `tfsdk:"charset"`
	Columns types.List   `tfsdk:"columns"`
}

type LtsvParserColumn struct {
	Name   types.String `tfsdk:"name"`
	Type   types.String `tfsdk:"type"`
	Format types.String `tfsdk:"format"`
}

func NewLtsvParser(ctx context.Context, ltsvParser *job_definitions.LtsvParser) *LtsvParser {
	if ltsvParser == nil {
		return nil
	}

	columnElements := make([]LtsvParserColumn, 0, len(ltsvParser.Columns))
	for _, input := range ltsvParser.Columns {
		column := LtsvParserColumn{
			Name:   types.StringValue(input.Name),
			Type:   types.StringValue(input.Type),
			Format: types.StringPointerValue(input.Format),
		}
		columnElements = append(columnElements, column)
	}

	columns, diags := types.ListValueFrom(
		ctx,
		types.ObjectType{
			AttrTypes: map[string]attr.Type{
				"name":   types.StringType,
				"type":   types.StringType,
				"format": types.StringType,
			},
		},
		columnElements,
	)
	if diags.HasError() {
		return nil
	}

	return &LtsvParser{
		Newline: types.StringPointerValue(ltsvParser.Newline),
		Charset: types.StringPointerValue(ltsvParser.Charset),
		Columns: columns,
	}
}

func (ltsvParser *LtsvParser) ToLtsvParserInput(ctx context.Context) *params.LtsvParserInput {
	if ltsvParser == nil {
		return nil
	}

	var columnElements []LtsvParserColumn
	diags := ltsvParser.Columns.ElementsAs(ctx, &columnElements, false)
	if diags.HasError() {
		return nil
	}

	columns := make([]params.LtsvParserColumnInput, 0, len(columnElements))
	for _, input := range columnElements {
		column := params.LtsvParserColumnInput{
			Name:   input.Name.ValueString(),
			Type:   input.Type.ValueString(),
			Format: input.Format.ValueStringPointer(),
		}
		columns = append(columns, column)
	}

	return &params.LtsvParserInput{
		Newline: model.NewNullableString(ltsvParser.Newline),
		Charset: model.NewNullableString(ltsvParser.Charset),
		Columns: columns,
	}
}
