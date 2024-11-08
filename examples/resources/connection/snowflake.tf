resource "trocco_connection" "snowflake" {
  connection_type = "snowflake"

  name        = "Snowflake Example"
  description = "This is a Snowflake connection example"

  host        = "exmaple.snowflakecomputing.com"
  auth_method = "user_password"
  user_name   = "<User Name>"
  password    = "<Password>"
}
