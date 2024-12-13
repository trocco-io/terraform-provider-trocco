package job_definitions

type CsvParser struct {
	Delimiter            string            `json:"delimiter"`
	Quote                *string           `json:"quote"`
	Escape               *string           `json:"escape"`
	SkipHeaderLines      int64             `json:"skip_header_lines"`
	NullStringEnabled    *bool             `json:"null_string_enabled"`
	NullString           *string           `json:"null_string"`
	TrimIfNotQuoted      bool              `json:"trim_if_not_quoted"`
	QuotesInQuotedFields string            `json:"quotes_in_quoted_fields"`
	CommentLineMarker    *string           `json:"comment_line_marker"`
	AllowOptionalColumns bool              `json:"allow_optional_columns"`
	AllowExtraColumns    bool              `json:"allow_extra_columns"`
	MaxQuotedSizeLimit   int64             `json:"max_quoted_size_limit"`
	StopOnInvalidRecord  bool              `json:"stop_on_invalid_record"`
	DefaultTimeZone      string            `json:"default_time_zone"`
	DefaultDate          string            `json:"default_date"`
	Newline              string            `json:"newline"`
	Charset              *string           `json:"charset"`
	Columns              []CsvParserColumn `json:"columns"`
}

type CsvParserInput struct {
	Delimiter            string                 `json:"delimiter"`
	Quote                *string                `json:"quote,omitempty"`
	Escape               *string                `json:"escape,omitempty"`
	SkipHeaderLines      int64                  `json:"skip_header_lines"`
	NullStringEnabled    *bool                  `json:"null_string_enabled,omitempty"`
	NullString           *string                `json:"null_string,omitempty"`
	TrimIfNotQuoted      bool                   `json:"trim_if_not_quoted"`
	QuotesInQuotedFields string                 `json:"quotes_in_quoted_fields"`
	CommentLineMarker    *string                `json:"comment_line_marker,omitempty"`
	AllowOptionalColumns bool                   `json:"allow_optional_columns"`
	AllowExtraColumns    bool                   `json:"allow_extra_columns"`
	MaxQuotedSizeLimit   int64                  `json:"max_quoted_size_limit"`
	StopOnInvalidRecord  bool                   `json:"stop_on_invalid_record"`
	DefaultTimeZone      string                 `json:"default_time_zone"`
	DefaultDate          string                 `json:"default_date"`
	Newline              string                 `json:"newline"`
	Charset              *string                `json:"charset,omitempty"`
	Columns              []CsvParserColumnInput `json:"columns"`
}

type CsvParserColumn struct {
	Name        string  `json:"name"`
	Type        string  `json:"type"`
	Format      *string `json:"format"`
	Date        *string `json:"date"`
	ColumnOrder int64   `json:"column_order"`
}

type CsvParserColumnInput struct {
	Name   string  `json:"name"`
	Type   string  `json:"type"`
	Format *string `json:"format,omitempty"`
	Date   *string `json:"date,omitempty"`
}

type JsonlParser struct {
	StopOnInvalidRecord *bool               `json:"stop_on_invalid_record"`
	DefaultTimeZone     string              `json:"default_time_zone"`
	Newline             *string             `json:"newline"`
	Charset             *string             `json:"charset"`
	Columns             []JsonlParserColumn `json:"columns"`
}

type JsonlParserInput struct {
	StopOnInvalidRecord *bool                    `json:"stop_on_invalid_record,omitempty"`
	DefaultTimeZone     string                   `json:"default_time_zone"`
	Newline             *string                  `json:"newline,omitempty"`
	Charset             *string                  `json:"charset,omitempty"`
	Columns             []JsonlParserColumnInput `json:"columns"`
}

type JsonlParserColumn struct {
	Name        string  `json:"name"`
	Type        string  `json:"type"`
	TimeZone    *string `json:"time_zone"`
	Format      *string `json:"format"`
	ColumnOrder int64   `json:"column_order"`
}

type JsonlParserColumnInput struct {
	Name     string  `json:"name"`
	Type     string  `json:"type"`
	TimeZone *string `json:"time_zone,omitempty"`
	Format   *string `json:"format,omitempty"`
}

type JsonpathParser struct {
	Root            string                 `json:"root"`
	DefaultTimeZone string                 `json:"default_time_zone"`
	Columns         []JsonpathParserColumn `json:"columns"`
}

type JsonpathParserInput struct {
	Root            string                      `json:"root"`
	DefaultTimeZone string                      `json:"default_time_zone"`
	Columns         []JsonpathParserColumnInput `json:"columns"`
}

type JsonpathParserColumn struct {
	Name        string  `json:"name"`
	Type        string  `json:"type"`
	TimeZone    *string `json:"time_zone"`
	Format      *string `json:"format"`
	ColumnOrder int64   `json:"column_order"`
}

type JsonpathParserColumnInput struct {
	Name     string  `json:"name"`
	Type     string  `json:"type"`
	TimeZone *string `json:"time_zone,omitempty"`
	Format   *string `json:"format,omitempty"`
}

type LtsvParser struct {
	Newline *string            `json:"newline"`
	Charset *string            `json:"charset"`
	Columns []LtsvParserColumn `json:"columns"`
}

type LtsvParserInput struct {
	Newline *string                 `json:"newline,omitempty"`
	Charset *string                 `json:"charset,omitempty"`
	Columns []LtsvParserColumnInput `json:"columns"`
}

type LtsvParserColumn struct {
	Name        string  `json:"name"`
	Type        string  `json:"type"`
	Format      *string `json:"format"`
	ColumnOrder int64   `json:"column_order"`
}

type LtsvParserColumnInput struct {
	Name   string  `json:"name"`
	Type   string  `json:"type"`
	Format *string `json:"format,omitempty"`
}

type ExcelParser struct {
	DefaultTimeZone string              `json:"default_time_zone"`
	SheetName       string              `json:"sheet_name"`
	SkipHeaderLines int64               `json:"skip_header_lines"`
	Columns         []ExcelParserColumn `json:"columns"`
}

type ExcelParserInput struct {
	DefaultTimeZone string                   `json:"default_time_zone"`
	SheetName       string                   `json:"sheet_name"`
	SkipHeaderLines int64                    `json:"skip_header_lines"`
	Columns         []ExcelParserColumnInput `json:"columns"`
}

type ExcelParserColumn struct {
	Name            string  `json:"name"`
	Type            string  `json:"type"`
	Format          *string `json:"format"`
	FormulaHandling *string `json:"formula_handling"`
	ColumnOrder     int64   `json:"column_order"`
}

type ExcelParserColumnInput struct {
	Name            string  `json:"name"`
	Type            string  `json:"type"`
	Format          *string `json:"format,omitempty"`
	FormulaHandling *string `json:"formula_handling,omitempty"`
}

type XmlParser struct {
	Root    string            `json:"root"`
	Columns []XmlParserColumn `json:"columns"`
}

type XmlParserInput struct {
	Root    string                 `json:"root"`
	Columns []XmlParserColumnInput `json:"columns"`
}

type XmlParserColumn struct {
	Name        string  `json:"name"`
	Type        string  `json:"type"`
	Path        string  `json:"path"`
	Format      *string `json:"format"`
	ColumnOrder int64   `json:"column_order"`
}

type XmlParserColumnInput struct {
	Name   string  `json:"name"`
	Type   string  `json:"type"`
	Path   string  `json:"path"`
	Format *string `json:"format,omitempty"`
}

type ParquetParser struct {
	Columns []ParquetParserColumn `json:"columns"`
}

type ParquetParserInput struct {
	Columns []ParquetParserColumnInput `json:"columns"`
}

type ParquetParserColumn struct {
	Name        string  `json:"name"`
	Type        string  `json:"type"`
	Format      *string `json:"format"`
	ColumnOrder int64   `json:"column_order"`
}

type ParquetParserColumnInput struct {
	Name   string  `json:"name"`
	Type   string  `json:"type"`
	Format *string `json:"format,omitempty"`
}
