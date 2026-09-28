resource "trocco_connection" "test_oracle_tns" {
  connection_type = "oracle"
  name            = "Oracle TNS Example"
  user_name       = "system"
  password        = "password"
  driver          = "19c-ojdbc8.jar"
  tns_admin_ora   = <<-EOT
    ORCL_high = (DESCRIPTION =
      (ADDRESS = (PROTOCOL = TCPS)(HOST = adb.example.oraclecloud.com)(PORT = 1522))
      (CONNECT_DATA = (SERVICE_NAME = orcl_high.adb.oraclecloud.com))
    )
  EOT
}

resource "trocco_team" "test_tns" {
  name = "test_tns"
  members = [
    {
      user_id = 10626
      role    = "team_admin"
    },
  ]
}

resource "trocco_resource_group" "test_tns" {
  name        = "test_tns"
  description = "test"
  teams = [
    {
      team_id = trocco_team.test_tns.id
      role    = "administrator"
    },
  ]
}

resource "trocco_job_definition" "oracle_tns_to_bigquery" {
  name                     = "Oracle TNS to BigQuery Test"
  description              = "Test job definition for transferring data from Oracle (TNS) to BigQuery"
  resource_enhancement     = "medium"
  resource_group_id        = trocco_resource_group.test_tns.id
  retry_limit              = 3
  is_runnable_concurrently = true

  input_option_type = "oracle"
  input_option = {
    oracle_input_option = {
      oracle_connection_id        = trocco_connection.test_oracle_tns.id
      net_service_name            = "orcl_high"
      incremental_loading_enabled = false
      default_time_zone           = "UTC"
      query                       = <<-EOT
        select
            *
        from
            example_table;
      EOT
      input_option_columns = [
        {
          name = "id"
          type = "long"
        },
      ]
    }
  }

  filter_columns = [
    {
      default                      = ""
      format                       = "%Y"
      json_expand_enabled          = false
      json_expand_keep_base_column = false
      name                         = "id"
      src                          = "id"
      type                         = "long"
    },
  ]

  output_option_type = "bigquery"
  output_option = {
    bigquery_output_option = {
      dataset                                  = "test_dataset"
      table                                    = "oracle_tns_to_bigquery_test_table"
      mode                                     = "append"
      auto_create_dataset                      = true
      timeout_sec                              = 300
      open_timeout_sec                         = 300
      read_timeout_sec                         = 300
      send_timeout_sec                         = 300
      retries                                  = 3
      bigquery_connection_id                   = trocco_connection.bigquery.id
      location                                 = "US"
      bigquery_output_option_clustering_fields = []
      bigquery_output_option_column_options    = []
      bigquery_output_option_merge_keys        = []
    }
  }
}
