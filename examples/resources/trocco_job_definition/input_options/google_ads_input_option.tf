resource "trocco_job_definition" "google_ads_input_example" {
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
}
