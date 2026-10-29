resource "trocco_connection" "test_oracle_service_name" {
  connection_type = "oracle"
  name            = "Oracle Service Name Example"
  host            = "db.example.com"
  port            = 1521
  user_name       = "system"
  password        = "password"
  driver          = "19c-ojdbc8.jar"
}

resource "trocco_team" "test_service_name" {
  name = "test_service_name"
  members = [
    {
      user_id = 10626
      role    = "team_admin"
    },
  ]
}

resource "trocco_resource_group" "test_service_name" {
  name        = "test_service_name"
  description = "test"
  teams = [
    {
      team_id = trocco_team.test_service_name.id
      role    = "administrator"
    },
  ]
}

resource "trocco_job_definition" "oracle_service_name_to_bigquery" {
  name                     = "Oracle Service Name to BigQuery Test"
  description              = "Test job definition for transferring data from Oracle (service name) to BigQuery"
  resource_enhancement     = "medium"
  resource_group_id        = trocco_resource_group.test_service_name.id
  retry_limit              = 3
  is_runnable_concurrently = true

  input_option_type = "oracle"
  input_option = {
    oracle_input_option = {
      oracle_connection_id        = trocco_connection.test_oracle_service_name.id
      database                    = "test_database"
      connection_type             = "service_name"
      incremental_loading_enabled = false
      query                       = <<-EOT
        select
            *
        from
            example_table
        where
            category = $category$;
      EOT
      input_option_columns = [
        {
          name = "id"
          type = "long"
        },
      ]
      custom_variable_settings = [
        {
          name  = "$category$"
          type  = "string"
          value = "electronics"
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
      table                                    = "oracle_service_name_to_bigquery_test_table"
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
