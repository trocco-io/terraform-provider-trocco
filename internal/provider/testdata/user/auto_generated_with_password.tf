resource "trocco_user" "test" {
  email                  = "test@example.com"
  password               = "3XRambMkp-Hw"
  role                   = "admin"
  password_auto_generated = true
}
