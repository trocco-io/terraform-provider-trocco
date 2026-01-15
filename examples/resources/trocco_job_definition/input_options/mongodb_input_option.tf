resource "trocco_job_definition" "mongodb_input_example" {
  input_option_type = "mongodb"
  input_option = {
    mongodb_input_option = {
      mongodb_connection_id       = 1 # require your mongodb connection id
      database                    = "test_database"
      collection                  = "test_collection"
      query                       = "{\"status\": \"active\"}"
      incremental_loading_enabled = true
      incremental_columns         = "created_at"
      last_record                 = "{\"created_at\":\"2024-01-01 00:00:00\"}"
      input_option_columns = [
        {
          name = "_id"
          type = "string"
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
          name     = "created_at"
          type     = "timestamp"
          format   = "%Y-%m-%d %H:%M:%S"
          timezone = "UTC"
        },
      ]
      custom_variable_settings = [
        {
          name  = "$collection_name$"
          type  = "string"
          value = "test_collection"
        }
      ]
    }
  }
}
