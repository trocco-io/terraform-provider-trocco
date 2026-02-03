resource "trocco_job_definition" "s3_output_example_csv" {
  output_option_type = "s3"
  output_option = {
    s3_output_option = {
      s3_connection_id         = 1 # require your s3 connection id
      bucket                   = "example-bucket"
      path_prefix              = "output/data"
      region                   = "ap-northeast-1"
      file_ext                 = "csv.gz"
      sequence_format          = ".%03d"
      canned_acl               = "Private"
      is_minimum_output_tasks  = false
      multipart_upload_enabled = true
      formatter_type           = "csv"
      encoder_type             = "gzip"
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

resource "trocco_job_definition" "s3_output_example_jsonl" {
  output_option_type = "s3"
  output_option = {
    s3_output_option = {
      s3_connection_id         = 1 # require your s3 connection id
      bucket                   = "example-bucket"
      path_prefix              = "output/json"
      region                   = "us-west-2"
      file_ext                 = "jsonl.gz"
      sequence_format          = ".%03d"
      is_minimum_output_tasks  = false
      multipart_upload_enabled = true
      formatter_type           = "jsonl"
      encoder_type             = "gzip"
      jsonl_formatter = {
        encoding    = "utf_8"
        newline     = "LF"
        date_format = "%Y-%m-%d"
        timezone    = "UTC"
      }
    }
  }
}
