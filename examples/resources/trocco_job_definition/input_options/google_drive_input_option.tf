resource "trocco_job_definition" "google_drive_input_example" {
  input_option_type = "google_drive"
  input_option = {
    google_drive_input_option = {
      google_drive_connection_id = 1 # please set your google drive connection id
      folder_id                  = "1BxiMVs0XRA5nFMdKvBdBZjgmUUqptlbs"
      file_match_pattern         = ""
      is_skip_header_line        = false
      stop_when_file_not_found   = false
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
        quote                   = "\""
        escape                  = "\""
        null_string             = ""
        comment_line_marker     = ""
        charset                 = "UTF-8"
        columns = [
          {
            name = "id"
            type = "long"
          },
          {
            name = "name"
            type = "string"
          },
          {
            name   = "created_at"
            type   = "timestamp"
            format = "%Y-%m-%d %H:%M:%S.%N %z"
          },
        ]
      }
    }
  }
}
