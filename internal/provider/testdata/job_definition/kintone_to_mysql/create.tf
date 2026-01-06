# resource "trocco_connection" "kintone" {
#   connection_type     = "kintone"
#   name                = "Kintone Example"
#   description         = "This is a Kintone connection example"
#   domain              = "test_domain"
#   login_method        = "username_and_password"
#   username            = "username"
#   password            = "password"
#   basic_auth_username = "basic_auth_username"
#   basic_auth_password = "basic_auth_password"
# }

# resource "trocco_connection" "mysql" {
#   connection_type = "mysql"
#   name            = "MySQL Example"
#   host            = "db.example.com"
#   port            = 3306
#   user_name       = "root"
#   password        = "password"
# }

resource "trocco_job_definition" "kintone_to_mysql" {
  name                     = "Kintone to Mysql Test"
  description              = "Test job definition for Kintone to Mysql transfer"
  is_runnable_concurrently = false
  retry_limit              = 0

  filter_columns = [
    {
      default                      = null
      json_expand_enabled          = false
      json_expand_keep_base_column = false
      name                         = "record_id"
      src                          = "record_id"
      type                         = "long"
    },
    {
      default                      = null
      json_expand_enabled          = false
      json_expand_keep_base_column = false
      name                         = "name"
      src                          = "name"
      type                         = "string"
    },
    {
      default                      = null
      json_expand_enabled          = false
      json_expand_keep_base_column = false
      name                         = "email"
      src                          = "email"
      type                         = "string"
    }
  ]

  input_option_type = "kintone"
  input_option = {
    kintone_input_option = {
      kintone_connection_id = 2
      app_id                = "403"
      guest_space_id        = null
      query                 = null
      expand_subtable       = false
      custom_variable_settings = [
        {
          name  = "$app_id$"
          type  = "string"
          value = "403"
        }
      ]
      input_option_columns = [
        {
          name = "record_id"
          type = "long"
        },
        {
          name = "name"
          type = "string"
        },
        {
          name = "email"
          type = "string"
        }
      ]
    }
  }

  output_option_type = "mysql"
  output_option = {
    mysql_output_option = {
      mysql_connection_id = 2
      database            = "$db_name$"
      table               = "$table_name$"
      mode                = "insert"
      retry_limit         = 12
      retry_wait          = 1000
      max_retry_wait      = 1800000
      default_time_zone   = "UTC"
      before_load         = "DELETE FROM test_table WHERE status = 'pending';"
      after_load          = "UPDATE test_table SET updated_at = NOW();"
      mysql_output_option_column_options = [
        {
          name = "description"
          type = "TEXT"
        },
        {
          name      = "price"
          type      = "DECIMAL"
          scale     = 2
          precision = 10
        },
        {
          name = "notes"
          type = "LONGTEXT"
        }
      ]
      custom_variable_settings = [
        {
          name  = "$db_name$"
          type  = "string"
          value = "test_database"
        },
        {
          name  = "$table_name$"
          type  = "string"
          value = "test_table"
        }
      ]
    }
  }
}
