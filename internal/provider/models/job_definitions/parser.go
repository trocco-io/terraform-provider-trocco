package job_definitions

import "github.com/hashicorp/terraform-plugin-framework/types"

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

type JsonpathParser struct {
	Root            types.String           `tfsdk:"root"`
	DefaultTimeZone types.String           `tfsdk:"default_time_zone"`
	Columns         []JsonpathParserColumn `tfsdk:"columns"`
}

type JsonpathParserColumn struct {
	Name        types.String `tfsdk:"name"`
	Type        types.String `tfsdk:"type"`
	TimeZone    types.String `tfsdk:"time_zone"`
	Format      types.String `tfsdk:"format"`
	ColumnOrder types.Int64  `tfsdk:"column_order"`
}

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

type ExcelParser struct {
	DefaultTimeZone types.String        `tfsdk:"default_time_zone"`
	SheetName       types.String        `tfsdk:"sheet_name"`
	SkipHeaderLines types.Int64         `tfsdk:"skip_header_lines"`
	Columns         []ExcelParserColumn `tfsdk:"columns"`
}

type ExcelParserColumn struct {
	Name            types.String `tfsdk:"name"`
	Type            types.String `tfsdk:"type"`
	Format          types.String `tfsdk:"format"`
	FormulaHandling types.String `tfsdk:"formula_handling"`
	ColumnOrder     types.Int64  `tfsdk:"column_order"`
}

type XmlParser struct {
	Root    types.String      `tfsdk:"root"`
	Columns []XmlParserColumn `tfsdk:"columns"`
}

type XmlParserColumn struct {
	Name        types.String `tfsdk:"name"`
	Type        types.String `tfsdk:"type"`
	Path        types.String `tfsdk:"path"`
	Format      types.String `tfsdk:"format"`
	ColumnOrder types.Int64  `tfsdk:"column_order"`
}

type ParquetParser struct {
	Columns []ParquetParserColumn `tfsdk:"columns"`
}

type ParquetParserColumn struct {
	Name        types.String `tfsdk:"name"`
	Type        types.String `tfsdk:"type"`
	Format      types.String `tfsdk:"format"`
	ColumnOrder types.Int64  `tfsdk:"column_order"`
}
