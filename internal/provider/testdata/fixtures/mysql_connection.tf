resource "trocco_connection" "mysql" {
  connection_type = "mysql"
  name            = "MySQL Example"
  host            = "db.example.com"
  port            = 3306
  user_name       = "root"
  password        = "password"
}
