resource "trocco_label" "test" {
  name        = "Test Label"
  color       = "#FFFFFF"
  description = "This is a test label"
}

resource "trocco_label" "test2" {
  name        = "Test Label 2"
  color       = "#FFFFFF"
  description = "This is a test label"
}

resource "trocco_label" "test_omitted_description" {
  name  = "Test Label Using Omitted Description"
  color = "#FFFFFF"
}

resource "trocco_label" "test_empty_description" {
  name        = "Test Label Using Empty Description"
  color       = "#FFFFFF"
  description = ""
}
