resource "trocco_custom_connector_output" "test" {
  name      = "custom-connector-output-invalid-path-parameter"
  url       = "https://example.com"
  auth_type = "api_key"

  endpoints = [
    {
      name      = "update_user"
      method    = "PATCH"
      operation = "update"
      # The server matches `{user_id}` as a substring, so the Liquid
      # placeholder must not contain surrounding whitespace.
      path     = "/users/{{ user_id }}"
      template = "{\"name\": \"{{ name }}\"}"

      path_parameters = [
        {
          name  = "user_id"
          value = "user_id"
        },
      ]
    },
  ]
}
