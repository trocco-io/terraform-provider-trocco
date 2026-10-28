# Replaces the certificate with a self-signed test PFX that has no password.
# The client_certificate placeholder is filled by the test from kintone_client_certificate_without_password.pfx.b64.
resource "trocco_connection" "kintone_client_certificate_test" {
  connection_type    = "kintone"
  name               = "Kintone Client Certificate Test"
  description        = "This is a Kintone connection example (updated)"
  domain             = "test_domain.s.cybozu.com"
  login_method       = "username_and_password"
  password           = "test_password"
  username           = "test_username"
  client_certificate = "%s"
}
