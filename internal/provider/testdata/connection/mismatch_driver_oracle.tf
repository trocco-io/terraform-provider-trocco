resource "trocco_connection" "mismatch_driver_test_oracle" {
  connection_type = "oracle"
  name            = "invalid driver test"
  host            = "localhost"
  port            = 1521
  user_name       = "root"
  password        = "password"
  driver          = "postgresql_42_5_1"
}
