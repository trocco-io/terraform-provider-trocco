# ==============================================================================
# Example 1: Lead Transfer with Date Filter
# ==============================================================================

resource "trocco_job_definition" "marketo_lead_with_date_filter" {
  name        = "Marketo to BigQuery - Lead with Date Filter"
  description = "Transfer Marketo lead data with date range filter"

  input_option_type = "marketo"
  input_option = {
    marketo_input_option = {
      marketo_connection_id   = 123
      target                  = "lead"
      from_date               = "2025-03-01"
      end_date                = "2025-03-18"
      use_updated_at          = false
      polling_interval_second = 60
      bulk_job_timeout_second = 3600

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
    }
  }

  filter_columns = [
    {
      name = "id"
      src  = "id"
      type = "long"
    }
  ]

  output_option_type = "bigquery"
  output_option = {
    bigquery_output_option = {
      bigquery_connection_id            = 456
      dataset                           = "marketing_data"
      table                             = "marketo_leads"
      mode                              = "merge"
      auto_create_dataset               = false
      location                          = "US"
      bigquery_output_option_merge_keys = ["id"]
    }
  }
}

# ==============================================================================
# Example 2: Activity Transfer with Type Filter
# ==============================================================================

resource "trocco_job_definition" "marketo_activity_with_type_filter" {
  name        = "Marketo to BigQuery - Activity with Type Filter"
  description = "Transfer Marketo activity data with specific activity types"

  input_option_type = "marketo"
  input_option = {
    marketo_input_option = {
      marketo_connection_id   = 123
      target                  = "activity"
      from_date               = "2025-03-10"
      end_date                = "2025-03-18"
      activity_type_ids       = [1, 6, 12]
      polling_interval_second = 120
      bulk_job_timeout_second = 7200

      input_option_columns = [
        {
          name = "lead_id"
          type = "long"
        },
        {
          name = "activity_type"
          type = "string"
        }
      ]
    }
  }

  filter_columns = [
    {
      name = "id"
      src  = "id"
      type = "long"
    }
  ]

  output_option_type = "bigquery"
  output_option = {
    bigquery_output_option = {
      bigquery_connection_id            = 456
      dataset                           = "marketing_data"
      table                             = "marketo_activities"
      mode                              = "merge"
      auto_create_dataset               = false
      location                          = "US"
      bigquery_output_option_merge_keys = ["lead_id"]
    }
  }
}

# ==============================================================================
# Example 3: Custom Object with Filter
# ==============================================================================

resource "trocco_job_definition" "marketo_custom_object_with_filter" {
  name        = "Marketo to BigQuery - Custom Object with Filter"
  description = "Transfer Marketo custom object data with ID range filter"

  input_option_type = "marketo"
  input_option = {
    marketo_input_option = {
      marketo_connection_id           = 123
      target                          = "custom_object"
      custom_object_api_name          = "company"
      custom_object_filter_type       = "id"
      custom_object_filter_from_value = 1000
      custom_object_filter_to_value   = 2000

      custom_object_fields = [
        {
          name = "id"
        },
        {
          name = "name"
        },
        {
          name = "revenue"
        }
      ]

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
          name = "revenue"
          type = "double"
        }
      ]
    }
  }

  filter_columns = [
    {
      name = "id"
      src  = "id"
      type = "long"
    }
  ]

  output_option_type = "bigquery"
  output_option = {
    bigquery_output_option = {
      bigquery_connection_id            = 456
      dataset                           = "marketing_data"
      table                             = "marketo_custom_objects"
      mode                              = "merge"
      auto_create_dataset               = false
      location                          = "US"
      bigquery_output_option_merge_keys = ["id"]
    }
  }
}

# ==============================================================================
# Example 4: Folder Transfer
# ==============================================================================

resource "trocco_job_definition" "marketo_folder_transfer" {
  name        = "Marketo to BigQuery - Folder Transfer"
  description = "Transfer Marketo folder structure with folder type"

  input_option_type = "marketo"
  input_option = {
    marketo_input_option = {
      marketo_connection_id = 123
      target                = "folder"
      root_type             = "program"
      root_id               = 456
      max_depth             = 3
      workspace             = "Marketing"

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
          name = "path"
          type = "string"
        }
      ]
    }
  }

  filter_columns = [
    {
      name = "id"
      src  = "id"
      type = "long"
    }
  ]

  output_option_type = "bigquery"
  output_option = {
    bigquery_output_option = {
      bigquery_connection_id            = 456
      dataset                           = "marketing_data"
      table                             = "marketo_folders"
      mode                              = "merge"
      auto_create_dataset               = false
      location                          = "US"
      bigquery_output_option_merge_keys = ["id"]
    }
  }
}

# ==============================================================================
# Example 5: Dynamic Configuration with Custom Variables
# ==============================================================================

resource "trocco_job_definition" "marketo_with_custom_variables" {
  name        = "Marketo to BigQuery - Dynamic Configuration"
  description = "Transfer Marketo lead data with custom variables for dynamic date filtering"

  input_option_type = "marketo"
  input_option = {
    marketo_input_option = {
      marketo_connection_id   = 123
      target                  = "lead"
      from_date               = "$start_date$"
      end_date                = "$end_date$"
      polling_interval_second = 60
      bulk_job_timeout_second = 3600

      input_option_columns = [
        {
          name = "id"
          type = "long"
        },
        {
          name = "email"
          type = "string"
        }
      ]

      custom_variable_settings = [
        {
          name      = "$start_date$"
          type      = "timestamp_runtime"
          quantity  = 7
          unit      = "date"
          direction = "ago"
          format    = "%Y-%m-%d"
          time_zone = "Asia/Tokyo"
        },
        {
          name      = "$end_date$"
          type      = "timestamp_runtime"
          quantity  = 1
          unit      = "date"
          direction = "ago"
          format    = "%Y-%m-%d"
          time_zone = "Asia/Tokyo"
        }
      ]
    }
  }

  filter_columns = [
    {
      name = "id"
      src  = "id"
      type = "long"
    }
  ]

  output_option_type = "bigquery"
  output_option = {
    bigquery_output_option = {
      bigquery_connection_id            = 456
      dataset                           = "marketing_data"
      table                             = "marketo_leads_dynamic"
      mode                              = "merge"
      auto_create_dataset               = false
      location                          = "US"
      bigquery_output_option_merge_keys = ["id"]
    }
  }
}
