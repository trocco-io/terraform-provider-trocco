variable "team_name" {
  type = string
}
variable "team_description" {
  type = string
  default = ""
}
variable "team_members" {
  type = list(object({
    user_id = number
    role    = string
  }))
}
