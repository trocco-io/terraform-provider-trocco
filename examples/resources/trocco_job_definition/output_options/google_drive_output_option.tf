resource "trocco_job_definition" "google_drive_output_example" {
  output_option_type = "google_drive"
  output_option = {
    google_drive_output_option = {
      google_drive_connection_id = 1 # please set your google drive connection id
      main_folder_id             = "1BxiMVs0XRA5nFMdKvBdBZjgmUUqptlbs"
      file_name                  = "output.csv"
      formatter_type             = "csv"

      csv_formatter = {
        delimiter           = ","
        newline             = "CRLF"
        newline_in_field    = "LF"
        charset             = "UTF-8"
        quote_policy        = "MINIMAL"
        escape              = "\\"
        header_line         = true
        null_string_enabled = false
        default_time_zone   = "UTC"

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
