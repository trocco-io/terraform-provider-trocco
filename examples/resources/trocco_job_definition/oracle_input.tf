# Oracle transfer - Query-based transfer (host/port method)
resource "trocco_job_definition" "oracle_query_transfer" {
  name       = "Oracle Query Transfer"
  is_draft   = false
  account_id = 1

  input_option_type = "oracle"

  oracle_input_option = {
    oracle_connection_id = trocco_connection.oracle_basic.id
    database             = "ORCL"
    connection_type      = "sid"
    query                = "SELECT id, name, created_at FROM users WHERE created_at >= '$start_date$'"
    default_time_zone    = "Asia/Tokyo"

    input_option_columns = [
      { name = "id", type = "long" },
      { name = "name", type = "string" },
      { name = "created_at", type = "timestamp", format = "%Y-%m-%d %H:%M:%S", timezone = "Asia/Tokyo" },
    ]

    input_option_column_options = [
      { column_name = "amount", column_value_type = "string" },
    ]
  }

  output_option_type = "bigquery"

  bigquery_output_option = {
    bigquery_dataset_id = trocco_output.bigquery_dataset.id
    table_mode          = "replace"
    table_name          = "users"
  }
}

# Oracle transfer - Incremental load (TNS naming with Wallet)
resource "trocco_job_definition" "oracle_incremental_transfer" {
  name       = "Oracle Incremental Transfer"
  is_draft   = false
  account_id = 1

  input_option_type = "oracle"

  oracle_input_option = {
    oracle_connection_id       = trocco_connection.oracle_wallet.id
    net_service_name           = "orcl_high"
    schema                     = "APP"
    table                      = "orders"
    incremental_loading_enabled = true
    incremental_columns        = "id,updated_at"
    default_time_zone          = "UTC"

    input_option_columns = [
      { name = "id", type = "long" },
      { name = "order_id", type = "long" },
      { name = "amount", type = "double" },
      { name = "updated_at", type = "timestamp", format = "%Y-%m-%d %H:%M:%S" },
    ]
  }

  output_option_type = "bigquery"

  bigquery_output_option = {
    bigquery_dataset_id = trocco_output.bigquery_dataset.id
    table_mode          = "replace"
    table_name          = "orders"
  }
}

# Oracle transfer - Service name method with custom variable
resource "trocco_job_definition" "oracle_service_name_transfer" {
  name       = "Oracle Service Name Transfer"
  is_draft   = false
  account_id = 1

  input_option_type = "oracle"

  oracle_input_option = {
    oracle_connection_id = trocco_connection.oracle_basic.id
    database             = "orcl_prod"
    connection_type      = "service_name"
    query                = "SELECT * FROM products WHERE category = :category"
    default_time_zone    = "UTC"

    input_option_columns = [
      { name = "product_id", type = "long" },
      { name = "product_name", type = "string" },
      { name = "price", type = "double" },
      { name = "category", type = "string" },
    ]

    custom_variable_settings = [
      { key = "category", type = "string", default_value = "electronics" },
    ]
  }

  output_option_type = "bigquery"

  bigquery_output_option = {
    bigquery_dataset_id = trocco_output.bigquery_dataset.id
    table_mode          = "replace"
    table_name          = "products"
  }
}
