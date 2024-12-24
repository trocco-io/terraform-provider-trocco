resource "trocco_team" "example" {
  name        = "team name"
  description = "description"
  members = [
    {
      user_id = 1
      role    = "team_admin"
    },
    {
      user_id = 2
      role    = "team_member"
    }
  ]
}

