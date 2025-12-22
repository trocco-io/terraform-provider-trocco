# BigQuery to SFTP Csv Job Definition Example
resource "trocco_job_definition" "bigquery_to_sftp_csv" {
  name           = "BigQuery to SFTP CSV Export"
  description    = "Export BigQuery data to SFTP in CSV format"

  input_option_type  = "bigquery"
  output_option_type = "sftp"

  filter_columns = [
    {
      default                      = ""
      format                       = "%Y"
      json_expand_enabled          = false
      json_expand_keep_base_column = false
      name                         = "id"
      src                          = "id"
      type                         = "long"
    },
    {
      default                      = ""
      json_expand_enabled          = false
      json_expand_keep_base_column = false
      name                         = "name"
      src                          = "name"
      type                         = "string"
    },
  ]

  input_option = {
    bigquery_input_option = {
      bigquery_connection_id   = trocco_connection.bigquery.id
      gcs_uri                  = "export_bucket"
      gcs_uri_format           = "bucket"
      query                    = "SELECT user_id, name, email, created_at FROM `dataset.users` WHERE DATE(created_at) = CURRENT_DATE() - 1"
      temp_dataset             = "temp_export"
      location                 = "asia-northeast1"
      is_standard_sql          = true
      cleanup_gcs_files        = true
      file_format              = "CSV"
      cache                    = false
      bigquery_job_wait_second = 600

      columns = [
        {
          name = "user_id"
          type = "long"
        },
        {
          name = "name"
          type = "string"
        },
        {
          name = "email"
          type = "string"
        },
        {
          name = "created_at"
          type = "timestamp"
        }
      ]
    }
  }

  output_option = {
    sftp_output_option = {
      sftp_connection_id      = trocco_connection.sftp.id
      path_prefix             = "/exports/users/users_$export_date$"
      file_ext                = ".csv"
      is_minimum_output_tasks = false
      encoder_type            = "gzip"

      csv_formatter = {
        delimiter               = ","
        newline                 = "CRLF"
        newline_in_field        = "LF"
        charset                 = "UTF-8"
        quote_policy            = "MINIMAL"
        escape                  = "\\"
        header_line             = true
        null_string_enabled     = true
        null_string             = "NULL"
        default_time_zone       = "Asia/Tokyo"

        csv_formatter_column_options_attributes = [
          {
            name     = "created_at"
            format   = "%Y-%m-%d %H:%M:%S"
            timezone = "Asia/Tokyo"
          }
        ]
      }

      custom_variable_settings = [
        {
          name      = "$export_date$"
          type      = "timestamp"
          quantity  = 1
          unit      = "date"
          direction = "ago"
          format    = "%Y%m%d"
          time_zone = "Asia/Tokyo"
        }
      ]
    }
  }
}

# BigQuery source connection
resource "trocco_connection" "bigquery" {
  connection_type = "bigquery"
  name           = "Analytics BigQuery"
  description    = "BigQuery connection for analytics data"
  
  project_id               = "my-analytics-project"
  service_account_json_key = <<JSON
  {
    "type": "service_account",
    "project_id": "example-project-id",
    "private_key_id": "example-private-key-id",
    "private_key":"-----BEGIN PRIVATE KEY-----\n..."
  }
  JSON
}

# SFTP destination connection
resource "trocco_connection" "sftp" {
  connection_type = "sftp"
  name           = "Analytics SFTP"
  description    = "SFTP server for analytics exports"
  
  host                    = "analytics-sftp.example.com"
  port                    = 22
  user_name               = "analytics"
}