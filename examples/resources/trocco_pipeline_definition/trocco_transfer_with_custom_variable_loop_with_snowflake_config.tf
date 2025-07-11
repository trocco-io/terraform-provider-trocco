resource "trocco_pipeline_definition" "trocco_transfer_with_custom_variable_loop_with_snowflake_config" {
  name = "trocco_transfer_with_custom_variable_loop_with_snowflake_config"

  tasks = [
    {
      key  = "trocco_transfer"
      type = "trocco_transfer"

      trocco_transfer_config = {
        definition_id = 1
        custom_variable_loop = {
          type = "snowflake"
          snowflake_config = {
            connection_id = 1
            query         = "select foo, bar from sample"
            variables = [
              "$foo$",
              "$bar$"
            ]
            warehouse = "COMPUTE_WH"
          }
        }
      }
    }
  ]
}
