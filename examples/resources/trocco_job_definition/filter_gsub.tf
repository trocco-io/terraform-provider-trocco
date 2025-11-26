resource "trocco_job_definition" "filter_gsub_example" {
  filter_gsub = [
    {
      column_name = "regex_col"
      pattern     = "/regex/"
      to          = "replace_string"
    }
  ]
}
