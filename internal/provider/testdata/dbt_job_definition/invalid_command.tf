resource "trocco_dbt_job_definition" "test" {
  name                  = "dbt-job-command"
  dbt_git_repository_id = 1

  bigquery_setting = {
    connection_id = 1
    dataset       = "analytics"
  }

  commands = [
    { command = "destroy" },
  ]
}
