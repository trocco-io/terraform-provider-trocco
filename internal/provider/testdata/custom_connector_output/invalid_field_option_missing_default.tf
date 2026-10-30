resource "trocco_custom_connector_output" "test" {
  name      = "custom-connector-output-invalid-field-option"
  url       = "https://example.com"
  auth_type = "api_key"

  endpoints = [
    {
      name      = "create_user"
      operation = "create"
      path      = "/users"
      template  = "{\"name\": \"{{ name }}\"}"

      query_parameters = [
        {
          name         = "dry_run"
          display_name = "Dry Run"
          is_editable  = false
          is_required  = true
        },
      ]
    },
  ]
}
