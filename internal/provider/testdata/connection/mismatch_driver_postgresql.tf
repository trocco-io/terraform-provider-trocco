resource "trocco_connection" "mismatch_driver_test_postgresql" {
  connection_type = "postgresql"
  name            = "invalid driver test"
  host            = "localhost"
  user_name       = "root"
  password        = "password"
  port            = 5432
  driver          = "mysql_connector_java_5_1_49"
}