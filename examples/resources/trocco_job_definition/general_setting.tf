resource "trocco_job_definition" "general_example" {
  name                     = "example tranfer"
  description              = "example description"
  resource_group_id        = 1
  retry_limit              = 1
  is_runnable_concurrently = true

  # if your account is professional
  resource_enhancement = "medium"
}
