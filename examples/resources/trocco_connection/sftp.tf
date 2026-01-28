# SFTP connection with password authentication
resource "trocco_connection" "sftp_password" {
  connection_type = "sftp"

  name              = "SFTP Example with Password"
  description       = "This is an SFTP connection example using password"
  resource_group_id = 1

  host                   = "sftp.example.com"
  port                   = 22
  user_name              = "sftpuser"
  password               = "password123"
  user_directory_is_root = true
  windows_server         = false
}

# SFTP connection with private key authentication
resource "trocco_connection" "sftp_key" {
  connection_type = "sftp"

  name              = "SFTP Example with Private Key"
  description       = "This is an SFTP connection example using private key"
  resource_group_id = 1

  host                   = "sftp.example.com"
  port                   = 22
  user_name              = "sftpuser"
  secret_key             = <<-EOT
    -----BEGIN RSA PRIVATE KEY-----
    ...Your RSA Private Key...
    -----END RSA PRIVATE KEY-----
  EOT
  secret_key_passphrase  = "passphrase123"
  user_directory_is_root = false
  windows_server         = false
}

# SFTP connection with SSH tunnel
resource "trocco_connection" "sftp_tunnel" {
  connection_type = "sftp"

  name              = "SFTP Example with SSH Tunnel"
  description       = "This is an SFTP connection example using SSH tunnel"
  resource_group_id = 1

  host                   = "sftp.example.com"
  port                   = 22
  user_name              = "sftpuser"
  password               = "password123"
  ssh_tunnel_id          = 1
  user_directory_is_root = true
  windows_server         = false
}

# SFTP connection for Windows server
resource "trocco_connection" "sftp_windows" {
  connection_type = "sftp"

  name        = "SFTP Example for Windows Server"
  description = "This is an SFTP connection example for Windows server"

  host                   = "sftp.windows-server.com"
  port                   = 22
  user_name              = "administrator"
  password               = "windows_password"
  user_directory_is_root = false
  windows_server         = true
}

# SFTP connection with AWS PrivateLink
resource "trocco_connection" "sftp_privatelink" {
  connection_type = "sftp"

  name              = "SFTP Example with AWS PrivateLink"
  description       = "This is an SFTP connection example using AWS PrivateLink"
  resource_group_id = 1

  host                    = "sftp.example.com"
  port                    = 22
  user_name               = "sftpuser"
  password                = "password123"
  aws_privatelink_enabled = true
  ssh_tunnel_id           = 1
  user_directory_is_root  = true
  windows_server          = false
}
