resource "trocco_job_definition" "gcs_output_example_csv" {
  output_option_type = "gcs"
  output_option = {
    gcs_output_option = {
      gcs_connection_id       = 1 # require your gcs connection id
      bucket                  = "my-bucket"
      path_prefix             = "output/data/"
      file_ext                = ".csv"
      sequence_format         = ".%03d.%02d"
      is_minimum_output_tasks = false
      formatter_type          = "csv"
      encoder_type            = "gzip"
      csv_formatter = {
        delimiter           = ","
        escape              = "\\"
        header_line         = true
        charset             = "UTF-8"
        quote_policy        = "ALL"
        newline             = "LF"
        newline_in_field    = "LF"
        null_string_enabled = true
        null_string         = "NULL"
        default_time_zone   = "Asia/Tokyo"
        csv_formatter_column_options_attributes = [
          {
            name     = "created_at"
            format   = "%Y-%m-%d %H:%M:%S"
            timezone = "Asia/Tokyo"
          }
        ]
      }
    }
  }
}

resource "trocco_job_definition" "gcs_output_example_jsonl" {
  output_option_type = "gcs"
  output_option = {
    gcs_output_option = {
      gcs_connection_id       = 1 # require your gcs connection id
      bucket                  = "my-bucket"
      path_prefix             = "output/data/"
      file_ext                = ".jsonl"
      sequence_format         = ".%03d.%02d"
      is_minimum_output_tasks = false
      formatter_type          = "jsonl"
      encoder_type            = "gzip"
      jsonl_formatter = {
        encoding    = "UTF-8"
        newline     = "LF"
        date_format = "%Y-%m-%d"
        timezone    = "UTC"
      }
    }
  }
}
