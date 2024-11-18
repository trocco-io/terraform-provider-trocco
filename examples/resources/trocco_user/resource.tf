resource "trocco_user" "example" {
  email                           = "trocco@example.com"
  password                        = "Jb1p4f1uuC"
  role                            = "member"
  can_use_audit_log               = false
  is_restricted_connection_modify = false
}
