resource "trocco_job_definition" "google_ads_to_bigquery_example" {
  name                     = "google_ads_to_bigquery_example"
  description              = "Transfer data from Google Ads to BigQuery"
  is_runnable_concurrently = false
  retry_limit              = 0
  filter_columns = [
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
      name                         = "campaign_id"
      src                          = "campaign_id"
      type                         = "long"
    },
    {
      default                      = null
      format                       = "%Y-%m-%d %H:%M:%S"
      json_expand_columns          = []
      json_expand_enabled          = false
      json_expand_keep_base_column = false
      name                         = "created_at"
      src                          = "created_at"
      type                         = "timestamp"
    },
    {
      default                      = null
      json_expand_columns          = []
      json_expand_enabled          = false
      json_expand_keep_base_column = false
      name                         = "ctr"
      src                          = "ctr"
      type                         = "double"
    },
    {
      default                      = null
      json_expand_columns          = []
      json_expand_enabled          = false
      json_expand_keep_base_column = false
      name                         = "is_enabled"
      src                          = "is_enabled"
      type                         = "boolean"
    },
  ]
  input_option_type = "google_ads"
  input_option = {
    google_ads_input_option = {
      customer_id              = "1234567890"
      resource_type            = "campaign"
      start_date               = "2024-01-01"
      end_date                 = "2024-01-31"
      google_ads_connection_id = 1 # please set your google ads connection id
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
      table                                    = "google_ads_example_table"
      template_table                           = ""
      timeout_sec                              = 300
    }
  }
}
