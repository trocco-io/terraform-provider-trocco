package input_options

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"terraform-provider-trocco/internal/provider/models"
	"terraform-provider-trocco/internal/provider/models/job_definitions"
)

type GcsInputOption struct {
	GcsConnectionID           types.Int64                    `tfsdk:"gcs_connection_id"`
	Bucket                    types.String                   `tfsdk:"bucket"`
	PathPrefix                types.String                   `tfsdk:"path_prefix"`
	IncrementalLoadingEnabled types.Bool                     `tfsdk:"incremental_loading_enabled"`
	LastPath                  types.String                   `tfsdk:"last_path"`
	StopWhenFileNotFound      types.Bool                     `tfsdk:"stop_when_file_not_found"`
	DecompressionType         types.String                   `tfsdk:"decompression_type"`
	CsvParsers                job_definitions.CsvParser      `tfsdk:"csv_parsers"`
	JsonlParsers              job_definitions.JsonlParser    `tfsdk:"jsonl_parsers"`
	JsonpathParsers           job_definitions.JsonpathParser `tfsdk:"jsonpath_parsers"`
	LtsvParsers               job_definitions.LtsvParser     `tfsdk:"ltsv_parsers"`
	ExcelParsers              job_definitions.ExcelParser    `tfsdk:"excel_parsers"`
	XmlParsers                job_definitions.XmlParser      `tfsdk:"xml_parsers"`
	ParquetParsers            job_definitions.ParquetParser  `tfsdk:"parquet_parsers"`
	CustomVariableSettings    []models.CustomVariableSetting `tfsdk:"custom_variable_settings"`
	Decoder                   job_definitions.Decoder        `tfsdk:"decoder"`
}
