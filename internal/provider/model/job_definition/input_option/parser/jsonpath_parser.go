package parser

import (
	"context"
	job_definitions "terraform-provider-trocco/internal/client/entity/job_definition"
	params "terraform-provider-trocco/internal/client/parameter/job_definition"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type JsonpathParser struct {
	Root            types.String `tfsdk:"root"`
	DefaultTimeZone types.String `tfsdk:"default_time_zone"`
	Columns         types.List   `tfsdk:"columns"`
}

type JsonpathParserColumn struct {
	Name     types.String `tfsdk:"name"`
	Type     types.String `tfsdk:"type"`
	TimeZone types.String `tfsdk:"time_zone"`
	Format   types.String `tfsdk:"format"`
}

func NewJsonPathParser(ctx context.Context, jsonpathParser *job_definitions.JsonpathParser) *JsonpathParser {
	if jsonpathParser == nil {
		return nil
	}

	columnElements := make([]JsonpathParserColumn, 0, len(jsonpathParser.Columns))
	for _, input := range jsonpathParser.Columns {
		column := JsonpathParserColumn{
			Name:     types.StringValue(input.Name),
			Type:     types.StringValue(input.Type),
			TimeZone: types.StringPointerValue(input.TimeZone),
			Format:   types.StringPointerValue(input.Format),
		}
		columnElements = append(columnElements, column)
	}

	columns, diags := types.ListValueFrom(
		ctx,
		types.ObjectType{
			AttrTypes: map[string]attr.Type{
				"name":      types.StringType,
				"type":      types.StringType,
				"time_zone": types.StringType,
				"format":    types.StringType,
			},
		},
		columnElements,
	)
	if diags.HasError() {
		return nil
	}

	return &JsonpathParser{
		Root:            types.StringValue(jsonpathParser.Root),
		DefaultTimeZone: types.StringValue(jsonpathParser.DefaultTimeZone),
		Columns:         columns,
	}
}

func (jsonpathParser *JsonpathParser) ToJsonpathParserInput(ctx context.Context) *params.JsonpathParserInput {
	if jsonpathParser == nil {
		return nil
	}

	var columnElements []JsonpathParserColumn
	diags := jsonpathParser.Columns.ElementsAs(ctx, &columnElements, false)
	if diags.HasError() {
		return nil
	}

	columns := make([]params.JsonpathParserColumnInput, 0, len(columnElements))
	for _, input := range columnElements {
		column := params.JsonpathParserColumnInput{
			Name:     input.Name.ValueString(),
			Type:     input.Type.ValueString(),
			TimeZone: input.TimeZone.ValueStringPointer(),
			Format:   input.Format.ValueStringPointer(),
		}
		columns = append(columns, column)
	}

	return &params.JsonpathParserInput{
		Root:            jsonpathParser.Root.ValueString(),
		DefaultTimeZone: jsonpathParser.DefaultTimeZone.ValueString(),
		Columns:         columns,
	}
}
