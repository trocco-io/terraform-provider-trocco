resource "trocco_connection" "test_mysql" {
  connection_type = "mysql"
  name            = "MySQL Example"
  host            = "db.example.com"
  port            = 65535
  user_name       = "root"
  password        = "password"
}
resource "trocco_connection" "test_bq" {
  connection_type = "bigquery"
  name        = "BigQuery Example"
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
resource "trocco_team" "test" {
  name        = "test"
  members     = [
    {
      user_id = 10626
      role    = "team_admin"
    },
  ]
}
resource "trocco_resource_group" "test" {
  name        = "test"
  description = "test"
  teams     = [
    {
      team_id = trocco_team.test.id
      role    = "administrator"
    },
  ]
}
resource "trocco_job_definition" "mysql_to_bigquery" {
  name                        = "test job_definition"
  description                 = "test description"
  resource_enhancement        = "large"
  resource_group_id           = trocco_resource_group.test.id
  retry_limit                 = 1
  is_runnable_concurrently    = true
  filter_columns              = [
    {
      default                      = ""
      format = "%Y"
      json_expand_enabled          = false
      json_expand_keep_base_column = false
      name                         = "id"
      src                          = "id"
      type                         = "long"
    },
    {
      default                      = ""
      json_expand_enabled          = false
      json_expand_keep_base_column = false
      name                         = "name"
      src                          = "name"
      type                         = "string"
    },
    {
      default                      = ""
      format                       = ""
      json_expand_enabled          = false
      json_expand_keep_base_column = false
      name                         = "created_at"
      src                          = "created_at"
      type                         = "timestamp"
    },
    {
      default                      = ""
      json_expand_enabled          = true
      json_expand_keep_base_column = true
      name                         = "jsontekitou"
      src                          = ""
      type                         = "json"
      json_expand_columns          = [
        {
          json_path = "path"
          name      = "json_col"
          format = "%Y"
          timezone  = "UTC/ETC"
          type      = "string"
        },
      ]
    }
  ]
  filter_gsub                 = [
    {
      column_name = "regex_col"
      pattern     = "/regex/"
      to          = "replace_string"
    },
  ]
  filter_hashes               = [
    {
      name = "hash_col"
    },
  ]
  filter_masks                = [
    {
      length    = 9
      mask_type = "all"
      name      = "mask_all_string"
    },
    {
      length    = 10
      mask_type = "email"
      name      = "mask_email"
    },
    {
      mask_type = "regex"
      name      = "mask_regex"
      pattern   = "/regex/"
    },
    {
      start_index = 2
      end_index   = 2
      length      = 10
      mask_type   = "all"
      name        = "mail_partial_string"
    },
  ]
  filter_rows                 = {
    condition             = "or"
    filter_row_conditions = [
      {
        argument = "2"
        column   = "bbb"
        operator = "greater_equal"
      },
    ]
  }

  filter_string_transforms    = [
    {
      column_name = "transforms"
      type        = "normalize_nfkc"
    },
  ]
  filter_unixtime_conversions = [
    {
      column_name       = "unix_to_time"
      datetime_format   = "%Y-%m-%d %H:%M:%S.%N"
      datetime_timezone = "Asia/Tokyo"
      kind              = "unixtime_to_timestamp"
      unixtime_unit     = "second"
    },
    {
      column_name       = "unix_to_time_str"
      datetime_format   = "%Y-%m-%d %H:%M:%S.%N %z"
      datetime_timezone = "Etc/UTC"
      kind              = "unixtime_to_string"
      unixtime_unit     = "second"
    },
    {
      column_name       = "time_to_unix"
      datetime_format   = "%Y-%m-%d %H:%M:%S.%N %Z"
      datetime_timezone = "Etc/UTC"
      kind              = "timestamp_to_unixtime"
      unixtime_unit     = "second"
    },
    {
      column_name       = "time_str_to_unix"
      datetime_format   = "%Y-%m-%d %H:%M:%S.%N %z"
      datetime_timezone = "Etc/UTC"
      kind              = "string_to_unixtime"
      unixtime_unit     = "second"
    },
  ]
  input_option_type           = "mysql"
  input_option                = {
    mysql_input_option = {
      connect_timeout             = 300
      database                    = "test_database"
      fetch_rows                  = 1000
      incremental_loading_enabled = false
      table = "test_table"
      default_time_zone           = "Asia/Tokyo"
      use_legacy_datetime_code    = false
      input_option_columns        = [
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
      mysql_connection_id         = trocco_connection.test_mysql.id
      query                       = <<-EOT
                select
                    *
                from
      	          example_table;
      EOT
      socket_timeout              = 1801
    }
  }
  output_option_type          = "bigquery"
  output_option               = {
    bigquery_output_option = {
      dataset                                    = "test_dataset"
      table                                      = "test_table"
      mode                                       = "append"
      auto_create_dataset                        = true
      timeout_sec                                = 300
      open_timeout_sec                           = 300
      read_timeout_sec                           = 300
      send_timeout_sec                           = 300
      retries                                    = 2
      bigquery_connection_id                     = trocco_connection.test_bq.id
      location                                   = "us-west1"
      bigquery_output_option_clustering_fields   = []
      bigquery_output_option_column_options      = []
      bigquery_output_option_merge_keys          = []
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
