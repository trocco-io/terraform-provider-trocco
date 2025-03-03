resource "trocco_job_definition" "s3_input_example" {
  input_option_type = "s3"
  input_option = {
    s3_input_option = {
      bucket = "test_bucket"
      csv_parser = {
        allow_extra_columns    = false
        allow_optional_columns = false
        charset                = "UTF-8"
        columns = [
          {
            name = "col1"
            type = "string"
          },
          {
            name = "col2"
            type = "string"
          },
          {
            name = "col3"
            type = "string"
          },
          {
            name = "col4"
            type = "string"
          },
        ]
        comment_line_marker     = ""
        default_date            = "1970-01-01"
        default_time_zone       = "UTC"
        delimiter               = ","
        escape                  = "\""
        max_quoted_size_limit   = 131072
        newline                 = "CRLF"
        null_string             = ""
        null_string_enabled     = false
        quote                   = "\""
        quotes_in_quoted_fields = "ACCEPT_ONLY_RFC4180_ESCAPED"
        skip_header_lines       = 1
        stop_on_invalid_record  = true
        trim_if_not_quoted      = false
      }
      decompression_type          = "default"
      incremental_loading_enabled = false
      is_skip_header_line         = false
      path_match_pattern          = ""
      path_prefix                 = "dev/000.00.csv"
      region                      = "ap-northeast-1"
      s3_connection_id            = 1 # please set your s3 connection id
      stop_when_file_not_found    = false
    }
  }
}
