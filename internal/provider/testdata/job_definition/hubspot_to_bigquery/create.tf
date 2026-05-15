resource "trocco_job_definition" "hubspot_to_bigquery" {
  name                     = "Hubspot to BigQuery Test"
  description              = "Test job definition for transferring data from Hubspot to BigQuery"
  resource_enhancement     = "medium"
  retry_limit              = 2
  is_runnable_concurrently = true

  filter_columns = [
    {
      name                         = "contact_id"
      src                          = "id"
      type                         = "long"
      default                      = null
      json_expand_enabled          = false
      json_expand_keep_base_column = false
    },
    {
      name                         = "email_address"
      src                          = "email"
      type                         = "string"
      default                      = null
      json_expand_enabled          = false
      json_expand_keep_base_column = false
    },
    {
      name                         = "created_timestamp"
      src                          = "created_at"
      type                         = "timestamp"
      default                      = null
      json_expand_enabled          = false
      json_expand_keep_base_column = false
    }
  ]

  input_option_type = "hubspot"

  input_option = {
    hubspot_input_option = {
      hubspot_connection_id       = 388
      target                      = "object"
      object_type                 = "contact"
      incremental_loading_enabled = false
      input_option_columns = [
        {
          name = "id"
          type = "long"
        },
        {
          name = "email"
          type = "string"
        },
        {
          name = "created_at"
          type = "timestamp"
        }
      ]
      custom_variable_settings = [
        {
          name  = "$test_var$"
          type  = "string"
          value = "test_value"
        }
      ]
    }
  }

  output_option_type = "bigquery"

  output_option = {
    bigquery_output_option = {
      auto_create_dataset                      = false
      bigquery_connection_id                   = trocco_connection.bigquery.id
      bigquery_output_option_clustering_fields = []
      bigquery_output_option_column_options    = []
      bigquery_output_option_merge_keys        = []
      dataset                                  = "test_dataset"
      location                                 = "US"
      mode                                     = "append"
      open_timeout_sec                         = 300
      read_timeout_sec                         = 300
      retries                                  = 5
      send_timeout_sec                         = 300
      table                                    = "hubspot_to_bigquery_test_table"
      template_table                           = ""
      timeout_sec                              = 300
    }
  }
}
