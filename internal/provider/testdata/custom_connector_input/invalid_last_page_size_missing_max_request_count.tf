resource "trocco_custom_connector_input" "test" {
  name      = "custom-connector-input-last-page-size"
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

      paginator = {
        inject_into = "query"

        page_increment_strategy = {
          page_size       = 100
          start_from_page = 1
          last_page_size  = "$.meta.last_page_size"
        }

        page_token_option = {
          field_name = "page"
        }
      }
    },
  ]
}
