resource "trocco_job_definition" "filter_row_example" {
  filter_rows = {
    condition = "or"
    filter_row_conditions = [
      {
        argument = "2"
        column   = "col1"
        operator = "greater_equal"
      },
    ]
  }
}
