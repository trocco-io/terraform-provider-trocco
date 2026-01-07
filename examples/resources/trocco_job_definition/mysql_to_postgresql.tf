resource "trocco_job_definition" "mysql_to_postgresql_example" {
  name                     = "mysql_to_postgresql_example"
  description              = "Transfer data from MySQL to PostgreSQL"
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
      json_expand_columns          = null
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
      database                    = "source_database"
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
      query                    = "select * from source_table;"
      socket_timeout           = 1800
      use_legacy_datetime_code = false
    }
  }
  output_option_type = "postgresql"
  output_option = {
    postgresql_output_option = {
      postgresql_connection_id = 1 // please set your postgresql connection id
      database                 = "destination_database"
      schema                   = "public"
      table                    = "mysql_to_postgresql_example_table"
      mode                     = "insert"
      default_time_zone        = "UTC"
    }
  }
}
