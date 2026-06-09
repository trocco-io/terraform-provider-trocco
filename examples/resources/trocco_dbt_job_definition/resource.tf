resource "trocco_dbt_job_definition" "example" {
  name                  = "example dbt job"
  description           = "daily dbt run"
  dbt_git_repository_id = trocco_dbt_git_repository.example.id
  threads               = 4
  target                = "prod"

  bigquery_setting = {
    connection_id = trocco_connection.bigquery.id
    dataset       = "analytics"
    location      = "asia-northeast1"
  }

  commands = [
    {
      command = "run"
      options = [{ key = "--vars", value = "ds=2026-05-20" }]
    },
    { command = "test" },
  ]

  custom_variable_settings = [
    {
      name  = "$ds$"
      type  = "string"
      value = "2026-05-20"
    },
  ]
}
