resource "trocco_connection" "test_mongodb" {
  connection_type          = "mongodb"
  name                     = "MongoDB Test Connection"
  description              = "This is a MongoDB connection for testing"
  connection_string_format = "standard"
  host                     = "mongodb.example.com"
  port                     = 27017
  auth_method              = "scram-sha-1"
  user_name                = "testuser"
  password                 = "testpassword"
  auth_source              = "admin"
  read_preference          = "primary"
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

resource "trocco_job_definition" "mongodb_to_bigquery" {
  name                     = "MongoDB to BigQuery Test"
  description              = "Test job definition for transferring data from MongoDB to BigQuery"
  resource_enhancement     = "medium"
  resource_group_id        = trocco_resource_group.test.id
  retry_limit              = 2
  is_runnable_concurrently = false

  input_option_type = "mongodb"
  input_option = {
    mongodb_input_option = {
      mongodb_connection_id       = trocco_connection.test_mongodb.id
      database                    = "test_database"
      collection                  = "test_collection"
      query                       = "{\"status\": \"active\"}"
      incremental_loading_enabled = true
      incremental_columns         = "created_at"
      last_record                 = "{\"created_at\":\"2024-01-01 00:00:00\"}"
      input_option_columns = [
        {
          name = "_id"
          type = "long"
        },
        {
          name     = "created_at"
          type     = "timestamp"
          format   = "%Y-%m-%d %H:%M:%S"
          timezone = "UTC"
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
      name                         = "created_at"
      src                          = "created_at"
      type                         = "timestamp"
    }
  ]

  output_option_type = "bigquery"
  output_option = {
    bigquery_output_option = {
      dataset                = "test_dataset"
      table                  = "mongodb_test_table"
      mode                   = "append"
      auto_create_dataset    = false
      bigquery_connection_id = trocco_connection.test_bq.id
      location               = "US"
    }
  }
}
