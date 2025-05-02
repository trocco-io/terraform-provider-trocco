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

resource "trocco_connection" "snowflake" {
  connection_type = "snowflake"

  name        = "Snowflake Example"
  description = "This is a Snowflake connection example"

  host        = "exmaple.snowflakecomputing.com"
  auth_method = "user_password"
  user_name   = "dummy_name"
  password    = "dummy_password"
}


resource "trocco_job_definition" "s3_test" {
  description              = ""
  filter_columns           = [
    {
      default                      = null
      json_expand_enabled          = false
      json_expand_keep_base_column = false
      name                         = "name"
      src                          = "name"
      type                         = "string"
    },
    {
      default                      = null
      json_expand_enabled          = false
      json_expand_keep_base_column = false
      name                         = "url"
      src                          = "url"
      type                         = "string"
    },
    {
      default                      = null
      json_expand_enabled          = false
      json_expand_keep_base_column = false
      name                         = "test"
      src                          = "test"
      type                         = "string"
    },
    {
      default                      = null
      json_expand_enabled          = false
      json_expand_keep_base_column = false
      name                         = "asf"
      src                          = "asf"
      type                         = "string"
    },
  ]
  input_option             = {
    s3_input_option = {
      bucket                      = "test_bucket"
      csv_parser                  = {
        allow_extra_columns     = false
        allow_optional_columns  = false
        charset                 = "UTF-8"
        columns                 = [
          {
            name = "name"
            type = "string"
          },
          {
            name = "url"
            type = "string"
          },
          {
            name = "test"
            type = "string"
          },
          {
            name = "asf"
            type = "string"
          },
        ]
        comment_line_marker     = ""
        default_date            = "1970-01-01"
        default_time_zone       = "UTC"
        delimiter               = ","
        escape                  = "\""
        max_quoted_size_limit   = 131072
        newline                 = "CRLF"
        null_string             = ""
        null_string_enabled     = false
        quote                   = "\""
        quotes_in_quoted_fields = "ACCEPT_ONLY_RFC4180_ESCAPED"
        skip_header_lines       = 1
        stop_on_invalid_record  = true
        trim_if_not_quoted      = false
      }
      decompression_type          = "default"
      incremental_loading_enabled = false
      is_skip_header_line         = false
      path_match_pattern          = ""
      path_prefix                 = "dev/000.00.csv"
      region                      = "ap-northeast-1"
      s3_connection_id            = trocco_connection.s3.id
      stop_when_file_not_found    = false
    }
  }
  input_option_type        = "s3"
  is_runnable_concurrently = false
  name                     = "s3 to snowflake"
  output_option            = {
    snowflake_output_option = {
      batch_size              = 50
      database                = "test_database"
      default_time_zone       = "UTC"
      delete_stage_on_error   = false
      empty_field_as_null     = true
      max_retry_wait          = 1800000
      mode                    = "insert"
      retry_limit             = 12
      retry_wait              = 1000
      schema                  = "PUBLIC"
      snowflake_connection_id = trocco_connection.snowflake.id
      table                   = "ewaoiiowe"
      warehouse               = "COMPUTE_WH"
    }
  }
  output_option_type       = "snowflake"
  resource_enhancement     = "custom_spec"
  retry_limit              = 0
}
