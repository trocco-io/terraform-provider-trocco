resource "trocco_connection" "bigquery" {
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


resource "trocco_connection" "bigquery_wif" {
  connection_type = "bigquery"
  name            = "BigQuery with WIF"
  description     = "BigQuery connection using Workload Identity Federation authentication"

  # WIF authentication configuration
  project_id                          = "example"
  is_workload_identity_federation     = true
  workload_identity_federation_config = <<JSON
{
  "type": "external_account",
  "audience": "//iam.googleapis.com/projects/123456789/locations/global/workloadIdentityPools/aws-pool/providers/aws-provider",
  "service_account_impersonation_url": "https://iamcredentials.googleapis.com/v1/projects/-/serviceAccounts/sa-bigquery@my-gcp-project.iam.gserviceaccount.com:generateAccessToken",
  "subject_token_type": "urn:ietf:params:aws:token-type:aws4_request",
  "token_url": "https://sts.googleapis.com/v1/token",
  "credential_source": {
    "environment_id": "aws1",
    "regional_cred_verification_url": "https://sts.ap-northeast-1.amazonaws.com?Action=GetCallerIdentity&Version=2011-06-15"
  }
}
JSON
}
