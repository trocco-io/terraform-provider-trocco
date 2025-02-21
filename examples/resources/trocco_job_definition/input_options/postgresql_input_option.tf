resource "trocco_job_definition" "postgresql_input_example" {
  input_option_type = "postgresql"
  input_option = {
    postgresql_input_option = {
      postgresql_connection_id    = 1 # require your postgresql connection id
      database                    = "test_database"
      schema                      = "public"
      incremental_loading_enabled = false
      connect_timeout             = 300
      socket_timeout              = 1801
      fetch_rows                  = 1000
      default_time_zone           = "Asia/Tokyo"
      query                       = <<-EOT
        select
            *
        from
            example_table;
      EOT
      postgresql_input_option_column_options : [
        {
          column_name : "test"
          column_value_type : "string"
        },
      ]
    }
  }
}
