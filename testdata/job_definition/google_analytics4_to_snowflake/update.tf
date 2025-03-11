resource "trocco_connection" "ga4" {
  connection_type          = "google_analytics4"
  name                     = "GA4 Example"
  service_account_json_key = <<JSON
  {
    "type": "service_account",
    "project_id": "example-project-id",
    "private_key_id": "example-private-key-id",
    "client_email":"joe@example.com",
    "client_id":"1234567890",
    "private_key":"-----BEGIN PRIVATE KEY-----\n..."
  }
  JSON
}
resource "trocco_connection" "snowflake" {
  connection_type = "snowflake"
  name            = "Snowflake Example"

  host        = "exmaple.snowflakecomputing.com"
  auth_method = "user_password"
  user_name   = "dummy_name"
  password    = "dummy_password"
}
resource "trocco_job_definition" "ga4_to_snowflake" {
  name = "GA4 to Snowflake"

  filter_columns = [
    {
      default             = ""
      format              = ""
      json_expand_enabled = false
      name                = "date_hour"
      src                 = "date_hour"
      type                = "timestamp"
    },
    {
      default             = ""
      json_expand_enabled = false
      name                = "yyyymm"
      src                 = "yyyymm"
      type                = "string"
    },
    {
      default             = ""
      json_expand_enabled = false
      name                = "total_users"
      src                 = "total_users"
      type                = "long"
    },
    {
      default             = ""
      json_expand_enabled = false
      name                = "property_id"
      src                 = "property_id"
      type                = "string"
    }
  ]
  input_option_type = "google_analytics4"
  input_option = {
    google_analytics4_input_option = {
      google_analytics4_connection_id = trocco_connection.ga4.id
      property_id                     = "262596771"
      time_series                     = "dateHour"
      start_date                      = "2daysAgo"
      google_analytics4_input_option_dimensions = []
      google_analytics4_input_option_metrics = [
        {
          name       = "totalUsers",
          expression = null
        }
      ]
      incremental_loading_enabled = false
      input_option_columns = [
        {
          name = "date_hour"
          type = "timestamp"
        },
        {
          name = "total_users"
          type = "long"
        },
        {
          name = "property_id"
          type = "string"
        },
      ]
    }
  }
  output_option_type = "snowflake"
  output_option = {
    snowflake_output_option = {
      batch_size              = 50
      database                = "test_database"
      default_time_zone       = "UTC"
      delete_stage_on_error   = false
      empty_field_as_null     = true
      max_retry_wait          = 1800000
      mode                    = "insert"
      retry_limit             = 12
      retry_wait              = 1000
      schema                  = "PUBLIC"
      snowflake_connection_id = trocco_connection.snowflake.id
      table                   = "example"
      warehouse               = "COMPUTE_WH"
    }
  }
}
