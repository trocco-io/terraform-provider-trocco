resource "trocco_connection" "test_oracle_incremental" {
  connection_type = "oracle"
  name            = "Oracle Incremental Example"
  host            = "db.example.com"
  port            = 1521
  user_name       = "system"
  password        = "password"
  driver          = "19c-ojdbc8.jar"
}

resource "trocco_team" "test_incremental" {
  name = "test_incremental"
  members = [
    {
      user_id = 10626
      role    = "team_admin"
    },
  ]
}

resource "trocco_resource_group" "test_incremental" {
  name        = "test_incremental"
  description = "test"
  teams = [
    {
      team_id = trocco_team.test_incremental.id
      role    = "administrator"
    },
  ]
}

resource "trocco_job_definition" "oracle_incremental_to_bigquery" {
  name                     = "Oracle Incremental to BigQuery Test"
  description              = "Test job definition for incremental transfer from Oracle to BigQuery"
  resource_enhancement     = "medium"
  resource_group_id        = trocco_resource_group.test_incremental.id
  retry_limit              = 3
  is_runnable_concurrently = true

  input_option_type = "oracle"
  input_option = {
    oracle_input_option = {
      oracle_connection_id        = trocco_connection.test_oracle_incremental.id
      database                    = "test_database"
      connection_type             = "sid"
      schema                      = "test_schema"
      table                       = "test_table"
      incremental_loading_enabled = true
      incremental_columns         = "updated_at"
      default_time_zone           = "UTC"
      input_option_columns = [
        {
          name = "id"
          type = "long"
        },
        {
          name = "updated_at"
          type = "timestamp"
        },
      ]
    }
  }

  filter_columns = [
    {
      default                      = ""
      format                       = "%Y"
      json_expand_enabled          = false
      json_expand_keep_base_column = false
      name                         = "id"
      src                          = "id"
      type                         = "long"
    },
    {
      default                      = ""
      format                       = ""
      json_expand_enabled          = false
      json_expand_keep_base_column = false
      name                         = "updated_at"
      src                          = "updated_at"
      type                         = "timestamp"
    }
  ]

  output_option_type = "bigquery"
  output_option = {
    bigquery_output_option = {
      dataset                                  = "test_dataset"
      table                                    = "oracle_incremental_to_bigquery_test_table"
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
