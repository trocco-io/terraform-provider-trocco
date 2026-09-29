# Basic Oracle connection (host/port method with SSL)
resource "trocco_connection" "oracle_basic" {
  connection_type   = "oracle"
  name              = "Oracle Basic Example"
  description       = "This is a basic Oracle connection example using host/port"
  resource_group_id = 1
  host              = "oracle-db.example.com"
  port              = 1521
  user_name         = "admin"
  password          = "MySecurePassword123!"
  ssl_enabled       = true
  driver            = "19c-ojdbc8.jar"

  ssl_ca = <<-EOT
    -----BEGIN CERTIFICATE-----
    ...
    -----END CERTIFICATE-----
  EOT
}

# Oracle connection with TNS naming and Wallet (Autonomous Database)
resource "trocco_connection" "oracle_wallet" {
  connection_type   = "oracle"
  name              = "Oracle Autonomous Database"
  description       = "Oracle Autonomous Database connection using TNS naming and Wallet"
  resource_group_id = 1
  user_name         = "admin"
  password          = "MySecurePassword123!"
  driver            = "19c-ojdbc8.jar"

  tns_admin_ora = <<-EOT
    ORCL_high = (DESCRIPTION =
      (ADDRESS = (PROTOCOL = TCPS)(HOST = adb.example.oraclecloud.com)(PORT = 1522))
      (CONNECT_DATA = (SERVICE_NAME = orcl_high.adb.oraclecloud.com))
      (SECURITY = (SSL_SERVER_DN_MATCH = yes))
    )
  EOT

  wallet_file = filebase64("${path.module}/wallet/cwallet.sso")
}

# Oracle connection with SSH tunnel (gateway)
resource "trocco_connection" "oracle_gateway" {
  connection_type   = "oracle"
  name              = "Oracle SSH Tunnel Example"
  description       = "Oracle connection with SSH tunnel"
  resource_group_id = 1
  host              = "oracle-internal.example.local"
  port              = 1521
  user_name         = "admin"
  password          = "MySecurePassword123!"
  driver            = "19c-ojdbc8.jar"

  gateway = {
    host           = "bastion.example.com"
    port           = 22
    user_name      = "ec2-user"
    key            = <<-EOT
      -----BEGIN OPENSSH PRIVATE KEY-----
      ...
      -----END OPENSSH PRIVATE KEY-----
    EOT
    key_passphrase = ""
  }
}

# Oracle connection with AWS PrivateLink
resource "trocco_connection" "oracle_privatelink" {
  connection_type   = "oracle"
  name              = "Oracle PrivateLink Example"
  description       = "Oracle connection via AWS PrivateLink"
  resource_group_id = 1
  host              = "oracle-internal.example.local"
  port              = 1521
  user_name         = "admin"
  password          = "MySecurePassword123!"
  driver            = "19c-ojdbc8.jar"

  aws_privatelink_enabled = true
  ssh_tunnel_id           = 1
}
