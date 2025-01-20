resource "trocco_pipeline_definition" "trocco_transfer_bulk" {
  name = "trocco_transfer_bulk"

  tasks = [
    {
      key  = "trocco_transfer_bulk"
      type = "trocco_transfer_bulk"

      trocco_transfer_bulk_config = {
        definition_id                 = 1
        is_parallel_execution_allowed = false
        is_stopped_on_errors          = true
        max_errors                    = 1
      }
    }
  ]
}
