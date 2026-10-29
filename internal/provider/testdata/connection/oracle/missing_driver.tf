resource "trocco_connection" "oracle_missing_driver_test" {
  connection_type = "oracle"
  name            = "Test Oracle Missing driver Connection"
  host            = "oracle.example.com"
  port            = 1521
  user_name       = "test_user"
  password        = "test_password"
}
