resource "trocco_job_definition" "custom_connector_output_insert_example" {
  output_option_type = "custom_connector"
  output_option = {
    custom_connector_output_option = {
      custom_connector_connection_id = trocco_connection.custom_connector_example.id

      # Reference the endpoint through the definition resource so Terraform
      # orders operations correctly (e.g. it won't delete an endpoint that
      # this job definition still depends on).
      create_custom_connector_output_endpoint_id = trocco_custom_connector_output.example.endpoints[0].id

      # mode defaults to "insert", where every record is sent to the create
      # endpoint and no update endpoint is involved.

      create_endpoint_settings = {
        # Omitting a collection keeps its existing values; specifying `[]`
        # clears them; specifying values fully replaces them. Only headers and
        # query parameters marked `is_editable = true` on the endpoint can be
        # set here; non-editable ones always keep their defined default.
        query_parameters = [
          {
            name  = "dry_run"
            value = "false"
          },
        ]
      }
    }
  }
}

resource "trocco_job_definition" "custom_connector_output_upsert_example" {
  output_option_type = "custom_connector"
  output_option = {
    custom_connector_output_option = {
      custom_connector_connection_id = trocco_connection.custom_connector_example.id

      # mode = "upsert" sends existing records to the update endpoint instead,
      # and requires update_key and update_custom_connector_output_endpoint_id.
      mode       = "upsert"
      update_key = "id"

      create_custom_connector_output_endpoint_id = trocco_custom_connector_output.example.endpoints[0].id
      update_custom_connector_output_endpoint_id = trocco_custom_connector_output.example.endpoints[1].id

      update_endpoint_settings = {
        # Each path parameter is named after the placeholder declared by the
        # endpoint (its `value`), and holds the column to substitute.
        path_parameters = [
          {
            name  = "user_id"
            value = "id"
          },
        ]
      }
    }
  }
}
