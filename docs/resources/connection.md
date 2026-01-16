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

### Databricks

```terraform
resource "trocco_connection" "databricks_pat" {
  connection_type = "databricks"

  name                  = "Databricks Example with PAT Auth"
  description           = "This is a Databricks connection example"
  server_hostname       = "example.databricks.com"
  http_path             = "/sql/1.0/warehouses/xxxx-xxxx-xxxx-xxxx"
  auth_type             = "pat"
  personal_access_token = "dapiXXXXXXXXXXXXXXXXXXXX"
}

resource "trocco_connection" "databricks_oauth2" {
  connection_type = "databricks"

  name                 = "Databricks Example with OAuth2"
  description          = "This is a Databricks connection example using OAuth2"
  server_hostname      = "example.databricks.com"
  http_path            = "/sql/1.0/warehouses/xxxx-xxxx-xxxx-xxxx"
  auth_type            = "oauth-m2m"
  oauth2_client_id     = "your-oauth2-client-id"
  oauth2_client_secret = "your-oauth2-client-secret"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `connection_type` (String) The type of the connection. It must be one of `bigquery`, `snowflake`, `gcs`, `google_spreadsheets`, `mysql`, `salesforce`, `s3`, `postgresql`, `google_analytics4`, `kintone`, `databricks`, `mongodb`.
- `name` (String) The name of the connection.

### Optional

- `application_name` (String) GCS: Application name.
- `auth_end_point` (String) Salesforce: Authentication endpoint.
- `auth_method` (String) Snowflake: The authentication method for the Snowflake user. It must be one of `key_pair` or `user_password`. MongoDB: The authentication method. It must be one of `auto`, `mongodb-cr`, or `scram-sha-1`.
- `auth_source` (String) MongoDB: Authentication database name.
- `auth_type` (String) Databricks: The Auth Type for the Databricks connection. It must be one of `pat` or `oauth-m2m`.
- `aws_assume_role` (Attributes) S3: AssumeRole configuration. (see [below for nested schema](#nestedatt--aws_assume_role))
- `aws_auth_type` (String) S3: The authentication type for the S3 connection. It must be one of `iam_user` or `assume_role`.
- `aws_iam_user` (Attributes) S3: IAM User configuration. (see [below for nested schema](#nestedatt--aws_iam_user))
- `basic_auth_password` (String, Sensitive) Kintone: Basic Auth Password
- `basic_auth_username` (String) Kintone: Basic Auth Username
- `connection_string_format` (String) MongoDB: Connection string format. It must be one of `standard` or `dns_seed_list`.
- `description` (String) The description of the connection.
- `domain` (String) Kintone: Domain.
- `driver` (String) Snowflake, MySQL, PostgreSQL: The name of a Database driver.
  - MySQL: null, mysql_connector_java_5_1_49
  - Snowflake: null, snowflake_jdbc_3_14_2, snowflake_jdbc_3_17_0,
  - PostgreSQL: postgresql_42_5_1, postgresql_9_4_1205_jdbc41
- `gateway` (Attributes) MySQL, PostgreSQL, MongoDB: Whether to connect via SSH (see [below for nested schema](#nestedatt--gateway))
- `host` (String) Snowflake, PostgreSQL, MongoDB: The host of a (Snowflake, PostgreSQL, MongoDB) account.
- `http_path` (String) Databricks: The HTTP Path for the Databricks connection.
- `login_method` (String) Kintone: Login Method
- `oauth2_client_id` (String) Databricks: The OAuth2 Client ID for the Databricks connection.
- `oauth2_client_secret` (String, Sensitive) Databricks: The OAuth2 Client Secret for the Databricks connection.
- `password` (String, Sensitive) Snowflake, PostgreSQL, MongoDB: The password for the (Snowflake, PostgreSQL, MongoDB) user.
- `personal_access_token` (String, Sensitive) Databricks: The Personal Access Token for the Databricks connection.
- `port` (Number) MySQL, PostgreSQL, MongoDB: The port of the (MySQL, PostgreSQL, MongoDB) server.
- `private_key` (String, Sensitive) Snowflake: A private key for the Snowflake user.
- `project_id` (String) BigQuery, GCS: A GCP project ID.
- `read_preference` (String) MongoDB: Read preference. It must be one of `primary`, `primaryPreferred`, `secondary`, `secondaryPreferred`, or `nearest`.
- `resource_group_id` (Number) The ID of the resource group the connection belongs to.
- `role` (String) Snowflake: A role attached to the Snowflake user.
- `security_token` (String, Sensitive) Salesforce: Security token.
- `server_hostname` (String) Databricks: The host of a (Databricks) account.
- `service_account_email` (String, Sensitive) GCS: A GCP service account email.
- `service_account_json_key` (String, Sensitive) BigQuery, Google Sheets, Google Analytics4: A GCP service account key.
- `ssl` (Attributes) MySQL, PostgreSQL: SSL configuration. (see [below for nested schema](#nestedatt--ssl))
- `token` (String, Sensitive) Kintone: Token.
- `user_name` (String) Snowflake, PostgreSQL, MongoDB: The name of a (Snowflake, PostgreSQL, MongoDB) user.
- `username` (String) Kintone: The name of a user.

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

- `host` (String, Sensitive) MySQL, PostgreSQL, MongoDB: SSH Host
- `key` (String, Sensitive) MySQL, PostgreSQL, MongoDB: SSH Private Key
- `key_passphrase` (String, Sensitive) MySQL, PostgreSQL, MongoDB: SSH Private Key Passphrase
- `password` (String, Sensitive) MySQL, PostgreSQL, MongoDB, Kintone: SSH Password
- `port` (Number, Sensitive) MySQL, PostgreSQL, MongoDB: SSH Port
- `user_name` (String, Sensitive) MySQL, PostgreSQL, MongoDB: SSH User


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

