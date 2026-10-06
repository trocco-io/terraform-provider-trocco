resource "trocco_custom_connector_input" "test" {
  name      = "custom-connector-input-oauth2"
  url       = "https://example.com"
  auth_type = "oauth2"
}
