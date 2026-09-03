resource "trocco_custom_connector_input" "test" {
  name        = "custom-connector-input-test"
  description = "test custom connector input"
  url         = "https://example.com"
  auth_type   = "api_key"

  auth_header_name   = "Authorization"
  auth_header_scheme = "Bearer"

  endpoints = [
    {
      name                = "list_users"
      path                = "/users"
      method              = "GET"
      jsonpath_root       = "$.data"
      success_codes       = "200"
      not_retryable_codes = "400,401,403,404"
      request_timeout_sec = 60

      query_parameters = [
        {
          name          = "limit"
          display_name  = "Limit"
          default_value = "100"
          is_editable   = true
          is_required   = true
        },
      ]

      headers         = []
      path_parameters = []

      request_body = "{\"status\": \"{status}\"}"

      request_body_parameters = [
        {
          name          = "status"
          display_name  = "Status"
          default_value = "active"
          is_editable   = true
          is_required   = true
        },
      ]

      paginator = {
        inject_into = "query"

        page_increment_strategy = {
          page_size       = 100
          start_from_page = 1
          stop_on_page    = 10
        }

        page_token_option = {
          field_name = "page"
        }
      }
    },

    # Only the attributes without a server-side default are set here. The rest
    # fall back to the schema defaults, which mirror the server's, so the
    # request still carries a value for every attribute the server assigns
    # unconditionally.
    {
      name = "list_minimal"
      path = "/minimal"
    },
  ]
}
