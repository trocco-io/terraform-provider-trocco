# BigQuery to SFTP JSONL Job Definition Example
resource "trocco_job_definition" "bigquery_to_sftp_jsonl" {
  name           = "BigQuery to SFTP JSONL Export"
  description    = "Export BigQuery analytics data to SFTP in JSONL format"

  input_option_type = "bigquery"
  output_option_type = "sftp"

  filter_columns = [
    {
      name                         = "id",
      src                          = "id",
      type                         = "long",
      default                      = "",
      has_parser                   = true,
      json_expand_enabled          = false,
      json_expand_keep_base_column = false,
      json_expand_columns          = null
    },
  ]

  input_option = {
    bigquery_input_option = {
      bigquery_connection_id   = trocco_connection.bigquery.id
      gcs_uri                  = "test_bucket"
      gcs_uri_format           = "bucket"
      query                    = "SELECT * FROM `test_dataset.test_table`"
      temp_dataset             = "temp_dataset"
      location                 = "asia-northeast1"
      is_standard_sql          = true
      cleanup_gcs_files        = true
      file_format              = "CSV"
      cache                    = true
      bigquery_job_wait_second = 600

      columns = [
        {
          name = "col1__c"
          type = "string"
        }
      ]
    }
  }

  output_option = {
    sftp_output_option = {
      sftp_connection_id       = trocco_connection.sftp.id
      path_prefix              = "/analytics/events/$date$/events"
      file_ext                 = ".jsonl"
      is_minimum_output_tasks  = true
      encoder_type             = "gzip"
      sequence_format          = "%03d.%02d"

      jsonl_formatter = {
        encoding    = "UTF-8"
        newline     = "LF"
        date_format = "%Y-%m-%d %H:%M:%S"
        timezone    = "UTC"
      }
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
