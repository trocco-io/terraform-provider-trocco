terraform {
  required_providers {
    trocco = {
      source = "registry.terraform.io/trocco-io/trocco"
    }
  }
}

variable "trocco_api_key" {
  type      = string
  sensitive = true
}

provider "trocco" {
  api_key = var.trocco_api_key
  region  = "japan"
}
