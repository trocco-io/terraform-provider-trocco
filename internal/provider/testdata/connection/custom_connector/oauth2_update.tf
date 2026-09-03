resource "trocco_custom_connector_input" "oauth2_test" {
  name             = "custom-connector-connection-test-oauth2-definition"
  url              = "https://example.com"
  auth_type        = "oauth2"
  grant_type       = "client_credentials"
  access_token_uri = "https://example.com/oauth2/token"

  endpoints = [
    {
      name                = "list_items"
      path                = "/items"
      method              = "GET"
      jsonpath_root       = "$.data"
      success_codes       = "200"
      not_retryable_codes = "400,401,403,404"

      query_parameters = []
      headers          = []
      path_parameters  = []
    },
  ]
}

resource "trocco_connection" "custom_connector_oauth2_test" {
  connection_type = "custom_connector"

  name        = "Test Custom Connector OAuth2 Connection"
  description = "Updated description"

  custom_connector_id = trocco_custom_connector_input.oauth2_test.id
  oauth2_client_id    = "test-oauth2-client-id"
  # Rotating the secret makes the server reset `authorized`, so the plan has to
  # leave it unknown rather than carrying the prior value forward.
  oauth2_client_secret = "updated-test-oauth2-client-secret"
  scopes               = ["read", "write"]
}
