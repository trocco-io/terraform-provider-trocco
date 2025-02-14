resource "trocco_job_definition" "sheets_to_sheets_example" {
  name                     = "example_sheets_to_sheets"
  description              = "this is an example job definition for transferring data from Google Spreadsheets to Google Spreadsheets"
  is_runnable_concurrently = false
  retry_limit              = 0
  filter_columns = [
    {
      default                      = null
      json_expand_columns          = []
      json_expand_enabled          = false
      json_expand_keep_base_column = false
      name                         = "id"
      src                          = "id"
      type                         = "long"
    },
    {
      default                      = null
      json_expand_columns          = []
      json_expand_enabled          = false
      json_expand_keep_base_column = false
      name                         = "name"
      src                          = "name"
      type                         = "string"
    },
    {
      default                      = null
      format                       = null
      json_expand_columns          = []
      json_expand_enabled          = false
      json_expand_keep_base_column = false
      name                         = "created_at"
      src                          = "created_at"
      type                         = "timestamp"
    },
  ]
  input_option_type = "google_spreadsheets"
  input_option = {
    google_spreadsheets_input_option = {
      google_spreadsheets_connection_id = 1
      # ex "https://docs.google.com/spreadsheets/d/YOUR_SHEETS_ID/edit?gid=0"
      spreadsheets_id   = "YOUR_SHEETS_ID"
      worksheet_title   = "inputdata"
      start_row         = 2
      start_column      = "A"
      default_time_zone = "Asia/Tokyo"
      null_string       = ""
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
          name   = "created_at"
          type   = "timestamp"
          format = "%Y-%m-%d %H:%M:%S"
        },
      ]
    }
  }
  output_option_type = "google_spreadsheets"
  output_option = {
    google_spreadsheets_output_option = {
      google_spreadsheets_connection_id = 1
      spreadsheets_id                   = "YOUR_SHEETS_ID"
      worksheet_title                   = "outputdata"
      timezone                          = "Asia/Tokyo"
      value_input_option                = "USER_ENTERED"
      mode                              = "replace"
      google_spreadsheets_output_option_sorts = [
        {
          column = "created_at"
          order  = "ASCENDING"
        }
      ]
    }
  }
}
