resource "trocco_team" "this" {
  name        = var.team_name
  description = var.team_description
  members = [for member in var.team_members : {
    user_id = member.user_id
    role    = member.role
  }]
}

output "debug_print" {
  value = "${var.team_name} created with members: ${join(", ", [for member in var.team_members : member.user_id])}"
}
