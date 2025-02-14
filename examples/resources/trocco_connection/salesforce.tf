resource "trocco_connection" "salesforce" {
  connection_type = "salesforce"

  name        = "Salesforce Example"
  description = "This is a Salesforce connection example"

  auth_method    = "user_password"
  user_name      = "<User Name>"
  password       = "<Password>"
  security_token = "<Security Token>"
  auth_end_point = "https://login.salesforce.com/services/Soap/u/"
}
