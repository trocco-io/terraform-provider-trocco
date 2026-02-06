resource "trocco_job_definition" "hubspot_output_example" {
  output_option_type = "hubspot"

  output_option = {
    hubspot_output_option = {
      hubspot_connection_id = 1 # require your hubspot connection id
      object_type           = "task"
      mode                  = "merge"
      upsert_key            = "id"
      number_of_parallels   = 2
      associations = [
        {
          to_object_type  = "contact"
          from_object_key = "contact_email"
          to_object_key   = "email"
        },
        {
          to_object_type  = "deal"
          from_object_key = "deal_task"
          to_object_key   = "task"
        }
      ]
    }
  }
}
