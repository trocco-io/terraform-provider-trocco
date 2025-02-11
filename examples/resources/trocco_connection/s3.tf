resource "trocco_connection" "s3" {
  connection_type = "s3"

  name        = "S3 Example"
  description = "This is a AWS S3 connection example"

  aws_auth_type = "iam_user"
  aws_iam_user = {
    access_key_id     = "YOUR_ACCESS_KEY_ID"
    secret_access_key = "YOUR_SECRET_ACCESS_KEY"
  }
}

resource "trocco_connection" "s3_with_assume_role" {
  connection_type = "s3"

  name        = "S3 Example"
  description = "This is a AWS S3 connection example"

  aws_auth_type = "assume_role"
  aws_assume_role = {
    account_id = "123456789012"
    role_name  = "YOUR_ROLE_NAME"
  }
}
