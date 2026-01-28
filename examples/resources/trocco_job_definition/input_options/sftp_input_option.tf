resource "trocco_job_definition" "sftp_input_example" {
  input_option_type = "sftp"
  input_option = {
    sftp_input_option = {
      sftp_connection_id          = 1 # require your sftp connection id
      path_prefix                 = "/data/files/"
      path_match_pattern          = ".*\\.csv$"
      incremental_loading_enabled = false
      stop_when_file_not_found    = false
      decompression_type          = "guess"
      csv_parser = {
        delimiter = ","
        escape    = "\\"
        quote     = "\""
        columns = [
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
            name = "created_at"
            type = "timestamp"
          },
        ]
      }
    }
  }
}

# Example with incremental loading
resource "trocco_job_definition" "sftp_input_incremental_example" {
  input_option_type = "sftp"
  input_option = {
    sftp_input_option = {
      sftp_connection_id          = 1 # require your sftp connection id
      path_prefix                 = "/data/incremental/"
      path_match_pattern          = "data_{{YYYYMMDD}}\\.csv"
      incremental_loading_enabled = true
      last_path                   = "/data/incremental/data_20250101.csv"
      stop_when_file_not_found    = true
      decompression_type          = "guess"
      csv_parser = {
        delimiter = ","
        escape    = "\\"
        quote     = "\""
        columns = [
          {
            name = "id"
            type = "long"
          },
          {
            name = "value"
            type = "string"
          },
        ]
      }
    }
  }
}

# Example with ZIP decompression
resource "trocco_job_definition" "sftp_input_zip_example" {
  input_option_type = "sftp"
  input_option = {
    sftp_input_option = {
      sftp_connection_id          = 1 # require your sftp connection id
      path_prefix                 = "/data/archives/"
      path_match_pattern          = ".*\\.zip$"
      incremental_loading_enabled = false
      stop_when_file_not_found    = false
      decompression_type          = "zip"
      decoder = {
        match_name = ".*\\.csv$"
      }
      csv_parser = {
        delimiter = ","
        escape    = "\\"
        quote     = "\""
        columns = [
          {
            name = "id"
            type = "long"
          },
          {
            name = "data"
            type = "string"
          },
        ]
      }
    }
  }
}

# Example with JSONL parser
resource "trocco_job_definition" "sftp_input_jsonl_example" {
  input_option_type = "sftp"
  input_option = {
    sftp_input_option = {
      sftp_connection_id          = 1 # require your sftp connection id
      path_prefix                 = "/data/json/"
      path_match_pattern          = ".*\\.jsonl$"
      incremental_loading_enabled = false
      stop_when_file_not_found    = false
      decompression_type          = "guess"
      jsonl_parser = {
        default_time_zone      = "Asia/Tokyo"
        stop_on_invalid_record = true
        newline                = "\n"
        columns = [
          {
            name = "id"
            type = "long"
          },
          {
            name = "name"
            type = "string"
          },
          {
            name      = "timestamp"
            type      = "timestamp"
            time_zone = "Asia/Tokyo"
          },
        ]
      }
    }
  }
}
