resource "trocco_connection" "google_analytics4_test" {
  connection_type          = "google_analytics4"
  name                     = "test"
  description              = "test"
  service_account_json_key = "{\"type\":\"service_account\",\"project_id\":\"create_project_id\",\"private_key_id\":\"create_private_key_id\",\"private_key\":\"create_private_key\",\"client_email\":\"create_client_email\",\"client_id\":\"create_client_id\"}"
}