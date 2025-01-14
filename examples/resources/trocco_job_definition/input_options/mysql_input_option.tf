resource "trocco_job_definition" "mysql_input_example" {
  input_option_type             = "mysql"
  input_option                  = {
    mysql_input_option = {
      connect_timeout             = 300
      socket_timeout              = 1801
      database                    = "test_database"
      fetch_rows                  = 1000
      incremental_loading_enabled = false
      default_time_zone           = "Asia/Tokyo"
      use_legacy_datetime_code    = false
      mysql_connection_id         = <your mysql connection id>
      input_option_columns        = [
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
      query                       = <<-EOT
                select
                    *
                from
                    example_table;
      EOT
    }
  }
}