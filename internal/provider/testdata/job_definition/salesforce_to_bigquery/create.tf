resource "trocco_connection" "test_salesforce" {
  connection_type = "salesforce"
  name            = "Salesforce Example"
  description     = "This is a Salesforce connection example"
  auth_method     = "user_password"
  user_name       = "test_user"
  password        = "test_password"
  security_token  = "test_token"
  auth_end_point  = "https://login.salesforce.com/services/Soap/u/"
}

resource "trocco_connection" "test_bq" {
  connection_type          = "bigquery"
  name                     = "BigQuery Example"
  project_id               = "example"
  service_account_json_key = <<JSON
  {
    "type": "service_account",
    "project_id": "example-project-id",
    "private_key_id": "example-private-key-id",
    "private_key":"-----BEGIN PRIVATE KEY-----\n..."
  }
  JSON
}

resource "trocco_team" "test" {
  name = "test"
  members = [
    {
      user_id = 10626
      role    = "team_admin"
    },
  ]
}

resource "trocco_resource_group" "test" {
  name        = "test"
  description = "test"
  teams = [
    {
      team_id = trocco_team.test.id
      role    = "administrator"
    },
  ]
}

resource "trocco_job_definition" "salesforce_to_bigquery" {
  name                     = "Salesforce to BigQuery Test"
  description              = "Test job definition for transferring data from Salesforce to BigQuery"
  resource_enhancement     = "medium"
  resource_group_id        = trocco_resource_group.test.id
  retry_limit              = 2
  is_runnable_concurrently = true

  input_option_type = "salesforce"
  input_option = {
    salesforce_input_option = {
      columns = [
        {
          name = "Id"
          type = "string"
        },
        {
          name = "Name"
          type = "string"
        },
        {
          name = "Email"
          type = "string"
        },
        {
          name = "CreatedDate"
          type = "timestamp"
        }
      ]
      include_deleted_or_archived_records = false
      is_convert_type_custom_columns      = false
      object                              = "Contact"
      object_acquisition_method           = "soql"
      soql                                = "SELECT Id, Name, Email, CreatedDate FROM Contact"
      salesforce_connection_id            = trocco_connection.test_salesforce.id
    }
  }

  filter_columns = [
    {
      default                      = null
      json_expand_enabled          = false
      json_expand_keep_base_column = false
      name                         = "contact_id"
      src                          = "Id"
      type                         = "string"
    },
    {
      default                      = null
      json_expand_enabled          = false
      json_expand_keep_base_column = false
      name                         = "contact_name"
      src                          = "Name"
      type                         = "string"
    },
    {
      default                      = null
      json_expand_enabled          = false
      json_expand_keep_base_column = false
      name                         = "contact_email"
      src                          = "Email"
      type                         = "string"
    },
    {
      default                      = null
      format                       = "%Y-%m-%d"
      json_expand_enabled          = false
      json_expand_keep_base_column = false
      name                         = "created_date"
      src                          = "CreatedDate"
      type                         = "timestamp"
    }
  ]

  output_option_type = "bigquery"
  output_option = {
    bigquery_output_option = {
      dataset                                  = "test_dataset"
      table                                    = "salesforce_to_bigquery_test_table"
      mode                                     = "append"
      auto_create_dataset                      = true
      timeout_sec                              = 300
      open_timeout_sec                         = 300
      read_timeout_sec                         = 300
      send_timeout_sec                         = 300
      retries                                  = 3
      bigquery_connection_id                   = trocco_connection.test_bq.id
      location                                 = "US"
      bigquery_output_option_clustering_fields = []
      bigquery_output_option_column_options    = []
      bigquery_output_option_merge_keys        = []
    }
  }

  # please create labels if testing in local environment
  # see https://trocco.io/labels#side-nav-labels
  # labels = [
  #   {
  #     name = "test_label1"
  #   },
  #   {
  #     name = "test_label2"
  #   },
  # ]
}
