resource "trocco_connection" "custom_connector_api_key" {
  connection_type = "custom_connector"

  name                = "Custom Connector Example with API Key Auth"
  description         = "This is a Custom Connector connection example"
  custom_connector_id = 1
  api_key             = "your-api-key"
}

resource "trocco_connection" "custom_connector_oauth2" {
  connection_type = "custom_connector"

  name                 = "Custom Connector Example with OAuth2"
  description          = "This is a Custom Connector connection example using OAuth2"
  custom_connector_id  = 2
  oauth2_client_id     = "your-oauth2-client-id"
  oauth2_client_secret = "your-oauth2-client-secret"
  scopes               = ["read", "write"]
}
