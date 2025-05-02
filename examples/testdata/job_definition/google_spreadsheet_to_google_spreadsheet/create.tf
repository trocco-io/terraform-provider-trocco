resource "trocco_connection" "google_spreadsheets" {
  connection_type = "google_spreadsheets"
  name            = "Google Sheets Example"
  description     = "This is a Google Sheets connection example"

  service_account_json_key = <<JSON
  {
    "type": "service_account",
    "project_id": "example-project-id",
    "private_key_id": "example-private-key-id",
    "private_key":"-----BEGIN PRIVATE KEY-----\n..."
  }
  JSON
}

resource "trocco_job_definition" "sheets_to_sheets" {
  name                     = "Google Spreadsheets to Google Spreadsheets"
  description              = "Test job definition for transferring data from Google Spreadsheets to Google Spreadsheets"
  is_runnable_concurrently = false
  retry_limit              = 0
  resource_enhancement     = "medium"

  input_option_type = "google_spreadsheets"
  input_option = {
    google_spreadsheets_input_option = {
      google_spreadsheets_connection_id = trocco_connection.google_spreadsheets.id
      spreadsheets_url                  = "https://docs.google.com/spreadsheets/d/TEST_SHEETS_ID/edit?gid=0"
      worksheet_title                   = "input-data"
      start_row                         = 2
      start_column                      = "A"
      default_time_zone                 = "Asia/Tokyo"
      null_string                       = ""
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
      google_spreadsheets_connection_id = trocco_connection.google_spreadsheets.id
      spreadsheets_id                   = "TEST_SHEETS_ID"
      worksheet_title                   = "output-data"
      timezone                          = "Asia/Tokyo"
      value_input_option                = "USER_ENTERED"
      mode                              = "replace"
      google_spreadsheets_output_option_sorts = [
        {
          column = "created_at"
          order  = "ascending"
        }
      ]
    }
  }

  filter_columns = [
    {
      default                      = ""
      json_expand_enabled          = false
      json_expand_keep_base_column = false
      name                         = "id"
      src                          = "id"
      type                         = "long"
    },
    {
      default                      = ""
      json_expand_enabled          = false
      json_expand_keep_base_column = false
      name                         = "name"
      src                          = "name"
      type                         = "string"
    },
    {
      default                      = ""
      json_expand_enabled          = false
      json_expand_keep_base_column = false
      name                         = "created_at"
      src                          = "created_at"
      type                         = "timestamp"
    }
  ]
}
