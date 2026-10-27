resource "trocco_job_definition" "sqlserver_empty_column_options" {
  name                     = "SQL Server empty column options"
  is_runnable_concurrently = false
  retry_limit              = 0

  filter_columns = [
    {
      name = "id"
      src  = "id"
      type = "long"
    },
  ]

  input_option_type = "sqlserver"
  input_option = {
    sqlserver_input_option = {
      sqlserver_connection_id     = 1
      database                    = "test_database"
      query                       = "select id from dbo.example_table"
      input_option_column_options = []
      input_option_columns = [
        {
          name = "id"
          type = "long"
        },
      ]
    }
  }

  output_option_type = "bigquery"
  output_option = {
    bigquery_output_option = {
      dataset                                  = "test_dataset"
      table                                    = "test_table"
      bigquery_connection_id                   = 1
      bigquery_output_option_clustering_fields = []
      bigquery_output_option_column_options    = []
      bigquery_output_option_merge_keys        = []
    }
  }
}
