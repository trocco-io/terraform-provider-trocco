# Basic Redshift connection
resource "trocco_connection" "redshift_basic" {
  connection_type       = "redshift"
  name                  = "Redshift Basic Example"
  description           = "This is a basic Redshift connection example"
  resource_group_id     = 1
  host                  = "my-cluster.abc123.us-east-1.redshift.amazonaws.com"
  port                  = 5439
  user_name             = "admin"
  password              = "password"
  aws_access_key_id     = "AKIAIOSFODNN7EXAMPLE"
  aws_secret_access_key = "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY"
  ssl_enabled           = true
}

# Redshift connection with gateway (SSH tunnel)
resource "trocco_connection" "redshift_gateway" {
  connection_type       = "redshift"
  name                  = "Redshift Gateway Example"
  description           = "This is a Redshift connection example with SSH gateway"
  resource_group_id     = 1
  host                  = "my-cluster.abc123.us-east-1.redshift.amazonaws.com"
  port                  = 5439
  user_name             = "admin"
  password              = "password"
  aws_access_key_id     = "AKIAIOSFODNN7EXAMPLE"
  aws_secret_access_key = "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY"
  ssl_enabled           = true

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
}

# Redshift connection with AWS PrivateLink
resource "trocco_connection" "redshift_privatelink" {
  connection_type       = "redshift"
  name                  = "Redshift PrivateLink Example"
  description           = "This is a Redshift connection example with AWS PrivateLink"
  resource_group_id     = 1
  host                  = "my-cluster.abc123.us-east-1.redshift.amazonaws.com"
  port                  = 5439
  user_name             = "admin"
  password              = "password"
  aws_access_key_id     = "AKIAIOSFODNN7EXAMPLE"
  aws_secret_access_key = "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY"
  ssl_enabled           = true

  ssh_tunnel_id           = 1
  aws_privatelink_enabled = true
}
