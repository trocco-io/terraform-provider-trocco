resource "trocco_notification_destination" "http_test" {
  type = "http"
  http_config = {
    name        = "terraform-acc-http"
    url         = "https://example.com/webhook"
    description = "created by acceptance test"
    headers = [
      {
        key     = "Authorization"
        value   = "Bearer acc-secret"
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
        value   = "acc-token"
        masking = true
      }
    ]
  }
}
