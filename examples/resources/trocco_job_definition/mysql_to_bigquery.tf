resource "trocco_job_definition" "mysql_to_bigquery_example" {
  name                     = "mysql_to_bigquery_example"
  description              = ""
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
  input_option_type = "mysql"
  input_option = {
    mysql_input_option = {
      connect_timeout             = 300
      database                    = "example_database"
      default_time_zone           = ""
      fetch_rows                  = 10000
      incremental_loading_enabled = false
      input_option_columns = [
        {
          name = "id"
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
        },
      ]
      mysql_connection_id      = 1 // please set your mysql connection id
      query                    = "select * from example_table;"
      socket_timeout           = 1800
      use_legacy_datetime_code = false
    }
  }
  output_option_type = "bigquery"
  output_option = {
    bigquery_output_option = {
      auto_create_dataset                      = false
      auto_create_table                        = false
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
      table                                    = "mysql_to_bigquery_example_table"
      template_table                           = ""
      timeout_sec                              = 300
    }
  }
}
