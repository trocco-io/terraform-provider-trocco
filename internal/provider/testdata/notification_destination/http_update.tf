resource "trocco_notification_destination" "http_test" {
  type = "http"
  http_config = {
    name = "terraform-acc-http-updated"
    url  = "https://example.com/webhook/v2"
    headers = [
      {
        key     = "X-Api-Key"
        value   = "rotated-secret"
        masking = true
      }
    ]
  }
}
