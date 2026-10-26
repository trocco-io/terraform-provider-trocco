resource "trocco_connection" "kintone_client_certificate_password_without_cert_test" {
  connection_type             = "kintone"
  name                        = "Kintone Test"
  domain                      = "test_domain.s.cybozu.com"
  login_method                = "username_and_password"
  password                    = "test_password"
  username                    = "test_username"
  client_certificate_password = "test_pfx_password"
}
