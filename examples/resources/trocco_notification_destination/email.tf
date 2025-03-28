resource "trocco_notification_destination" "email" {
  type = "email"

  email_config {
    email = "notify@example.com"
  }
}
