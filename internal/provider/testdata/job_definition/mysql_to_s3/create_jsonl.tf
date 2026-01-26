resource "trocco_connection" "test_mysql" {
  connection_type = "mysql"
  name            = "MySQL Example"
  host            = "db.example.com"
  port            = 3306
  user_name       = "root"
  password        = "password"
}

resource "trocco_connection" "test_s3" {
  connection_type = "s3"
  name            = "S3 Example"
  description     = "S3 connection for testing"

  aws_auth_type = "iam_user"
  aws_iam_user = {
    access_key_id     = "AKIAIOSFODNN7EXAMPLE"
    secret_access_key = "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY"
  }
}
resource "trocco_job_definition" "mysql_to_s3_jsonl" {
  name                     = "MySQL to S3 JSONL Test"
  description              = "Test job definition for transferring data from MySQL to S3 with JSONL formatter"
  resource_enhancement     = "medium"
  retry_limit              = 2
  is_runnable_concurrently = true

  filter_columns = [
    {
      name = "id"
      src  = "id"
      type = "long"
    },
    {
      name = "name"
      src  = "name"
      type = "string"
    },
    {
      name = "email"
      src  = "email"
      type = "string"
    },
    {
      name = "created_at"
      src  = "created_at"
      type = "timestamp"
    },
  ]

  input_option_type = "mysql"
  input_option = {
    mysql_input_option = {
      mysql_connection_id         = trocco_connection.test_mysql.id
      database                    = "test_database"
      table                       = "test_table"
      connect_timeout             = 300
      socket_timeout              = 1800
      default_time_zone           = "Asia/Tokyo"
      fetch_rows                  = 1000
      incremental_loading_enabled = false
      input_option_columns = [
        {
          name = "id"
          type = "long"
        },
        {
          name = "name"
          type = "string"
        },
        {
          name = "email"
          type = "string"
        },
        {
          name = "created_at"
          type = "timestamp"
        },
      ]
    }
  }

  output_option_type = "s3"
  output_option = {
    s3_output_option = {
      s3_connection_id         = trocco_connection.test_s3.id
      bucket                   = "test-bucket"
      path_prefix              = "output/json"
      region                   = "us-west-2"
      file_ext                 = "jsonl.gz"
      sequence_format          = ".%03d"
      is_minimum_output_tasks  = false
      multipart_upload_enabled = true
      canned_acl               = "Private"
      formatter_type           = "jsonl"
      encoder_type             = "gzip"
      jsonl_formatter = {
        encoding    = "utf_8"
        newline     = "LF"
        date_format = "%Y-%m-%d"
        timezone    = "UTC"
      }
    }
  }
}
