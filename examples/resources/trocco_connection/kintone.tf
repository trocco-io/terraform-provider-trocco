# login_method: token
resource "trocco_connection" "kintone_login_method_token" {
  connection_type     = "kintone"
  name                = "Kintone Example"
  description         = "This is a Kintone connection example"
  resource_group_id   = 1
  domain              = "test_domain"
  login_method        = "token"
  token               = "token"
  username            = nil
  password            = nil
  basic_auth_username = "basic_auth_username"
  basic_auth_password = "basic_auth_password"
}

# login_method: username_and_password
resource "trocco_connection" "kintone_login_method_username_and_password" {
  connection_type     = "kintone"
  name                = "Kintone Example"
  description         = "This is a Kintone connection example"
  resource_group_id   = 1
  domain              = "test_domain"
  login_method        = "username_and_password"
  token               = ""
  username            = "username"
  password            = "password"
  basic_auth_username = "basic_auth_username"
  basic_auth_password = "basic_auth_password"
}

