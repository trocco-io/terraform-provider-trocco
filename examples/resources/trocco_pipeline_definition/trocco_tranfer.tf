resource "trocco_pipeline_definition" "trocco_transfer" {
  name = "trocco_transfer"

  tasks = [
    {
      key  = "trocco_transfer"
      type = "trocco_transfer"

      trocco_transfer_config = {
        definition_id = 1
      }
    }
  ]
}
