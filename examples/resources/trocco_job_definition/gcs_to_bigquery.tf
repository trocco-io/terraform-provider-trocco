resource "trocco_job_definition" "gcs_to_bigquery_example" {
  name                     = "example_gcs_to_bigquery"
  description              = ""
  is_runnable_concurrently = false
  retry_limit              = 0
  input_option_type        = "gcs"
  input_option = {
    gcs_input_option = {
      bucket                      = "example_bucket"
      gcs_connection_id           = 1 # please set your gcs connection id
      incremental_loading_enabled = false
      path_prefix                 = "path/to/your/csv_file"
      stop_when_file_not_found    = false
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
            name = "num_col"
            type = "long"
          },
          {
            name = "str_col"
            type = "string"
          },
          {
            format = "%Y-%m-%d %H:%M:%S.%N %z"
            name   = "date_col"
            type   = "timestamp"
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
    }
  }
  filter_columns = [
    {
      default                      = null
      json_expand_columns          = []
      json_expand_enabled          = false
      json_expand_keep_base_column = false
      name                         = "id"
      src                          = "id"
      type                         = "long"
    },
    {
      default                      = null
      json_expand_columns          = []
      json_expand_enabled          = false
      json_expand_keep_base_column = false
      name                         = "num_col"
      src                          = "num_col"
      type                         = "long"
    },
    {
      default                      = null
      json_expand_columns          = []
      json_expand_enabled          = false
      json_expand_keep_base_column = false
      name                         = "str_col"
      src                          = "str_col"
      type                         = "string"
    },
    {
      default                      = null
      format                       = "%Y-%m-%d %H:%M:%S.%N %z"
      json_expand_columns          = []
      json_expand_enabled          = false
      json_expand_keep_base_column = false
      name                         = "date_col"
      src                          = "date_col"
      type                         = "timestamp"
    },
  ]
  output_option_type = "bigquery"
  output_option = {
    bigquery_output_option = {
      auto_create_dataset                      = false
      bigquery_connection_id                   = 1
      bigquery_output_option_clustering_fields = []
      bigquery_output_option_column_options    = []
      bigquery_output_option_merge_keys        = nil
      dataset                                  = "example_dataset"
      location                                 = "US"
      mode                                     = "append"
      open_timeout_sec                         = 300
      read_timeout_sec                         = 300
      retries                                  = 5
      send_timeout_sec                         = 300
      table                                    = "gcs_to_bigquery_example_table"
      template_table                           = ""
      timeout_sec                              = 300
    }
  }
}
