resource "trocco_job_definition" "kintone_output_example" {
  output_option_type = "kintone"

  output_option = {
    kintone_output_option = {
      kintone_connection_id = 1
      app_id                = 1
      guest_space_id        = 1
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
