# Basic SQL Server connection
resource "trocco_connection" "sqlserver_basic" {
  connection_type   = "sqlserver"
  name              = "SQL Server Basic Example"
  description       = "This is a basic SQL Server connection example"
  resource_group_id = 1
  host              = "sqlserver.example.com"
  port              = 1433
  user_name         = "sa"
  password          = "password"
  driver            = "ms_sqlserver_jdbc_driver_13_4"
  ssl_enabled       = true
}

# SQL Server connection with gateway (SSH tunnel)
resource "trocco_connection" "sqlserver_gateway" {
  connection_type   = "sqlserver"
  name              = "SQL Server Gateway Example"
  description       = "This is a SQL Server connection example with SSH gateway"
  resource_group_id = 1
  host              = "sqlserver.example.com"
  port              = 1433
  user_name         = "sa"
  password          = "password"
  driver            = "ms_sqlserver_jdbc_driver_13_4"
  ssl_enabled       = false

  gateway = {
    host           = "gateway.example.com"
    port           = 22
    user_name      = "gateway-joe"
    password       = "gateway-joepass"
    key            = <<-GATEWAY_KEY
      -----BEGIN PRIVATE KEY-----
      ... GATEWAY KEY...
      -----END PRIVATE KEY-----
    GATEWAY_KEY
    key_passphrase = "sample_passphrase"
  }
}
