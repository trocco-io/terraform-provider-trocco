resource "trocco_connection" "kintone_test" {
  connection_type         = "kintone"
  name                    = "Kintone Test"
  description             = "This is a Kintone connection example"
  domain                  = "test_domain"
  login_method            = "username_and_password"
  password                = "test_password"
  username                = "test_username"
  token                   = null
  basic_auth_username     = "test_basic_auth_username"
  basic_auth_password     = "test_basic_auth_password"
}