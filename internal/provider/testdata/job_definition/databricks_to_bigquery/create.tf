resource "trocco_connection" "test_databricks" {
  connection_type = "databricks"

  name                  = "Databricks Example with PAT Auth "
  description           = "This is a Databricks connection example"
  server_hostname       = "example.databricks.com"
  http_path             = "/sql/1.0/warehouses/xxxx-xxxx-xxxx-xxxx"
  auth_type             = "pat"
  personal_access_token = "dapiXXXXXXXXXXXXXXXXXXXX"
}

resource "trocco_connection" "test_bq2" {
  connection_type = "bigquery"

  name        = "BigQuery Example"
  description = "This is a BigQuery connection example"

  project_id               = "systemn-playground"
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

resource "trocco_job_definition" "databricks_to_bigquery" {
  name                     = "test databricks_to_bigquery job"
  description              = "Test job definition for Databricks to BigQuery transfer"
  resource_enhancement     = "medium"
  resource_group_id        = trocco_resource_group.test.id
  retry_limit              = 2
  is_runnable_concurrently = false

  filter_columns = [
    {
      default                      = null
      json_expand_columns          = null
      json_expand_enabled          = false
      json_expand_keep_base_column = false
      name                         = "id"
      src                          = "id"
      type                         = "long"
    },
    {
      default                      = null
      json_expand_columns          = null
      json_expand_enabled          = false
      json_expand_keep_base_column = false
      name                         = "user_name"
      src                          = "user_name"
      type                         = "string"
    },
    {
      default                      = null
      json_expand_columns          = null
      json_expand_enabled          = false
      json_expand_keep_base_column = false
      name                         = "email"
      src                          = "email"
      type                         = "string"
    }
  ]

  input_option_type = "databricks"
  input_option = {
    databricks_input_option = {
      databricks_connection_id = trocco_connection.test_databricks.id
      catalog_name             = "test_catalog"
      schema_name              = "test_schema"
      query                    = <<-EOT
        SELECT
          id,
          user_name,
          email,
          created_at
        FROM test_catalog.test_schema.users
        WHERE created_at >= '2023-01-01'
        ORDER BY created_at DESC
        LIMIT 1000
      EOT
      input_option_columns = [
        {
          name = "id"
          type = "long"
        },
        {
          name = "user_name"
          type = "string"
        },
        {
          name = "email"
          type = "string"
        }
      ]
    }
  }

  output_option_type = "bigquery"
  output_option = {
    bigquery_output_option = {
      dataset                                  = "test_dataset"
      table                                    = "databricks_users"
      mode                                     = "replace"
      auto_create_dataset                      = true
      timeout_sec                              = 600
      open_timeout_sec                         = 300
      read_timeout_sec                         = 300
      send_timeout_sec                         = 300
      retries                                  = 3
      bigquery_connection_id                   = trocco_connection.test_bq2.id
      location                                 = "US"
      bigquery_output_option_clustering_fields = []
      bigquery_output_option_column_options    = []
      bigquery_output_option_merge_keys        = []
      template_table                           = ""
    }
  }
}
