resource "trocco_job_definition" "mysql_to_databricks" {
  name                     = "MySQL to Databricks Test"
  description              = "Test job definition for transferring data from MySQL to Databricks"
  resource_enhancement     = "medium"
  retry_limit              = 2
  is_runnable_concurrently = true

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
      mysql_connection_id         = 1
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

  output_option_type = "databricks"
  output_option = {
    databricks_output_option = {
      databricks_connection_id = 1
      catalog_name             = "test_catalog"
      schema_name              = "test_schema"
      table                    = "test_table"
      batch_size               = 40000
    }
  }
}