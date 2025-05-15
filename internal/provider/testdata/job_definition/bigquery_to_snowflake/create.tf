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

resource "trocco_connection" "test_snowflake" {
  connection_type = "snowflake"
  name            = "Snowflake Example"
  description     = "This is a Snowflake connection example"
  host            = "example.snowflakecomputing.com"
  auth_method     = "user_password"
  user_name       = "test_user"
  password        = "test_password"
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

resource "trocco_job_definition" "bigquery_to_snowflake" {
  name                     = "BigQuery to Snowflake Test"
  description              = "Test job definition for transferring data from BigQuery to Snowflake"
  resource_enhancement     = "medium"
  resource_group_id        = trocco_resource_group.test.id
  retry_limit              = 2
  is_runnable_concurrently = true

  input_option_type = "bigquery"
  input_option = {
    bigquery_input_option = {
      bigquery_connection_id   = trocco_connection.test_bq.id
      gcs_uri                  = "test_bucket"
      gcs_uri_format           = "bucket"
      query                    = "SELECT id, name, email, created_at FROM `test_dataset.test_table`"
      temp_dataset             = "temp_dataset"
      location                 = "US"
      is_standard_sql          = true
      cleanup_gcs_files        = true
      file_format              = "CSV"
      cache                    = true
      bigquery_job_wait_second = 600

      columns = [
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

  output_option_type = "snowflake"
  output_option = {
    snowflake_output_option = {
      batch_size              = 50
      database                = "test_database"
      default_time_zone       = "UTC"
      delete_stage_on_error   = false
      empty_field_as_null     = true
      max_retry_wait          = 1800000
      mode                    = "insert"
      retry_limit             = 12
      retry_wait              = 1000
      schema                  = "PUBLIC"
      snowflake_connection_id = trocco_connection.test_snowflake.id
      table                   = "bigquery_to_snowflake_test_table"
      warehouse               = "COMPUTE_WH"
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
