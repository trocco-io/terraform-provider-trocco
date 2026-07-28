resource "trocco_dbt_job_definition" "test" {
  name                  = "dbt-job-threads"
  dbt_git_repository_id = 1
  threads               = 99

  bigquery_setting = {
    connection_id = 1
    dataset       = "analytics"
  }
}
