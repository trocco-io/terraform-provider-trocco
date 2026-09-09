resource "trocco_custom_connector_output" "test" {
  name      = "custom-connector-output-empty-endpoint-name"
  url       = "https://example.com"
  auth_type = "api_key"

  endpoints = [
    {
      # The server has no presence validation on an output endpoint name, so
      # this is rejected by the provider rather than the API.
      name      = ""
      operation = "create"
      path      = "/users"
      template  = "{\"name\": \"{{ name }}\"}"
    },
  ]
}
