resource "trocco_pipeline_definition" "trocco_redshift_datamart" {
  name = "trocco_redshift_datamart"

  tasks = [
    {
      key  = "trocco_redshift_datamart"
      type = "trocco_redshift_datamart"

      trocco_redshift_datamart_config = {
        definition_id = 1
      }
    }
  ]
}
