resource "trocco_custom_connector_input" "test" {
  name      = "custom-connector-input-conflict"
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
        page_increment_strategy = {
          page_size       = 100
          start_from_page = 1
          stop_on_page    = 10
        }

        offset_increment_strategy = {
          page_size         = 100
          start_from_offset = 0
          stop_on_offset    = 1000
        }
      }
    },
  ]
}
