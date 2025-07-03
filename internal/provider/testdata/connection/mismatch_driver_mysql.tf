resource "trocco_connection" "mismatch_driver_test_mysql" {
  connection_type = "mysql"
  name            = "invalid driver test"
  host            = "localhost"
  user_name       = "root"
  password        = "password"
  port            = 3306
  driver          = "snowflake_jdbc_3_14_2"
}
