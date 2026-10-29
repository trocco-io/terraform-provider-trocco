resource "trocco_connection" "test_oracle_missing_columns" {
  connection_type = "oracle"
  name            = "Oracle Missing Columns Example"
  host            = "db.example.com"
  port            = 1521
  user_name       = "system"
  password        = "password"
  driver          = "19c-ojdbc8.jar"
}

resource "trocco_team" "test_missing_columns" {
  name = "test_missing_columns"
  members = [
    {
      user_id = 10626
      role    = "team_admin"
    },
  ]
}

resource "trocco_resource_group" "test_missing_columns" {
  name        = "test_missing_columns"
  description = "test"
  teams = [
    {
      team_id = trocco_team.test_missing_columns.id
      role    = "administrator"
    },
  ]
}

resource "trocco_job_definition" "oracle_missing_input_option_columns" {
  name                     = "Oracle Missing Columns Test"
  resource_enhancement     = "medium"
  resource_group_id        = trocco_resource_group.test_missing_columns.id
  retry_limit              = 3
  is_runnable_concurrently = true

  input_option_type = "oracle"
  input_option = {
    oracle_input_option = {
      oracle_connection_id        = trocco_connection.test_oracle_missing_columns.id
      database                    = "test_database"
      connection_type             = "sid"
      incremental_loading_enabled = false
      query                       = <<-EOT
        select
            *
        from
            example_table;
      EOT
    }
  }

  output_option_type = "bigquery"
  output_option = {
    bigquery_output_option = {
      dataset                                  = "test_dataset"
      table                                    = "oracle_missing_columns_test_table"
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
