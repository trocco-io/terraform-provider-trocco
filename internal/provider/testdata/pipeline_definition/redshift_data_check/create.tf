resource "trocco_pipeline_definition" "redshift_data_check_query_check" {
  name = "redshift_data_check"

  tasks = [
    {
      key  = "redshift_data_check"
      type = "redshift_data_check"

      redshift_data_check_config = {
        name          = "Example"
        connection_id = 256
        query         = <<SQL
          SELECT COUNT(*) FROM examples
        SQL
        operator      = "equal"
        query_result  = 1
        accepts_null  = false
        database      = "example_db"
      }
    }
  ]
}
