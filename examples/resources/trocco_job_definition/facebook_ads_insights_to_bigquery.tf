resource "trocco_job_definition" "facebook_ads_insights_to_bigquery_example" {
  name                     = "facebook_ads_insights_to_bigquery_example"
  description              = "Transfer data from Facebook Ads Insights to BigQuery"
  is_runnable_concurrently = false
  retry_limit              = 0

  filter_columns = [
    {
      default                      = null
      json_expand_columns          = []
      json_expand_enabled          = false
      json_expand_keep_base_column = false
      name                         = "campaign_id"
      src                          = "campaign_id"
      type                         = "string"
    },
    {
      default                      = null
      json_expand_columns          = []
      json_expand_enabled          = false
      json_expand_keep_base_column = false
      name                         = "campaign_name"
      src                          = "campaign_name"
      type                         = "string"
    },
    {
      default                      = null
      json_expand_columns          = []
      json_expand_enabled          = false
      json_expand_keep_base_column = false
      name                         = "impressions"
      src                          = "impressions"
      type                         = "long"
    },
    {
      default                      = null
      json_expand_columns          = []
      json_expand_enabled          = false
      json_expand_keep_base_column = false
      name                         = "clicks"
      src                          = "clicks"
      type                         = "long"
    },
    {
      default                      = null
      json_expand_columns          = []
      json_expand_enabled          = false
      json_expand_keep_base_column = false
      name                         = "spend"
      src                          = "spend"
      type                         = "double"
    },
  ]

  input_option_type = "facebook_ads_insights"
  input_option = {
    facebook_ads_insights_input_option = {
      facebook_ads_insights_connection_id = 1 # please set your facebook ads insights connection id
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

  output_option_type = "bigquery"
  output_option = {
    bigquery_output_option = {
      auto_create_dataset                      = false
      bigquery_connection_id                   = 1 # please set your bigquery connection id
      bigquery_output_option_clustering_fields = []
      bigquery_output_option_column_options    = []
      bigquery_output_option_merge_keys        = []
      dataset                                  = "example_dataset"
      location                                 = "US"
      mode                                     = "append"
      open_timeout_sec                         = 300
      read_timeout_sec                         = 300
      retries                                  = 5
      send_timeout_sec                         = 300
      table                                    = "facebook_ads_insights_example_table"
      template_table                           = ""
      timeout_sec                              = 300
    }
  }
}
