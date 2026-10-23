# Removing client_certificate disables secure access; the API deletes the certificate and password.
resource "trocco_connection" "kintone_client_certificate_test" {
  connection_type = "kintone"
  name            = "Kintone Client Certificate Test"
  description     = "This is a Kintone connection example"
  domain          = "test_domain"
  login_method    = "username_and_password"
  password        = "test_password"
  username        = "test_username"
}
