resource "trocco_connection" "bigquery" {
  connection_type = "bigquery"

  name        = "BigQuery Example"
  description = "This is a BigQuery connection example"

  service_account_json_key = "{\"type\":\"service_account\", ...}"
}
