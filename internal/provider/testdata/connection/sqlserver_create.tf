resource "trocco_connection" "sqlserver_test" {
  connection_type = "sqlserver"

  name        = "sqlserver test"
  host        = "localhost"
  user_name   = "sa"
  password    = "password"
  port        = 1433
  driver      = "ms_sqlserver_jdbc_driver_13_4"
  ssl_enabled = true
}
