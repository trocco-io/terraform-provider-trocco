resource "trocco_connection" "oracle_gateway_test" {
  connection_type = "oracle"
  name            = "Test Oracle Gateway Connection"
  host            = "oracle-internal.example.local"
  port            = 1521
  user_name       = "test_user"
  password        = "test_password"
  driver          = "19c-ojdbc8.jar"

  gateway = {
    host      = "bastion.example.com"
    port      = 22
    user_name = "ec2-user"
    password  = "gateway_password"
  }
}
