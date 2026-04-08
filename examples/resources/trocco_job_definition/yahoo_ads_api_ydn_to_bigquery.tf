resource "trocco_job_definition" "yahoo_ads_report_to_bigquery" {
  name               = "Yahoo Ads Report to BigQuery"
  description        = "Daily report data from Yahoo Ads YDN synced to BigQuery"
  input_option_type  = "yahoo_ads_api_ydn"
  output_option_type = "bigquery"

  input_option = {
    yahoo_ads_api_ydn_input_option = {
      yahoo_ads_api_connection_id = 1 # require your yahoo_ads_api connection id
      target                      = "report"
      base_account_id             = "1234567890"
      account_id                  = "1234567890"
      start_date                  = "2024-01-01"
      end_date                    = "2024-01-31"
      include_deleted             = false

      input_option_columns = [
        {
          name = "DATE"
          type = "string"
        },
        {
          name = "CAMPAIGN_ID"
          type = "long"
        },
        {
          name = "CAMPAIGN_NAME"
          type = "string"
        },
        {
          name = "IMPRESSIONS"
          type = "long"
        },
        {
          name = "CLICKS"
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
      bigquery_connection_id = 1 # require your bigquery connection id
      dataset                = "yahoo_ads"
      table                  = "report"
      mode                   = "merge"
      auto_create_dataset    = false
      location               = "US"
      timeout_sec            = 300
      open_timeout_sec       = 300
      read_timeout_sec       = 300
      send_timeout_sec       = 300
      retries                = 5
      bigquery_output_option_merge_keys = [
        "DATE",
        "CAMPAIGN_ID"
      ]
      bigquery_output_option_clustering_fields = []
      bigquery_output_option_column_options    = []
    }
  }

  filter_columns = [
    {
      name = "DATE"
      src  = "DATE"
      type = "string"
    }
  ]
}

resource "trocco_job_definition" "yahoo_ads_campaign_stats_to_bigquery" {
  name               = "Yahoo Ads Campaign Stats to BigQuery"
  description        = "Campaign stats data from Yahoo Ads YDN synced to BigQuery"
  input_option_type  = "yahoo_ads_api_ydn"
  output_option_type = "bigquery"

  input_option = {
    yahoo_ads_api_ydn_input_option = {
      yahoo_ads_api_connection_id = 1 # require your yahoo_ads_api connection id
      target                      = "stats"
      base_account_id             = "1234567890"
      account_id                  = "1234567890"
      report_type                 = "CAMPAIGN"
      start_date                  = "2024-01-01"
      end_date                    = "2024-01-31"

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
      bigquery_connection_id = 1 # require your bigquery connection id
      dataset                = "yahoo_ads"
      table                  = "campaign_stats"
      mode                   = "merge"
      auto_create_dataset    = false
      location               = "US"
      timeout_sec            = 300
      open_timeout_sec       = 300
      read_timeout_sec       = 300
      send_timeout_sec       = 300
      retries                = 5
      bigquery_output_option_merge_keys = [
        "CAMPAIGN_ID"
      ]
    }
  }

  filter_columns = [
    {
      name = "CAMPAIGN_NAME"
      src  = "CAMPAIGN_NAME"
      type = "string"
    }
  ]
}

resource "trocco_job_definition" "yahoo_ads_adgroup_stats_to_bigquery" {
  name               = "Yahoo Ads AdGroup Stats to BigQuery"
  description        = "Ad group stats data from Yahoo Ads YDN for analytics"
  input_option_type  = "yahoo_ads_api_ydn"
  output_option_type = "bigquery"

  input_option = {
    yahoo_ads_api_ydn_input_option = {
      yahoo_ads_api_connection_id = 1 # require your yahoo_ads_api connection id
      target                      = "stats"
      base_account_id             = "1234567890"
      account_id                  = "1234567890"
      report_type                 = "ADGROUP"
      start_date                  = "2024-01-01"
      end_date                    = "2024-01-31"

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
          name = "ADGROUP_NAME"
          type = "string"
        },
        {
          name = "ADGROUP_ID"
          type = "long"
        },
        {
          name = "IMPRESSIONS"
          type = "long"
        },
        {
          name = "CLICKS"
          type = "long"
        },
        {
          name = "COST"
          type = "double"
        },
        {
          name = "CONVERSIONS"
          type = "long"
        },
        {
          name = "CONVERSION_VALUE"
          type = "double"
        }
      ]
    }
  }

  output_option = {
    bigquery_output_option = {
      bigquery_connection_id = 1 # require your bigquery connection id
      dataset                = "yahoo_ads"
      table                  = "adgroup_stats"
      mode                   = "merge"
      auto_create_dataset    = false
      location               = "US"
      timeout_sec            = 300
      open_timeout_sec       = 300
      read_timeout_sec       = 300
      send_timeout_sec       = 300
      retries                = 5
      bigquery_output_option_merge_keys = [
        "ADGROUP_ID"
      ]
    }
  }

  filter_columns = [
    {
      name = "CAMPAIGN_NAME"
      src  = "CAMPAIGN_NAME"
      type = "string"
    }
  ]
}
