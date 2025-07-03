resource "trocco_team" "test" {
  name        = "test"
  description = "test"
  members = [
    {
      user_id = 1
      role    = "team_member"
    }
  ]
}
