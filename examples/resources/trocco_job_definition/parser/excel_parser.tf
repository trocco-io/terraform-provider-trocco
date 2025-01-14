resource "trocco_job_definition" "excel_parser_example" {

  input_option = {
    # The example is gcs, but it can be applied to file-based input.
    gcs_input_option = {
      "excel_parser" = {
        default_time_zone = "Asia/Tokyo"
        sheet_name        = "Sheet1"
        skip_header_lines = 1
        columns = [
          {
            name             = "id"
            type             = "long"
            formula_handling = "cashed_value"
          },
          {
            name             = "str_col"
            type             = "string"
            formula_handling = "cashed_value"
          },
          {
            name             = "date_col"
            type             = "timestamp"
            formula_handling = "evaluate"
            format           = "%Y %m %d"
          }
        ]
      }
    }
  }
}