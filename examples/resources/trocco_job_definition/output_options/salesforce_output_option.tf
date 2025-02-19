# upsert
resource "trocco_job_definition" "salesforce_output_example" {
  output_option_type = "salesforce"
  output_option = {
    salesforce_output_option = {
      action_type              = "upsert"
      upsert_key               = "id"
      api_version              = "55.0"
      ignore_nulls             = true
      object                   = "test_object"
      salesforce_connection_id = 1 # please set your salesforce connection id
      throw_if_failed          = false
    }
  }
}

# insert
resource "trocco_job_definition" "salesforce_output_example" {
  output_option_type = "salesforce"
  output_option = {
    salesforce_output_option = {
      action_type              = "insert"
      api_version              = "55.0"
      ignore_nulls             = true
      object                   = "test_object"
      salesforce_connection_id = 1 # please set your salesforce connection id
      throw_if_failed          = false
    }
  }
}

