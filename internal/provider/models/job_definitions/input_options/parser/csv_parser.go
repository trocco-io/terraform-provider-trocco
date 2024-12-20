package parser

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"terraform-provider-trocco/internal/client/entities/job_definitions"
)

type CsvParser struct {
	Delimiter            types.String      `tfsdk:"delimiter"`
	Quote                types.String      `tfsdk:"quote"`
	Escape               types.String      `tfsdk:"escape"`
	SkipHeaderLines      types.Int64       `tfsdk:"skip_header_lines"`
	NullStringEnabled    types.Bool        `tfsdk:"null_string_enabled"`
	NullString           types.String      `tfsdk:"null_string"`
	TrimIfNotQuoted      types.Bool        `tfsdk:"trim_if_not_quoted"`
	QuotesInQuotedFields types.String      `tfsdk:"quotes_in_quoted_fields"`
	CommentLineMarker    types.String      `tfsdk:"comment_line_marker"`
	AllowOptionalColumns types.Bool        `tfsdk:"allow_optional_columns"`
	AllowExtraColumns    types.Bool        `tfsdk:"allow_extra_columns"`
	MaxQuotedSizeLimit   types.Int64       `tfsdk:"max_quoted_size_limit"`
	StopOnInvalidRecord  types.Bool        `tfsdk:"stop_on_invalid_record"`
	DefaultTimeZone      types.String      `tfsdk:"default_time_zone"`
	DefaultDate          types.String      `tfsdk:"default_date"`
	Newline              types.String      `tfsdk:"newline"`
	Charset              types.String      `tfsdk:"charset"`
	Columns              []CsvParserColumn `tfsdk:"columns"`
}

type CsvParserColumn struct {
	Name        types.String `tfsdk:"name"`
	Type        types.String `tfsdk:"type"`
	Format      types.String `tfsdk:"format"`
	Date        types.String `tfsdk:"date"`
	ColumnOrder types.Int64  `tfsdk:"column_order"`
}

func NewCsvParser(csvParser *job_definitions.CsvParser) *CsvParser {
	if csvParser == nil {
		return nil
	}
	columns := make([]CsvParserColumn, 0, len(csvParser.Columns))
	for _, input := range csvParser.Columns {
		column := CsvParserColumn{
			Name:        types.StringValue(input.Name),
			Type:        types.StringValue(input.Type),
			Format:      types.StringPointerValue(input.Format),
			Date:        types.StringPointerValue(input.Date),
			ColumnOrder: types.Int64Value(input.ColumnOrder),
		}
		columns = append(columns, column)
	}
	return &CsvParser{
		Delimiter:            types.StringValue(csvParser.Delimiter),
		Quote:                types.StringPointerValue(csvParser.Quote),
		Escape:               types.StringPointerValue(csvParser.Escape),
		SkipHeaderLines:      types.Int64Value(csvParser.SkipHeaderLines),
		NullStringEnabled:    types.BoolPointerValue(csvParser.NullStringEnabled),
		NullString:           types.StringPointerValue(csvParser.NullString),
		TrimIfNotQuoted:      types.BoolValue(csvParser.TrimIfNotQuoted),
		QuotesInQuotedFields: types.StringValue(csvParser.QuotesInQuotedFields),
		CommentLineMarker:    types.StringPointerValue(csvParser.CommentLineMarker),
		AllowOptionalColumns: types.BoolValue(csvParser.AllowOptionalColumns),
		AllowExtraColumns:    types.BoolValue(csvParser.AllowExtraColumns),
		MaxQuotedSizeLimit:   types.Int64Value(csvParser.MaxQuotedSizeLimit),
		StopOnInvalidRecord:  types.BoolValue(csvParser.StopOnInvalidRecord),
		DefaultTimeZone:      types.StringValue(csvParser.DefaultTimeZone),
		DefaultDate:          types.StringValue(csvParser.DefaultDate),
		Newline:              types.StringValue(csvParser.Newline),
		Charset:              types.StringPointerValue(csvParser.Charset),
		Columns:              columns,
	}
}
