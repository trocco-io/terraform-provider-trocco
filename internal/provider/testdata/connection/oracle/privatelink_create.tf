resource "trocco_connection" "oracle_privatelink_test" {
  connection_type         = "oracle"
  name                    = "Test Oracle PrivateLink Connection"
  host                    = "oracle-pl.example.com"
  port                    = 1521
  user_name               = "test_user"
  password                = "test_password"
  driver                  = "19c-ojdbc8.jar"
  aws_privatelink_enabled = true
  ssh_tunnel_id           = 1
}
