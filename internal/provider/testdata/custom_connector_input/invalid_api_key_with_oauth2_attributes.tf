resource "trocco_custom_connector_input" "test" {
  name      = "custom-connector-input-api-key-with-oauth2"
  url       = "https://example.com"
  auth_type = "api_key"

  grant_type       = "client_credentials"
  auth_uri         = "https://example.com/authorize"
  access_token_uri = "https://example.com/token"
}
