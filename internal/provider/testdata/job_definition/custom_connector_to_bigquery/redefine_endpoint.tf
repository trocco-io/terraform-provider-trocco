resource "trocco_custom_connector_input" "custom_connector_to_bigquery_test" {
  name      = "custom-connector-to-bigquery-test-definition"
  url       = "https://example.com"
  auth_type = "api_key"

  endpoints = [
    {
      name                = "list_users"
      path                = "/users"
      method              = "GET"
      jsonpath_root       = "$.data"
      success_codes       = "200,201"
      not_retryable_codes = "400,401,403,404"
      request_timeout_sec = 60

      query_parameters = [
        {
          name          = "limit"
          display_name  = "Limit"
          default_value = "50"
          is_editable   = true
          is_required   = false
        },
      ]
      headers         = []
      path_parameters = []

      request_body = "{\"status\": \"{status}\"}"

      request_body_parameters = [
        {
          name          = "status"
          display_name  = "Status"
          default_value = "active"
          is_editable   = true
          is_required   = false
        },
      ]

      paginator = {
        inject_into = "query"

        page_increment_strategy = {
          page_size       = 100
          start_from_page = 1
          stop_on_page    = 20
        }

        page_token_option = {
          field_name = "page"
        }
      }
    },
  ]
}

resource "trocco_connection" "custom_connector_to_bigquery_test" {
  connection_type = "custom_connector"

  name                = "Custom Connector to BigQuery Test Connection"
  custom_connector_id = trocco_custom_connector_input.custom_connector_to_bigquery_test.id
  api_key             = "test-api-key"
}

resource "trocco_job_definition" "custom_connector_to_bigquery" {
  name                     = "Custom Connector to BigQuery Test (Redefined)"
  description              = "Test job definition updated in the same apply as the endpoint definition it snapshots"
  resource_enhancement     = "medium"
  retry_limit              = 0
  is_runnable_concurrently = false

  filter_columns = []

  input_option_type = "custom_connector"

  input_option = {
    custom_connector_input_option = {
      custom_connector_endpoint_id   = trocco_custom_connector_input.custom_connector_to_bigquery_test.endpoints[0].id
      custom_connector_connection_id = trocco_connection.custom_connector_to_bigquery_test.id

      # The endpoint definition above changes in this same apply, so the
      # API rebuilds the snapshot while updating this job definition. The
      # plan must not pin the snapshot attributes to their prior values.

      jsonpath_parser = {
        columns = [
          {
            name = "id"
            type = "long"
          },
          {
            name = "name"
            type = "string"
          },
        ]
      }
    }
  }

  output_option_type = "bigquery"

  output_option = {
    bigquery_output_option = {
      bigquery_connection_id                   = trocco_connection.bigquery.id
      dataset                                  = "test_dataset"
      table                                    = "custom_connector_to_bigquery_test_table"
      mode                                     = "append"
      bigquery_output_option_clustering_fields = []
      bigquery_output_option_column_options    = []
      bigquery_output_option_merge_keys        = []
    }
  }
}
