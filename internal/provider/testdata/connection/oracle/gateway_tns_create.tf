resource "trocco_connection" "oracle_gateway_tns_test" {
  connection_type = "oracle"
  name            = "Test Oracle Gateway + TNS Connection"
  user_name       = "test_user"
  password        = "test_password"
  driver          = "19c-ojdbc8.jar"
  tns_admin_ora   = <<-EOT
    ORCL_high = (DESCRIPTION =
      (ADDRESS = (PROTOCOL = TCPS)(HOST = adb.example.oraclecloud.com)(PORT = 1522))
      (CONNECT_DATA = (SERVICE_NAME = orcl_high.adb.oraclecloud.com))
    )
  EOT

  gateway = {
    host      = "bastion.example.com"
    port      = 22
    user_name = "ec2-user"
    password  = "gateway_password"
  }
}
