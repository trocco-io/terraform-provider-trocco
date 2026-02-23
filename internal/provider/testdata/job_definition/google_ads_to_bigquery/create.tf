resource "trocco_connection" "test_bq" {
  connection_type = "bigquery"

  name        = "BigQuery Example"
  description = "This is a BigQuery connection example"

  project_id               = "system-playground"
  service_account_json_key = <<JSON
  {
    "type": "service_account",
    "project_id": "example-project-id",
    "private_key_id": "example-private-key-id",
    "private_key":"-----BEGIN PRIVATE KEY-----\n..."
  }
  JSON
}

resource "trocco_job_definition" "google_ads_to_bigquery" {
  name                     = "Google Ads to BigQuery Test"
  description              = "Test job definition for transferring data from Google Ads to BigQuery"
  resource_enhancement     = "medium"
  retry_limit              = 2
  is_runnable_concurrently = true

  filter_columns = [
    {
      name                         = "campaign_name"
      src                          = "campaign_name"
      type                         = "string"
      default                      = null
      json_expand_enabled          = false
      json_expand_keep_base_column = false
    },
    {
      name                         = "campaign_id"
      src                          = "campaign_id"
      type                         = "long"
      default                      = null
      json_expand_enabled          = false
      json_expand_keep_base_column = false
    },
    {
      name                         = "created_at"
      src                          = "created_at"
      type                         = "timestamp"
      default                      = null
      json_expand_enabled          = false
      json_expand_keep_base_column = false
    },
    {
      name                         = "ctr"
      src                          = "ctr"
      type                         = "double"
      default                      = null
      json_expand_enabled          = false
      json_expand_keep_base_column = false
    },
    {
      name                         = "is_enabled"
      src                          = "is_enabled"
      type                         = "boolean"
      default                      = null
      json_expand_enabled          = false
      json_expand_keep_base_column = false
    },
  ]

  input_option_type = "google_ads"

  input_option = {
    google_ads_input_option = {
      customer_id              = "1234567890"
      resource_type            = "campaign"
      start_date               = "2024-01-01"
      end_date                 = "2024-01-31"
      google_ads_connection_id = 794
      input_option_columns = [
        {
          name = "campaign.name"
          type = "string"
        },
        {
          name = "campaign.id"
          type = "long"
        },
        {
          name   = "campaign.create_time"
          type   = "timestamp"
          format = "%Y-%m-%d %H:%M:%S"
        },
        {
          name = "metrics.ctr"
          type = "double"
        },
        {
          name = "campaign.experiment_type"
          type = "boolean"
        },
      ]
      conditions = [
        "campaign.status = 'ENABLED'",
        "metrics.impressions > 0",
      ]
    }
  }

  output_option_type = "bigquery"

  output_option = {
    bigquery_output_option = {
      bigquery_connection_id = trocco_connection.test_bq.id
      dataset                = "test_dataset"
      table                  = "google_ads_campaign_test"
      location               = "US"
      auto_create_dataset    = false
      auto_create_table      = false
      mode                   = "append"
    }
  }
}
