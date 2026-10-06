resource "trocco_connection" "invalid_custom_connector_id_zero_test" {
  connection_type     = "custom_connector"
  name                = "Custom Connector Test"
  custom_connector_id = 0
  api_key             = "test-api-key"
}
