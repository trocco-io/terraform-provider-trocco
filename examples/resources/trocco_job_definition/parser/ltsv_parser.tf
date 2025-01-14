resource "trocco_job_definition" "ltsv_parser_example" {

  input_option                  = {
    # The example is gcs, but it can be applied to file-based input.
    gcs_input_option            = {
      "ltsv_parser"            = {
        newline                = "CRLF"
        charset                = "UTF-8"
        columns                 = [
          {
            name                = "id"
            type                = "long"
          },
          {
            name                = "str_col"
            type                = "string"
          },
          {
            name                = "date_col",
            type                = "timestamp"
          }
        ]
      }
    }
  }
}