terraform {
  required_providers {
    google = {
      source                = "hashicorp/google"
      version               = ">= 4.1"
      configuration_aliases = [google, google.peer]
    }
  }
}
