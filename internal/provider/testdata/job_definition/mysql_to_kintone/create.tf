resource "trocco_connection" "test_mysql" {
  connection_type = "mysql"
  name            = "MySQL Example"
  host            = "db.example.com"
  port            = 65535
  user_name       = "root"
  password        = "password"
}

resource "trocco_connection" "test_kintone" {
  connection_type = "kintone"
  name            = "Kintone Example"
  domain          = "example.cybozu.com"
  login_method    = "username_and_password"
  username        = "test_user"
  password        = "test_password"
}

resource "trocco_job_definition" "mysql_to_kintone" {
  name                     = "MySQL to Kintone Test"
  description              = "Test job definition for transferring data from MySQL to Kintone"
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

  output_option_type = "kintone"
  output_option = {
    kintone_output_option = {
      kintone_connection_id = trocco_connection.test_kintone.id
      app_id                = "123"
      guest_space_id        = "1"
      mode                  = "upsert"
      update_key            = "id"
      ignore_nulls          = true
      reduce_key            = "email"
      chunk_size            = 150
      kintone_output_option_column_options = [
        {
          name       = "id"
          field_code = "record_id"
          type       = "NUMBER"
        },
        {
          name       = "created_date"
          field_code = "created_at"
          type       = "DATE"
          timezone   = "Asia/Tokyo"
        },
        {
          name       = "updated_time"
          field_code = "updated_at"
          type       = "TIME"
          timezone   = "Asia/Tokyo"
        },
        {
          name        = "sub_items"
          field_code  = "items_table"
          type        = "SUBTABLE"
          sort_column = "item_order"
        }
      ]
    }
  }
}
