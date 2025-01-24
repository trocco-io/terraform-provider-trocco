resource "trocco_job_definition" "filter_hash_example" {
  filter_hashes = [
    {
      name = "hash_col"
    }
  ]
}
