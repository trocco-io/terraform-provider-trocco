resource "trocco_connection" "marketo" {
  connection_type = "marketo"

  name        = "Marketo Example"
  description = "This is a Marketo connection example"

  account_id         = "123-ABC-456"
  client_id          = "client_xyz789"
  client_secret      = "abcdef1234567890xyz"
  api_max_call_count = 5000
}
