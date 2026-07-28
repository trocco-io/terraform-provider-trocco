resource "trocco_dbt_job_definition" "test" {
  name                  = "dbt-job-conflict"
  dbt_git_repository_id = 1

  bigquery_setting = {
    connection_id = 1
    dataset       = "analytics"
  }

  snowflake_setting = {
    connection_id = 2
    warehouse     = "COMPUTE_WH"
    database      = "ANALYTICS"
    schema        = "PUBLIC"
  }
}
