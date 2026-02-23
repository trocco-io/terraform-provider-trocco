# PostgreSQL Output Option Example
resource "trocco_job_definition" "postgresql_output_example" {
  output_option_type = "postgresql"

  output_option = {
    postgresql_output_option = {
      postgresql_connection_id = 1
      database                 = "example_database"
      schema                   = "public"
      table                    = "example_table"
      mode                     = "insert"
      default_time_zone        = "UTC"
    }
  }
}

# PostgreSQL Output Option with Merge Mode Example
resource "trocco_job_definition" "postgresql_output_merge_example" {
  output_option_type = "postgresql"

  output_option = {
    postgresql_output_option = {
      postgresql_connection_id = 1
      database                 = "example_database"
      schema                   = "public"
      table                    = "example_table"
      mode                     = "merge"
      default_time_zone        = "UTC"
      merge_keys               = ["id", "user_id"]
    }
  }
}

# PostgreSQL Output Option with Different Modes
resource "trocco_job_definition" "postgresql_output_truncate_insert_example" {
  output_option_type = "postgresql"

  output_option = {
    postgresql_output_option = {
      postgresql_connection_id = 1
      database                 = "analytics_db"
      schema                   = "staging"
      table                    = "daily_metrics"
      mode                     = "truncate_insert"
      default_time_zone        = "Asia/Tokyo"
    }
  }
}

# PostgreSQL Output Option with Replace Mode
resource "trocco_job_definition" "postgresql_output_replace_example" {
  output_option_type = "postgresql"

  output_option = {
    postgresql_output_option = {
      postgresql_connection_id = 1
      database                 = "production_db"
      schema                   = "public"
      table                    = "user_events"
      mode                     = "replace"
      default_time_zone        = "America/New_York"
    }
  }
}
