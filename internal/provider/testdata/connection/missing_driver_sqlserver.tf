resource "trocco_connection" "missing_driver_test_sqlserver" {
  connection_type = "sqlserver"
  name            = "missing driver test"
  host            = "localhost"
  user_name       = "sa"
  password        = "password"
  port            = 1433
}
