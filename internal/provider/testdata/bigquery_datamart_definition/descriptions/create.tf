resource "trocco_connection" "my_conn_descriptions" {
  connection_type = "bigquery"

  name        = "BigQuery Example Descriptions"
  description = "This is a BigQuery connection example for table and column descriptions"

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

resource "trocco_bigquery_datamart_definition" "test_descriptions" {
  name                     = "test_descriptions"
  is_runnable_concurrently = false
  bigquery_connection_id   = trocco_connection.my_conn_descriptions.id
  query                    = <<SQL
    SELECT * FROM examples
  SQL
  query_mode               = "insert"
  destination_dataset      = "dist_datasets"
  destination_table        = "dist_tables"
  write_disposition        = "append"
  table_description        = "Table description"
  column_descriptions = [
    {
      name        = "id"
      description = "Primary key"
    },
    {
      name        = "name"
      description = "Name of the record"
    },
  ]
}
