resource "trocco_connection" "bigquery" {
  connection_type = "bigquery"
  name            = "BigQuery Example"
  description     = "This is a BigQuery connection example"
  project_id      = "example"
  service_account_json_key = jsonencode({
    type                        = "service_account"
    project_id                  = "example-project-id"
    private_key_id              = "example-private-key-id"
    private_key                 = "-----BEGIN PRIVATE KEY-----\nyour-private-key\n-----END PRIVATE KEY-----\n"
    client_email                = "example@example-project.iam.gserviceaccount.com"
    client_id                   = "123456789"
    auth_uri                    = "https://accounts.google.com/o/oauth2/auth"
    token_uri                   = "https://oauth2.googleapis.com/token"
    auth_provider_x509_cert_url = "https://www.googleapis.com/oauth2/v1/certs"
    client_x509_cert_url        = "https://www.googleapis.com/robot/v1/metadata/x509/example%40example-project.iam.gserviceaccount.com"
  })
}

resource "trocco_job_definition" "hubspot_to_bigquery" {
  name                     = "Hubspot to BigQuery Test"
  description              = "Test job definition for transferring data from Hubspot to BigQuery"
  resource_enhancement     = "medium"
  retry_limit              = 2
  is_runnable_concurrently = true

  filter_columns = [
    {
      name                         = "contact_id"
      src                          = "id"
      type                         = "long"
      default                      = null
      json_expand_enabled          = false
      json_expand_keep_base_column = false
    },
    {
      name                         = "email_address"
      src                          = "email"
      type                         = "string"
      default                      = null
      json_expand_enabled          = false
      json_expand_keep_base_column = false
    },
    {
      name                         = "created_timestamp"
      src                          = "created_at"
      type                         = "timestamp"
      default                      = null
      json_expand_enabled          = false
      json_expand_keep_base_column = false
    }
  ]

  input_option_type = "hubspot"

  input_option = {
    hubspot_input_option = {
      hubspot_connection_id       = 1
      target                      = "object"
      object_type                 = "contact"
      incremental_loading_enabled = false
      input_option_columns = [
        {
          name = "id"
          type = "long"
        },
        {
          name = "email"
          type = "string"
        },
        {
          name = "created_at"
          type = "timestamp"
        }
      ]
      custom_variable_settings = [
        {
          name  = "$test_var$"
          type  = "string"
          value = "test_value"
        }
      ]
    }
  }

  output_option_type = "bigquery"

  output_option = {
    bigquery_output_option = {
      auto_create_dataset                      = false
      bigquery_connection_id                   = trocco_connection.bigquery.id
      bigquery_output_option_clustering_fields = []
      bigquery_output_option_column_options    = []
      bigquery_output_option_merge_keys        = []
      dataset                                  = "test_dataset"
      location                                 = "US"
      mode                                     = "append"
      open_timeout_sec                         = 300
      read_timeout_sec                         = 300
      retries                                  = 5
      send_timeout_sec                         = 300
      table                                    = "hubspot_to_bigquery_test_table"
      template_table                           = ""
      timeout_sec                              = 300
    }
  }
}
