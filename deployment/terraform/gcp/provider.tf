terraform {
  required_version = ">= 1.0"
  required_providers {
    google = {
      source                = "hashicorp/google"
      version               = "~>6.46.0"
      configuration_aliases = [google, google.peer]
    }
  }
}
