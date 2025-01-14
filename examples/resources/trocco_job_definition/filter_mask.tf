resource "trocco_job_definition" "filter_mask_example" {
  filter_masks                = [
    {
      name      = "mask_all_string_col"
      mask_type = "all"
      length    = 10
    },
    {
      name      = "mask_email_col"
      mask_type = "email"
      length    = 10
    },
    {
      name      = "mask_regex_col"
      mask_type = "regex"
      pattern   = "/regex/"
    },
    {
      name        = "partial_string"
      length      = 10
      start_index = 2
      end_index   = 2
      mask_type   = "substring"

    },
  ]
}
