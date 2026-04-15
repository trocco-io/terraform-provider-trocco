resource "trocco_job_definition" "facebook_ads_insights_input_example" {
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
}
