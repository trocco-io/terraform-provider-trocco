resource "trocco_connection" "mongodb" {
  connection_type          = "mongodb"
  name                     = "MongoDB Example"
  description              = "This is a MongoDB connection example"
  host                     = "mongodb.example.com"
  port                     = 27017
  user_name                = "admin"
  password                 = "password"
  connection_string_format = "standard"
  auth_method              = "scram-sha-1"
  auth_source              = "admin"
  read_preference          = "primary"
  gateway = {
    host      = "gateway.example.com"
    port      = 22
    user_name = "gateway-user"
    password  = "gateway-password"
    # Or use SSH key authentication instead of password:
    key            = <<-GATEWAY_KEY
      -----BEGIN PRIVATE KEY-----
      ... GATEWAY KEY...
      -----END PRIVATE KEY-----
    GATEWAY_KEY
    key_passphrase = "sample_passphrase"
  }

  resource_group_id = 1
}
