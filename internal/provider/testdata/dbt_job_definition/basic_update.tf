resource "trocco_connection" "bigquery" {
  connection_type          = "bigquery"
  name                     = "BigQuery for dbt job definition test"
  project_id               = "example"
  service_account_json_key = <<JSON
  {
    "type": "service_account",
    "project_id": "example-project-id",
    "private_key_id": "example-private-key-id",
    "private_key":"-----BEGIN PRIVATE KEY-----\n..."
  }
  JSON
}

resource "trocco_dbt_git_repository" "bigquery" {
  name         = "dbt repo for job definition test"
  adapter_type = "bigquery"
  dbt_version  = "1.11"
  url          = "git@github.com:example/dbt-test.git"
  ref_type     = "branch"
  branch       = "main"
}

resource "trocco_dbt_job_definition" "test" {
  name                  = "dbt-job-test-renamed"
  description           = "updated description"
  dbt_git_repository_id = trocco_dbt_git_repository.bigquery.id
  threads               = 8
  target                = "dev"

  bigquery_setting = {
    connection_id = trocco_connection.bigquery.id
    dataset       = "analytics_v2"
    location      = "asia-northeast1"
  }

  commands = [
    { command = "seed" },
    { command = "snapshot" },
  ]

  custom_variable_settings = []
}
