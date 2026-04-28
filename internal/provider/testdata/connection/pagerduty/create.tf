resource "trocco_connection" "pagerduty_test" {
  connection_type = "pagerduty"
  name            = "Test Pagerduty Connection"
  api_key         = "test-api_key"
}
