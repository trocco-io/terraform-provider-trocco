resource "trocco_connection" "test_s3" {
  connection_type = "s3"
  name            = "S3 Example"
  description     = "This is a AWS S3 connection example"
  aws_auth_type   = "iam_user"
  aws_iam_user = {
    access_key_id     = "YOUR_ACCESS_KEY_ID"
    secret_access_key = "YOUR_SECRET_ACCESS_KEY"
  }
}

resource "trocco_connection" "test_bq" {
  connection_type = "bigquery"
  name            = "BigQuery Example"
  project_id      = "example"
  service_account_json_key = <<JSON
  {
    "type": "service_account",
    "project_id": "example-project-id",
    "private_key_id": "example-private-key-id",
    "private_key":"-----BEGIN PRIVATE KEY-----\n..."
  }
  JSON
}

resource "trocco_team" "test" {
  name    = "test"
  members = [
    {
      user_id = 10626
      role    = "team_admin"
    },
  ]
}

resource "trocco_resource_group" "test" {
  name        = "test"
  description = "test"
  teams       = [
    {
      team_id = trocco_team.test.id
      role    = "administrator"
    },
  ]
}

resource "trocco_job_definition" "s3_to_bigquery" {
  name                     = "S3 to BigQuery Test"
  description              = "Test job definition for transferring data from S3 to BigQuery with filter_columns"
  resource_enhancement      = "custom_spec"
  resource_group_id        = trocco_resource_group.test.id
  retry_limit              = 2
  is_runnable_concurrently = true
  
  input_option_type = "s3"
  input_option = {
    s3_input_option = {
      bucket = "test_bucket"
      csv_parser = {
        allow_extra_columns    = false
        allow_optional_columns = false
        charset                = "UTF-8"
        columns = [
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
            format = "%Y-%m-%d %H:%M:%S"
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
      is_skip_header_line         = true
      path_match_pattern          = ""
      path_prefix                 = "data/users.csv"
      region                      = "ap-northeast-1"
      s3_connection_id            = trocco_connection.test_s3.id
      stop_when_file_not_found    = false
    }
  }
  
  filter_columns = [
    {
      default                      = ""
      format                       = ""
      json_expand_enabled          = false
      json_expand_keep_base_column = false
      name                         = "user_id"
      src                          = "id"
      type                         = "long"
    },
    {
      default                      = "Unknown"
      json_expand_enabled          = false
      json_expand_keep_base_column = false
      name                         = "user_name"
      src                          = "name"
      type                         = "string"
    },
    {
      default                      = ""
      json_expand_enabled          = false
      json_expand_keep_base_column = false
      name                         = "user_email"
      src                          = "email"
      type                         = "string"
    },
    {
      default                      = ""
      format                       = "%Y-%m-%d"
      json_expand_enabled          = false
      json_expand_keep_base_column = false
      name                         = "registration_date"
      src                          = "created_at"
      type                         = "timestamp"
    }
  ]
  
  output_option_type = "bigquery"
  output_option = {
    bigquery_output_option = {
      dataset                                  = "test_dataset"
      table                                    = "s3_to_bigquery_test_table"
      mode                                     = "append"
      auto_create_dataset                      = true
      timeout_sec                              = 300
      open_timeout_sec                         = 300
      read_timeout_sec                         = 300
      send_timeout_sec                         = 300
      retries                                  = 3
      bigquery_connection_id                   = trocco_connection.test_bq.id
      location                                 = "US"
      bigquery_output_option_clustering_fields = []
      bigquery_output_option_column_options    = []
      bigquery_output_option_merge_keys        = []
    }
  }
  
  # please create labels if testing in local environment
  # see https://trocco.io/labels#side-nav-labels
  labels = [
    {
      name = "label1"
    },
    {
      name = "label2"
    },
  ]
}
