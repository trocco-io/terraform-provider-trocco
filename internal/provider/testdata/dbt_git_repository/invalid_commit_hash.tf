resource "trocco_dbt_git_repository" "test" {
  name         = "test_repo"
  adapter_type = "bigquery"
  dbt_version  = "1.11"
  url          = "git@github.com:example/repo.git"
  ref_type     = "commit_hash"
  commit_hash  = "not-a-valid-hash"
}
