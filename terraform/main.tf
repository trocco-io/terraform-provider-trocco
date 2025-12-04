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

resource "trocco_connection" "databricks_pat" {
  connection_type = "databricks"

  name                  = "Databricks Example with PAT Auth "
  description           = "This is a Databricks connection example"
  host                  = "example.databricks.com"
  http_path             = "/sql/1.0/warehouses/xxxx-xxxx-xxxx-xxxx"
  auth_type             = "pat"
  personal_access_token = "dapiXXXXXXXXXXXXXXXXXXXX"
}

resource "trocco_connection" "databricks_oauth2" {
  connection_type = "databricks"

  name                 = "Databricks Example with OAuth2 modification"
  description          = "This is a Databricks connection example using OAuth2"
  host                 = "example.databricks.com"
  http_path            = "/sql/1.0/warehouses/xxxx-xxxx-xxxx-xxxx"
  auth_type            = "oauth-m2m"
  oauth2_client_id     = "your-oauth2-client-id"
  oauth2_client_secret = "your-oauth2-client-secret"
}

resource "trocco_job_definition" "databricks_to_bigquery" {
  name        = "Databricks2 to BigQuery Job"
  description = "This job transfers data from Databricks to BigQuery"
  filter_columns = [
    {
      name                         = "id",
      src                          = "id",
      type                         = "long",
      default                      = "",
      has_parser                   = true,
      json_expand_enabled          = false,
      json_expand_keep_base_column = false,
      json_expand_columns          = null
    },
  ]

  input_option_type = "databricks"

  input_option = {
    databricks_input_option = {
      databricks_connection_id = trocco_connection.databricks_pat.id
      catalog_name             = "catalog_name example"
      schema_name              = "schema_name example"
      query                    = "select * from example_table"
      input_option_columns = [
        {
          name = "id"
          type = "long"
        },
      ]
    }
  }

  output_option_type = "bigquery"

  output_option = {
    bigquery_output_option = {
      auto_create_dataset                      = false
      bigquery_connection_id                   = 1
      bigquery_output_option_clustering_fields = []
      bigquery_output_option_column_options    = []
      bigquery_output_option_merge_keys        = []
      dataset                                  = "example_dataset"
      location                                 = "US"
      mode                                     = "append"
      open_timeout_sec                         = 300
      read_timeout_sec                         = 300
      retries                                  = 5
      send_timeout_sec                         = 300
      table                                    = "example_table"
      template_table                           = ""
      timeout_sec                              = 300
    }
  }
}
