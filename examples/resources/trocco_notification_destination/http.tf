resource "trocco_notification_destination" "http" {
  type = "http"

  http_config = {
    name        = "my-webhook"
    url         = "https://example.com/webhook"
    description = "Sends job notifications to an internal webhook"
    headers = [
      {
        key     = "Authorization"
        value   = "Bearer XXXX"
        masking = true
      },
      {
        key   = "Content-Type"
        value = "application/json"
      }
    ]
    query_params = [
      {
        key     = "token"
        value   = "XXXX"
        masking = true
      }
    ]
  }
}
