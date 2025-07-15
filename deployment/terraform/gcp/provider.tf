terraform {
  required_version = ">= 1.0"
  required_providers {
    google = {
      source                = "hashicorp/google"
      version               = "~>4.1.0"
      configuration_aliases = [google, google.peer]
    }
  }
}
