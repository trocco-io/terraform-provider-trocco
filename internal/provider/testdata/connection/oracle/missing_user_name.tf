resource "trocco_connection" "oracle_missing_user_name_test" {
  connection_type = "oracle"
  name            = "Test Oracle Missing user_name Connection"
  host            = "oracle.example.com"
  port            = 1521
  password        = "test_password"
}
