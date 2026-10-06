resource "trocco_connection" "invalid_custom_connector_missing_id_test" {
  connection_type = "custom_connector"
  name            = "Custom Connector Test"
  api_key         = "test-api-key"
}
