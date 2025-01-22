resource "trocco_connection" "gcs" {
  connection_type = "gcs"

  name        = "GCS Example"
  description = "This is a Google Cloud Storage(GCS) connection example"

  project_id               = "example-project-id"
  service_account_json_key = <<JSON
  {
    "type": "service_account",
    "project_id": "example-project-id",
    "private_key_id": "example-private-key-id",
    "private_key":"-----BEGIN PRIVATE KEY-----\n..."
  }
  JSON
  service_account_email    = "joe@example-project.iam.gserviceaccount.com"
  application_name         = "example-application-name"
}
