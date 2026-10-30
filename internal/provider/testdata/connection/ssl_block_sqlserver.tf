resource "trocco_connection" "ssl_block_test_sqlserver" {
  connection_type = "sqlserver"
  name            = "ssl block test"
  host            = "localhost"
  user_name       = "sa"
  password        = "password"
  port            = 1433
  driver          = "ms_sqlserver_jdbc_driver_13_4"
  ssl = {
    ca = "dummy"
  }
}
