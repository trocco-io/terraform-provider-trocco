resource "trocco_job_definition" "jsonpath_parser_example" {
  input_option = {
    # The example is gcs, but it can be applied to file-based input.
    gcs_input_option = {
      "jsonpath_parser" = {
        root              = "$root"
        default_time_zone = "UTC"
        columns = [
          {
            name = "long_col"
            type = "long"
          },
          {
            name = "str_col"
            type = "string"
          },
          {
            name   = "timestamp_col"
            type   = "timestamp"
            format = "%Y-%m-%d %H:%M:%S.%N %z"
          }
        ]
      }
    }
  }
}
