resource "trocco_job_definition" "jsonl_parser_example" {

  input_option                  = {
    # The example is gcs, but it can be applied to file-based input.
    gcs_input_option            = {
      "jsonl_parser"            = {
        stop_on_invalid_record = true
        default_time_zone      = "UTC"
        newline                = "LF"
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
            name                = "date_col"
            type                = "timestamp"
            format              = "%Y-%m-%d %H:%M:%S.%N %z"
          }
        ]
      }
    }
  }
}