resource "trocco_pipeline_definition" "tableau_extract" {
  name = "tableau_extract"

  tasks = [
    {
      key  = "tableau_extract"
      type = "tableau_extract"

      tableau_data_extraction_config = {
        name          = "Example"
        connection_id = 1
        task_id       = "57f1fdc6-aef7-4d4d-a38a-3b73ed529de3"
      }
    }
  ]
}
