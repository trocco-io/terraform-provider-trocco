resource "trocco_job_definition" "mysql_to_gcs" {
  name                     = "MySQL to GCS Test"
  description              = "Test job definition for transferring data from MySQL to GCS"
  resource_enhancement     = "medium"
  retry_limit              = 2
  is_runnable_concurrently = false

  filter_columns = [
    {
      default                      = ""
      format                       = ""
      json_expand_enabled          = false
      json_expand_keep_base_column = false
      name                         = "id"
      src                          = "id"
      type                         = "long"
    },
    {
      default                      = ""
      format                       = ""
      json_expand_enabled          = false
      json_expand_keep_base_column = false
      name                         = "type"
      src                          = "type"
      type                         = "string"
    },
  ]

  input_option_type = "mysql"
  input_option = {
    mysql_input_option = {
      mysql_connection_id         = 1
      database                    = "test_database"
      table                       = "test_table"
      connect_timeout             = 300
      socket_timeout              = 1800
      incremental_loading_enabled = false
      default_time_zone           = "Asia/Tokyo"
      use_legacy_datetime_code    = false
      input_option_columns = [
        {
          name = "id"
          type = "long"
        },
        {
          name = "type"
          type = "string"
        },
      ]
      query = <<-EOT
        SELECT *
        FROM test_table
        WHERE id < 100
      EOT
    }
  }

  output_option_type = "gcs"
  output_option = {
    gcs_output_option = {
      gcs_connection_id       = 1
      bucket                  = "my-test-bucket"
      path_prefix             = "output/test/"
      file_ext                = ".csv"
      sequence_format         = ".%03d.%02d"
      is_minimum_output_tasks = false
      formatter_type          = "csv"
      encoder_type            = "gzip"
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