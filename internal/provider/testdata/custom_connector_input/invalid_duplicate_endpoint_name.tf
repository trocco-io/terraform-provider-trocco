resource "trocco_custom_connector_input" "test" {
  name      = "custom-connector-input-dup"
  url       = "https://example.com"
  auth_type = "api_key"

  endpoints = [
    {
      name                = "list_users"
      path                = "/users"
      method              = "GET"
      jsonpath_root       = "$.data"
      success_codes       = "200"
      not_retryable_codes = "400,401,403,404"
    },
    {
      name                = "list_users"
      path                = "/users/all"
      method              = "GET"
      jsonpath_root       = "$.data"
      success_codes       = "200"
      not_retryable_codes = "400,401,403,404"
    },
  ]
}
