# Minimum

resource "trocco_datamart_definition" "minimum" {
  name                     = "example_minimum"
  data_warehouse_type      = "bigquery"
  is_runnable_concurrently = false
  datamart_bigquery_option = {
    bigquery_connection_id = 1
    query                  = "SELECT * FROM tables"
    query_mode             = "insert"
    destination_dataset    = "dist_datasets"
    destination_table      = "dist_tables"
    write_disposition      = "append"
  }
}

# With Optionals
resource "trocco_datamart_definition" "with_optionals" {
  name                     = "example_with_optionals"
  description              = "This is an example with optional fields"
  data_warehouse_type      = "bigquery"
  is_runnable_concurrently = false
  resource_group_id        = 1
  custom_variable_settings = [
    {
      name  = "$string$",
      type  = "string",
      value = "foo",
    },
    {
      name      = "$timestamp$",
      type      = "timestamp",
      quantity  = 1,
      unit      = "hour",
      direction = "ago",
      format    = "%Y-%m-%d %H:%M:%S",
      time_zone = "Asia/Tokyo",
    }
  ]
  datamart_bigquery_option = {
    bigquery_connection_id = 1
    query                  = "SELECT * FROM tables"
    query_mode             = "insert"
    destination_dataset    = "dist_datasets"
    destination_table      = "dist_tables"
    write_disposition      = "append"
  }
}

# BigQuery Insert Mode
resource "trocco_datamart_definition" "bigquery_insert_mode" {
  name                     = "example_bigquery_insert_mode"
  data_warehouse_type      = "bigquery"
  is_runnable_concurrently = false
  datamart_bigquery_option = {
    bigquery_connection_id = 1
    query                  = "SELECT * FROM tables"
    query_mode             = "insert"
    destination_dataset    = "dist_datasets"
    destination_table      = "dist_tables"
    write_disposition      = "append"
    before_load            = "DELETE FROM tables WHERE created_at < '2024-01-01'"
    partitioning           = "time_unit_column"
    partitioning_time      = "DAY"
    partitioning_field     = "created_at"
    clustering_fields      = ["id", "name"]
  }
}

# BigQuery Query Mode
resource "trocco_datamart_definition" "bigquery_query_mode" {
  name                     = "example_bigquery_query_mode"
  data_warehouse_type      = "bigquery"
  is_runnable_concurrently = false
  datamart_bigquery_option = {
    bigquery_connection_id = 1
    query                  = "SELECT * FROM tables"
    query_mode             = "query"
    location               = "asia-northeast1"
  }
}

# With Schedules
resource "trocco_datamart_definition" "with_schedules" {
  name                     = "example_with_schedules"
  data_warehouse_type      = "bigquery"
  is_runnable_concurrently = false
  datamart_bigquery_option = {
    bigquery_connection_id = 1
    query                  = "SELECT * FROM tables"
    query_mode             = "insert"
    destination_dataset    = "dist_datasets"
    destination_table      = "dist_tables"
    write_disposition      = "append"
  }
  schedules = [
    {
      frequency = "hourly"
      minute    = 0
      time_zone = "Asia/Tokyo"
    },
    {
      frequency = "daily"
      hour      = 0
      minute    = 0
      time_zone = "Asia/Tokyo"
    },
    {
      frequency   = "weekly"
      day_of_week = 0
      hour        = 0
      minute      = 0
      time_zone   = "Asia/Tokyo"
    },
    {
      frequency = "monthly"
      day       = 1
      hour      = 0
      minute    = 0
      time_zone = "Asia/Tokyo"
    }
  ]
}

# With Notifications
resource "trocco_datamart_definition" "with_notifications" {
  name                     = "example_with_notifications"
  data_warehouse_type      = "bigquery"
  is_runnable_concurrently = false
  datamart_bigquery_option = {
    bigquery_connection_id = 1
    query                  = "SELECT * FROM tables"
    query_mode             = "insert"
    destination_dataset    = "dist_datasets"
    destination_table      = "dist_tables"
    write_disposition      = "append"
  }
  notifications = [
    {
      destination_type  = "slack"
      slack_channel_id  = 1
      notification_type = "job"
      notify_when       = "finished"
      message           = "@here Job finished."
    },
    {
      destination_type  = "email"
      email_id          = 1
      notification_type = "record"
      record_count      = 100
      record_operator   = "below"
      message           = "Record count is below 100."
    }
  ]
}

# With Labels
resource "trocco_datamart_definition" "with_labels" {
  name                     = "example_with_labels"
  data_warehouse_type      = "bigquery"
  is_runnable_concurrently = false
  datamart_bigquery_option = {
    bigquery_connection_id = 1
    query                  = "SELECT * FROM tables"
    query_mode             = "insert"
    destination_dataset    = "dist_datasets"
    destination_table      = "dist_tables"
    write_disposition      = "append"
  }
  labels = [
    {
      name = "test_label_1"
    },
    {
      name = "test_label_2"
    }
  ]
}
