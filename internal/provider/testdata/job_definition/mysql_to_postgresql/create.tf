resource "trocco_connection" "test_mysql" {
  connection_type = "mysql"
  name            = "MySQL Example"
  host            = "db.example.com"
  port            = 65535
  user_name       = "root"
  password        = "password"
}

resource "trocco_connection" "test_postgresql" {
  connection_type = "postgresql"
  name            = "PostgreSQL Example"
  host            = "postgres.example.com"
  port            = 5432
  user_name       = "postgres"
  password        = "password"
  driver          = "postgresql_42_5_1"
}

resource "trocco_job_definition" "mysql_to_postgresql" {
  name                     = "test job_definition"
  description              = "test description"
  resource_enhancement     = "medium"
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
    {
      default                      = ""
      format                       = ""
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
      database                    = "test_database"
      fetch_rows                  = 1000
      incremental_loading_enabled = false
      table                       = "test_table"
      default_time_zone           = "Asia/Tokyo"
      use_legacy_datetime_code    = false
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
      mysql_connection_id = trocco_connection.test_mysql.id
      query               = <<-EOT
                select
                    *
                from
      	          example_table;
      EOT
      socket_timeout      = 1801
    }
  }
  output_option_type = "postgresql"
  output_option = {
    postgresql_output_option = {
      database                 = "test_database"
      schema                   = "public"
      table                    = "test_table"
      mode                     = "insert"
      default_time_zone        = "UTC"
      postgresql_connection_id = trocco_connection.test_postgresql.id
    }
  }
}
