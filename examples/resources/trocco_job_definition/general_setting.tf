resource "trocco_job_definition" "general_example" {
  name                     = "example transfer"
  description              = "example description"
  resource_group_id        = 1
  retry_limit              = 1
  is_runnable_concurrently = true

  # if your account is professional
  resource_enhancement = "medium"

  # if your account can use self-hosted runners, the job runs on the specified cluster
  self_hosted_runner_cluster_id = 1
}
