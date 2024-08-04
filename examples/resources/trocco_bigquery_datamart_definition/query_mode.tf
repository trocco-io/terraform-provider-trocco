resource "trocco_bigquery_datamart_definition" "query_mode" {
  name                     = "example_query_mode"
  is_runnable_concurrently = false
  bigquery_connection_id   = 1
  query                    = "SELECT * FROM tables"
  query_mode               = "query"
  location                 = "asia-northeast1"
}
