resource "trocco_connection" "test_yahoo_ads_api_stats" {
  connection_type = "yahoo_ads_api"
  name            = "Yahoo Ads API Test Connection"
  description     = "Test connection for Yahoo Ads API"
}

resource "trocco_connection" "test_bq_yahoo_stats" {
  connection_type = "bigquery"
  name            = "BigQuery Test Connection"
  description     = "Test connection for BigQuery"

  project_id               = "test-project"
  service_account_json_key = <<JSON
{
  "type": "service_account",
  "project_id": "test-project",
  "private_key_id": "test-key-id",
  "private_key": "-----BEGIN PRIVATE KEY-----\ntest\n-----END PRIVATE KEY-----\n"
}
JSON
}

resource "trocco_job_definition" "yahoo_ads_stats_to_bigquery" {
  name                     = "Yahoo Ads Campaign Stats to BigQuery"
  description              = "Campaign stats data from Yahoo Ads YDN synced to BigQuery"
  resource_enhancement     = "large"
  retry_limit              = 3
  is_runnable_concurrently = false

  input_option_type  = "yahoo_ads_api_ydn"
  output_option_type = "bigquery"

  input_option = {
    yahoo_ads_api_ydn_input_option = {
      yahoo_ads_api_connection_id = trocco_connection.test_yahoo_ads_api_stats.id
      target                      = "stats"
      base_account_id             = "1234567890"
      account_id                  = "1234567890"
      report_type                 = "CAMPAIGN"
      start_date                  = "2024-01-01"
      end_date                    = "2024-01-31"
      include_deleted             = true

      input_option_columns = [
        {
          name = "CAMPAIGN_NAME"
          type = "string"
        },
        {
          name = "CAMPAIGN_ID"
          type = "long"
        },
        {
          name = "CLICKS"
          type = "long"
        },
        {
          name = "IMPRESSIONS"
          type = "long"
        },
        {
          name = "COST"
          type = "double"
        }
      ]
    }
  }

  output_option = {
    bigquery_output_option = {
      bigquery_connection_id = trocco_connection.test_bq_yahoo_stats.id
      dataset                = "yahoo_ads"
      table                  = "campaign_stats"
      mode                   = "replace"
      auto_create_dataset    = true
      location               = "US"
      timeout_sec            = 600
      open_timeout_sec       = 300
      read_timeout_sec       = 300
      send_timeout_sec       = 300
      retries                = 3
      bigquery_output_option_merge_keys = [
        "CAMPAIGN_ID"
      ]
      bigquery_output_option_clustering_fields = []
      bigquery_output_option_column_options    = []
    }
  }

  filter_columns = [
    {
      name                         = "CAMPAIGN_NAME"
      src                          = "CAMPAIGN_NAME"
      type                         = "string"
      default                      = null
      json_expand_enabled          = false
      json_expand_keep_base_column = false
    }
  ]
}
