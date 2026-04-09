# ==============================================================================
# Example 1: Marketo to BigQuery - Append Mode (Simple Append)
# ==============================================================================

resource "trocco_job_definition" "marketo_to_bigquery_append" {
  name        = "Marketo to BigQuery Example - Append"
  description = "Transfer Marketo lead data to BigQuery with append mode"

  input_option_type = "marketo"
  input_option = {
    marketo_input_option = {
      marketo_connection_id   = 123
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
      bigquery_connection_id = 456
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

# ==============================================================================
# Example 2: Marketo to BigQuery - Merge Mode (Upsert)
# ==============================================================================

resource "trocco_job_definition" "marketo_to_bigquery_merge" {
  name        = "Marketo to BigQuery Example - Merge"
  description = "Transfer Marketo lead data to BigQuery with merge mode (upsert)"

  input_option_type = "marketo"
  input_option = {
    marketo_input_option = {
      marketo_connection_id   = 123
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
      bigquery_connection_id            = 456
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

# ==============================================================================
# Example 3: Marketo Activity to BigQuery - Replace Mode
# ==============================================================================

resource "trocco_job_definition" "marketo_activity_to_bigquery_replace" {
  name        = "Marketo Activity to BigQuery Example - Replace"
  description = "Transfer Marketo activity data to BigQuery with replace mode"

  input_option_type = "marketo"
  input_option = {
    marketo_input_option = {
      marketo_connection_id   = 123
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
      bigquery_connection_id = 456
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
