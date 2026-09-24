# Create one job definition per SQL Server table with for_each.
# Each table needs its own column list, so the tables are described as a map keyed by table name.
locals {
  sqlserver_tables = {
    orders = [
      { name = "id", type = "long" },
      { name = "customer_id", type = "long" },
      { name = "ordered_at", type = "timestamp" },
    ]
    customers = [
      { name = "id", type = "long" },
      { name = "name", type = "string" },
      { name = "email", type = "string" },
    ]
  }
}

resource "trocco_job_definition" "sqlserver_to_bigquery_by_table" {
  for_each = local.sqlserver_tables

  name                     = "sqlserver_${each.key}_to_bigquery"
  description              = "Transfer dbo.${each.key} from SQL Server to BigQuery"
  is_runnable_concurrently = false
  retry_limit              = 0

  # Optional: run the job on a self-hosted runner cluster instead of the TROCCO-managed infrastructure.
  # self_hosted_runner_cluster_id = 1

  filter_columns = [
    for column in each.value : {
      name = column.name
      src  = column.name
      type = column.type
    }
  ]

  input_option_type = "sqlserver"
  input_option = {
    sqlserver_input_option = {
      sqlserver_connection_id     = 1 // please set your sqlserver connection id
      database                    = "example_database"
      schema                      = "dbo"
      incremental_loading_enabled = true
      table                       = each.key
      incremental_columns         = "id"
      default_time_zone           = "UTC"
      input_option_columns        = each.value
    }
  }

  output_option_type = "bigquery"
  output_option = {
    bigquery_output_option = {
      auto_create_dataset                      = true
      bigquery_connection_id                   = 1
      bigquery_output_option_clustering_fields = []
      bigquery_output_option_column_options    = []
      bigquery_output_option_merge_keys        = []
      dataset                                  = "example_dataset"
      location                                 = "US"
      mode                                     = "append"
      open_timeout_sec                         = 300
      read_timeout_sec                         = 300
      retries                                  = 5
      send_timeout_sec                         = 300
      table                                    = each.key
      template_table                           = ""
      timeout_sec                              = 300
    }
  }
}
