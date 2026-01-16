resource "trocco_job_definition" "sftp_to_bigquery_example" {
  name                     = "sftp_to_bigquery_example"
  description              = "transfer data from SFTP to BigQuery"
  is_runnable_concurrently = false
  retry_limit              = 0
  filter_columns = [
    {
      default                      = null
      json_expand_columns          = []
      json_expand_enabled          = false
      json_expand_keep_base_column = false
      name                         = "id"
      src                          = "id"
      type                         = "long"
    },
    {
      default                      = null
      json_expand_columns          = []
      json_expand_enabled          = false
      json_expand_keep_base_column = false
      name                         = "name"
      src                          = "name"
      type                         = "string"
    },
    {
      default                      = null
      format                       = null
      json_expand_columns          = []
      json_expand_enabled          = false
      json_expand_keep_base_column = false
      name                         = "created_at"
      src                          = "created_at"
      type                         = "timestamp"
    },
  ]
  input_option_type = "sftp"
  input_option = {
    sftp_input_option = {
      sftp_connection_id          = 1 # please set your sftp connection id
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
          {
            name = "created_at"
            type = "timestamp"
          },
        ]
      }
    }
  }
  output_option_type = "bigquery"
  output_option = {
    bigquery_output_option = {
      auto_create_dataset                      = false
      bigquery_connection_id                   = 1
      bigquery_output_option_clustering_fields = []
      bigquery_output_option_column_options    = []
      bigquery_output_option_merge_keys        = []
      dataset                                  = "example_dataset"
      location                                 = "US"
      mode                                     = "append"
      open_timeout_sec                         = 300
      read_timeout_sec                         = 300
      retries                                  = 5
      send_timeout_sec                         = 300
      table                                    = "sftp_to_bigquery_example_table"
      template_table                           = ""
      timeout_sec                              = 300
    }
  }
}
