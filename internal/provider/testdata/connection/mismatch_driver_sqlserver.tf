resource "trocco_connection" "mismatch_driver_test_sqlserver" {
  connection_type = "sqlserver"
  name            = "invalid driver test"
  host            = "localhost"
  user_name       = "sa"
  password        = "password"
  port            = 1433
  driver          = "mysql_connector_java_5_1_49"
}
