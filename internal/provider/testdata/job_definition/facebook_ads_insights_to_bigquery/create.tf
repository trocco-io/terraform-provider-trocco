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

resource "trocco_job_definition" "facebook_ads_insights_to_bigquery" {
  name                     = "Facebook Ads Insights to BigQuery Test"
  description              = "Test job definition for transferring data from Facebook Ads Insights to BigQuery"
  resource_enhancement     = "medium"
  retry_limit              = 2
  is_runnable_concurrently = false

  input_option_type = "facebook_ads_insights"
  input_option = {
    facebook_ads_insights_input_option = {
      facebook_ads_insights_connection_id = 922
      ad_account_id                       = "act_123456789"
      level                               = "campaign"
      time_range_since                    = "2024-01-01"
      time_range_until                    = "2024-01-31"
      use_unified_attribution_setting     = true
      fields = [
        {
          name = "campaign_id"
        },
        {
          name = "campaign_name"
        },
        {
          name = "impressions"
        },
        {
          name = "clicks"
        },
        {
          name = "spend"
        },
      ]
      breakdowns = [
        {
          name = "country"
        },
        {
          name = "device_platform"
        },
      ]
      action_attribution_windows = [
        {
          name = "1d_click"
        },
        {
          name = "7d_click"
        },
      ]
      action_breakdowns = [
        {
          name = "action_type"
        },
      ]
      custom_variable_settings = [
        {
          name  = "$window_start$"
          type  = "string"
          value = "2024-01-01"
        },
      ]
    }
  }

  filter_columns = [
    {
      name                         = "campaign_id"
      src                          = "campaign_id"
      type                         = "string"
      default                      = ""
      format                       = ""
      json_expand_enabled          = false
      json_expand_keep_base_column = false
    },
    {
      name                         = "impressions"
      src                          = "impressions"
      type                         = "long"
      default                      = ""
      format                       = ""
      json_expand_enabled          = false
      json_expand_keep_base_column = false
    },
  ]

  output_option_type = "bigquery"
  output_option = {
    bigquery_output_option = {
      dataset                                  = "test_dataset"
      table                                    = "facebook_ads_insights_test"
      mode                                     = "append"
      auto_create_dataset                      = false
      bigquery_connection_id                   = trocco_connection.test_bq.id
      location                                 = "US"
      bigquery_output_option_clustering_fields = []
      bigquery_output_option_column_options    = []
      bigquery_output_option_merge_keys        = []
    }
  }
}
