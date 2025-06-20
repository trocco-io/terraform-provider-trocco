---
page_title: "trocco_job_definition Resource - trocco"
subcategory: ""
description: |-
  Provides a TROCCO job definitions.
---

# trocco_job_definition (Resource)

Provides a TROCCO job definitions.

## Example Usage

### Gcs to Bigquery(CSV file)

Minimum configuration

```terraform
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
      bigquery_output_option_merge_keys        = []
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
```

### Mysql to Bigquery

Minimum configuration

```terraform
resource "trocco_job_definition" "mysql_to_bigquery_example" {
  name                     = "mysql_to_bigquery_example"
  description              = ""
  is_runnable_concurrently = false
  retry_limit              = 0
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
      name                         = "name"
      src                          = "name"
      type                         = "string"
    },
    {
      default                      = null
      format                       = null
      json_expand_columns          = []
      json_expand_enabled          = false
      json_expand_keep_base_column = false
      name                         = "created_at"
      src                          = "created_at"
      type                         = "timestamp"
    },
  ]
  input_option_type = "mysql"
  input_option = {
    mysql_input_option = {
      connect_timeout             = 300
      database                    = "example_database"
      default_time_zone           = ""
      fetch_rows                  = 10000
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
      mysql_connection_id      = 1 // please set your mysql connection id
      query                    = "select * from example_table;"
      socket_timeout           = 1800
      use_legacy_datetime_code = false
    }
  }
  output_option_type = "bigquery"
  output_option = {
    bigquery_output_option = {
      auto_create_dataset                      = false
      bigquery_connection_id                   = 1
      bigquery_output_option_clustering_fields = []
      bigquery_output_option_column_options    = []
      bigquery_output_option_merge_keys        = []
      dataset                                  = "example_dataset"
      location                                 = "US"
      mode                                     = "append"
      open_timeout_sec                         = 300
      read_timeout_sec                         = 300
      retries                                  = 5
      send_timeout_sec                         = 300
      table                                    = "mysql_to_bigquery_example_table"
      template_table                           = ""
      timeout_sec                              = 300
    }
  }
}
```

### Google Spreadsheets to Google Spreadsheets

```terraform
resource "trocco_job_definition" "sheets_to_sheets_example" {
  name                     = "example_sheets_to_sheets"
  description              = "this is an example job definition for transferring data from Google Spreadsheets to Google Spreadsheets"
  is_runnable_concurrently = false
  retry_limit              = 0
  input_option_type        = "google_spreadsheets"
  input_option = {
    google_spreadsheets_input_option = {
      google_spreadsheets_connection_id = 1
      spreadsheets_url                  = "https://docs.google.com/spreadsheets/d/YOUR_SHEETS_ID/edit?gid=0"
      worksheet_title                   = "input-data"
      start_row                         = 2
      start_column                      = "A"
      default_time_zone                 = "Asia/Tokyo"
      null_string                       = ""
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
          name   = "created_at"
          type   = "timestamp"
          format = "%Y-%m-%d %H:%M:%S"
        },
      ]
    }
  }
  output_option_type = "google_spreadsheets"
  output_option = {
    google_spreadsheets_output_option = {
      google_spreadsheets_connection_id = 1
      spreadsheets_id                   = "YOUR_SHEETS_ID"
      worksheet_title                   = "output-data"
      timezone                          = "Asia/Tokyo"
      value_input_option                = "USER_ENTERED"
      mode                              = "replace"
      google_spreadsheets_output_option_sorts = [
        {
          column = "created_at"
          order  = "ascending"
        }
      ]
    }
  }
}
```

### General Setting

```terraform
resource "trocco_job_definition" "general_example" {
  name                     = "example tranfer"
  description              = "example description"
  resource_group_id        = 1
  retry_limit              = 1
  is_runnable_concurrently = true

  # if your account is professional
  resource_enhancement = "medium"
}
```

### FilterColumn

```terraform
resource "trocco_job_definition" "filter_column_example" {
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
      default                      = "default value"
      json_expand_enabled          = false
      json_expand_keep_base_column = false
      name                         = "name"
      src                          = "name"
      type                         = "string"
    },
    {
      default                      = ""
      json_expand_enabled          = false
      json_expand_keep_base_column = false
      name                         = "timestamp"
      src                          = "timestamp"
      type                         = "timestamp"
      format                       = "%Y%m%d"
    },
    {
      default                      = ""
      json_expand_enabled          = true
      json_expand_keep_base_column = true
      name                         = "json_col"
      src                          = "json_src"
      type                         = "json"
      json_expand_columns = [
        {
          json_path = "person.name"
          name      = "json_expand_col"
          timezone  = "UTC/ETC"
          type      = "string"
        },
      ]
    }
  ]
}
```

### FilterRow

```terraform
resource "trocco_job_definition" "filter_row_example" {
  filter_rows = {
    condition = "or"
    filter_row_conditions = [
      {
        argument = "2"
        column   = "col1"
        operator = "greater_equal"
      },
    ]
  }
}
```

### FilterMask

```terraform
resource "trocco_job_definition" "filter_mask_example" {
  filter_masks = [
    {
      name      = "mask_all_string_col"
      mask_type = "all"
      length    = 10
    },
    {
      name      = "mask_email_col"
      mask_type = "email"
      length    = 10
    },
    {
      name      = "mask_regex_col"
      mask_type = "regex"
      pattern   = "/regex/"
    },
    {
      name        = "partial_string"
      length      = 10
      start_index = 2
      end_index   = 2
      mask_type   = "substring"

    },
  ]
}
```

### FilterAddTime

```terraform
resource "trocco_job_definition" "filter_add_time_example" {
  filter_add_time = {
    column_name      = "time"
    time_zone        = "Asia/Tokyo"
    timestamp_format = "%Y-%m-%d %H:%M:%S.%N"
    type             = "timestamp"
  }
}
```

### FilterGsub

```terraform
resource "trocco_job_definition" "filter_gsub_example" {
  filter_gsub = [
    {
      column_name = "regex_col"
      pattern     = "/regex/"
      to          = "replace_string"
    },
  ]
}
```


### FilterStringTransform

```terraform
resource "trocco_job_definition" "filter_string_transform_example" {
  filter_string_transforms = [
    {
      column_name = "transform_col"
      type        = "normalize_nfkc"
    }
  ]
}
```

### FilterHash

```terraform
resource "trocco_job_definition" "filter_hash_example" {
  filter_hashes = [
    {
      name = "hash_col"
    }
  ]
}
```

### FilterUnixtimeConversions

```terraform
resource "trocco_job_definition" "filter_unixtime_conversion_example" {
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
}
```

### CsvParser

```terraform
resource "trocco_job_definition" "csv_parser_example" {
  input_option = {
    # The example is gcs, but it can be applied to file-based input.
    gcs_input_option = {
      csv_parser = {
        delimiter               = ","
        quote                   = "\""
        escape                  = "\""
        skip_header_lines       = 1
        null_string_enabled     = false
        null_string             = ""
        trim_if_not_quoted      = false
        quotes_in_quoted_fields = "ACCEPT_ONLY_RFC4180_ESCAPED"
        comment_line_marker     = ""
        allow_optional_columns  = false
        allow_extra_columns     = false
        max_quoted_size_limit   = 131072
        stop_on_invalid_record  = true
        default_time_zone       = "UTC"
        default_date            = "1970-01-01"
        newline                 = "CRLF"
        charset                 = "UTF-8"
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
            name   = "date_col"
            type   = "timestamp"
            format = "%Y-%m-%d %H:%M:%S.%N %z"
          }
        ]
      }
    }
  }
}
```

### JsonlParser

```terraform
resource "trocco_job_definition" "jsonl_parser_example" {
  input_option = {
    # The example is gcs, but it can be applied to file-based input.
    gcs_input_option = {
      "jsonl_parser" = {
        stop_on_invalid_record = true
        default_time_zone      = "UTC"
        newline                = "LF"
        charset                = "UTF-8"
        columns = [
          {
            name = "id"
            type = "long"
          },
          {
            name = "str_col"
            type = "string"
          },
          {
            name   = "date_col"
            type   = "timestamp"
            format = "%Y-%m-%d %H:%M:%S.%N %z"
          }
        ]
      }
    }
  }
}
```


### LtsvParser

```terraform
resource "trocco_job_definition" "ltsv_parser_example" {
  input_option = {
    # The example is gcs, but it can be applied to file-based input.
    gcs_input_option = {
      "ltsv_parser" = {
        newline = "CRLF"
        charset = "UTF-8"
        columns = [
          {
            name = "id"
            type = "long"
          },
          {
            name = "str_col"
            type = "string"
          },
          {
            name = "date_col",
            type = "timestamp"
          }
        ]
      }
    }
  }
}
```

### ExcelParser

```terraform
resource "trocco_job_definition" "excel_parser_example" {
  input_option = {
    # The example is gcs, but it can be applied to file-based input.
    gcs_input_option = {
      "excel_parser" = {
        default_time_zone = "Asia/Tokyo"
        sheet_name        = "Sheet1"
        skip_header_lines = 1
        columns = [
          {
            name             = "id"
            type             = "long"
            formula_handling = "cashed_value"
          },
          {
            name             = "str_col"
            type             = "string"
            formula_handling = "cashed_value"
          },
          {
            name             = "date_col"
            type             = "timestamp"
            formula_handling = "evaluate"
            format           = "%Y %m %d"
          }
        ]
      }
    }
  }
}
```

### XmlParser

```terraform
resource "trocco_job_definition" "xml_parser_example" {
  input_option = {
    # The example is gcs, but it can be applied to file-based input.
    gcs_input_option = {
      "xml_parser" = {
        root = "root"
        columns = [
          {
            name = "long_col"
            type = "long"
            path = "path/to/long_col"
          },
          {
            name = "str_col"
            type = "string"
            path = "path/to/str_col"
          },
          {
            name     = "timestamp_col"
            type     = "timestamp"
            format   = "%Y-%m-%d %H:%M:%S.%N %z"
            timezone = "UTC"
          }
        ]
      }
    }
  }
}
```


### JsonPathParser

```terraform
resource "trocco_job_definition" "jsonpath_parser_example" {
  input_option = {
    # The example is gcs, but it can be applied to file-based input.
    gcs_input_option = {
      "jsonpath_parser" = {
        root              = "$root"
        default_time_zone = "UTC"
        columns = [
          {
            name = "long_col"
            type = "long"
          },
          {
            name = "str_col"
            type = "string"
          },
          {
            name   = "timestamp_col"
            type   = "timestamp"
            format = "%Y-%m-%d %H:%M:%S.%N %z"
          }
        ]
      }
    }
  }
}
```

### ParquetParser

```terraform
resource "trocco_job_definition" "parquet_parser_example" {
  input_option = {
    # The example is gcs, but it can be applied to file-based input.
    gcs_input_option = {
      "parquet_parser" = {
        columns = [
          {
            name = "long_col"
            type = "long"
          },
          {
            name = "str_col"
            type = "string"
          },
          {
            name   = "timestamp_col"
            type   = "timestamp"
            format = "%Y-%m-%d %H:%M:%S.%N %z"
          }
        ]
      }
    }
  }
}
```

### Decoder

```terraform
resource "trocco_job_definition" "decoder_example" {

  input_option = {
    # The example is gcs, but it can be applied to file-based input.
    gcs_input_option = {
      decoder = {
        match_name = "regex"
      }
    }
  }
}
```

### InputOptions

#### BigqueryInputOption

```terraform
resource "trocco_job_definition" "bigquery_input_example" {
  input_option_type = "bigquery"
  input_option = {
    bigquery_input_option = {
      bigquery_connection_id   = 1
      gcs_uri                  = "test_bucket"
      gcs_uri_format           = "bucket"
      query                    = "SELECT * FROM `test_dataset.test_table`"
      temp_dataset             = "temp_dataset"
      location                 = "asia-northeast1"
      is_standard_sql          = true
      cleanup_gcs_files        = true
      file_format              = "CSV"
      cache                    = true
      bigquery_job_wait_second = 600

      columns = [
        {
          name = "col1__c"
          type = "string"
        }
      ]
    }
  }
}
```

#### MysqlInputOption

```terraform
resource "trocco_job_definition" "mysql_input_example" {
  input_option_type = "mysql"
  input_option = {
    mysql_input_option = {
      connect_timeout             = 300
      socket_timeout              = 1801
      database                    = "test_database"
      fetch_rows                  = 1000
      incremental_loading_enabled = false
      default_time_zone           = "Asia/Tokyo"
      use_legacy_datetime_code    = false
      mysql_connection_id         = 1 # require your mysql connection id
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
      query = <<-EOT
        select
            *
        from
            example_table;
      EOT
    }
  }
}
```

#### GcsInputOption(csv format)

```terraform
resource "trocco_job_definition" "gcs_input_example" {
  input_option_type = "gcs"
  input_option = {
    gcs_input_option = {
      bucket                      = "test-bucket"
      path_prefix                 = "path/to/your_file.csv"
      gcs_connection_id           = 1 # require your gcs connection id
      incremental_loading_enabled = false
      stop_when_file_not_found    = true
      csv_parser = {
        delimiter               = ","
        skip_header_lines       = 1
        trim_if_not_quoted      = false
        quotes_in_quoted_fields = "ACCEPT_ONLY_RFC4180_ESCAPED"
        allow_extra_columns     = false
        allow_optional_columns  = false
        stop_on_invalid_record  = true
        default_date            = "1970-01-01"
        default_time_zone       = "UTC"
        newline                 = "CRLF"
        max_quoted_size_limit   = 131072
        null_string_enabled     = false
        quote                   = "\""
        escape                  = "\""
        null_string             = ""
        comment_line_marker     = ""
        charset                 = "UTF-8"
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
            name   = "date_col"
            type   = "timestamp"
            format = "%Y-%m-%d %H:%M:%S.%N %z"
          },
        ]
      }
    }
  }
}
```

#### SalesforceInputOption
```terraform
# acquisition_method: soql
resource "trocco_job_definition" "salesforce_input_example" {
  input_option_type = "salesforce"
  input_option = {
    salesforce_input_option = {
      columns = [
        {
          name = "col1__c"
          type = "string"
        }
      ]
      include_deleted_or_archived_records = true
      is_convert_type_custom_columns      = false
      object                              = "test_object"
      object_acquisition_method           = "soql"
      soql                                = "select * from test_object"
      salesforce_connection_id            = 1 # pelase set your salesforce connection id
    }
  }
}

# acquisition_method: all_columns
resource "trocco_job_definition" "salesforce_input_example" {
  input_option_type = "salesforce"
  input_option = {
    salesforce_input_option = {
      columns = [
        {
          name = "col1__c"
          type = "string"
        }
      ]
      include_deleted_or_archived_records = true
      is_convert_type_custom_columns      = false
      object                              = "test_object"
      object_acquisition_method           = "all_columns"
      salesforce_connection_id            = 1 # pelase set your salesforce connection id
    }
  }
}
```

#### S3InputOption

```terraform
resource "trocco_job_definition" "s3_input_example" {
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
            name = "col1"
            type = "string"
          },
          {
            name = "col2"
            type = "string"
          },
          {
            name = "col3"
            type = "string"
          },
          {
            name = "col4"
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
      s3_connection_id            = 1 # please set your s3 connection id
      stop_when_file_not_found    = false
    }
  }
}
```

#### PostgresqlInputOption

```terraform
resource "trocco_job_definition" "postgresql_input_example" {
  input_option_type = "postgresql"
  input_option = {
    postgresql_input_option = {
      postgresql_connection_id    = 1 # require your postgresql connection id
      database                    = "test_database"
      schema                      = "public"
      incremental_loading_enabled = false
      connect_timeout             = 300
      socket_timeout              = 1801
      fetch_rows                  = 1000
      default_time_zone           = "Asia/Tokyo"
      query                       = <<-EOT
        select
            *
        from
            example_table;
      EOT
      input_option_column_options : [
        {
          column_name : "test"
          column_value_type : "string"
        }
      ]
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
          name = "test"
          type = "string"
        },
      ]
    }
  }
}
```

#### GoogleAnalytics4InputOption

```terraform
resource "trocco_job_definition" "ga4_input_example" {
  input_option_type = "google_analytics4"
  input_option = {
    google_analytics4_input_option = {
      google_analytics4_connection_id = 1
      property_id                     = "123456789"
      time_series                     = "dateHour"
      start_date                      = "2daysAgo"
      end_date                        = "today"
      google_analytics4_input_option_dimensions = [
        {
          name       = "yyyymm",
          expression = <<-DIM
            {
              "concatenate": {
                "dimensionNames": ["year","month"],
                "delimiter": "-"
              }
            }
          DIM

        }
      ]
      google_analytics4_input_option_metrics = [
        {
          name       = "totalUsers",
          expression = null
        }
      ]
      incremental_loading_enabled = false
      retry_limit                 = 5
      retry_sleep                 = 2
      raise_on_other_row          = false
      limit_of_rows               = 10000
    }
  }
}
```

#### HttpInputOption

```terraform
resource "trocco_job_definition" "http_to_bigquery_maximum" {
  input_option_type = "http"
  input_option = {
    http_input_option = {
      method                                     = "POST"
      url                                        = "http://example.com"
      user_agent                                 = "user-agent-example"
      pager_type                                 = "cursor"
      cursor_request_parameter_cursor_name       = "next_cursor"
      cursor_response_parameter_cursor_json_path = "$.next_cursor"
      request_headers = [
        { key = "Content-Type", value = "application/json", masking = false },
        { key = "Authorization", value = "Bearer example_token", masking = true },
      ]
      request_params = [
        { key = "foo", value = "bar" },
      ]
      success_code = "200"
      jsonl_parser = {
        stop_on_invalid_record = true
        default_time_zone      = "UTC"
        newline                = "LF"
        charset                = "UTF-8"
        columns = [
          {
            name = "id"
            type = "long"
          },
          {
            name = "name"
            type = "string"
          }
        ]
      }
    }
  }
}
```

#### KintoneInputOption

```terraform
resource "trocco_job_definition" "kintone_input_example" {
  input_option_type = "kintone"
  input_option = {
    kintone_input_option = {
      kintone_connection_id = 1 # require your kintone connection id
      app_id                = "236"
      guest_space_id        = "1"
      expand_subtable       = false
      query                 = <<-EOT
        select
            *
        from
            example_table;
      EOT
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
          name   = "test"
          type   = "timestamp"
          format = "%Y-%m-%d %H:%M:%S"
        },
      ]
    }
  }
}
```

#### YahooAdsApiYssInputOption

```terraform
resource "trocco_job_definition" "yahoo_ads_api_yss_input_example" {
  input_option_type = "yahoo_ads_api_yss"
  input_option = {
    yahoo_ads_api_yss_input_option = {
      yahoo_ads_api_connection_id = 1
      base_account_id             = "base-1234"
      account_id                  = "acc-5678"
      service                     = "report_definition_service"
      report_type                 = "ADGROUP"
      start_date                  = "20240501"
      end_date                    = "20240531"
      exclude_zero_impressions    = true
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
          name   = "test"
          type   = "timestamp"
          format = "%Y-%m-%d %H:%M:%S"
        },
      ]
    }
  }
}
```

### OutputOptions

#### BigqueryOutputOption

```terraform
resource "trocco_job_definition" "bigquery_output_example" {
  output_option_type = "bigquery"
  output_option = {
    bigquery_output_option = {
      dataset                         = "test_dataset"
      table                           = "test_table"
      mode                            = "merge"
      auto_create_dataset             = true
      timeout_sec                     = 300
      open_timeout_sec                = 300
      read_timeout_sec                = 300
      send_timeout_sec                = 300
      retries                         = 0
      bigquery_connection_id          = 1 # require your bigquery connection id
      partitioning_type               = "time_unit_column"
      time_partitioning_type          = "DAY"
      time_partitioning_field         = "created_at"
      time_partitioning_expiration_ms = 10000
      location                        = "US"
      bigquery_output_option_merge_keys = [
        "id"
      ]
    }
  }
}
```

#### SalesforceOutputOption

```terraform
# upsert
resource "trocco_job_definition" "salesforce_output_example" {
  output_option_type = "salesforce"
  output_option = {
    salesforce_output_option = {
      action_type              = "upsert"
      upsert_key               = "id"
      api_version              = "55.0"
      ignore_nulls             = true
      object                   = "test_object"
      salesforce_connection_id = 1 # please set your salesforce connection id
      throw_if_failed          = false
    }
  }
}

# insert
resource "trocco_job_definition" "salesforce_output_example" {
  output_option_type = "salesforce"
  output_option = {
    salesforce_output_option = {
      action_type              = "insert"
      api_version              = "55.0"
      ignore_nulls             = true
      object                   = "test_object"
      salesforce_connection_id = 1 # please set your salesforce connection id
      throw_if_failed          = false
    }
  }
}
```


### Label

```terraform
resource "trocco_job_definition" "labels" {
  labels = [
    {
      name = "aaa"
    }
  ]
}
```


### Notification

```terraform
resource "trocco_job_definition" "notifications" {
  notifications = [
    {
      destination_type  = "email"
      email_id          = 1 # require your email id
      message           = "email failed"
      notification_type = "job"
      notify_when       = "failed"
    },
    {
      destination_type  = "email"
      email_id          = 1 # require your email id
      message           = "email1"
      notification_type = "job"
      notify_when       = "finished"
    },
    {
      destination_type  = "email"
      email_id          = 1 # require your email id
      message           = "record count email skipped"
      notification_type = "record"
      record_count      = 10
      record_operator   = "below"
      record_type       = "skipped"
    },
    {
      destination_type  = "email"
      email_id          = 1 # require your email id
      message           = "time alert email"
      minutes           = 10
      notification_type = "exec_time"
    },
    {
      destination_type  = "slack"
      message           = "record count slack transfer"
      notification_type = "record"
      record_count      = 10
      record_operator   = "below"
      record_type       = "transfer"
      slack_channel_id  = 1 # require your slack id
    },
    {
      destination_type  = "slack"
      message           = "slack 1"
      notification_type = "job"
      notify_when       = "finished"
      slack_channel_id  = 1 # require your slack id
    },
    {
      destination_type  = "slack"
      message           = "slack failed"
      notification_type = "job"
      notify_when       = "failed"
      slack_channel_id  = 1 # require your slack id
    },
    {
      destination_type  = "slack"
      message           = "time alert slack"
      minutes           = 10
      notification_type = "exec_time"
      slack_channel_id  = 1 # require your slack id
    },
  ]
}
```

### Schedule

```terraform
resource "trocco_job_definition" "schedules" {
  schedules = [
    {
      day       = 1
      frequency = "monthly"
      hour      = 1
      minute    = 1
      time_zone = "Australia/Sydney"
    },
    {
      day_of_week = 0
      frequency   = "weekly"
      hour        = 1
      minute      = 1
      time_zone   = "Australia/Sydney"
    },
    {
      frequency = "daily"
      hour      = 1
      minute    = 1
      time_zone = "Australia/Sydney"
    },
    {
      frequency = "hourly"
      minute    = 1
      time_zone = "Australia/Sydney"
    },
  ]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `filter_columns` (Attributes List) (see [below for nested schema](#nestedatt--filter_columns))
- `input_option` (Attributes) (see [below for nested schema](#nestedatt--input_option))
- `input_option_type` (String) Input option type.
- `name` (String) Name of the job definition. It must be less than 256 characters
- `output_option` (Attributes) (see [below for nested schema](#nestedatt--output_option))
- `output_option_type` (String) Output option type.

### Optional

- `description` (String) Description of the job definition.
- `filter_add_time` (Attributes) Transfer Date Column Setting (see [below for nested schema](#nestedatt--filter_add_time))
- `filter_gsub` (Attributes List) String Regular Expression Replacement (see [below for nested schema](#nestedatt--filter_gsub))
- `filter_hashes` (Attributes List) Column hashing (see [below for nested schema](#nestedatt--filter_hashes))
- `filter_masks` (Attributes List) Filter masks to be attached to the job definition (see [below for nested schema](#nestedatt--filter_masks))
- `filter_rows` (Attributes) Filter settings (see [below for nested schema](#nestedatt--filter_rows))
- `filter_string_transforms` (Attributes List) Character string conversion (see [below for nested schema](#nestedatt--filter_string_transforms))
- `filter_unixtime_conversions` (Attributes List) UNIX time conversion (see [below for nested schema](#nestedatt--filter_unixtime_conversions))
- `is_runnable_concurrently` (Boolean) Specifies whether or not to run a job if another job with the same job definition is running at the time the job is run
- `labels` (Attributes Set) Labels to be attached to the job definition (see [below for nested schema](#nestedatt--labels))
- `notifications` (Attributes Set) Notifications to be attached to the job definition (see [below for nested schema](#nestedatt--notifications))
- `resource_enhancement` (String) Resource size to be used when executing the job. If not specified, the resource size specified in the transfer settings is applied. The value that can be specified varies depending on the connector. (This parameter is available only in the Professional plan.
- `resource_group_id` (Number) ID of the resource group to which the job definition belongs
- `retry_limit` (Number) Maximum number of retries. if set 0, the job will not be retried
- `schedules` (Attributes Set) Schedules to be attached to the job definition (see [below for nested schema](#nestedatt--schedules))

### Read-Only

- `id` (Number) The ID of the job definition

<a id="nestedatt--filter_columns"></a>
### Nested Schema for `filter_columns`

Required:

- `name` (String) Column name
- `src` (String) Column name in source
- `type` (String) column type

Optional:

- `default` (String) Default value. For existing columns, this value will be inserted only if input is null. For new columns, this value is inserted for all.
- `format` (String) date/time format
- `json_expand_columns` (Attributes List) (see [below for nested schema](#nestedatt--filter_columns--json_expand_columns))
- `json_expand_enabled` (Boolean) Flag whether to expand JSON
- `json_expand_keep_base_column` (Boolean) Flag whether to keep the base column

<a id="nestedatt--filter_columns--json_expand_columns"></a>
### Nested Schema for `filter_columns.json_expand_columns`

Required:

- `json_path` (String) JSON path. To extract id and age from a JSON column such as {'{“id”: 10, “person”: {“age”: 30}}'}, specify id and person.age in the JSON path, respectively.
- `name` (String) Column name
- `type` (String) Column type

Optional:

- `format` (String) date/time format
- `timezone` (String) time zone



<a id="nestedatt--input_option"></a>
### Nested Schema for `input_option`

Optional:

- `bigquery_input_option` (Attributes) Attributes about source bigquery (see [below for nested schema](#nestedatt--input_option--bigquery_input_option))
- `gcs_input_option` (Attributes) Attributes about source GCS (see [below for nested schema](#nestedatt--input_option--gcs_input_option))
- `google_analytics4_input_option` (Attributes) Attributes about source Google Analytics 4 (see [below for nested schema](#nestedatt--input_option--google_analytics4_input_option))
- `google_spreadsheets_input_option` (Attributes) Attributes about source Google Spreadsheets (see [below for nested schema](#nestedatt--input_option--google_spreadsheets_input_option))
- `http_input_option` (Attributes) Attributes about source HTTP (see [below for nested schema](#nestedatt--input_option--http_input_option))
- `kintone_input_option` (Attributes) Attributes of source kintone (see [below for nested schema](#nestedatt--input_option--kintone_input_option))
- `mysql_input_option` (Attributes) Attributes of source mysql (see [below for nested schema](#nestedatt--input_option--mysql_input_option))
- `postgresql_input_option` (Attributes) Attributes of source postgresql (see [below for nested schema](#nestedatt--input_option--postgresql_input_option))
- `s3_input_option` (Attributes) Attributes about source S3 (see [below for nested schema](#nestedatt--input_option--s3_input_option))
- `salesforce_input_option` (Attributes) Attributes about source Salesforce (see [below for nested schema](#nestedatt--input_option--salesforce_input_option))
- `snowflake_input_option` (Attributes) Attributes about source snowflake (see [below for nested schema](#nestedatt--input_option--snowflake_input_option))
- `yahoo_ads_api_yss_input_option` (Attributes) Attributes of source yahoo_ads_api_yss (see [below for nested schema](#nestedatt--input_option--yahoo_ads_api_yss_input_option))

<a id="nestedatt--input_option--bigquery_input_option"></a>
### Nested Schema for `input_option.bigquery_input_option`

Required:

- `bigquery_connection_id` (Number) Id of bigquery connection
- `columns` (Attributes List) List of columns to be retrieved and their types (see [below for nested schema](#nestedatt--input_option--bigquery_input_option--columns))
- `gcs_uri` (String) GCS URI
- `query` (String) Query
- `temp_dataset` (String) Temporary dataset name

Optional:

- `bigquery_job_wait_second` (Number) Wait time in seconds until bigquery job is completed
- `cache` (Boolean) Flag whether query cache is enabled
- `cleanup_gcs_files` (Boolean) Flag whether temporary GCS files should be cleaned up
- `custom_variable_settings` (Attributes List) (see [below for nested schema](#nestedatt--input_option--bigquery_input_option--custom_variable_settings))
- `decoder` (Attributes) (see [below for nested schema](#nestedatt--input_option--bigquery_input_option--decoder))
- `file_format` (String) File format of temporary GCS files
- `gcs_uri_format` (String) Format of GCS URI
- `is_standard_sql` (Boolean) Flag whether standard SQL is enabled
- `location` (String) Location of bigquery job

<a id="nestedatt--input_option--bigquery_input_option--columns"></a>
### Nested Schema for `input_option.bigquery_input_option.columns`

Required:

- `name` (String) Column name
- `type` (String) Column type.

Optional:

- `format` (String) format


<a id="nestedatt--input_option--bigquery_input_option--custom_variable_settings"></a>
### Nested Schema for `input_option.bigquery_input_option.custom_variable_settings`

Required:

- `name` (String) Custom variable name. It must start and end with `$`
- `type` (String) Custom variable type. The following types are supported: `string`, `timestamp`, `timestamp_runtime`

Optional:

- `direction` (String) Direction of the diff from context_time. The following directions are supported: `ago`, `later`. Required in `timestamp` and `timestamp_runtime` types
- `format` (String) Format used to replace variables. Required in `timestamp` and `timestamp_runtime` types
- `quantity` (Number) Quantity used to calculate diff from context_time. Required in `timestamp` and `timestamp_runtime` types
- `time_zone` (String) Time zone used to format the timestamp. Required in `timestamp` and `timestamp_runtime` types
- `unit` (String) Time unit used to calculate diff from context_time. The following units are supported: `hour`, `date`, `month`. Required in `timestamp` and `timestamp_runtime` types
- `value` (String) Fixed string which will replace variables at runtime. Required in `string` type


<a id="nestedatt--input_option--bigquery_input_option--decoder"></a>
### Nested Schema for `input_option.bigquery_input_option.decoder`

Optional:

- `match_name` (String) Relative path after decompression (regular expression). If not entered, all data in the compressed file will be transferred.



<a id="nestedatt--input_option--gcs_input_option"></a>
### Nested Schema for `input_option.gcs_input_option`

Required:

- `bucket` (String) Bucket name
- `gcs_connection_id` (Number) Id of GCS connection
- `path_prefix` (String) Path prefix

Optional:

- `csv_parser` (Attributes) For files in CSV format, this parameter is required (see [below for nested schema](#nestedatt--input_option--gcs_input_option--csv_parser))
- `custom_variable_settings` (Attributes List) (see [below for nested schema](#nestedatt--input_option--gcs_input_option--custom_variable_settings))
- `decoder` (Attributes) (see [below for nested schema](#nestedatt--input_option--gcs_input_option--decoder))
- `decompression_type` (String) Decompression type
- `excel_parser` (Attributes) For files in excel format, this parameter is required. (see [below for nested schema](#nestedatt--input_option--gcs_input_option--excel_parser))
- `incremental_loading_enabled` (Boolean) If it is true, to be incremental loading. If it is false, to be all record loading
- `jsonl_parser` (Attributes) For files in JSONL format, this parameter is required (see [below for nested schema](#nestedatt--input_option--gcs_input_option--jsonl_parser))
- `jsonpath_parser` (Attributes) For files in jsonpath format, this parameter is required. (see [below for nested schema](#nestedatt--input_option--gcs_input_option--jsonpath_parser))
- `last_path` (String) Last path transferred. It is only enabled when incremental loading is true. When updating differences, data behind in lexicographic order from the path specified here is transferred. If the form is blank, the data is transferred from the beginning. Do not change this value unless there is a special reason. Duplicate data may occur.
- `ltsv_parser` (Attributes) For files in LTSV format, this parameter is required. (see [below for nested schema](#nestedatt--input_option--gcs_input_option--ltsv_parser))
- `parquet_parser` (Attributes) For files in parquet format, this parameter is required. (see [below for nested schema](#nestedatt--input_option--gcs_input_option--parquet_parser))
- `stop_when_file_not_found` (Boolean) Flag whether the transfer should continue if the file does not exist in the specified path
- `xml_parser` (Attributes) For files in xml format, this parameter is required. (see [below for nested schema](#nestedatt--input_option--gcs_input_option--xml_parser))

<a id="nestedatt--input_option--gcs_input_option--csv_parser"></a>
### Nested Schema for `input_option.gcs_input_option.csv_parser`

Required:

- `columns` (Attributes List) (see [below for nested schema](#nestedatt--input_option--gcs_input_option--csv_parser--columns))

Optional:

- `allow_extra_columns` (Boolean) If true, ignore the column. If false, treat as invalid record.
- `allow_optional_columns` (Boolean) If true, NULL-complete the missing columns. If false, treat as invalid record.
- `charset` (String) Character set
- `comment_line_marker` (String) Comment line marker. Skip if this character is at the beginning of a line
- `default_date` (String) Default date
- `default_time_zone` (String) Default time zone
- `delimiter` (String) Delimiter
- `escape` (String) Escape character
- `max_quoted_size_limit` (Number) Maximum amount of data that can be enclosed in quotation marks.
- `newline` (String) Newline character
- `null_string` (String) Replacement source string to be converted to NULL
- `null_string_enabled` (Boolean) Flag whether or not to set the string to be replaced by NULL
- `quote` (String) Quote character
- `quotes_in_quoted_fields` (String) Processing method for irregular quarts
- `skip_header_lines` (Number) Number of header lines to skip
- `stop_on_invalid_record` (Boolean) Flag whether or not to abort the transfer if an invalid record is found.
- `trim_if_not_quoted` (Boolean) Flag whether or not to remove spaces from the value if it is not quoted

<a id="nestedatt--input_option--gcs_input_option--csv_parser--columns"></a>
### Nested Schema for `input_option.gcs_input_option.csv_parser.columns`

Required:

- `name` (String) Column name
- `type` (String) Column type

Optional:

- `date` (String) Date
- `format` (String) Format of the column



<a id="nestedatt--input_option--gcs_input_option--custom_variable_settings"></a>
### Nested Schema for `input_option.gcs_input_option.custom_variable_settings`

Required:

- `name` (String) Custom variable name. It must start and end with `$`
- `type` (String) Custom variable type. The following types are supported: `string`, `timestamp`, `timestamp_runtime`

Optional:

- `direction` (String) Direction of the diff from context_time. The following directions are supported: `ago`, `later`. Required in `timestamp` and `timestamp_runtime` types
- `format` (String) Format used to replace variables. Required in `timestamp` and `timestamp_runtime` types
- `quantity` (Number) Quantity used to calculate diff from context_time. Required in `timestamp` and `timestamp_runtime` types
- `time_zone` (String) Time zone used to format the timestamp. Required in `timestamp` and `timestamp_runtime` types
- `unit` (String) Time unit used to calculate diff from context_time. The following units are supported: `hour`, `date`, `month`. Required in `timestamp` and `timestamp_runtime` types
- `value` (String) Fixed string which will replace variables at runtime. Required in `string` type


<a id="nestedatt--input_option--gcs_input_option--decoder"></a>
### Nested Schema for `input_option.gcs_input_option.decoder`

Optional:

- `match_name` (String) Relative path after decompression (regular expression). If not entered, all data in the compressed file will be transferred.


<a id="nestedatt--input_option--gcs_input_option--excel_parser"></a>
### Nested Schema for `input_option.gcs_input_option.excel_parser`

Required:

- `columns` (Attributes List) List of columns to be retrieved and their types (see [below for nested schema](#nestedatt--input_option--gcs_input_option--excel_parser--columns))
- `sheet_name` (String) Sheet name

Optional:

- `default_time_zone` (String) Default time zone
- `skip_header_lines` (Number) Number of header lines to skip

<a id="nestedatt--input_option--gcs_input_option--excel_parser--columns"></a>
### Nested Schema for `input_option.gcs_input_option.excel_parser.columns`

Required:

- `formula_handling` (String) Formula handling
- `name` (String) Column name
- `type` (String) Column type

Optional:

- `format` (String) Format of the column.



<a id="nestedatt--input_option--gcs_input_option--jsonl_parser"></a>
### Nested Schema for `input_option.gcs_input_option.jsonl_parser`

Required:

- `columns` (Attributes List) List of columns to be retrieved and their types (see [below for nested schema](#nestedatt--input_option--gcs_input_option--jsonl_parser--columns))

Optional:

- `charset` (String) Character set
- `default_time_zone` (String) Default time zone
- `newline` (String) Newline character
- `stop_on_invalid_record` (Boolean) Flag whether the transfer should stop if an invalid record is found

<a id="nestedatt--input_option--gcs_input_option--jsonl_parser--columns"></a>
### Nested Schema for `input_option.gcs_input_option.jsonl_parser.columns`

Required:

- `name` (String) Column name
- `type` (String) Column type

Optional:

- `format` (String) Format of the column
- `time_zone` (String) time zone



<a id="nestedatt--input_option--gcs_input_option--jsonpath_parser"></a>
### Nested Schema for `input_option.gcs_input_option.jsonpath_parser`

Required:

- `columns` (Attributes List) (see [below for nested schema](#nestedatt--input_option--gcs_input_option--jsonpath_parser--columns))
- `root` (String) JSONPath

Optional:

- `default_time_zone` (String) Default time zone

<a id="nestedatt--input_option--gcs_input_option--jsonpath_parser--columns"></a>
### Nested Schema for `input_option.gcs_input_option.jsonpath_parser.columns`

Required:

- `name` (String) Column name
- `type` (String) Column type

Optional:

- `format` (String) Format of the column.
- `time_zone` (String) time zone



<a id="nestedatt--input_option--gcs_input_option--ltsv_parser"></a>
### Nested Schema for `input_option.gcs_input_option.ltsv_parser`

Required:

- `columns` (Attributes List) List of columns to be retrieved and their types (see [below for nested schema](#nestedatt--input_option--gcs_input_option--ltsv_parser--columns))

Optional:

- `charset` (String) Character set
- `newline` (String) Newline character

<a id="nestedatt--input_option--gcs_input_option--ltsv_parser--columns"></a>
### Nested Schema for `input_option.gcs_input_option.ltsv_parser.columns`

Required:

- `name` (String) Column name
- `type` (String) Column type

Optional:

- `format` (String) Format of the column.



<a id="nestedatt--input_option--gcs_input_option--parquet_parser"></a>
### Nested Schema for `input_option.gcs_input_option.parquet_parser`

Required:

- `columns` (Attributes List) List of columns to be retrieved and their types (see [below for nested schema](#nestedatt--input_option--gcs_input_option--parquet_parser--columns))

<a id="nestedatt--input_option--gcs_input_option--parquet_parser--columns"></a>
### Nested Schema for `input_option.gcs_input_option.parquet_parser.columns`

Required:

- `name` (String) Column name
- `type` (String) Column type

Optional:

- `format` (String) Format of the column.



<a id="nestedatt--input_option--gcs_input_option--xml_parser"></a>
### Nested Schema for `input_option.gcs_input_option.xml_parser`

Required:

- `columns` (Attributes List) (see [below for nested schema](#nestedatt--input_option--gcs_input_option--xml_parser--columns))
- `root` (String) Root element

<a id="nestedatt--input_option--gcs_input_option--xml_parser--columns"></a>
### Nested Schema for `input_option.gcs_input_option.xml_parser.columns`

Required:

- `name` (String) Column name
- `path` (String) XPath
- `type` (String) Column type

Optional:

- `format` (String) Format of the column.
- `timezone` (String) time zone




<a id="nestedatt--input_option--google_analytics4_input_option"></a>
### Nested Schema for `input_option.google_analytics4_input_option`

Required:

- `google_analytics4_connection_id` (Number) Id of Google Analytics 4 connection
- `google_analytics4_input_option_metrics` (Attributes List) (see [below for nested schema](#nestedatt--input_option--google_analytics4_input_option--google_analytics4_input_option_metrics))
- `input_option_columns` (Attributes List) List of columns to be retrieved and their types (see [below for nested schema](#nestedatt--input_option--google_analytics4_input_option--input_option_columns))
- `property_id` (String)
- `time_series` (String)

Optional:

- `custom_variable_settings` (Attributes List) (see [below for nested schema](#nestedatt--input_option--google_analytics4_input_option--custom_variable_settings))
- `end_date` (String)
- `google_analytics4_input_option_dimensions` (Attributes List) (see [below for nested schema](#nestedatt--input_option--google_analytics4_input_option--google_analytics4_input_option_dimensions))
- `incremental_loading_enabled` (Boolean)
- `limit_of_rows` (Number)
- `raise_on_other_row` (Boolean)
- `retry_limit` (Number)
- `retry_sleep` (Number) Sleep time in seconds between retries
- `start_date` (String)

<a id="nestedatt--input_option--google_analytics4_input_option--google_analytics4_input_option_metrics"></a>
### Nested Schema for `input_option.google_analytics4_input_option.google_analytics4_input_option_metrics`

Required:

- `name` (String) Metric name

Optional:

- `expression` (String)


<a id="nestedatt--input_option--google_analytics4_input_option--input_option_columns"></a>
### Nested Schema for `input_option.google_analytics4_input_option.input_option_columns`

Required:

- `name` (String) Column name
- `type` (String) Column type


<a id="nestedatt--input_option--google_analytics4_input_option--custom_variable_settings"></a>
### Nested Schema for `input_option.google_analytics4_input_option.custom_variable_settings`

Required:

- `name` (String) Custom variable name. It must start and end with `$`
- `type` (String) Custom variable type. The following types are supported: `string`, `timestamp`, `timestamp_runtime`

Optional:

- `direction` (String) Direction of the diff from context_time. The following directions are supported: `ago`, `later`. Required in `timestamp` and `timestamp_runtime` types
- `format` (String) Format used to replace variables. Required in `timestamp` and `timestamp_runtime` types
- `quantity` (Number) Quantity used to calculate diff from context_time. Required in `timestamp` and `timestamp_runtime` types
- `time_zone` (String) Time zone used to format the timestamp. Required in `timestamp` and `timestamp_runtime` types
- `unit` (String) Time unit used to calculate diff from context_time. The following units are supported: `hour`, `date`, `month`. Required in `timestamp` and `timestamp_runtime` types
- `value` (String) Fixed string which will replace variables at runtime. Required in `string` type


<a id="nestedatt--input_option--google_analytics4_input_option--google_analytics4_input_option_dimensions"></a>
### Nested Schema for `input_option.google_analytics4_input_option.google_analytics4_input_option_dimensions`

Required:

- `name` (String) Dimension name

Optional:

- `expression` (String)



<a id="nestedatt--input_option--google_spreadsheets_input_option"></a>
### Nested Schema for `input_option.google_spreadsheets_input_option`

Required:

- `default_time_zone` (String) Default time zone
- `google_spreadsheets_connection_id` (Number) Id of Snowflake connection
- `input_option_columns` (Attributes List) List of columns to be retrieved and their types (see [below for nested schema](#nestedatt--input_option--google_spreadsheets_input_option--input_option_columns))
- `spreadsheets_url` (String) URL of the Google Sheets
- `start_column` (String) Column to start reading data
- `start_row` (Number) Row number to start reading data
- `worksheet_title` (String) Title of the worksheet

Optional:

- `custom_variable_settings` (Attributes List) (see [below for nested schema](#nestedatt--input_option--google_spreadsheets_input_option--custom_variable_settings))
- `null_string` (String) String to be treated as NULL

<a id="nestedatt--input_option--google_spreadsheets_input_option--input_option_columns"></a>
### Nested Schema for `input_option.google_spreadsheets_input_option.input_option_columns`

Required:

- `name` (String) Column name
- `type` (String) Column type

Optional:

- `format` (String) Column format


<a id="nestedatt--input_option--google_spreadsheets_input_option--custom_variable_settings"></a>
### Nested Schema for `input_option.google_spreadsheets_input_option.custom_variable_settings`

Required:

- `name` (String) Custom variable name. It must start and end with `$`
- `type` (String) Custom variable type. The following types are supported: `string`, `timestamp`, `timestamp_runtime`

Optional:

- `direction` (String) Direction of the diff from context_time. The following directions are supported: `ago`, `later`. Required in `timestamp` and `timestamp_runtime` types
- `format` (String) Format used to replace variables. Required in `timestamp` and `timestamp_runtime` types
- `quantity` (Number) Quantity used to calculate diff from context_time. Required in `timestamp` and `timestamp_runtime` types
- `time_zone` (String) Time zone used to format the timestamp. Required in `timestamp` and `timestamp_runtime` types
- `unit` (String) Time unit used to calculate diff from context_time. The following units are supported: `hour`, `date`, `month`. Required in `timestamp` and `timestamp_runtime` types
- `value` (String) Fixed string which will replace variables at runtime. Required in `string` type



<a id="nestedatt--input_option--http_input_option"></a>
### Nested Schema for `input_option.http_input_option`

Required:

- `method` (String) HTTP method (GET or POST)
- `url` (String) URL to fetch data from

Optional:

- `charset` (String) Character set of the response
- `csv_parser` (Attributes) For files in CSV format, this parameter is required (see [below for nested schema](#nestedatt--input_option--http_input_option--csv_parser))
- `cursor_request_parameter_cursor_name` (String) Parameter name for cursor-based pagination
- `cursor_request_parameter_limit_name` (String) Parameter name for limit in cursor-based pagination
- `cursor_request_parameter_limit_value` (String) Value for limit parameter in cursor-based pagination
- `cursor_response_parameter_cursor_json_path` (String) JSONPath to extract cursor value from response
- `custom_variable_settings` (Attributes List) (see [below for nested schema](#nestedatt--input_option--http_input_option--custom_variable_settings))
- `excel_parser` (Attributes) For files in excel format, this parameter is required. (see [below for nested schema](#nestedatt--input_option--http_input_option--excel_parser))
- `jsonl_parser` (Attributes) For files in JSONL format, this parameter is required (see [below for nested schema](#nestedatt--input_option--http_input_option--jsonl_parser))
- `jsonpath_parser` (Attributes) For files in jsonpath format, this parameter is required. (see [below for nested schema](#nestedatt--input_option--http_input_option--jsonpath_parser))
- `ltsv_parser` (Attributes) For files in LTSV format, this parameter is required. (see [below for nested schema](#nestedatt--input_option--http_input_option--ltsv_parser))
- `max_retries` (Number) Maximum number of retry attempts
- `open_timeout` (Number) Timeout for opening connection in seconds
- `pager_from_param` (String) Parameter name for the starting offset/page
- `pager_pages` (Number) Number of requests to fetch
- `pager_start` (Number) Starting page number
- `pager_step` (Number) Step size for pagination
- `pager_to_param` (String) Parameter name for the ending offset/page
- `pager_type` (String) Type of pagination (offset, cursor, disable)
- `parquet_parser` (Attributes) For files in parquet format, this parameter is required. (see [below for nested schema](#nestedatt--input_option--http_input_option--parquet_parser))
- `read_timeout` (Number) Timeout for reading response in seconds
- `request_body` (String) Request body for POST/PUT requests
- `request_headers` (Attributes Set) HTTP headers to include in the request (see [below for nested schema](#nestedatt--input_option--http_input_option--request_headers))
- `request_interval` (Number) Interval between requests in seconds
- `request_params` (Attributes Set) Request parameters to include in the URL (see [below for nested schema](#nestedatt--input_option--http_input_option--request_params))
- `retry_interval` (Number) Interval between retry attempts in seconds
- `success_code` (String) HTTP status code that indicate success
- `user_agent` (String) User agent string to use for requests
- `xml_parser` (Attributes) For files in xml format, this parameter is required. (see [below for nested schema](#nestedatt--input_option--http_input_option--xml_parser))

<a id="nestedatt--input_option--http_input_option--csv_parser"></a>
### Nested Schema for `input_option.http_input_option.csv_parser`

Required:

- `columns` (Attributes List) (see [below for nested schema](#nestedatt--input_option--http_input_option--csv_parser--columns))

Optional:

- `allow_extra_columns` (Boolean) If true, ignore the column. If false, treat as invalid record.
- `allow_optional_columns` (Boolean) If true, NULL-complete the missing columns. If false, treat as invalid record.
- `charset` (String) Character set
- `comment_line_marker` (String) Comment line marker. Skip if this character is at the beginning of a line
- `default_date` (String) Default date
- `default_time_zone` (String) Default time zone
- `delimiter` (String) Delimiter
- `escape` (String) Escape character
- `max_quoted_size_limit` (Number) Maximum amount of data that can be enclosed in quotation marks.
- `newline` (String) Newline character
- `null_string` (String) Replacement source string to be converted to NULL
- `null_string_enabled` (Boolean) Flag whether or not to set the string to be replaced by NULL
- `quote` (String) Quote character
- `quotes_in_quoted_fields` (String) Processing method for irregular quarts
- `skip_header_lines` (Number) Number of header lines to skip
- `stop_on_invalid_record` (Boolean) Flag whether or not to abort the transfer if an invalid record is found.
- `trim_if_not_quoted` (Boolean) Flag whether or not to remove spaces from the value if it is not quoted

<a id="nestedatt--input_option--http_input_option--csv_parser--columns"></a>
### Nested Schema for `input_option.http_input_option.csv_parser.columns`

Required:

- `name` (String) Column name
- `type` (String) Column type

Optional:

- `date` (String) Date
- `format` (String) Format of the column



<a id="nestedatt--input_option--http_input_option--custom_variable_settings"></a>
### Nested Schema for `input_option.http_input_option.custom_variable_settings`

Required:

- `name` (String) Custom variable name. It must start and end with `$`
- `type` (String) Custom variable type. The following types are supported: `string`, `timestamp`, `timestamp_runtime`

Optional:

- `direction` (String) Direction of the diff from context_time. The following directions are supported: `ago`, `later`. Required in `timestamp` and `timestamp_runtime` types
- `format` (String) Format used to replace variables. Required in `timestamp` and `timestamp_runtime` types
- `quantity` (Number) Quantity used to calculate diff from context_time. Required in `timestamp` and `timestamp_runtime` types
- `time_zone` (String) Time zone used to format the timestamp. Required in `timestamp` and `timestamp_runtime` types
- `unit` (String) Time unit used to calculate diff from context_time. The following units are supported: `hour`, `date`, `month`. Required in `timestamp` and `timestamp_runtime` types
- `value` (String) Fixed string which will replace variables at runtime. Required in `string` type


<a id="nestedatt--input_option--http_input_option--excel_parser"></a>
### Nested Schema for `input_option.http_input_option.excel_parser`

Required:

- `columns` (Attributes List) List of columns to be retrieved and their types (see [below for nested schema](#nestedatt--input_option--http_input_option--excel_parser--columns))
- `sheet_name` (String) Sheet name

Optional:

- `default_time_zone` (String) Default time zone
- `skip_header_lines` (Number) Number of header lines to skip

<a id="nestedatt--input_option--http_input_option--excel_parser--columns"></a>
### Nested Schema for `input_option.http_input_option.excel_parser.columns`

Required:

- `formula_handling` (String) Formula handling
- `name` (String) Column name
- `type` (String) Column type

Optional:

- `format` (String) Format of the column.



<a id="nestedatt--input_option--http_input_option--jsonl_parser"></a>
### Nested Schema for `input_option.http_input_option.jsonl_parser`

Required:

- `columns` (Attributes List) List of columns to be retrieved and their types (see [below for nested schema](#nestedatt--input_option--http_input_option--jsonl_parser--columns))

Optional:

- `charset` (String) Character set
- `default_time_zone` (String) Default time zone
- `newline` (String) Newline character
- `stop_on_invalid_record` (Boolean) Flag whether the transfer should stop if an invalid record is found

<a id="nestedatt--input_option--http_input_option--jsonl_parser--columns"></a>
### Nested Schema for `input_option.http_input_option.jsonl_parser.columns`

Required:

- `name` (String) Column name
- `type` (String) Column type

Optional:

- `format` (String) Format of the column
- `time_zone` (String) time zone



<a id="nestedatt--input_option--http_input_option--jsonpath_parser"></a>
### Nested Schema for `input_option.http_input_option.jsonpath_parser`

Required:

- `columns` (Attributes List) (see [below for nested schema](#nestedatt--input_option--http_input_option--jsonpath_parser--columns))
- `root` (String) JSONPath

Optional:

- `default_time_zone` (String) Default time zone

<a id="nestedatt--input_option--http_input_option--jsonpath_parser--columns"></a>
### Nested Schema for `input_option.http_input_option.jsonpath_parser.columns`

Required:

- `name` (String) Column name
- `type` (String) Column type

Optional:

- `format` (String) Format of the column.
- `time_zone` (String) time zone



<a id="nestedatt--input_option--http_input_option--ltsv_parser"></a>
### Nested Schema for `input_option.http_input_option.ltsv_parser`

Required:

- `columns` (Attributes List) List of columns to be retrieved and their types (see [below for nested schema](#nestedatt--input_option--http_input_option--ltsv_parser--columns))

Optional:

- `charset` (String) Character set
- `newline` (String) Newline character

<a id="nestedatt--input_option--http_input_option--ltsv_parser--columns"></a>
### Nested Schema for `input_option.http_input_option.ltsv_parser.columns`

Required:

- `name` (String) Column name
- `type` (String) Column type

Optional:

- `format` (String) Format of the column.



<a id="nestedatt--input_option--http_input_option--parquet_parser"></a>
### Nested Schema for `input_option.http_input_option.parquet_parser`

Required:

- `columns` (Attributes List) List of columns to be retrieved and their types (see [below for nested schema](#nestedatt--input_option--http_input_option--parquet_parser--columns))

<a id="nestedatt--input_option--http_input_option--parquet_parser--columns"></a>
### Nested Schema for `input_option.http_input_option.parquet_parser.columns`

Required:

- `name` (String) Column name
- `type` (String) Column type

Optional:

- `format` (String) Format of the column.



<a id="nestedatt--input_option--http_input_option--request_headers"></a>
### Nested Schema for `input_option.http_input_option.request_headers`

Required:

- `key` (String) Header key
- `value` (String, Sensitive) Header value

Optional:

- `masking` (Boolean) Whether to mask this header


<a id="nestedatt--input_option--http_input_option--request_params"></a>
### Nested Schema for `input_option.http_input_option.request_params`

Required:

- `key` (String) Parameter key
- `value` (String, Sensitive) Parameter value

Optional:

- `masking` (Boolean) Whether to mask this parameter


<a id="nestedatt--input_option--http_input_option--xml_parser"></a>
### Nested Schema for `input_option.http_input_option.xml_parser`

Required:

- `columns` (Attributes List) (see [below for nested schema](#nestedatt--input_option--http_input_option--xml_parser--columns))
- `root` (String) Root element

<a id="nestedatt--input_option--http_input_option--xml_parser--columns"></a>
### Nested Schema for `input_option.http_input_option.xml_parser.columns`

Required:

- `name` (String) Column name
- `path` (String) XPath
- `type` (String) Column type

Optional:

- `format` (String) Format of the column.
- `timezone` (String) time zone




<a id="nestedatt--input_option--kintone_input_option"></a>
### Nested Schema for `input_option.kintone_input_option`

Required:

- `app_id` (String) app id
- `input_option_columns` (Attributes List) List of columns to be retrieved and their types (see [below for nested schema](#nestedatt--input_option--kintone_input_option--input_option_columns))
- `kintone_connection_id` (Number) ID of kintone connection

Optional:

- `custom_variable_settings` (Attributes List) (see [below for nested schema](#nestedatt--input_option--kintone_input_option--custom_variable_settings))
- `expand_subtable` (Boolean) If enabled and the target Kintone app contains a table, data will be transferred by table row instead of per record.
- `guest_space_id` (String) guest space id
- `query` (String) If you want to use all record loading, specify it.

<a id="nestedatt--input_option--kintone_input_option--input_option_columns"></a>
### Nested Schema for `input_option.kintone_input_option.input_option_columns`

Required:

- `name` (String) Column name
- `type` (String) Column type

Optional:

- `format` (String) Column format


<a id="nestedatt--input_option--kintone_input_option--custom_variable_settings"></a>
### Nested Schema for `input_option.kintone_input_option.custom_variable_settings`

Required:

- `name` (String) Custom variable name. It must start and end with `$`
- `type` (String) Custom variable type. The following types are supported: `string`, `timestamp`, `timestamp_runtime`

Optional:

- `direction` (String) Direction of the diff from context_time. The following directions are supported: `ago`, `later`. Required in `timestamp` and `timestamp_runtime` types
- `format` (String) Format used to replace variables. Required in `timestamp` and `timestamp_runtime` types
- `quantity` (Number) Quantity used to calculate diff from context_time. Required in `timestamp` and `timestamp_runtime` types
- `time_zone` (String) Time zone used to format the timestamp. Required in `timestamp` and `timestamp_runtime` types
- `unit` (String) Time unit used to calculate diff from context_time. The following units are supported: `hour`, `date`, `month`. Required in `timestamp` and `timestamp_runtime` types
- `value` (String) Fixed string which will replace variables at runtime. Required in `string` type



<a id="nestedatt--input_option--mysql_input_option"></a>
### Nested Schema for `input_option.mysql_input_option`

Required:

- `database` (String) database name
- `input_option_columns` (Attributes List) List of columns to be retrieved and their types (see [below for nested schema](#nestedatt--input_option--mysql_input_option--input_option_columns))
- `mysql_connection_id` (Number) ID of MySQL connection

Optional:

- `connect_timeout` (Number) Connection timeout (sec)
- `custom_variable_settings` (Attributes List) (see [below for nested schema](#nestedatt--input_option--mysql_input_option--custom_variable_settings))
- `default_time_zone` (String) Default time zone. enter the server-side time zone setting for MySQL. If the time zone is set to Japan, enter “Asia/Tokyo”.
- `fetch_rows` (Number) Number of records processed by the cursor at one time
- `incremental_columns` (String) Columns to determine incremental data
- `incremental_loading_enabled` (Boolean) If it is true, to be incremental loading. If it is false, to be all record loading
- `last_record` (String) Last record transferred. The value of the column specified here is stored in “Last Transferred Record” for each transfer, and for the second and subsequent transfers, only records for which the value of the “Column for Determining Incremental Data” is greater than the value of the previous transfer (= “Last Transferred Record”) are transferred. If you wish to specify multiple columns, specify them separated by commas. If not specified, the primary key is used.
- `query` (String) If you want to use all record loading, specify it.
- `socket_timeout` (Number) Socket timeout (seconds)
- `table` (String) table name. If you want to use incremental loading, specify it.
- `use_legacy_datetime_code` (Boolean) Legacy time code setting. setting the useLegacyDatetimeCode option in the JDBC driver

<a id="nestedatt--input_option--mysql_input_option--input_option_columns"></a>
### Nested Schema for `input_option.mysql_input_option.input_option_columns`

Required:

- `name` (String) Column name
- `type` (String) Column type


<a id="nestedatt--input_option--mysql_input_option--custom_variable_settings"></a>
### Nested Schema for `input_option.mysql_input_option.custom_variable_settings`

Required:

- `name` (String) Custom variable name. It must start and end with `$`
- `type` (String) Custom variable type. The following types are supported: `string`, `timestamp`, `timestamp_runtime`

Optional:

- `direction` (String) Direction of the diff from context_time. The following directions are supported: `ago`, `later`. Required in `timestamp` and `timestamp_runtime` types
- `format` (String) Format used to replace variables. Required in `timestamp` and `timestamp_runtime` types
- `quantity` (Number) Quantity used to calculate diff from context_time. Required in `timestamp` and `timestamp_runtime` types
- `time_zone` (String) Time zone used to format the timestamp. Required in `timestamp` and `timestamp_runtime` types
- `unit` (String) Time unit used to calculate diff from context_time. The following units are supported: `hour`, `date`, `month`. Required in `timestamp` and `timestamp_runtime` types
- `value` (String) Fixed string which will replace variables at runtime. Required in `string` type



<a id="nestedatt--input_option--postgresql_input_option"></a>
### Nested Schema for `input_option.postgresql_input_option`

Required:

- `database` (String) database name
- `input_option_columns` (Attributes List) List of columns to be retrieved and their types (see [below for nested schema](#nestedatt--input_option--postgresql_input_option--input_option_columns))
- `postgresql_connection_id` (Number) ID of Postgresql connection

Optional:

- `connect_timeout` (Number) Connection timeout (sec)
- `custom_variable_settings` (Attributes List) (see [below for nested schema](#nestedatt--input_option--postgresql_input_option--custom_variable_settings))
- `default_time_zone` (String) Default time zone. enter the server-side time zone setting for PostgreSQL. If the time zone is set to Japan, enter “Asia/Tokyo”.
- `fetch_rows` (Number) Number of records processed by the cursor at one time
- `incremental_columns` (String) Columns to determine incremental data
- `incremental_loading_enabled` (Boolean) If it is true, to be incremental loading. If it is false, to be all record loading
- `input_option_column_options` (Attributes List) List of unsupported data types and their convertible types (see [below for nested schema](#nestedatt--input_option--postgresql_input_option--input_option_column_options))
- `last_record` (String) Last record transferred. The value of the column specified here is stored in “Last Transferred Record” for each transfer, and for the second and subsequent transfers, only records for which the value of the “Column for Determining Incremental Data” is greater than the value of the previous transfer (= “Last Transferred Record”) are transferred. If you wish to specify multiple columns, specify them separated by commas. If not specified, the primary key is used.
- `query` (String) If you want to use all record loading, specify it.
- `schema` (String) schema name
- `socket_timeout` (Number) Socket timeout (seconds)
- `table` (String) table name. If you want to use incremental loading, specify it.

<a id="nestedatt--input_option--postgresql_input_option--input_option_columns"></a>
### Nested Schema for `input_option.postgresql_input_option.input_option_columns`

Required:

- `name` (String) Column name
- `type` (String) Column type


<a id="nestedatt--input_option--postgresql_input_option--custom_variable_settings"></a>
### Nested Schema for `input_option.postgresql_input_option.custom_variable_settings`

Required:

- `name` (String) Custom variable name. It must start and end with `$`
- `type` (String) Custom variable type. The following types are supported: `string`, `timestamp`, `timestamp_runtime`

Optional:

- `direction` (String) Direction of the diff from context_time. The following directions are supported: `ago`, `later`. Required in `timestamp` and `timestamp_runtime` types
- `format` (String) Format used to replace variables. Required in `timestamp` and `timestamp_runtime` types
- `quantity` (Number) Quantity used to calculate diff from context_time. Required in `timestamp` and `timestamp_runtime` types
- `time_zone` (String) Time zone used to format the timestamp. Required in `timestamp` and `timestamp_runtime` types
- `unit` (String) Time unit used to calculate diff from context_time. The following units are supported: `hour`, `date`, `month`. Required in `timestamp` and `timestamp_runtime` types
- `value` (String) Fixed string which will replace variables at runtime. Required in `string` type


<a id="nestedatt--input_option--postgresql_input_option--input_option_column_options"></a>
### Nested Schema for `input_option.postgresql_input_option.input_option_column_options`

Required:

- `column_name` (String) Column name
- `column_value_type` (String) Column value type



<a id="nestedatt--input_option--s3_input_option"></a>
### Nested Schema for `input_option.s3_input_option`

Required:

- `bucket` (String) Bucket name
- `s3_connection_id` (Number) Id of S3 connection

Optional:

- `csv_parser` (Attributes) For files in CSV format, this parameter is required (see [below for nested schema](#nestedatt--input_option--s3_input_option--csv_parser))
- `custom_variable_settings` (Attributes List) (see [below for nested schema](#nestedatt--input_option--s3_input_option--custom_variable_settings))
- `decoder` (Attributes) (see [below for nested schema](#nestedatt--input_option--s3_input_option--decoder))
- `decompression_type` (String) Decompression type
- `excel_parser` (Attributes) For files in excel format, this parameter is required. (see [below for nested schema](#nestedatt--input_option--s3_input_option--excel_parser))
- `incremental_loading_enabled` (Boolean) If it is true, to be incremental loading. If it is false, to be all record loading
- `is_skip_header_line` (Boolean) Flag whether or not to skip header columns For CSV/TSV files that do not contain header columns, a temporary header name generated on the TROCCO side is assigned.
- `jsonl_parser` (Attributes) For files in JSONL format, this parameter is required (see [below for nested schema](#nestedatt--input_option--s3_input_option--jsonl_parser))
- `jsonpath_parser` (Attributes) For files in jsonpath format, this parameter is required. (see [below for nested schema](#nestedatt--input_option--s3_input_option--jsonpath_parser))
- `ltsv_parser` (Attributes) For files in LTSV format, this parameter is required. (see [below for nested schema](#nestedatt--input_option--s3_input_option--ltsv_parser))
- `parquet_parser` (Attributes) For files in parquet format, this parameter is required. (see [below for nested schema](#nestedatt--input_option--s3_input_option--parquet_parser))
- `path_match_pattern` (String) Path regular expression. If not entered, all files matching the path prefix will be included.
- `path_prefix` (String) Path prefix. If not entered, all files under the bucket will be targeted.
- `region` (String) Region
- `stop_when_file_not_found` (Boolean) Flag whether the transfer should continue if the file does not exist in the specified path
- `xml_parser` (Attributes) For files in xml format, this parameter is required. (see [below for nested schema](#nestedatt--input_option--s3_input_option--xml_parser))

<a id="nestedatt--input_option--s3_input_option--csv_parser"></a>
### Nested Schema for `input_option.s3_input_option.csv_parser`

Required:

- `columns` (Attributes List) (see [below for nested schema](#nestedatt--input_option--s3_input_option--csv_parser--columns))

Optional:

- `allow_extra_columns` (Boolean) If true, ignore the column. If false, treat as invalid record.
- `allow_optional_columns` (Boolean) If true, NULL-complete the missing columns. If false, treat as invalid record.
- `charset` (String) Character set
- `comment_line_marker` (String) Comment line marker. Skip if this character is at the beginning of a line
- `default_date` (String) Default date
- `default_time_zone` (String) Default time zone
- `delimiter` (String) Delimiter
- `escape` (String) Escape character
- `max_quoted_size_limit` (Number) Maximum amount of data that can be enclosed in quotation marks.
- `newline` (String) Newline character
- `null_string` (String) Replacement source string to be converted to NULL
- `null_string_enabled` (Boolean) Flag whether or not to set the string to be replaced by NULL
- `quote` (String) Quote character
- `quotes_in_quoted_fields` (String) Processing method for irregular quarts
- `skip_header_lines` (Number) Number of header lines to skip
- `stop_on_invalid_record` (Boolean) Flag whether or not to abort the transfer if an invalid record is found.
- `trim_if_not_quoted` (Boolean) Flag whether or not to remove spaces from the value if it is not quoted

<a id="nestedatt--input_option--s3_input_option--csv_parser--columns"></a>
### Nested Schema for `input_option.s3_input_option.csv_parser.columns`

Required:

- `name` (String) Column name
- `type` (String) Column type

Optional:

- `date` (String) Date
- `format` (String) Format of the column



<a id="nestedatt--input_option--s3_input_option--custom_variable_settings"></a>
### Nested Schema for `input_option.s3_input_option.custom_variable_settings`

Required:

- `name` (String) Custom variable name. It must start and end with `$`
- `type` (String) Custom variable type. The following types are supported: `string`, `timestamp`, `timestamp_runtime`

Optional:

- `direction` (String) Direction of the diff from context_time. The following directions are supported: `ago`, `later`. Required in `timestamp` and `timestamp_runtime` types
- `format` (String) Format used to replace variables. Required in `timestamp` and `timestamp_runtime` types
- `quantity` (Number) Quantity used to calculate diff from context_time. Required in `timestamp` and `timestamp_runtime` types
- `time_zone` (String) Time zone used to format the timestamp. Required in `timestamp` and `timestamp_runtime` types
- `unit` (String) Time unit used to calculate diff from context_time. The following units are supported: `hour`, `date`, `month`. Required in `timestamp` and `timestamp_runtime` types
- `value` (String) Fixed string which will replace variables at runtime. Required in `string` type


<a id="nestedatt--input_option--s3_input_option--decoder"></a>
### Nested Schema for `input_option.s3_input_option.decoder`

Optional:

- `match_name` (String) Relative path after decompression (regular expression). If not entered, all data in the compressed file will be transferred.


<a id="nestedatt--input_option--s3_input_option--excel_parser"></a>
### Nested Schema for `input_option.s3_input_option.excel_parser`

Required:

- `columns` (Attributes List) List of columns to be retrieved and their types (see [below for nested schema](#nestedatt--input_option--s3_input_option--excel_parser--columns))
- `sheet_name` (String) Sheet name

Optional:

- `default_time_zone` (String) Default time zone
- `skip_header_lines` (Number) Number of header lines to skip

<a id="nestedatt--input_option--s3_input_option--excel_parser--columns"></a>
### Nested Schema for `input_option.s3_input_option.excel_parser.columns`

Required:

- `formula_handling` (String) Formula handling
- `name` (String) Column name
- `type` (String) Column type

Optional:

- `format` (String) Format of the column.



<a id="nestedatt--input_option--s3_input_option--jsonl_parser"></a>
### Nested Schema for `input_option.s3_input_option.jsonl_parser`

Required:

- `columns` (Attributes List) List of columns to be retrieved and their types (see [below for nested schema](#nestedatt--input_option--s3_input_option--jsonl_parser--columns))

Optional:

- `charset` (String) Character set
- `default_time_zone` (String) Default time zone
- `newline` (String) Newline character
- `stop_on_invalid_record` (Boolean) Flag whether the transfer should stop if an invalid record is found

<a id="nestedatt--input_option--s3_input_option--jsonl_parser--columns"></a>
### Nested Schema for `input_option.s3_input_option.jsonl_parser.columns`

Required:

- `name` (String) Column name
- `type` (String) Column type

Optional:

- `format` (String) Format of the column
- `time_zone` (String) time zone



<a id="nestedatt--input_option--s3_input_option--jsonpath_parser"></a>
### Nested Schema for `input_option.s3_input_option.jsonpath_parser`

Required:

- `columns` (Attributes List) (see [below for nested schema](#nestedatt--input_option--s3_input_option--jsonpath_parser--columns))
- `root` (String) JSONPath

Optional:

- `default_time_zone` (String) Default time zone

<a id="nestedatt--input_option--s3_input_option--jsonpath_parser--columns"></a>
### Nested Schema for `input_option.s3_input_option.jsonpath_parser.columns`

Required:

- `name` (String) Column name
- `type` (String) Column type

Optional:

- `format` (String) Format of the column.
- `time_zone` (String) time zone



<a id="nestedatt--input_option--s3_input_option--ltsv_parser"></a>
### Nested Schema for `input_option.s3_input_option.ltsv_parser`

Required:

- `columns` (Attributes List) List of columns to be retrieved and their types (see [below for nested schema](#nestedatt--input_option--s3_input_option--ltsv_parser--columns))

Optional:

- `charset` (String) Character set
- `newline` (String) Newline character

<a id="nestedatt--input_option--s3_input_option--ltsv_parser--columns"></a>
### Nested Schema for `input_option.s3_input_option.ltsv_parser.columns`

Required:

- `name` (String) Column name
- `type` (String) Column type

Optional:

- `format` (String) Format of the column.



<a id="nestedatt--input_option--s3_input_option--parquet_parser"></a>
### Nested Schema for `input_option.s3_input_option.parquet_parser`

Required:

- `columns` (Attributes List) List of columns to be retrieved and their types (see [below for nested schema](#nestedatt--input_option--s3_input_option--parquet_parser--columns))

<a id="nestedatt--input_option--s3_input_option--parquet_parser--columns"></a>
### Nested Schema for `input_option.s3_input_option.parquet_parser.columns`

Required:

- `name` (String) Column name
- `type` (String) Column type

Optional:

- `format` (String) Format of the column.



<a id="nestedatt--input_option--s3_input_option--xml_parser"></a>
### Nested Schema for `input_option.s3_input_option.xml_parser`

Required:

- `columns` (Attributes List) (see [below for nested schema](#nestedatt--input_option--s3_input_option--xml_parser--columns))
- `root` (String) Root element

<a id="nestedatt--input_option--s3_input_option--xml_parser--columns"></a>
### Nested Schema for `input_option.s3_input_option.xml_parser.columns`

Required:

- `name` (String) Column name
- `path` (String) XPath
- `type` (String) Column type

Optional:

- `format` (String) Format of the column.
- `timezone` (String) time zone




<a id="nestedatt--input_option--salesforce_input_option"></a>
### Nested Schema for `input_option.salesforce_input_option`

Required:

- `columns` (Attributes List) List of columns to be retrieved and their types (see [below for nested schema](#nestedatt--input_option--salesforce_input_option--columns))
- `object` (String) Object name
- `salesforce_connection_id` (Number) Id of Salesforce connection

Optional:

- `api_version` (String) Api version
- `custom_variable_settings` (Attributes List) (see [below for nested schema](#nestedatt--input_option--salesforce_input_option--custom_variable_settings))
- `include_deleted_or_archived_records` (Boolean) Extraction of deleted and archived records
- `is_convert_type_custom_columns` (Boolean) false: transfer based on custom columns. true: Change custom columns to STRING type. If you select change custom columns to STRING type, all custom columns except BOOLEAN type will be changed to STRING type.
- `object_acquisition_method` (String) Object Acquisition Method. If 'all_columns' is specified, soql is automatically completed.
- `soql` (String) SOQL. If object_acquisition_method is 'soql', this field is required.

<a id="nestedatt--input_option--salesforce_input_option--columns"></a>
### Nested Schema for `input_option.salesforce_input_option.columns`

Required:

- `name` (String) Column name
- `type` (String) Column type.

Optional:

- `format` (String) format


<a id="nestedatt--input_option--salesforce_input_option--custom_variable_settings"></a>
### Nested Schema for `input_option.salesforce_input_option.custom_variable_settings`

Required:

- `name` (String) Custom variable name. It must start and end with `$`
- `type` (String) Custom variable type. The following types are supported: `string`, `timestamp`, `timestamp_runtime`

Optional:

- `direction` (String) Direction of the diff from context_time. The following directions are supported: `ago`, `later`. Required in `timestamp` and `timestamp_runtime` types
- `format` (String) Format used to replace variables. Required in `timestamp` and `timestamp_runtime` types
- `quantity` (Number) Quantity used to calculate diff from context_time. Required in `timestamp` and `timestamp_runtime` types
- `time_zone` (String) Time zone used to format the timestamp. Required in `timestamp` and `timestamp_runtime` types
- `unit` (String) Time unit used to calculate diff from context_time. The following units are supported: `hour`, `date`, `month`. Required in `timestamp` and `timestamp_runtime` types
- `value` (String) Fixed string which will replace variables at runtime. Required in `string` type



<a id="nestedatt--input_option--snowflake_input_option"></a>
### Nested Schema for `input_option.snowflake_input_option`

Required:

- `database` (String) Database name
- `input_option_columns` (Attributes List) List of columns to be retrieved and their types (see [below for nested schema](#nestedatt--input_option--snowflake_input_option--input_option_columns))
- `query` (String) Query
- `snowflake_connection_id` (Number) Id of Snowflake connection
- `warehouse` (String) Warehouse name

Optional:

- `connect_timeout` (Number) Connection timeout (sec)
- `custom_variable_settings` (Attributes List) (see [below for nested schema](#nestedatt--input_option--snowflake_input_option--custom_variable_settings))
- `fetch_rows` (Number) Number of records processed by the cursor at one time
- `schema` (String) Schema name
- `socket_timeout` (Number) Socket timeout (seconds)

<a id="nestedatt--input_option--snowflake_input_option--input_option_columns"></a>
### Nested Schema for `input_option.snowflake_input_option.input_option_columns`

Required:

- `name` (String) Column name
- `type` (String) Column type


<a id="nestedatt--input_option--snowflake_input_option--custom_variable_settings"></a>
### Nested Schema for `input_option.snowflake_input_option.custom_variable_settings`

Required:

- `name` (String) Custom variable name. It must start and end with `$`
- `type` (String) Custom variable type. The following types are supported: `string`, `timestamp`, `timestamp_runtime`

Optional:

- `direction` (String) Direction of the diff from context_time. The following directions are supported: `ago`, `later`. Required in `timestamp` and `timestamp_runtime` types
- `format` (String) Format used to replace variables. Required in `timestamp` and `timestamp_runtime` types
- `quantity` (Number) Quantity used to calculate diff from context_time. Required in `timestamp` and `timestamp_runtime` types
- `time_zone` (String) Time zone used to format the timestamp. Required in `timestamp` and `timestamp_runtime` types
- `unit` (String) Time unit used to calculate diff from context_time. The following units are supported: `hour`, `date`, `month`. Required in `timestamp` and `timestamp_runtime` types
- `value` (String) Fixed string which will replace variables at runtime. Required in `string` type



<a id="nestedatt--input_option--yahoo_ads_api_yss_input_option"></a>
### Nested Schema for `input_option.yahoo_ads_api_yss_input_option`

Required:

- `account_id` (String) account id
- `base_account_id` (String) base account id
- `input_option_columns` (Attributes List) List of columns to be retrieved and their types (see [below for nested schema](#nestedatt--input_option--yahoo_ads_api_yss_input_option--input_option_columns))
- `yahoo_ads_api_connection_id` (Number) ID of yahoo_ads_api connection

Optional:

- `custom_variable_settings` (Attributes List) (see [below for nested schema](#nestedatt--input_option--yahoo_ads_api_yss_input_option--custom_variable_settings))
- `end_date` (String) end_date
- `exclude_zero_impressions` (Boolean) exclude_zero_impressions
- `report_type` (String) report_type
- `service` (String) service
- `start_date` (String) start_date

<a id="nestedatt--input_option--yahoo_ads_api_yss_input_option--input_option_columns"></a>
### Nested Schema for `input_option.yahoo_ads_api_yss_input_option.input_option_columns`

Required:

- `name` (String) Column name
- `type` (String) Column type

Optional:

- `format` (String) Column format


<a id="nestedatt--input_option--yahoo_ads_api_yss_input_option--custom_variable_settings"></a>
### Nested Schema for `input_option.yahoo_ads_api_yss_input_option.custom_variable_settings`

Required:

- `name` (String) Custom variable name. It must start and end with `$`
- `type` (String) Custom variable type. The following types are supported: `string`, `timestamp`, `timestamp_runtime`

Optional:

- `direction` (String) Direction of the diff from context_time. The following directions are supported: `ago`, `later`. Required in `timestamp` and `timestamp_runtime` types
- `format` (String) Format used to replace variables. Required in `timestamp` and `timestamp_runtime` types
- `quantity` (Number) Quantity used to calculate diff from context_time. Required in `timestamp` and `timestamp_runtime` types
- `time_zone` (String) Time zone used to format the timestamp. Required in `timestamp` and `timestamp_runtime` types
- `unit` (String) Time unit used to calculate diff from context_time. The following units are supported: `hour`, `date`, `month`. Required in `timestamp` and `timestamp_runtime` types
- `value` (String) Fixed string which will replace variables at runtime. Required in `string` type




<a id="nestedatt--output_option"></a>
### Nested Schema for `output_option`

Optional:

- `bigquery_output_option` (Attributes) Attributes of destination BigQuery settings (see [below for nested schema](#nestedatt--output_option--bigquery_output_option))
- `google_spreadsheets_output_option` (Attributes) Attributes of destination Snowflake settings (see [below for nested schema](#nestedatt--output_option--google_spreadsheets_output_option))
- `salesforce_output_option` (Attributes) Attributes of destination Salesforce settings (see [below for nested schema](#nestedatt--output_option--salesforce_output_option))
- `snowflake_output_option` (Attributes) Attributes of destination Snowflake settings (see [below for nested schema](#nestedatt--output_option--snowflake_output_option))

<a id="nestedatt--output_option--bigquery_output_option"></a>
### Nested Schema for `output_option.bigquery_output_option`

Required:

- `bigquery_connection_id` (Number) Id of BigQuery connection
- `dataset` (String) Dataset name
- `table` (String) Table name

Optional:

- `auto_create_dataset` (Boolean) Option for automatic data set generation
- `bigquery_output_option_clustering_fields` (Set of String) Clustered column. Clustering can only be set when creating a new table. A maximum of four clustered columns can be specified.
- `bigquery_output_option_column_options` (Attributes List) (see [below for nested schema](#nestedatt--output_option--bigquery_output_option--bigquery_output_option_column_options))
- `bigquery_output_option_merge_keys` (Set of String) Merge key. The column to be used as the merge key. Required when mode is 'merge'.
- `custom_variable_settings` (Attributes List) (see [below for nested schema](#nestedatt--output_option--bigquery_output_option--custom_variable_settings))
- `location` (String) Location
- `mode` (String) Transfer mode
- `open_timeout_sec` (Number) Timeout to start connection (seconds)
- `partitioning_type` (String) Partitioning type. If params is null, No partitions. ingestion_time: Partitioning by acquisition time. time_unit_column: Partitioning by time unit column
- `read_timeout_sec` (Number) Read timeout (seconds)
- `retries` (Number) Number of retries
- `send_timeout_sec` (Number) Transmission timeout (sec)
- `template_table` (String) Template table. Generate schema information for inclusion in Google BigQuery from schema information in this table
- `time_partitioning_expiration_ms` (Number) Duration of partition(milliseconds). Duration of the partition (in milliseconds). There is no minimum value. The date of the partition plus this integer value is the expiration date. The default value is unspecified (keep forever).
- `time_partitioning_field` (String) If partitioning_type is time_unit_column, this parameter is required
- `time_partitioning_type` (String) Time partitioning type. If you specify anything for partitioning_type, this parameter is required
- `timeout_sec` (Number) Time out (seconds)

<a id="nestedatt--output_option--bigquery_output_option--bigquery_output_option_column_options"></a>
### Nested Schema for `output_option.bigquery_output_option.bigquery_output_option_column_options`

Required:

- `mode` (String) Mode
- `name` (String) Column name
- `type` (String) Column type

Optional:

- `description` (String) Description
- `timestamp_format` (String) Timestamp format
- `timezone` (String) Time zone


<a id="nestedatt--output_option--bigquery_output_option--custom_variable_settings"></a>
### Nested Schema for `output_option.bigquery_output_option.custom_variable_settings`

Required:

- `name` (String) Custom variable name. It must start and end with `$`
- `type` (String) Custom variable type. The following types are supported: `string`, `timestamp`, `timestamp_runtime`

Optional:

- `direction` (String) Direction of the diff from context_time. The following directions are supported: `ago`, `later`. Required in `timestamp` and `timestamp_runtime` types
- `format` (String) Format used to replace variables. Required in `timestamp` and `timestamp_runtime` types
- `quantity` (Number) Quantity used to calculate diff from context_time. Required in `timestamp` and `timestamp_runtime` types
- `time_zone` (String) Time zone used to format the timestamp. Required in `timestamp` and `timestamp_runtime` types
- `unit` (String) Time unit used to calculate diff from context_time. The following units are supported: `hour`, `date`, `month`. Required in `timestamp` and `timestamp_runtime` types
- `value` (String) Fixed string which will replace variables at runtime. Required in `string` type



<a id="nestedatt--output_option--google_spreadsheets_output_option"></a>
### Nested Schema for `output_option.google_spreadsheets_output_option`

Required:

- `google_spreadsheets_connection_id` (Number) Snowflake connection ID
- `spreadsheets_id` (String) Spreadsheet ID
- `worksheet_title` (String) Worksheet title

Optional:

- `custom_variable_settings` (Attributes List) (see [below for nested schema](#nestedatt--output_option--google_spreadsheets_output_option--custom_variable_settings))
- `google_spreadsheets_output_option_sorts` (Attributes List) (see [below for nested schema](#nestedatt--output_option--google_spreadsheets_output_option--google_spreadsheets_output_option_sorts))
- `mode` (String) Transfer mode
- `timezone` (String) Time zone
- `value_input_option` (String) Value input option

<a id="nestedatt--output_option--google_spreadsheets_output_option--custom_variable_settings"></a>
### Nested Schema for `output_option.google_spreadsheets_output_option.custom_variable_settings`

Required:

- `name` (String) Custom variable name. It must start and end with `$`
- `type` (String) Custom variable type. The following types are supported: `string`, `timestamp`, `timestamp_runtime`

Optional:

- `direction` (String) Direction of the diff from context_time. The following directions are supported: `ago`, `later`. Required in `timestamp` and `timestamp_runtime` types
- `format` (String) Format used to replace variables. Required in `timestamp` and `timestamp_runtime` types
- `quantity` (Number) Quantity used to calculate diff from context_time. Required in `timestamp` and `timestamp_runtime` types
- `time_zone` (String) Time zone used to format the timestamp. Required in `timestamp` and `timestamp_runtime` types
- `unit` (String) Time unit used to calculate diff from context_time. The following units are supported: `hour`, `date`, `month`. Required in `timestamp` and `timestamp_runtime` types
- `value` (String) Fixed string which will replace variables at runtime. Required in `string` type


<a id="nestedatt--output_option--google_spreadsheets_output_option--google_spreadsheets_output_option_sorts"></a>
### Nested Schema for `output_option.google_spreadsheets_output_option.google_spreadsheets_output_option_sorts`

Required:

- `column` (String) Column name
- `order` (String) Data type



<a id="nestedatt--output_option--salesforce_output_option"></a>
### Nested Schema for `output_option.salesforce_output_option`

Required:

- `object` (String) Object name
- `salesforce_connection_id` (Number) Salesforce connection ID. Only connection information with authentication method user_password can be selected.

Optional:

- `action_type` (String) Transfer mode
- `api_version` (String) Api version
- `ignore_nulls` (Boolean) Update processing when NULL is included. Even if true, the record update process itself is performed.
- `throw_if_failed` (Boolean) Status of records that could not be sent
- `upsert_key` (String) Upsert key. If action_type is 'upsert', this field can be set.


<a id="nestedatt--output_option--snowflake_output_option"></a>
### Nested Schema for `output_option.snowflake_output_option`

Required:

- `database` (String) Database name
- `schema` (String) Schema name
- `snowflake_connection_id` (Number) Snowflake connection ID
- `table` (String) Table name
- `warehouse` (String) Warehouse name

Optional:

- `batch_size` (Number) Batch size (MB)
- `custom_variable_settings` (Attributes List) (see [below for nested schema](#nestedatt--output_option--snowflake_output_option--custom_variable_settings))
- `default_time_zone` (String) Default time zone
- `delete_stage_on_error` (Boolean) Delete temporary stage on error
- `empty_field_as_null` (Boolean) Replace empty string with NULL
- `max_retry_wait` (Number) Maximum retry wait time (milliseconds)
- `mode` (String) Transfer mode
- `retry_limit` (Number) Maximum retry limit
- `retry_wait` (Number) Retry wait time (milliseconds)
- `snowflake_output_option_column_options` (Attributes List) (see [below for nested schema](#nestedatt--output_option--snowflake_output_option--snowflake_output_option_column_options))
- `snowflake_output_option_merge_keys` (Set of String) Merge keys (only applicable if mode is 'merge')

<a id="nestedatt--output_option--snowflake_output_option--custom_variable_settings"></a>
### Nested Schema for `output_option.snowflake_output_option.custom_variable_settings`

Required:

- `name` (String) Custom variable name. It must start and end with `$`
- `type` (String) Custom variable type. The following types are supported: `string`, `timestamp`, `timestamp_runtime`

Optional:

- `direction` (String) Direction of the diff from context_time. The following directions are supported: `ago`, `later`. Required in `timestamp` and `timestamp_runtime` types
- `format` (String) Format used to replace variables. Required in `timestamp` and `timestamp_runtime` types
- `quantity` (Number) Quantity used to calculate diff from context_time. Required in `timestamp` and `timestamp_runtime` types
- `time_zone` (String) Time zone used to format the timestamp. Required in `timestamp` and `timestamp_runtime` types
- `unit` (String) Time unit used to calculate diff from context_time. The following units are supported: `hour`, `date`, `month`. Required in `timestamp` and `timestamp_runtime` types
- `value` (String) Fixed string which will replace variables at runtime. Required in `string` type


<a id="nestedatt--output_option--snowflake_output_option--snowflake_output_option_column_options"></a>
### Nested Schema for `output_option.snowflake_output_option.snowflake_output_option_column_options`

Required:

- `name` (String) Column name
- `type` (String) Data type

Optional:

- `timestamp_format` (String) Timestamp format
- `timezone` (String) Time zone
- `value_type` (String) Value type




<a id="nestedatt--filter_add_time"></a>
### Nested Schema for `filter_add_time`

Required:

- `column_name` (String) Column name
- `type` (String) Column type

Optional:

- `time_zone` (String) Time zone
- `timestamp_format` (String) Timestamp format


<a id="nestedatt--filter_gsub"></a>
### Nested Schema for `filter_gsub`

Required:

- `column_name` (String) Target column name
- `pattern` (String) Regular expression pattern
- `to` (String) String to be replaced


<a id="nestedatt--filter_hashes"></a>
### Nested Schema for `filter_hashes`

Required:

- `name` (String) Target column name. Replaces the string in the set column with a hashed version using SHA-256.


<a id="nestedatt--filter_masks"></a>
### Nested Schema for `filter_masks`

Required:

- `mask_type` (String) Masking type
- `name` (String) Target column name

Optional:

- `end_index` (Number) Mask end position
- `length` (Number) Number of mask symbols
- `pattern` (String) regular expression pattern
- `start_index` (Number) Mask start position


<a id="nestedatt--filter_rows"></a>
### Nested Schema for `filter_rows`

Required:

- `condition` (String) Conditions for applying multiple filtering
- `filter_row_conditions` (Attributes List) (see [below for nested schema](#nestedatt--filter_rows--filter_row_conditions))

<a id="nestedatt--filter_rows--filter_row_conditions"></a>
### Nested Schema for `filter_rows.filter_row_conditions`

Required:

- `argument` (String) Argument
- `column` (String) Target column name
- `operator` (String) Operator



<a id="nestedatt--filter_string_transforms"></a>
### Nested Schema for `filter_string_transforms`

Required:

- `column_name` (String) Column name

Optional:

- `type` (String) Transformation type


<a id="nestedatt--filter_unixtime_conversions"></a>
### Nested Schema for `filter_unixtime_conversions`

Required:

- `column_name` (String) Target column name
- `datetime_format` (String) Date and tim format after conversion
- `datetime_timezone` (String) Time zon after conversion
- `kind` (String) Conversion Type
- `unixtime_unit` (String) UNIX time units before conversion


<a id="nestedatt--labels"></a>
### Nested Schema for `labels`

Required:

- `name` (String) The name of the label

Read-Only:

- `id` (Number) The ID of the label


<a id="nestedatt--notifications"></a>
### Nested Schema for `notifications`

Required:

- `destination_type` (String) Destination service where the notification will be sent. The following types are supported: `slack`, `email`
- `message` (String) The message to be sent with the notification
- `notification_type` (String) Category of condition. The following types are supported: `job`, `record`, `exec_time`

Optional:

- `email_id` (Number) ID of the email used to send notifications. Required when `destination_type` is `email`
- `minutes` (Number)
- `notify_when` (String) Specifies the job status that trigger a notification. The following types are supported: `finished`, `failed`. Required when `notification_type` is `job`
- `record_count` (Number) The number of records to be used for condition. Required when `notification_type` is `record`
- `record_operator` (String) Operator to be used for condition. The following operators are supported: `above`, `below`. Required when `notification_type` is `record`
- `record_type` (String) Condition for number of records to be notified
- `slack_channel_id` (Number) ID of the slack channel used to send notifications. Required when `destination_type` is `slack`


<a id="nestedatt--schedules"></a>
### Nested Schema for `schedules`

Required:

- `frequency` (String) Frequency of automatic execution. The following frequencies are supported: `hourly`, `daily`, `weekly`, `monthly`
- `minute` (Number) Value of minute. Required for all schedules
- `time_zone` (String) Time zone to be used for calculation

Optional:

- `day` (Number) Value of day. Required in `monthly` schedule
- `day_of_week` (Number) Value of day of week. Sunday - Saturday is represented as 0 - 6. Required in `weekly` schedule
- `hour` (Number) Value of hour. Required in `daily`, `weekly`, and `monthly` schedules




## Import

Import is supported using the following syntax:

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import trocco_job_definition (Resource). For example:

```terraform
import {
  id = 1
  to = trocco_job_definition.example
}
```

Using the [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import):

```shell
terraform import trocco_job_definition.example <job_definition_id>
```

