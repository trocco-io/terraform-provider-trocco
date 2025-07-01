resource "trocco_connection" "snowflake_test" {
  connection_type = "snowflake"
  auth_method     = "user_password"

  name      = "snowflake test"
  host      = "example.snowflakecomputing.com"
  user_name = "root"
  password  = "password"
}