resource "trocco_job_definition" "mysql_output_example" {
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
      before_load         = "DELETE FROM test_table WHERE status = 'old';"
      after_load          = "UPDATE test_table SET updated_at = NOW();"
      mysql_output_option_column_options = [
        {
          name = "description"
          type = "TEXT"
        },
        {
          name      = "price"
          type      = "DECIMAL"
          scale     = 2
          precision = 10
        },
        {
          name = "notes"
          type = "LONGTEXT"
        }
      ]
      custom_variable_settings = [
        {
          name  = "$db_name$"
          type  = "string"
          value = "production_db"
        },
        {
          name  = "$table_name$"
          type  = "string"
          value = "orders"
        }
      ]
    }
  }
}
