resource "trocco_connection" "google_drive" {
  connection_type          = "google_drive"
  name                     = "Google Drive for Output"
  description              = "Google Drive connection for output testing"
  service_account_json_key = <<JSON
  {
    "type": "service_account",
    "project_id": "example-project-id",
    "private_key_id": "example-private-key-id",
    "private_key":"-----BEGIN PRIVATE KEY-----\n..."
  }
  JSON
}

resource "trocco_connection" "bigquery" {
  connection_type          = "bigquery"
  name                     = "BigQuery for Input"
  project_id               = "example"
  service_account_json_key = <<JSON
  {
    "type": "service_account",
    "project_id": "example-project-id",
    "private_key_id": "example-private-key-id",
    "private_key":"-----BEGIN PRIVATE KEY-----\n..."
  }
  JSON
}

resource "trocco_job_definition" "bigquery_to_google_drive" {
  name                     = "BigQuery to Google Drive CSV Export"
  description              = "Export BigQuery data to Google Drive in CSV format"
  resource_enhancement     = "medium"
  retry_limit              = 0
  is_runnable_concurrently = false

  input_option_type = "bigquery"
  input_option = {
    bigquery_input_option = {
      bigquery_connection_id   = trocco_connection.bigquery.id
      gcs_uri                  = "export_bucket"
      gcs_uri_format           = "bucket"
      query                    = "SELECT id, name, created_at FROM `dataset.users`"
      temp_dataset             = "temp_export"
      location                 = "asia-northeast1"
      is_standard_sql          = true
      cleanup_gcs_files        = true
      file_format              = "CSV"
      cache                    = false
      bigquery_job_wait_second = 600

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

  output_option_type = "google_drive"
  output_option = {
    google_drive_output_option = {
      google_drive_connection_id = trocco_connection.google_drive.id
      main_folder_id             = "1BxiMVs0XRA5nFMdKvBdBZjgmUUqptlbs"
      file_name                  = "users_export.csv"
      formatter_type             = "csv"

      csv_formatter = {
        delimiter           = ","
        newline             = "CRLF"
        newline_in_field    = "LF"
        charset             = "UTF-8"
        quote_policy        = "MINIMAL"
        escape              = "\\"
        header_line         = true
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
