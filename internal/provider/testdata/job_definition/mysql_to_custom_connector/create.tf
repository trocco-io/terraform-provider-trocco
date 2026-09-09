resource "trocco_custom_connector_output" "mysql_to_custom_connector_test" {
  name      = "mysql-to-custom-connector-test-definition"
  url       = "https://example.com"
  auth_type = "api_key"

  auth_header_name   = "Authorization"
  auth_header_scheme = "Bearer"

  endpoints = [
    {
      name                = "create_user"
      request_type        = "single"
      method              = "POST"
      operation           = "create"
      path                = "/users"
      payload_type        = "json"
      success_codes       = "200,201"
      not_retryable_codes = "400,401,403,404"
      request_timeout_sec = 45
      template            = "{\"name\": \"{{ name }}\"}"

      headers = [
        {
          name          = "X-Api-Version"
          display_name  = "API Version"
          default_value = "2"
          is_editable   = false
          is_required   = true
        },
      ]

      path_parameters = []

      query_parameters = [
        {
          name          = "dry_run"
          display_name  = "Dry Run"
          default_value = "false"
          is_editable   = true
          is_required   = false
        },
      ]
    },
    {
      name                = "update_user"
      request_type        = "single"
      method              = "PATCH"
      operation           = "update"
      path                = "/users/{{user_id}}"
      payload_type        = "json"
      success_codes       = "200,204"
      not_retryable_codes = "400,401,403,404"
      request_timeout_sec = 45
      template            = "{\"name\": \"{{ name }}\"}"

      headers = []

      path_parameters = [
        {
          name  = "user_id"
          value = "user_id"
        },
      ]

      query_parameters = []
    },
  ]
}

resource "trocco_connection" "mysql_to_custom_connector_test" {
  connection_type = "custom_connector"

  name                = "MySQL to Custom Connector Test Connection"
  custom_connector_id = trocco_custom_connector_output.mysql_to_custom_connector_test.id
  api_key             = "test-api-key"
}

resource "trocco_job_definition" "mysql_to_custom_connector" {
  name                     = "MySQL to Custom Connector Test"
  description              = "Test job definition for transferring data from MySQL to a custom connector"
  resource_enhancement     = "medium"
  retry_limit              = 0
  is_runnable_concurrently = false

  filter_columns = []

  input_option_type = "mysql"
  input_option = {
    mysql_input_option = {
      mysql_connection_id         = trocco_connection.mysql.id
      database                    = "test_database"
      table                       = "test_table"
      connect_timeout             = 300
      socket_timeout              = 1800
      incremental_loading_enabled = false
      default_time_zone           = "Asia/Tokyo"
      use_legacy_datetime_code    = false
      input_option_columns = [
        {
          name = "id"
          type = "long"
        },
        {
          name = "name"
          type = "string"
        },
      ]
      query = <<-EOT
        SELECT *
        FROM test_table
      EOT
    }
  }

  output_option_type = "custom_connector"
  output_option = {
    custom_connector_output_option = {
      custom_connector_connection_id             = trocco_connection.mysql_to_custom_connector_test.id
      create_custom_connector_output_endpoint_id = trocco_custom_connector_output.mysql_to_custom_connector_test.endpoints[0].id

      create_endpoint_settings = {
        query_parameters = [
          {
            name  = "dry_run"
            value = "true"
          },
        ]
      }
    }
  }
}
