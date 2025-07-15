terraform {
  required_providers {
    trocco = {
      source = "registry.terraform.io/trocco-io/trocco"
      version = "0.2.0"
    }
  }
}

provider "trocco" {
  region = "japan"
  api_key = "zFrff9dvkiCYHS1TEnvm9emWHJTazTgbRkzkxBrw3ZquaEm2KmBJi9Vp92"
  dev_base_url = "https://localhost:32777"
}


resource "trocco_user" "test_2" {
  email                  = "terraform_test_2@example.com"
  role                   = "admin"
  password_auto_generated = true
  password               = "3XRambMkp-HwHw"
}

