resource "trocco_connection" "test_oracle_conflict" {
  connection_type = "oracle"
  name            = "Oracle Conflict Example"
  host            = "db.example.com"
  port            = 1521
  user_name       = "system"
  password        = "password"
  driver          = "19c-ojdbc8.jar"
}

resource "trocco_connection" "test_mysql_conflict" {
  connection_type = "mysql"
  name            = "MySQL Conflict Example"
  host            = "mysql.example.com"
  port            = 3306
  user_name       = "root"
  password        = "password"
}

resource "trocco_team" "test_conflict" {
  name = "test_conflict"
  members = [
    {
      user_id = 10626
      role    = "team_admin"
    },
  ]
}

resource "trocco_resource_group" "test_conflict" {
  name        = "test_conflict"
  description = "test"
  teams = [
    {
      team_id = trocco_team.test_conflict.id
      role    = "administrator"
    },
  ]
}

resource "trocco_job_definition" "oracle_input_option_conflict" {
  name                     = "Oracle Input Option Conflict Test"
  resource_enhancement     = "medium"
  resource_group_id        = trocco_resource_group.test_conflict.id
  retry_limit              = 3
  is_runnable_concurrently = true

  input_option_type = "oracle"
  input_option = {
    oracle_input_option = {
      oracle_connection_id        = trocco_connection.test_oracle_conflict.id
      database                    = "test_database"
      incremental_loading_enabled = false
      query                       = "select * from example_table;"
      input_option_columns = [
        {
          name = "id"
          type = "long"
        },
      ]
    }
    mysql_input_option = {
      mysql_connection_id         = trocco_connection.test_mysql_conflict.id
      database                    = "test_database"
      table                       = "test_table"
      incremental_loading_enabled = false
      query                       = "select * from example_table;"
      input_option_columns = [
        {
          name = "id"
          type = "long"
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
  ]

  output_option_type = "bigquery"
  output_option = {
    bigquery_output_option = {
      dataset                                  = "test_dataset"
      table                                    = "oracle_input_option_conflict_test_table"
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
