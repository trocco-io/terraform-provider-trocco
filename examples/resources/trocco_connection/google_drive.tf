resource "trocco_connection" "google_drive" {
  connection_type = "google_drive"
  name            = "Google Drive Example"
  description     = "This is a Google Drive connection example"

  service_account_json_key = <<JSON
  {
    "type": "service_account",
    "project_id": "example-project-id",
    "private_key_id": "example-private-key-id",
    "private_key":"-----BEGIN PRIVATE KEY-----\n..."
  }
  JSON
}
