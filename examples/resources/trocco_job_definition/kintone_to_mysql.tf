resource "trocco_job_definition" "kintone_to_mysql_example" {
  name                     = "kintone_to_mysql_example"
  description              = "Transfer data from Kintone to Mysql"
  is_runnable_concurrently = false
  retry_limit              = 0

  filter_columns = [
    {
      default                      = null
      json_expand_columns          = []
      json_expand_enabled          = false
      json_expand_keep_base_column = false
      name                         = "record_id"
      src                          = "record_id"
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
      json_expand_columns          = []
      json_expand_enabled          = false
      json_expand_keep_base_column = false
      name                         = "email"
      src                          = "email"
      type                         = "string"
    },
  ]

  input_option_type = "kintone"
  input_option = {
    kintone_input_option = {
      kintone_connection_id = 1 # require your kintone connection id
      app_id                = "123"
      guest_space_id        = null
      query                 = null
      expand_subtable       = false
      custom_variable_settings = [
        {
          name  = "$app_id$"
          type  = "string"
          value = "123"
        }
      ]
      input_option_columns = [
        {
          name = "record_id"
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
      ]
    }
  }

  output_option_type = "mysql"
  output_option = {
    mysql_output_option = {
      mysql_connection_id = 1 # require your mysql connection id
      database            = "$db_name$"
      table               = "$table_name$"
      mode                = "insert"
      retry_limit         = 12
      retry_wait          = 1000
      max_retry_wait      = 1800000
      default_time_zone   = "UTC"
      before_load         = "DELETE FROM customer_data WHERE status = 'pending';"
      after_load          = "UPDATE customer_data SET updated_at = NOW();"
      mysql_output_option_column_options = [
        {
          name = "description"
          type = "TEXT"
        },
        {
          name      = "amount"
          type      = "DECIMAL"
          scale     = 2
          precision = 10
        },
        {
          name = "comments"
          type = "LONGTEXT"
        }
      ]
      custom_variable_settings = [
        {
          name  = "$db_name$"
          type  = "string"
          value = "customer_db"
        },
        {
          name  = "$table_name$"
          type  = "string"
          value = "customer_data"
        }
      ]
    }
  }
}
