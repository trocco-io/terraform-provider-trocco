resource "trocco_connection" "sftp_test" {
  connection_type = "sftp"

  name      = "sftp test"
  host      = "sftp.example.com"
  user_name = "testuser"
  password  = "password"
  port      = 22
}
