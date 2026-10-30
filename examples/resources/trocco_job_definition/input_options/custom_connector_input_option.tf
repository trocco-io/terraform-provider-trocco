resource "trocco_job_definition" "custom_connector_input_example" {
  input_option_type = "custom_connector"
  input_option = {
    custom_connector_input_option = {
      # Reference the endpoint through the definition resource so Terraform
      # orders operations correctly (e.g. it won't delete an endpoint that
      # this job definition still depends on).
      custom_connector_endpoint_id   = trocco_custom_connector_input.example.endpoints[0].id
      custom_connector_connection_id = trocco_connection.custom_connector_example.id

      # Omitting `query_parameters`/`headers` keeps their existing values;
      # specifying `[]` clears them; specifying values fully replaces them.
      query_parameters = [
        {
          name  = "limit"
          value = "50"
        },
      ]

      # request_body_parameters fill in the `{name}` placeholders declared by
      # the endpoint's request_body_parameters.
      request_body_parameters = [
        {
          name  = "status"
          value = "pending"
        },
      ]

      # path_parameters must include every placeholder in the endpoint's
      # path, since it's not partially updatable.
      path_parameters = [
        {
          name  = "user_id"
          value = "12345"
        },
      ]

      jsonpath_parser = {
        columns = [
          {
            name = "id"
            type = "long"
          },
          {
            name = "name"
            type = "string"
          },
        ]
      }
    }
  }
}
