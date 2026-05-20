resource "trocco_dbt_git_repository" "test" {
  name         = "test_repo_renamed"
  description  = "updated description"
  adapter_type = "bigquery"
  dbt_version  = "1.10"
  url          = "git@github.com:example/repo-new.git"
  branch       = "develop"
  subdirectory = "dbt/"
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
