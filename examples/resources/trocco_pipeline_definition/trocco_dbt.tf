resource "trocco_pipeline_definition" "trocco_dbt" {
  name = "trocco_dbt"

  tasks = [
    {
      key  = "trocco_dbt"
      type = "trocco_dbt"

      trocco_dbt_config = {
        definition_id = 1
      }
    }
  ]
}
