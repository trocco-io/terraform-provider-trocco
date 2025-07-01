resource "trocco_notification_destination" "email_test" {
  type = "email"
  email_config = {
    email = "test@example.com"
  }
}