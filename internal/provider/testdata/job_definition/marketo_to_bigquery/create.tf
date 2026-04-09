resource "trocco_connection" "marketo_test" {
  connection_type = "marketo"

  name                       = "Test Marketo Connection"
  description                = "Test Marketo connection for job definitions"
  account_id         = "123-ABC-456"
  client_id          = "client_test_123"
  client_secret      = "secret_test_123"
  api_max_call_count = 5000
}

resource "trocco_connection" "bigquery_test" {
  connection_type = "bigquery"
  name            = "Test BigQuery Connection"
  project_id      = "test-project"
  service_account_json_key = jsonencode({
    type                        = "service_account"
    project_id                  = "test-project"
    private_key_id              = "test-key-id"
    private_key                 = "-----BEGIN PRIVATE KEY-----\nMIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQC7W8jN\n-----END PRIVATE KEY-----\n"
    client_email                = "test@test-project.iam.gserviceaccount.com"
    client_id                   = "123456789"
    auth_uri                    = "https://accounts.google.com/o/oauth2/auth"
    token_uri                   = "https://oauth2.googleapis.com/token"
    auth_provider_x509_cert_url = "https://www.googleapis.com/oauth2/v1/certs"
  })
}

# Marketo Append Example
resource "trocco_job_definition" "marketo_to_bigquery_append" {
  name        = "Marketo to BigQuery Example - Append"
  description = "Transfer Marketo lead data to BigQuery with append mode"

  input_option_type = "marketo"
  input_option = {
    marketo_input_option = {
      marketo_connection_id   = trocco_connection.marketo_test.id
      target                  = "lead"
      from_date               = "2025-03-01"
      end_date                = "2025-03-18"
      use_updated_at          = false
      polling_interval_second = 60
      bulk_job_timeout_second = 3600

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
        },
        {
          name = "first_name"
          type = "string"
        },
        {
          name = "last_name"
          type = "string"
        }
      ]
    }
  }

  output_option_type = "bigquery"
  output_option = {
    bigquery_output_option = {
      bigquery_connection_id = trocco_connection.bigquery_test.id
      dataset                = "marketing_data"
      table                  = "marketo_leads"
      mode                   = "append"
    }
  }

  filter_columns = [
    {
      name = "id"
      src  = "id"
      type = "long"
    }
  ]
}

# Marketo Merge Example
resource "trocco_job_definition" "marketo_to_bigquery_merge" {
  name        = "Marketo to BigQuery Example - Merge"
  description = "Transfer Marketo lead data to BigQuery with merge mode (upsert)"

  input_option_type = "marketo"
  input_option = {
    marketo_input_option = {
      marketo_connection_id   = trocco_connection.marketo_test.id
      target                  = "lead"
      from_date               = "2025-03-01"
      end_date                = "2025-03-18"
      use_updated_at          = true
      polling_interval_second = 60
      bulk_job_timeout_second = 3600

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
        },
        {
          name = "updated_at"
          type = "timestamp"
        },
        {
          name = "first_name"
          type = "string"
        },
        {
          name = "last_name"
          type = "string"
        }
      ]
    }
  }

  output_option_type = "bigquery"
  output_option = {
    bigquery_output_option = {
      bigquery_connection_id            = trocco_connection.bigquery_test.id
      dataset                           = "marketing_data"
      table                             = "marketo_leads"
      mode                              = "merge"
      auto_create_dataset               = false
      location                          = "US"
      timeout_sec                       = 300
      open_timeout_sec                  = 300
      read_timeout_sec                  = 300
      send_timeout_sec                  = 300
      retries                           = 5
      bigquery_output_option_merge_keys = ["id"]
    }
  }

  filter_columns = [
    {
      name = "id"
      src  = "id"
      type = "long"
    }
  ]
}

# Marketo Activity Example
resource "trocco_job_definition" "marketo_activity_to_bigquery_replace" {
  name        = "Marketo Activity to BigQuery Example - Replace"
  description = "Transfer Marketo activity data to BigQuery with replace mode"

  input_option_type = "marketo"
  input_option = {
    marketo_input_option = {
      marketo_connection_id   = trocco_connection.marketo_test.id
      target                  = "activity"
      from_date               = "2025-03-01"
      end_date                = "2025-03-18"
      activity_type_ids       = [1, 6, 12]
      polling_interval_second = 120
      bulk_job_timeout_second = 7200

      input_option_columns = [
        {
          name = "lead_id"
          type = "long"
        },
        {
          name = "activity_type"
          type = "string"
        },
        {
          name = "activity_date"
          type = "timestamp"
        },
        {
          name = "activity_attributes"
          type = "json"
        }
      ]
    }
  }

  output_option_type = "bigquery"
  output_option = {
    bigquery_output_option = {
      bigquery_connection_id = trocco_connection.bigquery_test.id
      dataset                = "marketing_data"
      table                  = "marketo_activities"
      mode                   = "replace"
      auto_create_dataset    = false
      location               = "US"
    }
  }

  filter_columns = [
    {
      name = "lead_id"
      src  = "lead_id"
      type = "long"
    }
  ]
}
