resource "trocco_connection" "pagerduty_example" {
  connection_type = "pagerduty"
  name            = "My Pagerduty Connection"
  description     = "Example Pagerduty connection"
  api_key         = "your-api_key"
}
