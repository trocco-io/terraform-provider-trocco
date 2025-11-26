terraform {
  required_providers {
    trocco = {
      source  = "registry.terraform.io/trocco-io/trocco"
      version = "0.1.0"
    }
  }
}

variable "trocco_api_key" {
  type      = string
  sensitive = true
}

variable "trocco_dev_base_url" {
  type    = string
  default = "https://localhost:4000"
}

provider "trocco" {
  api_key      = var.trocco_api_key
  dev_base_url = var.trocco_dev_base_url
  region       = "japan"
}

resource "trocco_connection" "bigquery" {
  connection_type = "bigquery"

  name        = "BigQuery Example"
  description = "This is a BigQuery connection example"

  project_id               = "example"
  service_account_json_key = <<JSON
  {
    "type": "service_account",
    "project_id": "example-project-id",
    "private_key_id": "example-private-key-id",
    "private_key": "-----BEGIN PRIVATE KEY-----\nyour-private-key\n-----END PRIVATE KEY-----\n"
  }
  JSON
}
