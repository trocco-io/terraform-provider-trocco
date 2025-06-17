resource "trocco_job_definition" "http_to_bigquery_maximum" {
  input_option_type = "http"
  input_option = {
    http_input_option = {
      method                                     = "POST"
      url                                        = "http://example.com"
      user_agent                                 = "user-agent-example"
      pager_type                                 = "cursor"
      cursor_request_parameter_cursor_name       = "next_cursor"
      cursor_response_parameter_cursor_json_path = "$.next_cursor"
      request_headers = [
        { key = "Content-Type", value = "application/json", masking = false },
        { key = "Authorization", value = "Bearer example_token", masking = true },
      ]
      request_params = [
        { key = "foo", value = "bar" },
      ]
      success_code = "200"
      jsonl_parser = {
        stop_on_invalid_record = true
        default_time_zone      = "UTC"
        newline                = "LF"
        charset                = "UTF-8"
        columns = [
          {
            name = "id"
            type = "long"
          },
          {
            name = "name"
            type = "string"
          }
        ]
      }
    }
  }
}
