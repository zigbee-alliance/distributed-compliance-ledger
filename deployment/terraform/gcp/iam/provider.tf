terraform {
  required_providers {
    google = {
      source  = "hashicorp/google"
      // TODO: Initially set the same as AWS one. What should be here?
      version = ">= 4.1"
    }
  }
}
