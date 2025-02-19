# acquisition_method: soql
resource "trocco_job_definition" "salesforce_input_example" {
  input_option_type = "salesforce"
  input_option = {
    salesforce_input_option = {
      columns = [
        {
          name = "col1__c"
          type = "string"
        }
      ]
      include_deleted_or_archived_records = true
      is_convert_type_custom_columns      = false
      object                              = "test_object"
      object_acquisition_method           = "soql"
      soql                                = "select * from test_object"
      salesforce_connection_id            = 1 # pelase set your salesforce connection id
    }
  }
}

# acquisition_method: all_columns
resource "trocco_job_definition" "salesforce_input_example" {
  input_option_type = "salesforce"
  input_option = {
    salesforce_input_option = {
      columns = [
        {
          name = "col1__c"
          type = "string"
        }
      ]
      include_deleted_or_archived_records = true
      is_convert_type_custom_columns      = false
      object                              = "test_object"
      object_acquisition_method           = "all_columns"
      salesforce_connection_id            = 1 # pelase set your salesforce connection id
    }
  }
}
