resource "trocco_connection" "mysql" {
  connection_type        = "mysql"
  name                   = "MySQL Example"
  description            = "This is a MySQL connection example"
  host                   = "db.example.com"
  port                   = 3306
  user_name              = "root"
  password               = "password"
  ssl                    = true
  ssl_ca                 = <<-SSL_CA
    -----BEGIN PRIVATE KEY-----
    ... SSL CA...
    -----END PRIVATE KEY-----
  SSL_CA
  ssl_cert               = <<-SSL_CERT
    -----BEGIN CERTIFICATE-----
    ... SSL CRT...
    -----END CERTIFICATE-----
  SSL_CERT
  ssl_key                = <<-SSL_KEY
    -----BEGIN PRIVATE KEY-----
    ... SSL KEY...
    -----END PRIVATE KEY-----
  SSL_KEY
  gateway_enabled        = true
  gateway_host           = "gateway.example.com"
  gateway_port           = 1234
  gateway_user_name      = "root"
  gateway_password       = "password"
  gateway_key            = <<-GATEWAY_KEY
    -----BEGIN PRIVATE KEY-----
    ... GATEWAY KEY...
    -----END
  GATEWAY_KEY
  gateway_key_passphrase = "sample_passphrase"
  resource_group_id      = 1
}
