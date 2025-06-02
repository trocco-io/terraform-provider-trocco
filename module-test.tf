locals {
  test_members = [
    {
      user_id = 2
      role    = "team_admin"
    },
    {
      user_id = 1
      role    = "team_admin"
    }
  ]
}
module "module_sample"{
  source = "./modules/team"
  team_name        = "test"
  team_description = "module test"
  team_members = local.test_members
}

output "debug_print" {
  description = "Debug output to verify team creation"
  value       = module.module_sample.debug_print
}
