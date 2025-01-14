resource "trocco_job_definition" "gcs_input_example" {
  input_option_type             = "gcs"
  input_option                  = {
      gcs_input_option = {
        bucket                      = "test-bucket"
        path_prefix                 = "path/to/your_file.csv"
        gcs_connection_id           = 1 # require your gcs connection id
        incremental_loading_enabled = false
        stop_when_file_not_found    = true
        csv_parser = {
          delimiter               = ","
          skip_header_lines       = 1
          trim_if_not_quoted      = false
          quotes_in_quoted_fields = "ACCEPT_ONLY_RFC4180_ESCAPED"
          allow_extra_columns     = false
          allow_optional_columns  = false
          stop_on_invalid_record  = true
          default_date            = "1970-01-01"
          default_time_zone       = "UTC"
          newline                 = "CRLF"
          max_quoted_size_limit   = 131072
          null_string_enabled     = false
          quote  = "\""
          escape = "\""
          null_string         = ""
          comment_line_marker = ""
          charset             = "UTF-8"
          columns = [
            {
              name   = "id"
              type   = "long"
            },
            {
              name = "num_col"
              type = "long"
            },
            {
              name = "str_col"
              type = "string"
            },
            {
              name   = "date_col"
              type   = "timestamp"
              format = "%Y-%m-%d %H:%M:%S.%N %z"
            },
          ]
        }
      }
    }
}