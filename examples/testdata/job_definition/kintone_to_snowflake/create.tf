resource "trocco_connection" "kintone" {
  connection_type     = "kintone"
  name                = "Kintone Example"
  description         = "This is a Kintone connection example"
  domain              = "test_domain"
  login_method        = "username_and_password"
  username            = "username"
  password            = "password"
  basic_auth_username = "basic_auth_username"
  basic_auth_password = "basic_auth_password"
}

resource "trocco_connection" "snowflake" {
  connection_type = "snowflake"
  name            = "Snowflake Example"
  description     = "This is a Snowflake connection example"
  host            = "exmaple.snowflakecomputing.com"
  auth_method     = "user_password"
  user_name       = "dummy_name"
  password        = "dummy_password"
}

resource "trocco_job_definition" "kintone_to_snowflake" {
  description              = "Test job definition"
  filter_columns           = [
    {
      default                      = null
      json_expand_enabled          = false
      json_expand_keep_base_column = false
      name                         = "name"
      src                          = "name"
      type                         = "string"
    },
    {
      default                      = null
      json_expand_enabled          = false
      json_expand_keep_base_column = false
      name                         = "url"
      src                          = "url"
      type                         = "string"
    }
  ]
  input_option = {
    kintone_input_option = {
      kintone_connection_id     = trocco_connection.kintone.id
      app_id                    = "123"
      guest_space_id            = null
      query                     = null
      expand_subtable           = false
      custom_variable_settings = [
        {
          name  = "$string$"
          type  = "string"
          value = "foo"
        }
      ]
      input_option_columns = [
        {
          name = "duration"
          type = "string"
        },
        {
          name = "date"
          type = "timestamp"
          format = "%Y%m%d"
        }
      ]
    }
  }
  input_option_type        = "kintone"
  is_runnable_concurrently = false
  name                     = "kintone to snowflake"
  output_option            = {
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
      table                   = "ewaoiiowe"
      warehouse               = "COMPUTE_WH"
    }
  }
  output_option_type       = "snowflake"
  resource_enhancement     = "medium"
  retry_limit              = 0
}
