resource "trocco_connection" "mismatch_driver_test_snowflake" {
  connection_type = "snowflake"
  name            = "invalid driver test"
  auth_method     = "user_password"
  host            = "example.snowflakecomputing.com"
  user_name       = "root"
  password        = "password"
  driver          = "mysql_connector_java_5_1_49"
}