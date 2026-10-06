resource "trocco_custom_connector_output" "example" {
  name        = "My Output Connector"
  description = "Example custom connector (transfer destination)"
  url         = "https://api.example.com"
  auth_type   = "api_key"

  auth_header_name   = "Authorization"
  auth_header_scheme = "Bearer"

  endpoints = [
    {
      # A job definition references one `create` endpoint, and one `update`
      # endpoint when the transfer runs in upsert mode.
      name                = "create_user"
      request_type        = "single"
      method              = "POST"
      operation           = "create"
      path                = "/users"
      payload_type        = "json"
      success_codes       = "200,201,204"
      not_retryable_codes = "400,401,403,404"
      request_timeout_sec = 30 # optional, defaults to 30 (1-60)
      template            = "{\"name\": \"{{ name }}\", \"email\": \"{{ email }}\"}"

      headers = [
        {
          name          = "X-Api-Version"
          display_name  = "API Version"
          default_value = "2"
          is_editable   = false
          is_required   = true
        },
      ]
    },
    {
      # `path_parameters.value` must appear in `path` as `{value}`, so the
      # Liquid placeholder is written without surrounding whitespace.
      name                = "update_user"
      request_type        = "single"
      method              = "PATCH"
      operation           = "update"
      path                = "/users/{{user_id}}"
      payload_type        = "json"
      success_codes       = "200,204"
      not_retryable_codes = "400,401,403,404"
      template            = "{\"name\": \"{{ name }}\"}"

      path_parameters = [
        {
          name  = "user_id"
          value = "user_id"
        },
      ]

      query_parameters = [
        {
          name          = "dry_run"
          display_name  = "Dry Run"
          default_value = "false"
          is_editable   = true
          is_required   = false
        },
      ]
    },
    {
      # `request_type = "multiple"` sends records in batches; `batch_size` must
      # stay 1 for `single`.
      name                = "bulk_create_users"
      request_type        = "multiple"
      batch_size          = 100
      method              = "POST"
      operation           = "create"
      path                = "/users/bulk"
      payload_type        = "json"
      success_codes       = "200,201,204"
      not_retryable_codes = "400,401,403,404"
      template            = "{\"users\": [{% for record in records %}{\"name\": \"{{ record.name }}\"}{% unless forloop.last %},{% endunless %}{% endfor %}]}"
    },
  ]
}

# OAuth2 with `client_credentials` works end-to-end via Terraform/API.
resource "trocco_custom_connector_output" "oauth2_example" {
  name      = "OAuth2 Output Connector"
  url       = "https://api.example.com"
  auth_type = "oauth2"

  grant_type       = "client_credentials"
  access_token_uri = "https://auth.example.com/token"

  endpoints = [
    {
      name      = "create_item"
      method    = "POST"
      operation = "create"
      path      = "/items"
      template  = "{\"name\": \"{{ name }}\"}"
    },
  ]
}
