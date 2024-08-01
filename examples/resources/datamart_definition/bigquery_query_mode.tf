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
