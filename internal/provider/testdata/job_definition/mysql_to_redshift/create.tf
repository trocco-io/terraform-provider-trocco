resource "trocco_connection" "test_mysql" {
  connection_type = "mysql"
  name            = "MySQL Example"
  host            = "db.example.com"
  port            = 3306
  user_name       = "root"
  password        = "password"
}

resource "trocco_job_definition" "mysql_to_redshift" {
  name                     = "MySQL to Redshift Test"
  description              = "Test job definition for transferring data from MySQL to Redshift"
  resource_enhancement     = "medium"
  retry_limit              = 2
  is_runnable_concurrently = false

  filter_columns = [
    {
      default                      = ""
      format                       = ""
      json_expand_enabled          = false
      json_expand_keep_base_column = false
      name                         = "id"
      src                          = "id"
      type                         = "long"
    },
    {
      default                      = ""
      format                       = ""
      json_expand_enabled          = false
      json_expand_keep_base_column = false
      name                         = "type"
      src                          = "type"
      type                         = "string"
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
      incremental_loading_enabled = false
      default_time_zone           = "Asia/Tokyo"
      use_legacy_datetime_code    = false
      input_option_columns = [
        {
          name = "id"
          type = "long"
        },
        {
          name = "type"
          type = "string"
        },
      ]
      query = <<-EOT
        SELECT *
        FROM test_table
        WHERE id < 100
      EOT
    }
  }

  output_option_type = "redshift"
  output_option = {
    redshift_output_option = {
      redshift_connection_id  = 301
      database                = "analytics"
      schema                  = "$schema$"
      table                   = "users"
      mode                    = "merge"
      s3_bucket               = "my-redshift-bucket"
      s3_key_prefix           = "/redshift-temp"
      delete_s3_temp_file     = true
      retry_limit             = 12
      retry_wait              = 1000
      max_retry_wait          = 1800000
      default_time_zone       = "UTC"
      batch_size              = 1024
      copy_iam_role_name      = "ADMIN"
      create_table_constraint = "PRIMARY KEY (user_id);"
      create_table_option     = "DISTSTYLE KEY DISTKEY (user_id) SORTKEY (created_timestamp);"
      before_load             = "DELETE FROM public.users WHERE created_timestamp < DATEADD(day, -90, GETDATE());"
      after_load              = "ANALYZE public.users;"
      custom_variable_settings = [
        {
          name  = "$schema$"
          type  = "string"
          value = "public"
        },
      ]
      redshift_output_option_merge_keys = [
        "user_id",
        "user_name",
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
          name       = "created_timestamp"
          type       = "TIMESTAMP"
          value_type = "timestamp"
          timezone   = "Asia/Tokyo"
        },
      ]
    }
  }
}
