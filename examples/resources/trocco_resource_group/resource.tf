resource "trocco_resource_group" "example" {
  name        = "resource group name"
  description = "description"
  teams = [
    {
      team_id = 1
      role    = "administrator"
    },
    {
      team_id = 2
      role    = "operator"
    }
  ]
}

