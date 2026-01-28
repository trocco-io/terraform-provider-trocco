# SFTP Output Option Example
resource "trocco_job_definition" "sftp_output_example" {
  name = "SFTP Output Option Example"

  output_option {
    sftp_output_option {
      sftp_connection_id      = 1
      path_prefix             = "/data/exports/users_$start_time$.csv"
      file_ext                = ".csv"
      is_minimum_output_tasks = false
      formatter_type          = "csv"
      encoder_type            = "gzip"

      csv_formatter {
        delimiter           = ","
        newline             = "CRLF"
        newline_in_field    = "LF"
        charset             = "UTF-8"
        quote_policy        = "MINIMAL"
        escape              = "\\"
        header_line         = true
        null_string_enabled = false
        null_string         = ""
        default_time_zone   = "UTC"

        csv_formatter_column_options_attributes = [
          {
            name     = "created_at"
            format   = "%Y-%m-%d %H:%M:%S.%6N %z"
            timezone = "UTC"
          },
          {
            name     = "updated_at"
            format   = "%Y-%m-%d %H:%M:%S.%6N %z"
            timezone = "UTC"
          }
        ]
      }
    }
  }
}

# SFTP Output Option with JSONL formatter
resource "trocco_job_definition" "sftp_jsonl_output_example" {
  name = "SFTP JSONL Output Option Example"

  output_option {
    sftp_output_option {
      sftp_connection_id      = 1
      path_prefix             = "/analytics/events/events_$date$.jsonl"
      file_ext                = ".jsonl"
      is_minimum_output_tasks = true
      formatter_type          = "jsonl"
      encoder_type            = ""

      jsonl_formatter {
        encoding    = "UTF-8"
        newline     = "LF"
        date_format = "%Y-%m-%d %H:%M:%S"
        timezone    = "UTC"
      }
    }
  }
}
