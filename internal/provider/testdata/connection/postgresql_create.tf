resource "trocco_connection" "postgresql_test" {
  connection_type = "postgresql"
  name            = "postgresql test"
  host            = "localhost"
  user_name       = "root"
  password        = "password"
  port            = 5432
  driver          = "postgresql_42_5_1"
}