resource "trocco_pipeline_definition" "trocco_pipeline" {
  name = "trocco_pipeline"

  tasks = [
    {
      key  = "trocco_pipeline"
      type = "trocco_pipeline"

      trocco_pipeline_config = {
        definition_id = 1
      }
    }
  ]
}
