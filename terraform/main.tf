terraform {
  required_providers {
    trocco = {
      source = "trocco-io/trocco"
    }
  }
}

variable "trocco_api_key" {
  type      = string
  sensitive = true
}

variable "trocco_dev_base_url" {
  type    = string
  default = "https://localhost:4000"
}

provider "trocco" {
  api_key      = "tra_TbJFyRvBBnesvjirCX4TiUbFhb7V3BSaSZKvkkYSyNrQBzWuhJkGTTQVUB"
  dev_base_url = var.trocco_dev_base_url
  region       = "japan"
}

resource "trocco_connection" "test_sftp1" {
  connection_type = "sftp"
  name            = "SFTP Example aaa"
  host            = "sftp.example.com"
  port            = 22
  user_name       = "testuser"
  password        = "password"
}

resource "trocco_job_definition" "sftp_to_bigquery" {
  name                     = "aaa test job_definition"
  description              = "test description"
  resource_group_id        = 2
  retry_limit              = 1
  is_runnable_concurrently = true
  filter_columns = [
    {
      default                      = ""
      format                       = "%Y"
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
  ]
  input_option_type = "sftp"
  input_option = {
    sftp_input_option = {
      sftp_connection_id          = 1
      path_prefix                 = "/data/files/"
      path_match_pattern          = ".*\\.csv$"
      incremental_loading_enabled = true
      stop_when_file_not_found    = false
      decompression_type          = "guess"
      csv_parser = {
        delimiter = ","
        escape    = "\\"
        quote     = "\""
        columns = [
          {
            name = "id"
            type = "long"
          },
          {
            name = "name"
            type = "string"
          },
        ]
      }
    }
  }
  output_option_type = "bigquery"
  output_option = {
    bigquery_output_option = {
      dataset                                  = "test_dataset"
      table                                    = "test_table"
      mode                                     = "append"
      auto_create_dataset                      = true
      timeout_sec                              = 300
      open_timeout_sec                         = 300
      read_timeout_sec                         = 300
      send_timeout_sec                         = 300
      retries                                  = 2
      bigquery_connection_id                   = 2
      location                                 = "us-west1"
      bigquery_output_option_clustering_fields = []
      bigquery_output_option_column_options    = []
      bigquery_output_option_merge_keys        = []
    }
  }
}


resource "trocco_job_definition" "bigquery_to_sftp_csv2" {
  name           = "[updated] BigQuery to SFTP CSV Export"
  description    = "Export BigQuery data to SFTP in CSV format"

  input_option_type  = "bigquery"
  output_option_type = "sftp"

  filter_columns = [
    {
      default                      = ""
      format                       = "%Y"
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
  ]

  input_option = {
    bigquery_input_option = {
      bigquery_connection_id   = 2
      gcs_uri                  = "export_bucket"
      gcs_uri_format           = "bucket"
      query                    = "SELECT user_id, name, email, created_at FROM `dataset.users` WHERE DATE(created_at) = CURRENT_DATE() - 1"
      temp_dataset             = "temp_export"
      location                 = "asia-northeast1"
      is_standard_sql          = true
      cleanup_gcs_files        = true
      file_format              = "CSV"
      cache                    = false
      bigquery_job_wait_second = 600

      columns = [
        {
          name = "user_id"
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
        }
      ]
    }
  }

  output_option = {
    sftp_output_option = {
      sftp_connection_id      = 1
      path_prefix             = "/exports/users/users_$export_date$"
      file_ext                = ".csv"
      is_minimum_output_tasks = false
      encoder_type            = "gzip"

      csv_formatter = {
        delimiter               = ","
        newline                 = "CRLF"
        newline_in_field        = "LF"
        charset                 = "UTF-8"
        quote_policy            = "MINIMAL"
        escape                  = "\\"
        header_line             = true
        null_string_enabled     = true
        null_string             = "NULL"
        default_time_zone       = "Asia/Tokyo"

        csv_formatter_column_options_attributes = [
          {
            name     = "created_at"
            format   = "%Y-%m-%d %H:%M:%S"
            timezone = "Asia/Tokyo"
          }
        ]
      }

      custom_variable_settings = [
        {
          name      = "$export_date$"
          type      = "timestamp"
          quantity  = 1
          unit      = "date"
          direction = "ago"
          format    = "%Y%m%d"
          time_zone = "Asia/Tokyo"
        }
      ]
    }
  }
}
