resource "trocco_custom_connector_output" "test" {
  name      = "custom-connector-output-duplicate-endpoint-name"
  url       = "https://example.com"
  auth_type = "api_key"

  endpoints = [
    {
      name      = "create_user"
      operation = "create"
      path      = "/users"
      template  = "{\"name\": \"{{ name }}\"}"
    },
    {
      name      = "create_user"
      operation = "create"
      path      = "/users"
      template  = "{\"name\": \"{{ name }}\"}"
    },
  ]
}
