resource "trocco_connection" "oracle_test" {
  connection_type = "oracle"
  name            = "Test Oracle Connection"
  description     = "Test Oracle connection"
  host            = "oracle.example.com"
  port            = 1521
  user_name       = "test_user"
  password        = "test_password"
  driver          = "19c-ojdbc8.jar"
  ssl_enabled     = true
}
