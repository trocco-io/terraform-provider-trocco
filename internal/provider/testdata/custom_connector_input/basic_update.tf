resource "trocco_custom_connector_input" "test" {
  name        = "custom-connector-input-test"
  description = "updated description"
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
      success_codes       = "200,201"
      not_retryable_codes = "400,401,403,404"
      request_timeout_sec = 90

      query_parameters = [
        {
          name          = "limit"
          display_name  = "Limit"
          default_value = "100"
          is_editable   = true
          is_required   = true
        },
      ]

      headers = [
        {
          name          = "X-Api-Version"
          display_name  = "API Version"
          default_value = "2"
          is_editable   = false
          is_required   = true
        },
      ]

      path_parameters = []

      request_body = "{\"status\": \"{status}\"}"

      request_body_parameters = [
        {
          name          = "status"
          display_name  = "Status"
          default_value = "inactive"
          is_editable   = false
          is_required   = true
        },
      ]

      paginator = {
        inject_into = "query"

        offset_increment_strategy = {
          page_size         = 50
          start_from_offset = 0
          stop_on_offset    = 500
        }

        page_size_option = {
          field_name = "limit"
        }
      }
    },

    # Kept minimal across the update as well, to confirm the defaults are sent
    # again rather than nulled out on PATCH.
    {
      name = "list_minimal"
      path = "/minimal"
    },
  ]
}
