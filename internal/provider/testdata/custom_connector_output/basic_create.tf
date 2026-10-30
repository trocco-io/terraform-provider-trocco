resource "trocco_custom_connector_output" "test" {
  name        = "custom-connector-output-test"
  description = "test custom connector output"
  url         = "https://example.com"
  auth_type   = "api_key"

  auth_header_name   = "Authorization"
  auth_header_scheme = "Bearer"

  endpoints = [
    {
      name                = "create_user"
      request_type        = "single"
      method              = "POST"
      operation           = "create"
      path                = "/users"
      payload_type        = "json"
      success_codes       = "200,201,204"
      not_retryable_codes = "400,401,403,404"
      request_timeout_sec = 45
      template            = "{\"name\": \"{{ name }}\"}"

      headers = [
        {
          name          = "X-Api-Version"
          display_name  = "API Version"
          default_value = "2"
          is_editable   = false
          is_required   = true
        },
      ]

      path_parameters  = []
      query_parameters = []
    },
  ]
}
