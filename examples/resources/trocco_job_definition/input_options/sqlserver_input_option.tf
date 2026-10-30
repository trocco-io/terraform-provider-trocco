resource "trocco_job_definition" "sqlserver_input_example" {
  input_option_type = "sqlserver"
  input_option = {
    sqlserver_input_option = {
      sqlserver_connection_id     = 1 # require your sqlserver connection id
      database                    = "test_database"
      schema                      = "dbo"
      incremental_loading_enabled = false
      connect_timeout             = 300
      socket_timeout              = 1801
      fetch_rows                  = 1000
      default_time_zone           = "Asia/Tokyo"
      query                       = <<-EOT
        select
            *
        from
            dbo.example_table;
      EOT
      # Columns whose data type is not supported by TROCCO can be loaded as string.
      input_option_column_options = [
        {
          column_name       = "created_at"
          column_value_type = "string"
        }
      ]
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
          type = "string"
        },
      ]
    }
  }
}
