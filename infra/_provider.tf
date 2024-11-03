terraform {
  required_version = ">= 0.12"

  required_providers {
    adguard = {
      source  = "gmichels/adguard"
      version = "1.3.0"
    }
  }
}

provider "adguard" {
  host     = var.adguard_host
  username = var.adguard_username
  password = var.adguard_password
  scheme   = var.adguard_scheme
  timeout  = 10
}
