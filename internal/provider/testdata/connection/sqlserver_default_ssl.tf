resource "trocco_connection" "sqlserver_default_ssl_test" {
  connection_type = "sqlserver"

  name      = "sqlserver default ssl test"
  host      = "localhost"
  user_name = "sa"
  password  = "password"
  port      = 1433
  driver    = "jtds_driver_1_3_1"
}
