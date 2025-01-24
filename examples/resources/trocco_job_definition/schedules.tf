resource "trocco_job_definition" "schedules" {
  schedules = [
    {
      day       = 1
      frequency = "monthly"
      hour      = 1
      minute    = 1
      time_zone = "Australia/Sydney"
    },
    {
      day_of_week = 0
      frequency   = "weekly"
      hour        = 1
      minute      = 1
      time_zone   = "Australia/Sydney"
    },
    {
      frequency = "daily"
      hour      = 1
      minute    = 1
      time_zone = "Australia/Sydney"
    },
    {
      frequency = "hourly"
      minute    = 1
      time_zone = "Australia/Sydney"
    },
  ]
}
