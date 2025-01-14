resource "trocco_job_definition" "xml_parser_example" {

  input_option                  = {
    # The example is gcs, but it can be applied to file-based input.
    gcs_input_option            = {
      "xml_parser"            = {
        root                    = "root"
        columns                 = [
          {
            name                = "long_col"
            type                = "long"
            path                = "path/to/long_col"
          },
          {
            name                = "str_col"
            type                = "string"
            path                = "path/to/str_col"
          },
          {
            name                = "timestamp_col"
            type                = "timestamp"
            format              = "%Y-%m-%d %H:%M:%S.%N %z"
            timezone            = "UTC"
          }
        ]
      }
    }
  }
}