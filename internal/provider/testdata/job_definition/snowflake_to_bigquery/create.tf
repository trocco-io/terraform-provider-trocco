resource "trocco_connection" "test_snowflake" {
  connection_type = "snowflake"
  name            = "Snowflake Example"
  description     = "This is a Snowflake connection example"
  host            = "example.snowflakecomputing.com"
  auth_method     = "user_password"
  user_name       = "test_user"
  password        = "test_password"
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

resource "trocco_job_definition" "snowflake_to_bigquery" {
  name                     = "Snowflake to BigQuery Test"
  description              = "Test job definition for transferring data from Snowflake to BigQuery"
  resource_enhancement     = "medium"
  resource_group_id        = trocco_resource_group.test.id
  retry_limit              = 2
  is_runnable_concurrently = true

  input_option_type = "snowflake"
  input_option = {
    snowflake_input_option = {
      snowflake_connection_id = trocco_connection.test_snowflake.id
      database                = "test_database"
      schema                  = "PUBLIC"
      warehouse               = "COMPUTE_WH"
      query                   = "SELECT id, name, email, created_at FROM test_table"
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
        }
      ]
    }
  }

  filter_columns = [
    {
      default                      = null
      json_expand_enabled          = false
      json_expand_keep_base_column = false
      name                         = "user_id"
      src                          = "id"
      type                         = "long"
    },
    {
      default                      = null
      json_expand_enabled          = false
      json_expand_keep_base_column = false
      name                         = "user_name"
      src                          = "name"
      type                         = "string"
    },
    {
      default                      = null
      json_expand_enabled          = false
      json_expand_keep_base_column = false
      name                         = "user_email"
      src                          = "email"
      type                         = "string"
    },
    {
      default                      = null
      format                       = "%Y-%m-%d"
      json_expand_enabled          = false
      json_expand_keep_base_column = false
      name                         = "registration_date"
      src                          = "created_at"
      type                         = "timestamp"
    }
  ]

  output_option_type = "bigquery"
  output_option = {
    bigquery_output_option = {
      dataset                = "test_dataset"
      table                  = "snowflake_to_bigquery_test_table"
      mode                   = "append"
      auto_create_dataset    = true
      timeout_sec            = 300
      open_timeout_sec       = 300
      read_timeout_sec       = 300
      send_timeout_sec       = 300
      retries                = 3
      bigquery_connection_id = trocco_connection.test_bq.id
      location               = "US"
    }
  }

  # please create labels if testing in local environment
  # see https://trocco.io/labels#side-nav-labels
  # labels = [
  #   {
  #     name = "test_label1"
  #   },
  #   {
  #     name = "test_label2"
  #   },
  # ]
}
