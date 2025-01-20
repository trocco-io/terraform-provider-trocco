resource "trocco_pipeline_definition" "trocco_bigquery_datamart" {
  name = "trocco_bigquery_datamart"

  tasks = [
    {
      key  = "trocco_bigquery_datamart"
      type = "trocco_bigquery_datamart"

      trocco_bigquery_datamart_config = {
        definition_id = 1
      }
    }
  ]
}
