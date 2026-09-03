resource "trocco_custom_connector_input" "test" {
  name      = "custom-connector-input-invalid-path-parameter"
  url       = "https://example.com"
  auth_type = "api_key"

  endpoints = [
    {
      name = "get_user"
      # `{user_id}` is missing from the path, so the server would reject the
      # path parameter below.
      path = "/users"

      path_parameters = [
        {
          name  = "user_id"
          value = "user_id"
        },
      ]
    },
  ]
}
