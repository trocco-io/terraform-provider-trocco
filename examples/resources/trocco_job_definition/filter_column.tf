
resource "trocco_job_definition" "filter_column_example" {
  filter_columns = [
    {
      default                      = ""
      format                       = "%Y"
      json_expand_columns          = []
      json_expand_enabled          = false
      json_expand_keep_base_column = false
      name                         = "id"
      src                          = "id"
      type                         = "long"
    },
    {
      default                      = "default value"
      json_expand_columns          = []
      json_expand_enabled          = false
      json_expand_keep_base_column = false
      name                         = "name"
      src                          = "name"
      type                         = "string"
    },
    {
      default                      = ""
      json_expand_columns          = []
      json_expand_enabled          = false
      json_expand_keep_base_column = false
      name                         = "timestamp"
      src                          = "timestamp"
      type                         = "timestamp"
      format                       = "%Y%m%d"
    },
    {
      default                      = ""
      json_expand_enabled          = true
      json_expand_keep_base_column = true
      name                         = "json_col"
      src                          = "json_src"
      type                         = "json"
      json_expand_columns = [
        {
          json_path = "person.name"
          name      = "json_expand_col"
          timezone  = "UTC/ETC"
          type      = "string"
        },
      ]
    }
  ]
}

