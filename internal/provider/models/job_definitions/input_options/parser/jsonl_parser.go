package parser

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"terraform-provider-trocco/internal/client/entities/job_definitions"
	job_definitions2 "terraform-provider-trocco/internal/client/parameters/job_definitions"
)

type JsonlParser struct {
	StopOnInvalidRecord types.Bool          `tfsdk:"stop_on_invalid_record"`
	DefaultTimeZone     types.String        `tfsdk:"default_time_zone"`
	Newline             types.String        `tfsdk:"newline"`
	Charset             types.String        `tfsdk:"charset"`
	Columns             []JsonlParserColumn `tfsdk:"columns"`
}

type JsonlParserColumn struct {
	Name        types.String `tfsdk:"name"`
	Type        types.String `tfsdk:"type"`
	TimeZone    types.String `tfsdk:"time_zone"`
	Format      types.String `tfsdk:"format"`
	ColumnOrder types.Int64  `tfsdk:"column_order"`
}

func NewJsonlParser(jsonlParser *job_definitions.JsonlParser) *JsonlParser {
	if jsonlParser == nil {
		return nil
	}
	columns := make([]JsonlParserColumn, 0, len(jsonlParser.Columns))
	for _, input := range jsonlParser.Columns {
		column := JsonlParserColumn{
			Name:        types.StringValue(input.Name),
			Type:        types.StringValue(input.Type),
			TimeZone:    types.StringPointerValue(input.TimeZone),
			Format:      types.StringPointerValue(input.Format),
			ColumnOrder: types.Int64Value(input.ColumnOrder),
		}
		columns = append(columns, column)
	}
	return &JsonlParser{
		StopOnInvalidRecord: types.BoolPointerValue(jsonlParser.StopOnInvalidRecord),
		DefaultTimeZone:     types.StringValue(jsonlParser.DefaultTimeZone),
		Newline:             types.StringPointerValue(jsonlParser.Newline),
		Charset:             types.StringPointerValue(jsonlParser.Charset),
		Columns:             columns,
	}
}

func (jsonlParser *JsonlParser) ToJsonlParserInput() *job_definitions2.JsonlParserInput {
	if jsonlParser == nil {
		return nil
	}
	columns := make([]job_definitions2.JsonlParserColumnInput, 0, len(jsonlParser.Columns))
	for _, input := range jsonlParser.Columns {
		column := job_definitions2.JsonlParserColumnInput{
			Name:     input.Name.String(),
			Type:     input.Type.String(),
			TimeZone: input.TimeZone.ValueStringPointer(),
			Format:   input.Format.ValueStringPointer(),
		}
		columns = append(columns, column)
	}

	return &job_definitions2.JsonlParserInput{
		StopOnInvalidRecord: jsonlParser.StopOnInvalidRecord.ValueBoolPointer(),
		DefaultTimeZone:     jsonlParser.DefaultTimeZone.String(),
		Newline:             jsonlParser.Newline.ValueStringPointer(),
		Charset:             jsonlParser.Charset.ValueStringPointer(),
		Columns:             columns,
	}
}

func ToJsonlParserModel(jsonlParser *job_definitions.JsonlParser) *JsonlParser {
	if jsonlParser == nil {
		return nil
	}

	columns := make([]JsonlParserColumn, 0, len(jsonlParser.Columns))
	for _, input := range jsonlParser.Columns {
		column := JsonlParserColumn{
			Name:        types.StringValue(input.Name),
			Type:        types.StringValue(input.Type),
			TimeZone:    types.StringPointerValue(input.TimeZone),
			Format:      types.StringPointerValue(input.Format),
			ColumnOrder: types.Int64Value(input.ColumnOrder),
		}
		columns = append(columns, column)
	}
	return &JsonlParser{
		StopOnInvalidRecord: types.BoolPointerValue(jsonlParser.StopOnInvalidRecord),
		DefaultTimeZone:     types.StringValue(jsonlParser.DefaultTimeZone),
		Newline:             types.StringPointerValue(jsonlParser.Newline),
		Charset:             types.StringPointerValue(jsonlParser.Charset),
		Columns:             columns,
	}
}