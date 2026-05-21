resource "trocco_dbt_git_repository" "example" {
  name         = "example repository"
  description  = "example repository description"
  adapter_type = "bigquery"
  dbt_version  = "1.11"
  url          = "git@github.com:example/repo.git"
  ref_type     = "branch"
  branch       = "main"
  subdirectory = "dbt/"
}

resource "trocco_dbt_git_repository" "example_tag" {
  name         = "example repository (tag)"
  adapter_type = "bigquery"
  dbt_version  = "1.11"
  url          = "git@github.com:example/repo.git"
  ref_type     = "tag"
  tag          = "v1.0.0"
}

resource "trocco_dbt_git_repository" "example_commit_hash" {
  name         = "example repository (commit)"
  adapter_type = "bigquery"
  dbt_version  = "1.11"
  url          = "git@github.com:example/repo.git"
  ref_type     = "commit_hash"
  commit_hash  = "0123456789abcdef0123456789abcdef01234567"
}
