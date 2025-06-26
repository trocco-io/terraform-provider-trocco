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


resource "trocco_job_definition" "http_to_bigquery" {
  name              = "HTTP to BigQuery Test"
  description       = "Test job definition for transferring data from HTTP to BigQuery"
  input_option_type = "http"
  input_option = {
    http_input_option = {
      method     = "GET"
      url        = "https://example.com"
      pager_type = "offset"
      jsonl_parser = {
        stop_on_invalid_record = true
        default_time_zone      = "UTC"
        newline                = "LF"
        charset                = "UTF-8"
        columns = [
          {
            name = "id"
            type = "long"
          },
          {
            name = "name"
            type = "string"
          }
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
    }
  ]
  output_option_type = "bigquery"
  output_option = {
    bigquery_output_option = {
      dataset                                  = "test_dataset"
      table                                    = "http_to_bigquery_table"
      mode                                     = "append"
      auto_create_dataset                      = true
      timeout_sec                              = 300
      open_timeout_sec                         = 300
      read_timeout_sec                         = 300
      send_timeout_sec                         = 300
      retries                                  = 3
      bigquery_connection_id                   = trocco_connection.test_bq.id
      location                                 = "US"
      bigquery_output_option_clustering_fields = []
      bigquery_output_option_column_options    = []
      bigquery_output_option_merge_keys        = []
    }
  }
}
