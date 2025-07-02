resource "trocco_connection" "invalid_login_method_test" {
  connection_type     = "kintone"
  name                = "Kintone Test"
  description         = "This is a Kintone connection example"
  domain              = "test_domain"
  login_method        = "invalid_login_method"
  password            = "test_password"
  username            = "test_username"
  token               = null
  basic_auth_username = "test_basic_auth_username"
  basic_auth_password = "test_basic_auth_password"
}
