resource "trocco_connection" "test" {
  connection_type = "bigquery"
  name            = "test"
  description     = "The quick brown fox jumps over the lazy dog."
  project_id      = "test"

  service_account_json_key = "{\"type\":\"service_account\",\"project_id\":\"\",\"private_key_id\":\"\",\"private_key\":\"\"}"
}