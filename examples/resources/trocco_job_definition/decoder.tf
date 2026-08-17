resource "trocco_job_definition" "decoder_example" {

  input_option = {
    # The example is gcs, but it can be applied to file-based input.
    gcs_input_option = {
      # `decoder` only configures the relative path inside zip / tar.gz archives.
      # To decompress gzip or bzip2 files, set `decompression_type` instead.
      decoder = {
        match_name = "regex"
      }
    }
  }
}
