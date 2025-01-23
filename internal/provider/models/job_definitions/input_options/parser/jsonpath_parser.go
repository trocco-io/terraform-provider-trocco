package parser

import (
	"terraform-provider-trocco/internal/client/entities/job_definitions"
	params "terraform-provider-trocco/internal/client/parameter/job_definitions"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

type JsonpathParser struct {
	Root            types.String           `tfsdk:"root"`
	DefaultTimeZone types.String           `tfsdk:"default_time_zone"`
	Columns         []JsonpathParserColumn `tfsdk:"columns"`
}

type JsonpathParserColumn struct {
	Name     types.String `tfsdk:"name"`
	Type     types.String `tfsdk:"type"`
	TimeZone types.String `tfsdk:"time_zone"`
	Format   types.String `tfsdk:"format"`
}

func NewJsonPathParser(jsonpathParser *job_definitions.JsonpathParser) *JsonpathParser {
	if jsonpathParser == nil {
		return nil
	}
	columns := make([]JsonpathParserColumn, 0, len(jsonpathParser.Columns))
	for _, input := range jsonpathParser.Columns {
		column := JsonpathParserColumn{
			Name:     types.StringValue(input.Name),
			Type:     types.StringValue(input.Type),
			TimeZone: types.StringPointerValue(input.TimeZone),
			Format:   types.StringPointerValue(input.Format),
		}
		columns = append(columns, column)
	}
	return &JsonpathParser{
		Root:            types.StringValue(jsonpathParser.Root),
		DefaultTimeZone: types.StringValue(jsonpathParser.DefaultTimeZone),
		Columns:         columns,
	}
}

func (jsonpathParser *JsonpathParser) ToJsonpathParserInput() *params.JsonpathParserInput {
	if jsonpathParser == nil {
		return nil
	}
	columns := make([]params.JsonpathParserColumnInput, 0, len(jsonpathParser.Columns))
	for _, input := range jsonpathParser.Columns {
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
