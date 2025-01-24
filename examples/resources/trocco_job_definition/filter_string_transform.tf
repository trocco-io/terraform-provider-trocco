resource "trocco_job_definition" "filter_string_transform_example" {
  filter_string_transforms = [
    {
      column_name = "transform_col"
      type        = "normalize_nfkc"
    }
  ]
}
