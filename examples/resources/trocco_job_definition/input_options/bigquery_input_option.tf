resource "trocco_job_definition" "bigquery_input_example" {
  input_option_type = "bigquery"
  input_option = {
    bigquery_input_option = {
      bigquery_connection_id   = 1
      gcs_uri                  = "test_bucket"
      gcs_uri_format           = "bucket"
      query                    = "SELECT * FROM `test_dataset.test_table`"
      temp_dataset             = "temp_dataset"
      location                 = "asia-northeast1"
      is_standard_sql          = true
      cleanup_gcs_files        = true
      file_format              = "CSV"
      cache                    = true
      bigquery_job_wait_second = 600

      columns = [
        {
          name = "col1__c"
          type = "string"
        }
      ]
    }
  }
}
