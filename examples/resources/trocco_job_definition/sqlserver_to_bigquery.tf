resource "trocco_job_definition" "sqlserver_to_bigquery_example" {
  name                     = "sqlserver_to_bigquery_example"
  description              = ""
  is_runnable_concurrently = false
  retry_limit              = 0
  filter_columns = [
    {
      default                      = null
      json_expand_enabled          = false
      json_expand_keep_base_column = false
      name                         = "id"
      src                          = "id"
      type                         = "long"
    },
    {
      default                      = null
      json_expand_enabled          = false
      json_expand_keep_base_column = false
      name                         = "name"
      src                          = "name"
      type                         = "string"
    },
  ]
  input_option_type = "sqlserver"
  input_option = {
    sqlserver_input_option = {
      sqlserver_connection_id     = 1 // please set your sqlserver connection id
      database                    = "example_database"
      schema                      = "dbo"
      incremental_loading_enabled = false
      query                       = "select * from dbo.example_table;"
      connect_timeout             = 300
      socket_timeout              = 1800
      fetch_rows                  = 10000
      default_time_zone           = "UTC"
      input_option_columns = [
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
      table                                    = "sqlserver_to_bigquery_example_table"
      template_table                           = ""
      timeout_sec                              = 300
    }
  }
}
