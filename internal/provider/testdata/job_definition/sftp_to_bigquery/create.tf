resource "trocco_connection" "sftp" {
  connection_type = "sftp"
  name            = "SFTP Example"
  host            = "sftp.example.com"
  port            = 22
  user_name       = "testuser"
  password        = "password"
}

resource "trocco_connection" "bigquery" {
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

resource "trocco_job_definition" "sftp_to_bigquery" {
  name                     = "test job_definition"
  description              = "test description"
  resource_group_id        = 2
  retry_limit              = 1
  is_runnable_concurrently = true
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
  input_option_type = "sftp"
  input_option = {
    sftp_input_option = {
      sftp_connection_id          = trocco_connection.sftp.id
      path_prefix                 = "/data/files/"
      path_match_pattern          = ".*\\.csv$"
      incremental_loading_enabled = false
      stop_when_file_not_found    = false
      decompression_type          = "guess"
      csv_parser = {
        delimiter = ","
        escape    = "\\"
        quote     = "\""
        columns = [
          {
            name = "id"
            type = "long"
          },
          {
            name = "name"
            type = "string"
          },
        ]
      }
    }
  }
  output_option_type = "bigquery"
  output_option = {
    bigquery_output_option = {
      dataset                                  = "test_dataset"
      table                                    = "test_table"
      mode                                     = "append"
      auto_create_dataset                      = true
      timeout_sec                              = 300
      open_timeout_sec                         = 300
      read_timeout_sec                         = 300
      send_timeout_sec                         = 300
      retries                                  = 2
      bigquery_connection_id                   = trocco_connection.bigquery.id
      location                                 = "us-west1"
      bigquery_output_option_clustering_fields = []
      bigquery_output_option_column_options    = []
      bigquery_output_option_merge_keys        = []
    }
  }
}
