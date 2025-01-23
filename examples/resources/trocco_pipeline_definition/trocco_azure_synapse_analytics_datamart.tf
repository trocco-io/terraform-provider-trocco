resource "trocco_pipeline_definition" "trocco_azure_synapse_analytics_datamart" {
  name = "trocco_azure_synapse_analytics_datamart"

  tasks = [
    {
      key  = "trocco_azure_synapse_analytics_datamart"
      type = "trocco_azure_synapse_analytics_datamart"

      trocco_azure_synapse_analytics_datamart_config = {
        definition_id = 26
      }
    }
  ]
}
