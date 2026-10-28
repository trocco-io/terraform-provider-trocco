# The certificate is a self-signed test PFX (password: test_pfx_password) generated with openssl.
# The client_certificate placeholder is filled by the test from kintone_client_certificate_with_password.pfx.b64.
resource "trocco_connection" "kintone_client_certificate_test" {
  connection_type             = "kintone"
  name                        = "Kintone Client Certificate Test"
  description                 = "This is a Kintone connection example"
  domain                      = "test_domain.s.cybozu.com"
  login_method                = "username_and_password"
  password                    = "test_password"
  username                    = "test_username"
  client_certificate          = "%s"
  client_certificate_password = "test_pfx_password"
}
