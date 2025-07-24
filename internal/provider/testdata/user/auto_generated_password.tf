resource "trocco_user" "test" {
  email                  = "test@example.com"
  role                   = "admin"
  password_auto_generated = true
}
