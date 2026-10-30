resource "trocco_custom_connector_output" "test" {
  name        = "custom-connector-output-test"
  description = "updated description"
  url         = "https://example.com"
  auth_type   = "api_key"

  auth_header_name   = "Authorization"
  auth_header_scheme = "Bearer"

  endpoints = [
    {
      name                = "create_user"
      request_type        = "multiple"
      batch_size          = 50
      method              = "POST"
      operation           = "create"
      path                = "/users/bulk"
      payload_type        = "json"
      success_codes       = "200,201"
      not_retryable_codes = "400,401,403,404"
      request_timeout_sec = 60
      template            = "{\"users\": [{\"name\": \"{{ name }}\"}]}"

      headers = [
        {
          name          = "X-Api-Version"
          display_name  = "API Version"
          default_value = "3"
          is_editable   = false
          is_required   = true
        },
      ]

      path_parameters = []

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
      name                = "update_user"
      request_type        = "single"
      method              = "PATCH"
      operation           = "update"
      path                = "/users/{{user_id}}"
      payload_type        = "json"
      success_codes       = "200,204"
      not_retryable_codes = "400,401,403,404"
      template            = "{\"name\": \"{{ name }}\"}"

      headers = []

      path_parameters = [
        {
          name  = "user_id"
          value = "user_id"
        },
      ]

      query_parameters = []
    },
  ]
}
