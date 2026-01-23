terraform {
  required_providers {
    trocco = {
      source = "registry.terraform.io/trocco-io/trocco"
    }
  }
}

provider "trocco" {
  api_key = "tra_5KvhFRX4fL3eCq2yta9Z23TKk3BTMgbURFodfFVkvYFoUnGhSwEnBdMHUC"
  region  = "japan" # Options: japan, india, korea

  # For development/testing only:
  dev_base_url = "https://localhost:4000"
}

# MySQL Connection
resource "trocco_connection" "mysql" {
  connection_type = "mysql"
  name            = "MySQL Connection"
  host            = "db.example.com"
  port            = 3306
  user_name       = "root"
  password        = "password"
}

# S3 Connection
resource "trocco_connection" "s3" {
  connection_type = "s3"
  name            = "S3 Connection"
  description     = "AWS S3 connection for data export"

  aws_auth_type = "iam_user"
  aws_iam_user = {
    access_key_id     = "AKIAIOSFODNN7EXAMPLE"
    secret_access_key = "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY"
  }
}

# MySQL to S3 Job Definition with CSV Formatter
resource "trocco_job_definition" "mysql_to_s3_csv" {
  name                     = "MySQL to S3 CSV"
  description              = "Transfer data from MySQL to S3 with CSV formatter"
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
      mysql_connection_id         = trocco_connection.mysql.id
      database                    = "production_db"
      table                       = "users"
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
      s3_connection_id         = trocco_connection.s3.id
      bucket                   = "my-data-bucket"
      path_prefix              = "exports/users"
      region                   = "ap-northeast-1"
      file_ext                 = "csv.gz"
      sequence_format          = ".%03d"
      canned_acl               = "private"
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

# MySQL to S3 Job Definition with JSONL Formatter
# resource "trocco_job_definition" "mysql_to_s3_jsonl" {
#   name                     = "MySQL to S3 JSONL"
#   description              = "Transfer data from MySQL to S3 with JSONL formatter"
#   resource_enhancement     = "medium"
#   retry_limit              = 2
#   is_runnable_concurrently = true

#   filter_columns = [
#     {
#       name = "id"
#       src  = "id"
#       type = "long"
#     },
#     {
#       name = "name"
#       src  = "name"
#       type = "string"
#     },
#     {
#       name = "email"
#       src  = "email"
#       type = "string"
#     },
#     {
#       name = "created_at"
#       src  = "created_at"
#       type = "timestamp"
#     },
#   ]

#   input_option_type = "mysql"
#   input_option = {
#     mysql_input_option = {
#       mysql_connection_id         = trocco_connection.mysql.id
#       database                    = "production_db"
#       table                       = "users"
#       connect_timeout             = 300
#       socket_timeout              = 1800
#       default_time_zone           = "Asia/Tokyo"
#       fetch_rows                  = 1000
#       incremental_loading_enabled = false
#       input_option_columns = [
#         {
#           name = "id"
#           type = "long"
#         },
#         {
#           name = "name"
#           type = "string"
#         },
#         {
#           name = "email"
#           type = "string"
#         },
#         {
#           name = "created_at"
#           type = "timestamp"
#         },
#       ]
#     }
#   }

#   output_option_type = "s3"
#   output_option = {
#     s3_output_option = {
#       s3_connection_id         = trocco_connection.s3.id
#       bucket                   = "my-data-bucket"
#       path_prefix              = "exports/users-json"
#       region                   = "us-west-2"
#       file_ext                 = "jsonl.gz"
#       sequence_format          = ".%03d"
#       is_minimum_output_tasks  = false
#       multipart_upload_enabled = true
#       formatter_type           = "jsonl"
#       encoder_type             = "gzip"
#       jsonl_formatter = {
#         encoding    = "UTF-8"
#         newline     = "LF"
#         date_format = "%Y-%m-%d"
#         timezone    = "UTC"
#       }
#     }
#   }
# }
