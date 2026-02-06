resource "trocco_job_definition" "mysql_to_hubspot_example" {
  name                     = "mysql_to_hubspot_example"
  description              = "Example job definition to transfer data from MySQL to HubSpot"
  is_runnable_concurrently = false
  retry_limit              = 0

  filter_columns = [
    {
      name                         = "id"
      src                          = "id"
      type                         = "long"
      default                      = null
      has_parser                   = false
      json_expand_enabled          = false
      json_expand_keep_base_column = false
      json_expand_columns          = []
    },
    {
      name                         = "task_name"
      src                          = "task_name"
      type                         = "string"
      default                      = null
      has_parser                   = false
      json_expand_enabled          = false
      json_expand_keep_base_column = false
      json_expand_columns          = []
    },
    {
      name                         = "contact_email"
      src                          = "contact_email"
      type                         = "string"
      default                      = null
      has_parser                   = false
      json_expand_enabled          = false
      json_expand_keep_base_column = false
      json_expand_columns          = []
    },
    {
      name                         = "deal_task"
      src                          = "deal_task"
      type                         = "string"
      default                      = null
      has_parser                   = false
      json_expand_enabled          = false
      json_expand_keep_base_column = false
      json_expand_columns          = []
    }
  ]

  input_option_type = "mysql"
  input_option = {
    mysql_input_option = {
      mysql_connection_id         = 1 # require your mysql connection id
      database                    = "example_database"
      table                       = "example_table"
      query                       = "SELECT * FROM example_table WHERE id < 100;"
      incremental_loading_enabled = false
      connect_timeout             = 300
      socket_timeout              = 1800
      fetch_rows                  = 10000
      default_time_zone           = ""
      use_legacy_datetime_code    = false
      input_option_columns = [
        {
          name = "id"
          type = "long"
        },
        {
          name = "task_name"
          type = "string"
        },
        {
          name = "contact_email"
          type = "string"
        },
        {
          name = "deal_task"
          type = "string"
        }
      ]
    }
  }

  output_option_type = "hubspot"
  output_option = {
    hubspot_output_option = {
      hubspot_connection_id = 1 # require your hubspot connection id
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
