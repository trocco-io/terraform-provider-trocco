resource "trocco_job_definition" "mysql_to_kintone_example" {
  name        = "MySQL to Kintone Job Example"
  description = "This job transfers data from MySQL to Kintone"
  filter_columns = [
    {
      name                         = "id",
      src                          = "id",
      type                         = "long",
      default                      = "",
      has_parser                   = true,
      json_expand_enabled          = false,
      json_expand_keep_base_column = false,
      json_expand_columns          = null
    },
  ]

  input_option_type = "mysql"

  input_option = {
    mysql_input_option = {
      mysql_connection_id         = 1
      database                    = "database_name"
      table                       = "table_name"
      incremental_loading_enabled = false
      query                       = "select * from  table_name;"
      input_option_columns = [
        {
          name = "id"
          type = "long"
        },
        {
          name = "type"
          type = "string"
        }
      ]
    }
  }

  output_option_type = "kintone"

  output_option = {
    kintone_output_option = {
      kintone_connection_id = 1
      app_id                = "1"
      guest_space_id        = "1"
      mode                  = "upsert"
      update_key            = "id"
      ignore_nulls          = true
      reduce_key            = "email"
      chunk_size            = 100
      kintone_output_option_column_options = [
        {
          name       = "id"
          field_code = "record_id"
          type       = "NUMBER"
        },
        {
          name       = "created_date"
          field_code = "created_at"
          type       = "DATE"
          timezone   = "Asia/Tokyo"
        },
        {
          name       = "updated_time"
          field_code = "updated_at"
          type       = "TIME"
          timezone   = "Asia/Tokyo"
        },
        {
          name        = "sub_items"
          field_code  = "items_table"
          type        = "SUBTABLE"
          sort_column = "item_order"
        }
      ]
    }
  }
}