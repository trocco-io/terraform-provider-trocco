resource "trocco_connection" "test_postgresql" {
  connection_type = "postgresql"
  name            = "PostgreSQL Example"
  description     = "This is a PostgreSQL connection example"
  host            = "db.example.com"
  port            = 5432
  user_name       = "postgres"
  password        = "password"
  driver          = "postgresql_42_5_1"
}

resource "trocco_connection" "test_bq" {
  connection_type          = "bigquery"
  name                     = "BigQuery Example"
  project_id               = "example"
  service_account_json_key = <<JSON
  {
    "type": "service_account",
    "project_id": "example-project-id",
    "private_key_id": "example-private-key-id",
    "private_key":"-----BEGIN PRIVATE KEY-----\n..."
  }
  JSON
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

resource "trocco_job_definition" "postgresql_to_bigquery" {
  name                     = "PostgreSQL to BigQuery Test"
  description              = "Test job definition for transferring data from PostgreSQL to BigQuery"
  resource_enhancement     = "medium"
  resource_group_id        = trocco_resource_group.test.id
  retry_limit              = 3
  is_runnable_concurrently = true

  input_option_type = "postgresql"
  input_option = {
    postgresql_input_option = {
      postgresql_connection_id    = trocco_connection.test_postgresql.id
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
          name = "email"
          type = "string"
        },
        {
          name = "created_at"
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
      table                                    = "postgresql_to_bigquery_test_table"
      mode                                     = "append"
      auto_create_dataset                      = true
      timeout_sec                              = 300
      open_timeout_sec                         = 300
      read_timeout_sec                         = 300
      send_timeout_sec                         = 300
      retries                                  = 3
      bigquery_connection_id                   = trocco_connection.test_bq.id
      location                                 = "US"
      bigquery_output_option_clustering_fields = []
      bigquery_output_option_column_options    = []
      bigquery_output_option_merge_keys        = []
    }
  }

  # please create labels if testing in local environment
  # see https://trocco.io/labels#side-nav-labels
  labels = [
    {
      name = "label1"
    },
    {
      name = "label2"
    },
  ]
}
