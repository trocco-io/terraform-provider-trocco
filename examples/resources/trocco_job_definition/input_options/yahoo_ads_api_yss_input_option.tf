resource "trocco_job_definition" "yahoo_ads_api_yss_input_example" {
  input_option_type = "yahoo_ads_api_yss"
  input_option = {
    yahoo_ads_api_yss_input_option = {
      yahoo_ads_api_connection_id = 1
      base_account_id             = "base-1234"
      account_id                  = "acc-5678"
      service                     = "report_definition_service"
      report_type                 = "ADGROUP"
      start_date                  = "20240501"
      end_date                    = "20240531"
      exclude_zero_impressions    = true
      input_option_columns = [
        {
          name = "id"
          type = "long"
        },
        {
          name = "name"
          type = "string"
        },
        {
          name = "email"
          type = "string"
        },
        {
          name   = "test"
          type   = "timestamp"
          format = "%Y-%m-%d %H:%M:%S"
        },
      ]
    }
  }
}
