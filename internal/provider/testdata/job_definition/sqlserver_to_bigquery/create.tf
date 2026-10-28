resource "trocco_connection" "test_sqlserver" {
  connection_type = "sqlserver"
  name            = "SQL Server Example"
  description     = "This is a SQL Server connection example"
  host            = "db.example.com"
  port            = 1433
  user_name       = "sa"
  password        = "password"
  driver          = "ms_sqlserver_jdbc_driver_13_4"
  ssl_enabled     = false
}

resource "trocco_job_definition" "sqlserver_to_bigquery" {
  name                     = "SQL Server to BigQuery Test"
  description              = "Test job definition for transferring data from SQL Server to BigQuery"
  retry_limit              = 3
  is_runnable_concurrently = true

  input_option_type = "sqlserver"
  input_option = {
    sqlserver_input_option = {
      sqlserver_connection_id     = trocco_connection.test_sqlserver.id
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

  filter_columns = [
    {
      name = "id"
      src  = "id"
      type = "long"
    },
    {
      name = "name"
      src  = "name"
      type = "string"
    },
    {
      name = "email"
      src  = "email"
      type = "string"
    },
    {
      name = "created_at"
      src  = "created_at"
      type = "string"
    },
  ]

  output_option_type = "bigquery"
  output_option = {
    bigquery_output_option = {
      dataset                                  = "test_dataset"
      table                                    = "sqlserver_to_bigquery_test_table"
      mode                                     = "append"
      auto_create_dataset                      = true
      timeout_sec                              = 300
      open_timeout_sec                         = 300
      read_timeout_sec                         = 300
      send_timeout_sec                         = 300
      retries                                  = 3
      bigquery_connection_id                   = trocco_connection.bigquery.id
      location                                 = "US"
      bigquery_output_option_clustering_fields = []
      bigquery_output_option_column_options    = []
      bigquery_output_option_merge_keys        = []
    }
  }
}
