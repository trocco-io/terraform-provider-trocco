resource "trocco_custom_connector_input" "example" {
  name        = "My API Connector"
  description = "Example custom connector (transfer source)"
  url         = "https://api.example.com"
  auth_type   = "api_key"

  auth_header_name   = "Authorization"
  auth_header_scheme = "Bearer"

  endpoints = [
    {
      # page_increment_strategy: the server requires `start_from_page` and
      # exactly one of `last_page_size` / `total_pages` / `stop_on_page`.
      name                = "list_users"
      path                = "/users"
      method              = "GET"
      jsonpath_root       = "$.data"
      success_codes       = "200"
      not_retryable_codes = "400,401,403,404"
      request_timeout_sec = 60 # optional, defaults to 30 (1-1800)

      query_parameters = [
        {
          name          = "limit"
          display_name  = "Limit"
          default_value = "100"
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
    {
      # page_increment_strategy alternative: stop via `last_page_size`
      # (requires `max_request_count`) instead of `stop_on_page`.
      name                = "list_products"
      path                = "/products"
      method              = "GET"
      jsonpath_root       = "$.data"
      success_codes       = "200"
      not_retryable_codes = "400,401,403,404"

      paginator = {
        inject_into = "query"

        page_increment_strategy = {
          page_size         = 100
          start_from_page   = 1
          last_page_size    = "$.meta.last_page_size"
          max_request_count = 100
        }

        page_token_option = {
          field_name = "page"
        }
      }
    },
    {
      # page_increment_strategy alternative: stop via `total_pages`.
      #
      # Only `name` and `path` are required. `method`, `jsonpath_root`,
      # `success_codes`, `not_retryable_codes` and `request_timeout_sec` are
      # omitted here and fall back to the defaults shown on the other
      # endpoints (`GET`, `$.*`, `200`, `400,401,403,404` and `30`).
      name = "list_categories"
      path = "/categories"

      paginator = {
        inject_into = "query"

        page_increment_strategy = {
          page_size       = 100
          start_from_page = 1
          total_pages     = "$.meta.total_pages"
        }

        page_token_option = {
          field_name = "page"
        }
      }
    },
    {
      # offset_increment_strategy: the server requires `start_from_offset` and
      # exactly one of `last_page_size` / `total_records` / `stop_on_offset`.
      name                = "list_orders"
      path                = "/orders"
      method              = "GET"
      jsonpath_root       = "$.data"
      success_codes       = "200"
      not_retryable_codes = "400,401,403,404"

      paginator = {
        inject_into = "query"

        offset_increment_strategy = {
          page_size         = 100
          start_from_offset = 0
          stop_on_offset    = 1000
        }

        page_size_option = {
          field_name = "limit"
        }
        page_token_option = {
          field_name = "offset"
        }
      }
    },
    {
      # offset_increment_strategy alternative: stop via `last_page_size`
      # (requires `max_request_count`) instead of `stop_on_offset`.
      name                = "list_invoices"
      path                = "/invoices"
      method              = "GET"
      jsonpath_root       = "$.data"
      success_codes       = "200"
      not_retryable_codes = "400,401,403,404"

      paginator = {
        inject_into = "query"

        offset_increment_strategy = {
          page_size         = 100
          start_from_offset = 0
          last_page_size    = "$.meta.last_page_size"
          max_request_count = 100
        }

        page_size_option = {
          field_name = "limit"
        }
        page_token_option = {
          field_name = "offset"
        }
      }
    },
    {
      # offset_increment_strategy alternative: stop via `total_records`.
      name                = "list_payments"
      path                = "/payments"
      method              = "GET"
      jsonpath_root       = "$.data"
      success_codes       = "200"
      not_retryable_codes = "400,401,403,404"

      paginator = {
        inject_into = "query"

        offset_increment_strategy = {
          page_size         = 100
          start_from_offset = 0
          total_records     = "$.meta.total_records"
        }

        page_size_option = {
          field_name = "limit"
        }
        page_token_option = {
          field_name = "offset"
        }
      }
    },
    {
      # `request_body_parameters` declares the `{placeholder}` names that
      # `request_body` may contain; a transfer setting fills in their values.
      name                = "search_users"
      path                = "/users/search"
      method              = "POST"
      jsonpath_root       = "$.data"
      success_codes       = "200"
      not_retryable_codes = "400,401,403,404"

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
    },
    {
      # cursor_based_strategy: the server requires `cursor_value`.
      name                = "list_events"
      path                = "/events"
      method              = "GET"
      jsonpath_root       = "$.data"
      success_codes       = "200"
      not_retryable_codes = "400,401,403,404"

      paginator = {
        inject_into = "query"

        cursor_based_strategy = {
          page_size    = 100
          cursor_value = "$.meta.next_cursor"
        }

        page_token_option = {
          field_name = "cursor"
        }
      }
    },
  ]
}

# OAuth2 with `client_credentials` works end-to-end via Terraform/API.
resource "trocco_custom_connector_input" "oauth2_example" {
  name      = "OAuth2 Connector"
  url       = "https://api.example.com"
  auth_type = "oauth2"

  grant_type       = "client_credentials"
  access_token_uri = "https://auth.example.com/token"

  endpoints = [
    {
      name                = "list_items"
      path                = "/items"
      method              = "GET"
      jsonpath_root       = "$.items"
      success_codes       = "200"
      not_retryable_codes = "400,401,403,404"
    },
  ]
}
