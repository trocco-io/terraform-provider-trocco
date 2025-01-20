resource "trocco_pipeline_definition" "trocco_snowflake_datamart" {
  name = "trocco_snowflake_datamart"

  tasks = [
    {
      key  = "trocco_snowflake_datamart"
      type = "trocco_snowflake_datamart"

      trocco_snowflake_datamart_config = {
        definition_id = 1
      }
    }
  ]
}
