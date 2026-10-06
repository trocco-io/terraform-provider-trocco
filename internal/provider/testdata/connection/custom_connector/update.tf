resource "trocco_custom_connector_input" "test" {
  name      = "custom-connector-connection-test-definition"
  url       = "https://example.com"
  auth_type = "api_key"

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

resource "trocco_connection" "custom_connector_test" {
  connection_type = "custom_connector"

  name                = "Test Custom Connector Connection"
  description         = "Updated description"
  custom_connector_id = trocco_custom_connector_input.test.id
  api_key             = "updated-test-api-key"
}
