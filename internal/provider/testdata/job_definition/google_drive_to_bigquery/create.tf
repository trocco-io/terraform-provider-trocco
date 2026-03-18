resource "trocco_connection" "test_google_drive" {
  connection_type          = "google_drive"
  name                     = "Google Drive Example"
  description              = "This is a Google Drive connection example"
  service_account_json_key = <<JSON
  {
    "type": "service_account",
    "project_id": "example-project-id",
    "private_key_id": "example-private-key-id",
    "private_key":"-----BEGIN PRIVATE KEY-----\n..."
  }
  JSON
}

resource "trocco_connection" "test_bq" {
  connection_type          = "bigquery"
  name                     = "BigQuery Example"
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

resource "trocco_job_definition" "google_drive_to_bigquery" {
  name                     = "Google Drive to BigQuery Test"
  description              = "Test job definition for transferring data from Google Drive to BigQuery"
  resource_enhancement     = "medium"
  retry_limit              = 0
  is_runnable_concurrently = false

  input_option_type = "google_drive"
  input_option = {
    google_drive_input_option = {
      google_drive_connection_id = trocco_connection.test_google_drive.id
      folder_id                  = "1BxiMVs0XRA5nFMdKvBdBZjgmUUqptlbs"
      file_match_pattern         = ""
      is_skip_header_line        = false
      stop_when_file_not_found   = false
      csv_parser = {
        allow_extra_columns     = false
        allow_optional_columns  = false
        charset                 = "UTF-8"
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
            format = "%Y-%m-%d %H:%M:%S.%N %z"
            name   = "created_at"
            type   = "timestamp"
          },
        ]
      }
    }
  }

  filter_columns = [
    {
      name = "id"
      src  = "id"
      type = "long"
    },
    {
      name = "name"
      src  = "name"
      type = "string"
    },
    {
      name   = "created_at"
      src    = "created_at"
      type   = "timestamp"
      format = "%Y-%m-%d %H:%M:%S.%N %z"
    },
  ]

  output_option_type = "bigquery"
  output_option = {
    bigquery_output_option = {
      auto_create_dataset                      = true
      bigquery_connection_id                   = trocco_connection.test_bq.id
      bigquery_output_option_clustering_fields = []
      bigquery_output_option_column_options    = []
      bigquery_output_option_merge_keys        = []
      dataset                                  = "test_dataset"
      location                                 = "US"
      mode                                     = "append"
      open_timeout_sec                         = 300
      read_timeout_sec                         = 300
      retries                                  = 3
      send_timeout_sec                         = 300
      table                                    = "google_drive_to_bigquery_test_table"
      template_table                           = ""
      timeout_sec                              = 300
    }
  }
}
