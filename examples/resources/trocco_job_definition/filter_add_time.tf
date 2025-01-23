resource "trocco_job_definition" "filter_add_time_example" {
  filter_add_time = {
    column_name      = "time"
    time_zone        = "Asia/Tokyo"
    timestamp_format = "%Y-%m-%d %H:%M:%S.%N"
    type             = "timestamp"
  }
}