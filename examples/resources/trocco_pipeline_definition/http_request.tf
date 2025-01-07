resource "trocco_pipeline_definition" "http_request" {
  name = "http_request"

  tasks = [
    {
      key  = "http_request"
      type = "http_request"

      http_request_config = {
        name = "Example"

        http_method = "GET"

        url = "https://example.com"

        request_headers = [
          { key : "Authorization", value : "Bearer example", masking : true },
          { key : "Content-Type", value : "application/json", masking : false },
        ]

        request_parameters = [
          { key : "foo", value : "bar", masking : true },
        ]
      }
    }
  ]
}
