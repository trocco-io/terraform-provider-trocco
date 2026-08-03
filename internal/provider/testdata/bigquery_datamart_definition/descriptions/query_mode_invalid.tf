resource "trocco_connection" "my_conn_descriptions_invalid" {
  connection_type = "bigquery"

  name        = "BigQuery Example Descriptions Invalid"
  description = "This is a BigQuery connection example for invalid descriptions in query mode"

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

resource "trocco_bigquery_datamart_definition" "test_descriptions_invalid" {
  name                     = "test_descriptions_invalid"
  is_runnable_concurrently = false
  bigquery_connection_id   = trocco_connection.my_conn_descriptions_invalid.id
  query                    = <<SQL
    SELECT * FROM examples
  SQL
  query_mode               = "query"
  table_description        = "Table description"
}
