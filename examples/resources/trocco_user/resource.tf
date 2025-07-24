resource "trocco_user" "example_with_password" {
  email                           = "trocco@example.com"
  password                        = "Jb1p4f1uuC"
  password_auto_generated         = false
  role                            = "member"
  can_use_audit_log               = false
  is_restricted_connection_modify = false
}

resource "trocco_user" "example_with_auto_generated_password" {
  email                           = "auto-generated@example.com"
  password_auto_generated         = true
  role                            = "member"
  can_use_audit_log               = false
  is_restricted_connection_modify = false
}
