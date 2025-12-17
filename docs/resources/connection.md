---
page_title: "trocco_connection Resource - trocco"
subcategory: ""
description: |-
  Provides a TROCCO connection resource.
---

# trocco_connection (Resource)

Provides a TROCCO connection resource.

## Example Usage

### Bigquery

```terraform
resource "trocco_connection" "bigquery" {
  connection_type = "bigquery"

  name        = "BigQuery Example"
  description = "This is a BigQuery connection example"

  project_id               = "example"
  service_account_json_key = <<JSON
  {
    "type": "service_account",
    "project_id": "example-project-id",
    "private_key_id": "example-private-key-id",
    "private_key":"-----BEGIN PRIVATE KEY-----\n..."
  }
  JSON
}
```

### Snowflake

```terraform
resource "trocco_connection" "snowflake" {
  connection_type = "snowflake"

  name        = "Snowflake Example"
  description = "This is a Snowflake connection example"

  host        = "exmaple.snowflakecomputing.com"
  auth_method = "user_password"
  user_name   = "<User Name>"
  password    = "<Password>"
}
```

### Google Cloud Storage(GCS)

```terraform
resource "trocco_connection" "gcs" {
  connection_type = "gcs"

  name        = "GCS Example"
  description = "This is a Google Cloud Storage(GCS) connection example"

  project_id               = "example-project-id"
  service_account_json_key = <<JSON
  {
    "type": "service_account",
    "project_id": "example-project-id",
    "private_key_id": "example-private-key-id",
    "private_key":"-----BEGIN PRIVATE KEY-----\n..."
  }
  JSON
  service_account_email    = "joe@example-project.iam.gserviceaccount.com"
  application_name         = "example-application-name"
}
```

### Google Sheets

```terraform
resource "trocco_connection" "google_spreadsheets" {
  connection_type = "google_spreadsheets"
  name            = "Google Sheets Example"
  description     = "This is a Google Sheets connection example"

  service_account_json_key = <<JSON
  {
    "type": "service_account",
    "project_id": "example-project-id",
    "private_key_id": "example-private-key-id",
    "private_key":"-----BEGIN PRIVATE KEY-----\n..."
  }
  JSON
}
```

### MySQL

```terraform
resource "trocco_connection" "mysql" {
  connection_type = "mysql"
  name            = "MySQL Example"
  description     = "This is a MySQL connection example"
  host            = "db.example.com"
  port            = 65535
  user_name       = "root"
  password        = "password"
  ssl = {
    ca   = <<-SSL_CA
      -----BEGIN PRIVATE KEY-----
      ...SSL CA...
      -----END PRIVATE KEY-----
    SSL_CA
    cert = <<-SSL_CERT
      -----BEGIN CERTIFICATE-----
      ...SSL CRT...
      -----END CERTIFICATE-----
    SSL_CERT
    key  = <<-SSL_KEY
      -----BEGIN PRIVATE KEY-----
      ...SSL KEY...
      -----END PRIVATE KEY-----
    SSL_KEY
  }
  gateway = {
    host           = "gateway.example.com"
    port           = 1234
    user_name      = "gateway-joe"
    password       = "gateway-joepass"
    key            = <<-GATEWAY_KEY
      -----BEGIN PRIVATE KEY-----
      ... GATEWAY KEY...
      -----END PRIVATE KEY-----
    GATEWAY_KEY
    key_passphrase = "sample_passphrase"
  }
  resource_group_id = 1
}
```

### Salesforce

```terraform
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
```

### S3

```terraform
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
```

### PostgreSQL

```terraform
resource "trocco_connection" "postgresql" {
  connection_type = "postgresql"
  name            = "PostgreSQL Example"
  description     = "This is a PostgreSQL connection example"
  host            = "db.example.com"
  port            = 5432
  user_name       = "root"
  password        = "password"
  driver          = "postgresql_42_5_1"
  ssl = {
    ca       = <<-SSL_CA
      -----BEGIN PRIVATE KEY-----
      ...SSL CA...
      -----END PRIVATE KEY-----
    SSL_CA
    cert     = <<-SSL_CERT
      -----BEGIN CERTIFICATE-----
      ...SSL CRT...
      -----END CERTIFICATE-----
    SSL_CERT
    key      = <<-SSL_KEY
      -----BEGIN PRIVATE KEY-----
      ...SSL KEY...
      -----END PRIVATE KEY-----
    SSL_KEY
    ssl_mode = "require"
  }
  gateway = {
    host           = "gateway.example.com"
    port           = 1234
    user_name      = "gateway-joe"
    password       = "gateway-joepass"
    key            = <<-GATEWAY_KEY
      -----BEGIN PRIVATE KEY-----
      ... GATEWAY KEY...
      -----END PRIVATE KEY-----
    GATEWAY_KEY
    key_passphrase = "sample_passphrase"
  }
  resource_group_id = 1
}
```

### Google Analytics4

```terraform
resource "trocco_connection" "google_analytics4" {
  connection_type   = "google_analytics4"
  name              = "Google Analytics4 Example"
  description       = "This is a Google Analytics4 connection example"
  resource_group_id = 1

  service_account_json_key = <<JSON
  {
    "type": "service_account",
    "project_id": "example-project-id",
    "private_key_id": "example-private-key-id",
    "private_key":"-----BEGIN PRIVATE KEY-----\n..."
  }
  JSON
}
```

### Kintone

```terraform
# login_method: token
resource "trocco_connection" "kintone_login_method_token" {
  connection_type     = "kintone"
  name                = "Kintone Example"
  description         = "This is a Kintone connection example"
  resource_group_id   = 1
  domain              = "test_domain"
  login_method        = "token"
  token               = "token"
  username            = nil
  password            = nil
  basic_auth_username = "basic_auth_username"
  basic_auth_password = "basic_auth_password"
}

# login_method: username_and_password
resource "trocco_connection" "kintone_login_method_username_and_password" {
  connection_type     = "kintone"
  name                = "Kintone Example"
  description         = "This is a Kintone connection example"
  resource_group_id   = 1
  domain              = "test_domain"
  login_method        = "username_and_password"
  token               = ""
  username            = "username"
  password            = "password"
  basic_auth_username = "basic_auth_username"
  basic_auth_password = "basic_auth_password"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `connection_type` (String) The type of the connection. It must be one of `bigquery`, `snowflake`, `gcs`, `google_spreadsheets`, `mysql`, `salesforce`, `s3`, `postgresql`, `google_analytics4`, `kintone`, `sftp`.
- `name` (String) The name of the connection.

### Optional

- `application_name` (String) GCS: Application name.
- `auth_end_point` (String) Salesforce: Authentication endpoint.
- `auth_method` (String) Snowflake: The authentication method for the Snowflake user. It must be one of `key_pair` or `user_password`.
- `aws_assume_role` (Attributes) S3: AssumeRole configuration. (see [below for nested schema](#nestedatt--aws_assume_role))
- `aws_auth_type` (String) S3: The authentication type for the S3 connection. It must be one of `iam_user` or `assume_role`.
- `aws_iam_user` (Attributes) S3: IAM User configuration. (see [below for nested schema](#nestedatt--aws_iam_user))
- `aws_privatelink_enabled` (Boolean) SFTP: Whether AWS PrivateLink is enabled. Default is false.
- `basic_auth_password` (String, Sensitive) Kintone: Basic Auth Password
- `basic_auth_username` (String) Kintone: Basic Auth Username
- `description` (String) The description of the connection.
- `domain` (String) Kintone: Domain.
- `driver` (String) Snowflake, MySQL, PostgreSQL: The name of a Database driver.
  - MySQL: null, mysql_connector_java_5_1_49
  - Snowflake: null, snowflake_jdbc_3_14_2, snowflake_jdbc_3_17_0,
  - PostgreSQL: postgresql_42_5_1, postgresql_9_4_1205_jdbc41
- `gateway` (Attributes) MySQL, PostgreSQL: Whether to connect via SSH (see [below for nested schema](#nestedatt--gateway))
- `host` (String) Snowflake, PostgreSQL: The host of a (Snowflake, PostgreSQL) account.
- `login_method` (String) Kintone: Login Method
- `password` (String, Sensitive) Snowflake, PostgreSQL: The password for the (Snowflake, PostgreSQL) user.
- `port` (Number) MySQL, PostgreSQL: The port of the (MySQL, PostgreSQL) server.
- `private_key` (String, Sensitive) Snowflake: A private key for the Snowflake user.
- `project_id` (String) BigQuery, GCS: A GCP project ID.
- `resource_group_id` (Number) The ID of the resource group the connection belongs to.
- `role` (String) Snowflake: A role attached to the Snowflake user.
- `secret_key` (String, Sensitive) SFTP: RSA private key for authentication.
- `secret_key_passphrase` (String, Sensitive) SFTP: Passphrase for the RSA private key.
- `security_token` (String, Sensitive) Salesforce: Security token.
- `service_account_email` (String, Sensitive) GCS: A GCP service account email.
- `service_account_json_key` (String, Sensitive) BigQuery, Google Sheets, Google Analytics4: A GCP service account key.
- `ssh_tunnel_id` (Number) SFTP: SSH tunnel ID. Required when aws_privatelink_enabled is true.
- `ssl` (Attributes) MySQL, PostgreSQL: SSL configuration. (see [below for nested schema](#nestedatt--ssl))
- `token` (String, Sensitive) Kintone: Token.
- `user_directory_is_root` (Boolean) SFTP: Whether the user directory is root. Default is true.
- `user_name` (String) Snowflake, PostgreSQL: The name of a (Snowflake, PostgreSQL) user.
- `username` (String) Kintone: The name of a user.
- `windows_server` (Boolean) SFTP: Whether the server is a Windows server. Default is false.

### Read-Only

- `id` (Number) The ID of the connection.

<a id="nestedatt--aws_assume_role"></a>
### Nested Schema for `aws_assume_role`

Optional:

- `account_id` (String) S3: The account ID for the AssumeRole configuration.
- `role_name` (String) S3: The account role name for the AssumeRole configuration.


<a id="nestedatt--aws_iam_user"></a>
### Nested Schema for `aws_iam_user`

Optional:

- `access_key_id` (String) S3: The access key ID for the S3 connection.
- `secret_access_key` (String, Sensitive) S3: The secret access key for the S3 connection.


<a id="nestedatt--gateway"></a>
### Nested Schema for `gateway`

Optional:

- `host` (String, Sensitive) MySQL, PostgreSQL: SSH Host
- `key` (String, Sensitive) MySQL, PostgreSQL: SSH Private Key
- `key_passphrase` (String, Sensitive) MySQL, PostgreSQL: SSH Private Key Passphrase
- `password` (String, Sensitive) MySQL, PostgreSQL, Kintone: SSH Password
- `port` (Number, Sensitive) MySQL, PostgreSQL: SSH Port
- `user_name` (String, Sensitive) MySQL, PostgreSQL: SSH User


<a id="nestedatt--ssl"></a>
### Nested Schema for `ssl`

Optional:

- `ca` (String, Sensitive) MySQL, PostgreSQL: CA certificate
- `cert` (String, Sensitive) MySQL, PostgreSQL: Certificate (CRT file)
- `key` (String, Sensitive) MySQL, PostgreSQL: Key (KEY file)
- `ssl_mode` (String) PostgreSQL: SSL connection mode.




## Import

Import is supported using the following syntax:

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import trocco_connection (Resource). For example:

```terraform
import {
  id = "salesforce,1" # id should be <connection_type>,<id> format.
  to = trocco_connection.example
}
```

Using the [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import):

```shell
terraform import trocco_connection.example <connection_type>,<id>
```

