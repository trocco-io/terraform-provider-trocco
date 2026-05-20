resource "trocco_dbt_git_repository" "example" {
  name         = "example repository"
  description  = "example repository description"
  adapter_type = "bigquery"
  dbt_version  = "1.11"
  url          = "git@github.com:example/repo.git"
  branch       = "main"
  subdirectory = "dbt/"
}
