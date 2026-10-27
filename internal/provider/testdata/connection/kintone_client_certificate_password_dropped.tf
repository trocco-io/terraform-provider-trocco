# Same certificate as kintone_client_certificate_create.tf (it still needs a password) with
# client_certificate_password dropped, so the API rejects the re-sent certificate with 400.
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
