resource "trocco_job_definition" "redshift_input_option_example" {
  input_option_type = "redshift"
  input_option = {
    redshift_input_option = {
      redshift_connection_id = 1 # require your redshift connection id
      database               = "analytics"
      query                  = "SELECT * FROM test_table WHERE status = 'active'"
      schema                 = "public"
      fetch_rows             = 1000
      connect_timeout        = 30
      socket_timeout         = 60
    }
  }
}
