resource "trocco_connection" "databricks_pat" {
  connection_type = "databricks"

  name        = "Databricks Example with PAT Auth"
  description = "This is a Databricks connection example"
  host                   = "example.databricks.com"
  http_path              = "/sql/1.0/warehouses/xxxx-xxxx-xxxx-xxxx"
  auth_type             = "pat"
  personal_access_token = "dapiXXXXXXXXXXXXXXXXXXXX"
}

resource "trocco_connection" "databricks_oauth2" {
  connection_type = "databricks"

  name        = "Databricks Example with OAuth2"
  description = "This is a Databricks connection example using OAuth2"
  host                   = "example.databricks.com"
  http_path              = "/sql/1.0/warehouses/xxxx-xxxx-xxxx-xxxx"
  auth_type             = "oauth-m2m"
  oauth2_client_id      = "your-oauth2-client-id"
  oauth2_client_secret  = "your-oauth2-client-secret"
}
