resource "trocco_job_definition" "parquet_parser_example" {

  input_option = {
    # The example is gcs, but it can be applied to file-based input.
    gcs_input_option = {
      "parquet_parser" = {
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