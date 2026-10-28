resource "trocco_connection" "oracle_update_test" {
  connection_type = "oracle"
  name            = "Test Oracle Connection (Before Update)"
  host            = "oracle-before-update.example.com"
  port            = 1521
  user_name       = "test_user"
  password        = "test_password"
  driver          = "19c-ojdbc8.jar"
}
