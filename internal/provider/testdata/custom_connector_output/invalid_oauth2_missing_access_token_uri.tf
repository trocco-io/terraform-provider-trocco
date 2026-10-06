resource "trocco_custom_connector_output" "test" {
  name      = "custom-connector-output-invalid-oauth2"
  url       = "https://example.com"
  auth_type = "oauth2"

  grant_type = "client_credentials"
}
