package input_options

import "terraform-provider-trocco/internal/client/common"

type GcsInputOption struct {
	GcsConnectionID           int64                           `json:"gcs_connection_id"`
	Bucket                    string                          `json:"bucket"`
	PathPrefix                *string                         `json:"path_prefix"`
	IncrementalLoadingEnabled bool                            `json:"incremental_loading_enabled"`
	LastPath                  *string                         `json:"last_path"`
	StopWhenFileNotFound      bool                            `json:"stop_when_file_not_found"`
	DecompressionType         *string                         `json:"decompression_type"`
	CsvParsers                *common.CsvParser               `json:"csv_parsers"`
	JsonlParsers              *common.JsonlParser             `json:"jsonl_parsers"`
	JsonpathParsers           *common.JsonpathParser          `json:"jsonpath_parsers"`
	LtsvParsers               *common.LtsvParser              `json:"ltsv_parsers"`
	ExcelParsers              *common.ExcelParser             `json:"excel_parsers"`
	XmlParsers                *common.XmlParser               `json:"xml_parsers"`
	ParquetParsers            *common.ParquetParser           `json:"parquet_parsers"`
	CustomVariableSettings    *[]common.CustomVariableSetting `json:"custom_variable_settings"`
	Decoder                   *common.Decoder                 `json:"decoder"`
}

type GcsInputOptionInput struct {
	GcsConnectionID           int64                           `json:"gcs_connection_id"`
	Bucket                    string                          `json:"bucket"`
	PathPrefix                *string                         `json:"path_prefix,omitempty"`
	IncrementalLoadingEnabled bool                            `json:"incremental_loading_enabled"`
	LastPath                  *string                         `json:"last_path,omitempty"`
	StopWhenFileNotFound      bool                            `json:"stop_when_file_not_found"`
	DecompressionType         *string                         `json:"decompression_type,omitempty"`
	CsvParsers                *common.CsvParserInput          `json:"csv_parsers,omitempty"`
	JsonlParsers              *common.JsonlParserInput        `json:"jsonl_parsers,omitempty"`
	JsonpathParsers           *common.JsonpathParserInput     `json:"jsonpath_parsers,omitempty"`
	LtsvParsers               *common.LtsvParserInput         `json:"ltsv_parsers,omitempty"`
	ExcelParsers              *common.ExcelParserInput        `json:"excel_parsers,omitempty"`
	XmlParsers                *common.XmlParserInput          `json:"xml_parsers,omitempty"`
	ParquetParsers            *common.ParquetParserInput      `json:"parquet_parsers,omitempty"`
	CustomVariableSettings    *[]common.CustomVariableSetting `json:"custom_variable_settings,omitempty"`
	Decoder                   *common.Decoder                 `json:"decoder,omitempty"`
}
