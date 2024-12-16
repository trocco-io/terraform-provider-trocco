package job_definitions

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

type CsvParserColumnInput struct {
	Name   string  `json:"name"`
	Type   string  `json:"type"`
	Format *string `json:"format,omitempty"`
	Date   *string `json:"date,omitempty"`
}

type JsonlParserInput struct {
	StopOnInvalidRecord *bool                    `json:"stop_on_invalid_record,omitempty"`
	DefaultTimeZone     string                   `json:"default_time_zone"`
	Newline             *string                  `json:"newline,omitempty"`
	Charset             *string                  `json:"charset,omitempty"`
	Columns             []JsonlParserColumnInput `json:"columns"`
}

type JsonlParserColumnInput struct {
	Name     string  `json:"name"`
	Type     string  `json:"type"`
	TimeZone *string `json:"time_zone,omitempty"`
	Format   *string `json:"format,omitempty"`
}

type JsonpathParserInput struct {
	Root            string                      `json:"root"`
	DefaultTimeZone string                      `json:"default_time_zone"`
	Columns         []JsonpathParserColumnInput `json:"columns"`
}

type JsonpathParserColumnInput struct {
	Name     string  `json:"name"`
	Type     string  `json:"type"`
	TimeZone *string `json:"time_zone,omitempty"`
	Format   *string `json:"format,omitempty"`
}

type LtsvParserInput struct {
	Newline *string                 `json:"newline,omitempty"`
	Charset *string                 `json:"charset,omitempty"`
	Columns []LtsvParserColumnInput `json:"columns"`
}

type LtsvParserColumnInput struct {
	Name   string  `json:"name"`
	Type   string  `json:"type"`
	Format *string `json:"format,omitempty"`
}

type ExcelParserInput struct {
	DefaultTimeZone string                   `json:"default_time_zone"`
	SheetName       string                   `json:"sheet_name"`
	SkipHeaderLines int64                    `json:"skip_header_lines"`
	Columns         []ExcelParserColumnInput `json:"columns"`
}

type ExcelParserColumnInput struct {
	Name            string  `json:"name"`
	Type            string  `json:"type"`
	Format          *string `json:"format,omitempty"`
	FormulaHandling *string `json:"formula_handling,omitempty"`
}

type XmlParserInput struct {
	Root    string                 `json:"root"`
	Columns []XmlParserColumnInput `json:"columns"`
}

type XmlParserColumnInput struct {
	Name   string  `json:"name"`
	Type   string  `json:"type"`
	Path   string  `json:"path"`
	Format *string `json:"format,omitempty"`
}

type ParquetParserInput struct {
	Columns []ParquetParserColumnInput `json:"columns"`
}

type ParquetParserColumnInput struct {
	Name   string  `json:"name"`
	Type   string  `json:"type"`
	Format *string `json:"format,omitempty"`
}
