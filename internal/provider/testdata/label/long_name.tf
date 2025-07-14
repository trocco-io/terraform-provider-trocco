resource "trocco_label" "test_too_long" {
  name  = "This_is_a_very_long_label_name_that_has_exactly_101_characters_for_testing_max_length_validationX"
  color = "#FF0000"
}