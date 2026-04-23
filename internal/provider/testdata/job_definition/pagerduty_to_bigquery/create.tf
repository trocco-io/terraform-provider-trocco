resource "trocco_connection" "bigquery" {
  connection_type          = "bigquery"
  name                     = "BigQuery Example"
  project_id               = "example"
  service_account_json_key = jsonencode({
    type       = "service_account"
    project_id = "example"
  })
}

resource "trocco_connection" "pagerduty_test" {
  connection_type = "pagerduty"
  name            = "Test Pagerduty Connection"
  api_key = "test-api_key"
}

resource "trocco_job_definition" "pagerduty_to_bigquery" {
  name                     = "Pagerduty to BigQuery Test"
  description              = "Test job definition for transferring data from Pagerduty to BigQuery"
  resource_enhancement     = "medium"
  retry_limit              = 0
  is_runnable_concurrently = false

  filter_columns = []

  input_option_type = "pagerduty"

  input_option = {
    pagerduty_input_option = {
      pagerduty_connection_id = trocco_connection.pagerduty_test.id
      path                   = "escalation_policies"
    }
  }

  output_option_type = "bigquery"

  output_option = {
    bigquery_output_option = {
      bigquery_connection_id = trocco_connection.bigquery.id
      dataset                = "test_dataset"
      table                  = "pagerduty_to_bigquery_test_table"
      mode                   = "append"
    }
  }
}
