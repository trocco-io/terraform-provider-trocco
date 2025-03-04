resource "trocco_job_definition" "ga4_input_example" {
  input_option_type = "google_analytics4"
  input_option = {
    google_analytics4_input_option = {
      google_analytics4_connection_id = 1
      property_id                     = "123456789"
      time_series                     = "dateHour"
      start_date                      = "2daysAgo"
      end_date                        = "today"
      google_analytics4_input_option_dimensions = [
        {
          name       = "yyyymm",
          expression = <<-DIM
            {
              "concatenate": {
                "dimensionNames": ["year","month"],
                "delimiter": "-"
              }
            }
          DIM

        }
      ]
      google_analytics4_input_option_metrics = [
        {
          name       = "totalUsers",
          expression = null
        }
      ]
      incremental_loading_enabled = false
      retry_limit                 = 5
      retry_sleep                 = 2
      raise_on_other_row          = false
      limit_of_rows               = 10000
    }
  }
}
