resource "trocco_connection" "google_spreadsheets" {
  connection_type = "google_spreadsheets"
  name            = "Google Sheets Example"
  description     = "This is a Google Sheets connection example"

  service_account_json_key = <<JSON
  {
    "type": "service_account",
    "project_id": "example-project-id",
    "private_key_id": "example-private-key-id",
    "private_key":"-----BEGIN PRIVATE KEY-----\n..."
  }
  JSON
}
