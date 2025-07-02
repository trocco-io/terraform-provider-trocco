resource "trocco_team" "test" {
  name        = "test"
  description = "test"
  members = [
    {
      user_id = 10626
      role    = "team_admin"
    }
  ]
}
