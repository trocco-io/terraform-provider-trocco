resource "trocco_job_definition" "mysql_to_redshift_example" {
  name                     = "mysql_to_redshift_example"
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
      mysql_connection_id      = 1 # please set your mysql connection id
      query                    = "select * from example_table;"
      socket_timeout           = 1800
      use_legacy_datetime_code = false
    }
  }

  output_option_type = "redshift"
  output_option = {
    redshift_output_option = {
      redshift_connection_id = 1 # please set your redshift connection id
      database               = "$database$"
      schema                 = "$schema$"
      table                  = "users"
      mode                   = "merge"
      s3_bucket              = "example-redshift-temp-bucket"
      s3_key_prefix          = "trocco/tmp/redshift_output"
      delete_s3_temp_file    = true
      copy_iam_role_name     = "REDSHIFT_COPY_ROLE"
      retry_limit            = 12
      retry_wait             = 1000
      max_retry_wait         = 1800000
      default_time_zone      = "UTC"
      batch_size             = 16384

      create_table_constraint = "PRIMARY KEY (user_id);"
      create_table_option     = "DISTSTYLE KEY DISTKEY (user_id) SORTKEY (created_at);"
      before_load             = "DELETE FROM $schema$.users WHERE created_at < DATEADD(day, -90, GETDATE());"
      after_load              = "ANALYZE $schema$.users;"

      redshift_output_option_merge_keys = [
        "id",
      ]

      redshift_output_option_column_options = [
        {
          name       = "id"
          type       = "BIGINT"
          value_type = "long"
        },
        {
          name       = "name"
          type       = "VARCHAR"
          value_type = "string"
        },
        {
          name       = "created_at"
          type       = "TIMESTAMP"
          value_type = "timestamp"
          timezone   = "UTC"
        },
      ]

      custom_variable_settings = [
        {
          name  = "$database$"
          type  = "string"
          value = "analytics"
        },
        {
          name  = "$schema$"
          type  = "string"
          value = "public"
        },
      ]
    }
  }
}
