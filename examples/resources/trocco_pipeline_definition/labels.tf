resource "trocco_pipeline_definition" "labels" {
  name = "labels"

  labels = [
    "foo",
    "bar",
  ]
}