resource "trocco_connection" "marketo" {
  connection_type = "marketo"

  name               = "Test Marketo Connection"
  description        = "Test Marketo connection for API operations"
  account_id         = "123-ABC-456"
  client_id          = "client_test_123"
  client_secret      = "secret_test_123"
  api_max_call_count = 5000
}
