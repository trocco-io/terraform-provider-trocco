resource "trocco_job_definition" "kintone_input_example" {
  input_option_type = "kintone"
  input_option = {
    kintone_input_option = {
      kintone_connection_id = 1 # require your kintone connection id
      app_id                = "236"
      guest_space_id        = "1"
      expand_subtable       = false
      query                 = <<-EOT
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
        {
          name = "name"
          type = "string"
        },
        {
          name = "email"
          type = "string"
        },
        {
          name   = "test"
          type   = "timestamp"
          format = "%Y-%m-%d %H:%M:%S"
        },
      ]
    }
  }
}
