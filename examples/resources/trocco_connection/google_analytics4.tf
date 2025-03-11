resource "trocco_connection" "google_analytics4" {
  connection_type   = "google_analytics4"
  name              = "Google Analytics4 Example"
  description       = "This is a Google Analytics4 connection example"
  resource_group_id = 1

  service_account_json_key = <<JSON
  {
    "type": "service_account",
    "project_id": "example-project-id",
    "private_key_id": "example-private-key-id",
    "private_key":"-----BEGIN PRIVATE KEY-----\n..."
  }
  JSON
}
