resource "trocco_connection" "oracle_update_test" {
  connection_type = "oracle"
  name            = "Test Oracle Connection (Updated to TNS)"
  user_name       = "test_user"
  password        = "test_password"
  driver          = "26ai-ojdbc8.jar"
  tns_admin_ora   = <<-EOT
    ORCL_high = (DESCRIPTION =
      (ADDRESS = (PROTOCOL = TCPS)(HOST = adb.example.oraclecloud.com)(PORT = 1522))
      (CONNECT_DATA = (SERVICE_NAME = orcl_high.adb.oraclecloud.com))
    )
  EOT
}
