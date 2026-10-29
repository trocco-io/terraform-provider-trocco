# Same as kintone_client_certificate_create.tf with only the description changed.
# The certificate is re-sent on an unrelated update and the API must accept the same one again.
resource "trocco_connection" "kintone_client_certificate_test" {
  connection_type             = "kintone"
  name                        = "Kintone Client Certificate Test"
  description                 = "This is a Kintone connection example (updated)"
  domain                      = "test_domain.s.cybozu.com"
  login_method                = "username_and_password"
  password                    = "test_password"
  username                    = "test_username"
  client_certificate          = "%s"
  client_certificate_password = "test_pfx_password"
}
