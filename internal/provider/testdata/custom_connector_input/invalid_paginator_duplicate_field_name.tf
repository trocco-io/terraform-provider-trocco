resource "trocco_custom_connector_input" "test" {
  name      = "custom-connector-input-dup-field-name"
  url       = "https://example.com"
  auth_type = "api_key"

  endpoints = [
    {
      name                = "list_users"
      path                = "/users"
      method              = "POST"
      request_body        = "{}"
      jsonpath_root       = "$.data"
      success_codes       = "200"
      not_retryable_codes = "400,401,403,404"

      paginator = {
        inject_into = "request_body"

        page_increment_strategy = {
          page_size       = 100
          start_from_page = 1
          stop_on_page    = 10
        }

        page_size_option = {
          field_name = "page"
        }
        page_token_option = {
          field_name = "page"
        }
      }
    },
  ]
}
