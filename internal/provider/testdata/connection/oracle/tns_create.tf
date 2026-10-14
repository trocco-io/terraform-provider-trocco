resource "trocco_connection" "oracle_tns_test" {
  connection_type = "oracle"
  name            = "Test Oracle TNS Connection"
  user_name       = "test_user"
  password        = "test_password"
  tns_admin_ora   = <<-EOT
    ORCL_high = (DESCRIPTION =
      (ADDRESS = (PROTOCOL = TCPS)(HOST = adb.example.oraclecloud.com)(PORT = 1522))
      (CONNECT_DATA = (SERVICE_NAME = orcl_high.adb.oraclecloud.com))
      (SECURITY = (SSL_SERVER_DN_MATCH = yes))
    )
  EOT
  driver          = "19c-ojdbc8.jar"
}
