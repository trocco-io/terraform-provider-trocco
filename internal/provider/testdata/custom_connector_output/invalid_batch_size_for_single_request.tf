resource "trocco_custom_connector_output" "test" {
  name      = "custom-connector-output-invalid-batch-size"
  url       = "https://example.com"
  auth_type = "api_key"

  endpoints = [
    {
      name         = "create_user"
      request_type = "single"
      batch_size   = 10
      operation    = "create"
      path         = "/users"
      template     = "{\"name\": \"{{ name }}\"}"
    },
  ]
}
