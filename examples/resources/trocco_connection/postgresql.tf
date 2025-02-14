resource "trocco_connection" "postgresql" {
  connection_type = "postgresql"
  name            = "PostgreSQL Example"
  description     = "This is a PostgreSQL connection example"
  host            = "db.example.com"
  port            = 5432
  user_name       = "root"
  password        = "password"
  ssl_mode        = "require"
  driver          = "postgresql_42_5_1"
  ssl = {
    ca   = <<-SSL_CA
      -----BEGIN PRIVATE KEY-----
      ...SSL CA...
      -----END PRIVATE KEY-----
    SSL_CA
    cert = <<-SSL_CERT
      -----BEGIN CERTIFICATE-----
      ...SSL CRT...
      -----END CERTIFICATE-----
    SSL_CERT
    key  = <<-SSL_KEY
      -----BEGIN PRIVATE KEY-----
      ...SSL KEY...
      -----END PRIVATE KEY-----
    SSL_KEY
  }
  gateway = {
    host           = "gateway.example.com"
    port           = 1234
    user_name      = "gateway-joe"
    password       = "gateway-joepass"
    key            = <<-GATEWAY_KEY
      -----BEGIN PRIVATE KEY-----
      ... GATEWAY KEY...
      -----END PRIVATE KEY-----
    GATEWAY_KEY
    key_passphrase = "sample_passphrase"
  }
  resource_group_id = 1
}
