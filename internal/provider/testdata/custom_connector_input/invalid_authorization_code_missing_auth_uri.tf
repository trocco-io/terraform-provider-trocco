resource "trocco_custom_connector_input" "test" {
  name      = "custom-connector-input-authorization-code"
  url       = "https://example.com"
  auth_type = "oauth2"

  grant_type       = "authorization_code"
  access_token_uri = "https://example.com/token"
}
