resource "trocco_job_definition" "csv_parser_example" {

  input_option = {
    # The example is gcs, but it can be applied to file-based input.
    gcs_input_option = {
      csv_parser = {
        delimiter               = ","
        quote                   = "\""
        escape                  = "\""
        skip_header_lines       = 1
        null_string_enabled     = false
        null_string             = ""
        trim_if_not_quoted      = false
        quotes_in_quoted_fields = "ACCEPT_ONLY_RFC4180_ESCAPED"
        comment_line_marker     = ""
        allow_optional_columns  = false
        allow_extra_columns     = false
        max_quoted_size_limit   = 131072
        stop_on_invalid_record  = true
        default_time_zone       = "UTC"
        default_date            = "1970-01-01"
        newline                 = "CRLF"
        charset                 = "UTF-8"
        columns = [
          {
            name = "id"
            type = "long"
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
          }
        ]
      }
    }
  }
}