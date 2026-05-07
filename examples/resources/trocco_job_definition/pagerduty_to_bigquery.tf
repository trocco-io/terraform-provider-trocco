resource "trocco_connection" "bigquery_example" {
  connection_type          = "bigquery"
  name                     = "BigQuery Connection"
  project_id               = "your-project-id"
  service_account_json_key = "{}"
}

resource "trocco_connection" "pagerduty_example" {
  connection_type = "pagerduty"
  name            = "Pagerduty Connection"
  # fill in required connection fields
  api_key = "your-api-key"
}

# Example: pagerduty → bigquery (path = escalation_policies)
resource "trocco_job_definition" "pagerduty_EscalationPolicies_to_bigquery" {
  name               = "Pagerduty escalation_policies to BigQuery"
  input_option_type  = "pagerduty"
  output_option_type = "bigquery"

  filter_columns = []

  input_option = {
    pagerduty_input_option = {
      pagerduty_connection_id = trocco_connection.pagerduty_example.id
      path                    = "escalation_policies"
    }
  }

  output_option = {
    bigquery_output_option = {
      bigquery_connection_id = trocco_connection.bigquery_example.id
      dataset                = "your_dataset"
      table                  = "pagerduty_EscalationPolicies"
      mode                   = "append"
    }
  }
}

# Example: pagerduty → bigquery (path = incidents)
resource "trocco_job_definition" "pagerduty_Incidents_to_bigquery" {
  name               = "Pagerduty incidents to BigQuery"
  input_option_type  = "pagerduty"
  output_option_type = "bigquery"

  filter_columns = []

  input_option = {
    pagerduty_input_option = {
      pagerduty_connection_id = trocco_connection.pagerduty_example.id
      path                    = "incidents"
      since                   = "2023-01-01T00:00:00Z"
      until                   = "2023-02-01T00:00:00Z"
    }
  }

  output_option = {
    bigquery_output_option = {
      bigquery_connection_id = trocco_connection.bigquery_example.id
      dataset                = "your_dataset"
      table                  = "pagerduty_Incidents"
      mode                   = "append"
    }
  }
}

# Example: pagerduty → bigquery (path = log_entries)
resource "trocco_job_definition" "pagerduty_LogEntries_to_bigquery" {
  name               = "Pagerduty log_entries to BigQuery"
  input_option_type  = "pagerduty"
  output_option_type = "bigquery"

  filter_columns = []

  input_option = {
    pagerduty_input_option = {
      pagerduty_connection_id = trocco_connection.pagerduty_example.id
      path                    = "log_entries"
      # since = <string>
      # until = <string>
    }
  }

  output_option = {
    bigquery_output_option = {
      bigquery_connection_id = trocco_connection.bigquery_example.id
      dataset                = "your_dataset"
      table                  = "pagerduty_LogEntries"
      mode                   = "append"
    }
  }
}

# Example: pagerduty → bigquery (path = oncalls)
resource "trocco_job_definition" "pagerduty_Oncalls_to_bigquery" {
  name               = "Pagerduty oncalls to BigQuery"
  input_option_type  = "pagerduty"
  output_option_type = "bigquery"

  filter_columns = []

  input_option = {
    pagerduty_input_option = {
      pagerduty_connection_id = trocco_connection.pagerduty_example.id
      path                    = "oncalls"
      # earliest = <boolean>
      # since = <string>
      # until = <string>
    }
  }

  output_option = {
    bigquery_output_option = {
      bigquery_connection_id = trocco_connection.bigquery_example.id
      dataset                = "your_dataset"
      table                  = "pagerduty_Oncalls"
      mode                   = "append"
    }
  }
}

# Example: pagerduty → bigquery (path = priorities)
resource "trocco_job_definition" "pagerduty_Priorities_to_bigquery" {
  name               = "Pagerduty priorities to BigQuery"
  input_option_type  = "pagerduty"
  output_option_type = "bigquery"

  filter_columns = []

  input_option = {
    pagerduty_input_option = {
      pagerduty_connection_id = trocco_connection.pagerduty_example.id
      path                    = "priorities"
    }
  }

  output_option = {
    bigquery_output_option = {
      bigquery_connection_id = trocco_connection.bigquery_example.id
      dataset                = "your_dataset"
      table                  = "pagerduty_Priorities"
      mode                   = "append"
    }
  }
}

# Example: pagerduty → bigquery (path = schedules)
resource "trocco_job_definition" "pagerduty_Schedules_to_bigquery" {
  name               = "Pagerduty schedules to BigQuery"
  input_option_type  = "pagerduty"
  output_option_type = "bigquery"

  filter_columns = []

  input_option = {
    pagerduty_input_option = {
      pagerduty_connection_id = trocco_connection.pagerduty_example.id
      path                    = "schedules"
    }
  }

  output_option = {
    bigquery_output_option = {
      bigquery_connection_id = trocco_connection.bigquery_example.id
      dataset                = "your_dataset"
      table                  = "pagerduty_Schedules"
      mode                   = "append"
    }
  }
}

# Example: pagerduty → bigquery (path = services)
resource "trocco_job_definition" "pagerduty_Services_to_bigquery" {
  name               = "Pagerduty services to BigQuery"
  input_option_type  = "pagerduty"
  output_option_type = "bigquery"

  filter_columns = []

  input_option = {
    pagerduty_input_option = {
      pagerduty_connection_id = trocco_connection.pagerduty_example.id
      path                    = "services"
    }
  }

  output_option = {
    bigquery_output_option = {
      bigquery_connection_id = trocco_connection.bigquery_example.id
      dataset                = "your_dataset"
      table                  = "pagerduty_Services"
      mode                   = "append"
    }
  }
}

# Example: pagerduty → bigquery (path = teams)
resource "trocco_job_definition" "pagerduty_Teams_to_bigquery" {
  name               = "Pagerduty teams to BigQuery"
  input_option_type  = "pagerduty"
  output_option_type = "bigquery"

  filter_columns = []

  input_option = {
    pagerduty_input_option = {
      pagerduty_connection_id = trocco_connection.pagerduty_example.id
      path                    = "teams"
    }
  }

  output_option = {
    bigquery_output_option = {
      bigquery_connection_id = trocco_connection.bigquery_example.id
      dataset                = "your_dataset"
      table                  = "pagerduty_Teams"
      mode                   = "append"
    }
  }
}

# Example: pagerduty → bigquery (path = users)
resource "trocco_job_definition" "pagerduty_Users_to_bigquery" {
  name               = "Pagerduty users to BigQuery"
  input_option_type  = "pagerduty"
  output_option_type = "bigquery"

  filter_columns = []

  input_option = {
    pagerduty_input_option = {
      pagerduty_connection_id = trocco_connection.pagerduty_example.id
      path                    = "users"
    }
  }

  output_option = {
    bigquery_output_option = {
      bigquery_connection_id = trocco_connection.bigquery_example.id
      dataset                = "your_dataset"
      table                  = "pagerduty_Users"
      mode                   = "append"
    }
  }
}

