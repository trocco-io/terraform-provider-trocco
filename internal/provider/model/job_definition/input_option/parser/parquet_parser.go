package parser

import (
	"context"
	job_definitions "terraform-provider-trocco/internal/client/entity/job_definition"
	params "terraform-provider-trocco/internal/client/parameter/job_definition"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ParquetParser struct {
	Columns types.List `tfsdk:"columns"`
}

type ParquetParserColumn struct {
	Name   types.String `tfsdk:"name"`
	Type   types.String `tfsdk:"type"`
	Format types.String `tfsdk:"format"`
}

func NewParquetParser(ctx context.Context, parquetParser *job_definitions.ParquetParser) *ParquetParser {
	if parquetParser == nil {
		return nil
	}

	columnElements := make([]ParquetParserColumn, 0, len(parquetParser.Columns))
	for _, input := range parquetParser.Columns {
		column := ParquetParserColumn{
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
	return &ParquetParser{
		Columns: columns,
	}
}

func (parquetParser *ParquetParser) ToParquetParserInput(ctx context.Context) *params.ParquetParserInput {
	if parquetParser == nil {
		return nil
	}

	var columnElements []ParquetParserColumn
	diags := parquetParser.Columns.ElementsAs(ctx, &columnElements, false)
	if diags.HasError() {
		return nil
	}

	columns := make([]params.ParquetParserColumnInput, 0, len(columnElements))
	for _, input := range columnElements {
		column := params.ParquetParserColumnInput{
			Name:   input.Name.ValueString(),
			Type:   input.Type.ValueString(),
			Format: input.Format.ValueStringPointer(),
		}
		columns = append(columns, column)
	}

	return &params.ParquetParserInput{Columns: columns}
}
