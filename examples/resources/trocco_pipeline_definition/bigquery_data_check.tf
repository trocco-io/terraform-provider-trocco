resource "trocco_pipeline_definition" "bigquery_data_check" {
  name = "bigquery_data_check"

  tasks = [
    {
      key  = "bigquery_data_check"
      type = "bigquery_data_check"

      bigquery_data_check_config = {
        name          = "Example"
        connection_id = 1
        query         = "SELECT COUNT(id) FROM examples"
        operator      = "equal"
        query_result  = 1
        accepts_null  = false
      }
    }
  ]
}
