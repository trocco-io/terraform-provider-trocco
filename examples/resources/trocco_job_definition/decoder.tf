resource "trocco_job_definition" "decoder_example" {

  input_option                  = {
    # The example is gcs, but it can be applied to file-based input.
    gcs_input_option            = {
      decoder            = {
        match_name       = "regex"
      }
    }
  }
}