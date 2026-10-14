resource "trocco_connection" "test_oracle" {
  connection_type = "oracle"
  name            = "Oracle Example"
  description     = "This is an Oracle connection example"
  host            = "db.example.com"
  port            = 1521
  user_name       = "system"
  password        = "password"
  driver          = "19c-ojdbc8.jar"
}

resource "trocco_team" "test" {
  name = "test"
  members = [
    {
      user_id = 10626
      role    = "team_admin"
    },
  ]
}

resource "trocco_resource_group" "test" {
  name        = "test"
  description = "test"
  teams = [
    {
      team_id = trocco_team.test.id
      role    = "administrator"
    },
  ]
}

resource "trocco_job_definition" "oracle_to_bigquery" {
  name                     = "Oracle to BigQuery Test (Updated)"
  description              = "Test job definition for transferring data from Oracle to BigQuery (updated)"
  resource_enhancement     = "medium"
  resource_group_id        = trocco_resource_group.test.id
  retry_limit              = 5
  is_runnable_concurrently = true

  input_option_type = "oracle"
  input_option = {
    oracle_input_option = {
      oracle_connection_id        = trocco_connection.test_oracle.id
      database                    = "test_database_updated"
      connection_type             = "sid"
      incremental_loading_enabled = false
      source_time_zone            = "UTC"
      query                       = <<-EOT
        select
            *
        from
            example_table_updated;
      EOT
      input_option_column_options = [
        {
          column_name       = "test"
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
          name     = "created_at"
          type     = "timestamp"
          format   = "%Y-%m-%d %H:%M:%S"
          timezone = "Asia/Tokyo"
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
      json_expand_enabled          = false
      json_expand_keep_base_column = false
      name                         = "name"
      src                          = "name"
      type                         = "string"
    },
    {
      default                      = ""
      format                       = ""
      json_expand_enabled          = false
      json_expand_keep_base_column = false
      name                         = "created_at"
      src                          = "created_at"
      type                         = "timestamp"
    }
  ]

  output_option_type = "bigquery"
  output_option = {
    bigquery_output_option = {
      dataset                                  = "test_dataset"
      table                                    = "oracle_to_bigquery_test_table_updated"
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

  labels = [
    {
      name = "label1"
    },
  ]
}
