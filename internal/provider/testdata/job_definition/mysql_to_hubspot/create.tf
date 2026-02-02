resource "trocco_connection" "test_mysql" {
  connection_type = "mysql"
  name            = "MySQL Example"
  host            = "db.example.com"
  port            = 3306
  user_name       = "root"
  password        = "password"
}

resource "trocco_job_definition" "mysql_to_hubspot" {
  name                     = "MySQL to HubSpot Test"
  description              = "Test job definition for transferring data from MySQL to HubSpot"
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

  output_option_type = "hubspot"
  output_option = {
    hubspot_output_option = {
      hubspot_connection_id = 1
      object_type           = "task"
      mode                  = "merge"
      upsert_key            = "id"
      number_of_parallels   = 2
      associations = [
        {
          to_object_type  = "contact"
          from_object_key = "contact_email"
          to_object_key   = "email"
        },
        {
          to_object_type  = "deal"
          from_object_key = "deal_task"
          to_object_key   = "task"
        }
      ]
    }
  }
}
