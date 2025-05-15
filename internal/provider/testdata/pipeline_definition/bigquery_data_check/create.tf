resource "trocco_connection" "my_conn" {
  connection_type = "bigquery"

  name        = "BigQuery Example"
  description = "This is a BigQuery connection example"

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

resource "trocco_pipeline_definition" "bigquery_data_check_query_check" {
  name = "bigquery_data_check"

  tasks = [
    {
      key  = "bigquery_data_check"
      type = "bigquery_data_check"

      bigquery_data_check_config = {
        name          = "Example"
        connection_id = trocco_connection.my_conn.id
        query         = <<SQL
          SELECT COUNT(*) FROM examples
        SQL
        operator      = "equal"
        query_result  = 1
        accepts_null  = false
      }
    }
  ]
}
