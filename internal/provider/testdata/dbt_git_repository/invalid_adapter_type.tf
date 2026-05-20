resource "trocco_dbt_git_repository" "test" {
  name         = "test_repo"
  adapter_type = "mysql"
  dbt_version  = "1.11"
  url          = "git@github.com:example/repo.git"
  branch       = "main"
}
