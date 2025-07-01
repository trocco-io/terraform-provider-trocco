resource "trocco_connection" "invalid_driver_test" {
  connection_type = "postgresql"
  name            = "invalid driver test"
  host            = "localhost"
  user_name       = "root"
  password        = "password"
  port            = 5432
  driver          = "invalid_driver"
}