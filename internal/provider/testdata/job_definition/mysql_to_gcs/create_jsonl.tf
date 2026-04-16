resource "trocco_connection" "test_mysql" {
  connection_type = "mysql"
  name            = "MySQL Example"
  host            = "db.example.com"
  port            = 3306
  user_name       = "root"
  password        = "password"
}

resource "trocco_connection" "test_gcs" {
  connection_type = "gcs"

  name        = "GCS Example"
  description = "This is a Google Cloud Storage(GCS) connection example"

  project_id               = "example-project-id"
  service_account_json_key = <<JSON
  {
    "type": "service_account",
    "project_id": "example-project-id",
    "private_key_id": "example-private-key-id",
    "private_key":"-----BEGIN PRIVATE KEY-----\n..."
  }
  JSON
  service_account_email    = "joe@example-project.iam.gserviceaccount.com"
  application_name         = "example-application-name"
}

resource "trocco_job_definition" "mysql_to_gcs" {
  name                     = "MySQL to GCS Test"
  description              = "Test job definition for transferring data from MySQL to GCS"
  resource_enhancement     = "medium"
  retry_limit              = 2
  is_runnable_concurrently = false

  filter_columns = [
    {
      default                      = ""
      format                       = ""
      json_expand_enabled          = false
      json_expand_keep_base_column = false
      name                         = "id"
      src                          = "id"
      type                         = "long"
    },
    {
      default                      = ""
      format                       = ""
      json_expand_enabled          = false
      json_expand_keep_base_column = false
      name                         = "type"
      src                          = "type"
      type                         = "string"
    },
  ]

  input_option_type = "mysql"
  input_option = {
    mysql_input_option = {
      mysql_connection_id         = trocco_connection.test_mysql.id
      database                    = "test_database"
      table                       = "test_table"
      connect_timeout             = 300
      socket_timeout              = 1800
      incremental_loading_enabled = false
      default_time_zone           = "Asia/Tokyo"
      use_legacy_datetime_code    = false
      input_option_columns = [
        {
          name = "id"
          type = "long"
        },
        {
          name = "type"
          type = "string"
        },
      ]
      query = <<-EOT
        SELECT *
        FROM test_table
        WHERE id < 100
      EOT
    }
  }

  output_option_type = "gcs"
  output_option = {
    gcs_output_option = {
      gcs_connection_id       = trocco_connection.test_gcs.id
      bucket                  = "my-test-bucket"
      path_prefix             = "output/test/"
      file_ext                = ".jsonl"
      sequence_format         = ".%03d.%02d"
      is_minimum_output_tasks = false
      formatter_type          = "jsonl"
      encoder_type            = "gzip"
      jsonl_formatter = {
        encoding    = "UTF-8"
        newline     = "LF"
        date_format = "%Y-%m-%d %H:%M:%S"
        timezone    = "UTC"
      }
    }
  }
}
