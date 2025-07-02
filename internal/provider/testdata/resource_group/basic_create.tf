resource "trocco_team" "team1" {
  name = "test"
  members = [
    {
      user_id = 10626
      role    = "team_admin"
    }
  ]
}

resource "trocco_resource_group" "test" {
  name        = "test"
  description = "test"
  teams = [
    {
      team_id = trocco_team.team1.id
      role    = "operator"
    }
  ]
}
