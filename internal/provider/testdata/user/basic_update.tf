resource "trocco_user" "test" {
  email                           = "test@example.com"
  role                            = "member"
  can_use_audit_log               = true
  is_restricted_connection_modify = true
}
