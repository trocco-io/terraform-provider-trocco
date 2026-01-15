resource "trocco_job_definition" "mongodb_to_bigquery_example" {
  name                     = "mongodb_to_bigquery_example"
  description              = ""
  is_runnable_concurrently = false
  retry_limit              = 0
  filter_columns = [
    {
      default                      = null
      json_expand_columns          = []
      json_expand_enabled          = false
      json_expand_keep_base_column = false
      name                         = "_id"
      src                          = "_id"
      type                         = "string"
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
  input_option_type = "mongodb"
  input_option = {
    mongodb_input_option = {
      mongodb_connection_id       = 1 // please set your mongodb connection id
      database                    = "example_database"
      collection                  = "example_collection"
      query                       = "{\"status\": \"active\"}"
      incremental_loading_enabled = true
      incremental_columns         = "created_at"
      last_record                 = "{\"created_at\":\"2024-01-01 00:00:00\"}"
      input_option_columns = [
        {
          name = "_id"
          type = "string"
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
          name     = "created_at"
          type     = "timestamp"
          format   = "%Y-%m-%d %H:%M:%S"
          timezone = "UTC"
        },
      ]
      custom_variable_settings = [
        {
          name  = "$collection_name$"
          type  = "string"
          value = "example_collection"
        }
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
      table                                    = "mongodb_to_bigquery_example_table"
      template_table                           = ""
      timeout_sec                              = 300
    }
  }
}
