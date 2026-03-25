resource "trocco_job_definition" "redshift_output_example" {
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
        "user_id",
      ]

      redshift_output_option_column_options = [
        {
          name       = "user_id"
          type       = "BIGINT"
          value_type = "long"
        },
        {
          name       = "user_name"
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
