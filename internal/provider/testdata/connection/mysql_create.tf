resource "trocco_connection" "mysql_test" {
  connection_type = "mysql"

  name      = "mysql test"
  host      = "localhost"
  user_name = "root"
  password  = "password"
  port      = 3306
}
