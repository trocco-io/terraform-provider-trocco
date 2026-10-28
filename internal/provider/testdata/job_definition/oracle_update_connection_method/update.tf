resource "trocco_connection" "test_oracle_update_method" {
  connection_type = "oracle"
  name            = "Oracle Update Connection Method Example"
  host            = "db.example.com"
  port            = 1521
  user_name       = "system"
  password        = "password"
  driver          = "19c-ojdbc8.jar"
}

resource "trocco_team" "test_update_method" {
  name = "test_update_method"
  members = [
    {
      user_id = 10626
      role    = "team_admin"
    },
  ]
}

resource "trocco_resource_group" "test_update_method" {
  name        = "test_update_method"
  description = "test"
  teams = [
    {
      team_id = trocco_team.test_update_method.id
      role    = "administrator"
    },
  ]
}

resource "trocco_job_definition" "oracle_update_connection_method" {
  name                     = "Oracle Update Connection Method Test (Updated)"
  description              = "Test switching the input_option connection method on update"
  resource_enhancement     = "medium"
  resource_group_id        = trocco_resource_group.test_update_method.id
  retry_limit              = 3
  is_runnable_concurrently = true

  input_option_type = "oracle"
  input_option = {
    oracle_input_option = {
      oracle_connection_id        = trocco_connection.test_oracle_update_method.id
      net_service_name            = "orcl"
      incremental_loading_enabled = false
      query                       = <<-EOT
        select
            *
        from
            example_table;
      EOT
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
      table                                    = "oracle_update_connection_method_test_table"
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
