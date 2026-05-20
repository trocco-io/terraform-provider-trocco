resource "trocco_dbt_git_repository" "test" {
  name         = "test_repo"
  description  = "test description"
  adapter_type = "bigquery"
  dbt_version  = "1.11"
  url          = "git@github.com:example/repo.git"
  branch       = "main"
}

resource "trocco_dbt_git_repository" "test_minimal" {
  name         = "test_repo_minimal"
  adapter_type = "snowflake"
  dbt_version  = "1.10"
  url          = "git@github.com:example/minimal.git"
  branch       = "main"
}

resource "trocco_dbt_git_repository" "test_with_subdirectory" {
  name         = "test_repo_subdir"
  adapter_type = "redshift"
  dbt_version  = "1.9"
  url          = "git@github.com:example/subdir.git"
  branch       = "develop"
  subdirectory = "dbt/"
}

resource "trocco_dbt_git_repository" "test_with_tag" {
  name         = "test_repo_tag"
  adapter_type = "bigquery"
  dbt_version  = "1.11"
  url          = "git@github.com:example/tag.git"
  ref_type     = "tag"
  tag          = "v1.0.0"
}

resource "trocco_dbt_git_repository" "test_with_commit_hash" {
  name         = "test_repo_commit"
  adapter_type = "bigquery"
  dbt_version  = "1.11"
  url          = "git@github.com:example/commit.git"
  ref_type     = "commit_hash"
  commit_hash  = "0123456789abcdef0123456789abcdef01234567"
}
