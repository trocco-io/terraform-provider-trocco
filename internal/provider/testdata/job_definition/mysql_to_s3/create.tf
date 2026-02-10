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

resource "trocco_job_definition" "mysql_to_s3_csv" {
  name                     = "MySQL to S3 CSV Test"
  description              = "Test job definition for transferring data from MySQL to S3 with CSV formatter"
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
      path_prefix              = "output/data"
      region                   = "ap-northeast-1"
      file_ext                 = "csv.gz"
      sequence_format          = ".%03d"
      canned_acl               = "Private"
      is_minimum_output_tasks  = false
      multipart_upload_enabled = true
      formatter_type           = "csv"
      encoder_type             = "gzip"
      csv_formatter = {
        delimiter           = ","
        escape              = "\\"
        header_line         = true
        charset             = "UTF-8"
        quote_policy        = "ALL"
        newline             = "LF"
        newline_in_field    = "LF"
        null_string_enabled = true
        null_string         = "NULL"
        default_time_zone   = "Asia/Tokyo"
        csv_formatter_column_options_attributes = [
          {
            name     = "created_at"
            format   = "%Y-%m-%d %H:%M:%S"
            timezone = "Asia/Tokyo"
          }
        ]
      }
    }
  }
}
