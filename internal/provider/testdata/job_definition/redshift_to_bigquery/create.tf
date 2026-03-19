resource "trocco_connection" "test_redshift" {
  connection_type       = "redshift"
  name                  = "Redshift Basic Example"
  description           = "This is a basic Redshift connection example"
  host                  = "my-cluster.abc123.us-east-1.redshift.amazonaws.com"
  port                  = 5439
  user_name             = "admin"
  password              = "password"
  aws_access_key_id     = "ACCESS_KEY_ID_EXAMPLE"
  aws_secret_access_key = "SECRET_ACCESS_KEY_EXAMPLE"
  ssl_enabled           = true
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

resource "trocco_job_definition" "redshift_to_bigquery" {
  name                     = "Redshift to BigQuery Test"
  description              = "Test job definition for transferring data from Redshift to BigQuery"
  resource_enhancement     = "medium"
  retry_limit              = 2
  is_runnable_concurrently = false

  input_option_type = "redshift"
  input_option = {
    redshift_input_option = {
      redshift_connection_id = trocco_connection.test_redshift.id
      database                    = "analytics"
      query                       = "SELECT * FROM test_table WHERE status = 'active'"
      schema                      = "public"
      fetch_rows                  = 1000
      connect_timeout             = 30
      socket_timeout              = 60
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
      table                  = "redshift_test_table"
      mode                   = "append"
      auto_create_dataset    = false
      bigquery_connection_id = trocco_connection.test_bq.id
      location               = "US"
    }
  }
}
