resource "trocco_team" "test" {
  name        = "updated"
  description = "updated"
  members = [
    {
      user_id = 10626
      role    = "team_admin"
    },
    {
      user_id = 10652
      role    = "team_member"
    }
  ]
}