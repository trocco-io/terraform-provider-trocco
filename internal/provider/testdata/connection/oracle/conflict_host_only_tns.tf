resource "trocco_connection" "oracle_conflict_host_only_test" {
  connection_type = "oracle"
  name            = "Test Oracle Conflict Host Only Connection"
  host            = "oracle.example.com"
  user_name       = "test_user"
  password        = "test_password"
  driver          = "19c-ojdbc8.jar"
  tns_admin_ora   = <<-EOT
    ORCL_high = (DESCRIPTION =
      (ADDRESS = (PROTOCOL = TCPS)(HOST = adb.example.oraclecloud.com)(PORT = 1522))
      (CONNECT_DATA = (SERVICE_NAME = orcl_high.adb.oraclecloud.com))
    )
  EOT
}
