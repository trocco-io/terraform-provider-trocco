resource "trocco_connection" "missing_driver_test_snowflake" {
  connection_type = "snowflake"
  name            = "missing driver test"
  auth_method     = "user_password"
  host            = "example.snowflakecomputing.com"
  user_name       = "root"
  password        = "password"
}
