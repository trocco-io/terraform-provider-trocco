resource "trocco_job_definition" "notifications_empty_test" {
  name                     = "notifications_empty_test"
  description              = "Test job definition with an explicitly empty notifications list"
  resource_enhancement     = "medium"
  retry_limit              = 0
  is_runnable_concurrently = false

  filter_columns = [
    {
      name                         = "id"
      src                          = "id"
      type                         = "long"
      default                      = ""
      format                       = ""
      json_expand_enabled          = false
      json_expand_keep_base_column = false
    },
    {
      name                         = "name"
      src                          = "name"
      type                         = "string"
      default                      = ""
      json_expand_enabled          = false
      json_expand_keep_base_column = false
    }
  ]

  input_option_type = "mysql"
  input_option = {
    mysql_input_option = {
      database                    = "test_database"
      table                       = "test_table"
      fetch_rows                  = 1000
      default_time_zone           = "Asia/Tokyo"
      incremental_loading_enabled = false
      mysql_connection_id         = trocco_connection.mysql.id
      input_option_columns = [
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

  output_option_type = "bigquery"
  output_option = {
    bigquery_output_option = {
      dataset                                  = "test_dataset"
      table                                    = "test_table"
      mode                                     = "append"
      location                                 = "US"
      bigquery_connection_id                   = trocco_connection.bigquery.id
      bigquery_output_option_clustering_fields = []
      bigquery_output_option_column_options    = []
      bigquery_output_option_merge_keys        = []
    }
  }

  notifications = []
}
