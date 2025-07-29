terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = ">= 3.72"
    }
    google = {
      source  = "hashicorp/google"
      version = ">= 4.0"
    }
    azurerm = {
      source  = "hashicorp/azurerm"
      version = ">= 4.1"
    }
  }
}

provider "aws" {
  region     = var.region
  alias      = "aws"
  skip_credentials_validation = false
  skip_metadata_api_check     = false
  skip_region_validation      = false
  count      = var.cloud_provider == "aws" ? 1 : 0
}

provider "google" {
  project = var.gcp_project
  region  = var.gcp_regions[0]
  alias   = "gcp"
  count   = var.cloud_provider == "gcp" ? 1 : 0
}
provider "azurerm" {
  alias    = "azure"
  features {}
  count    = var.cloud_provider == "azure" ? 1 : 0
}

# AWS Infrastructure
module "aws_infra" {
  source           = "./aws"
  region           = var.region
  ssh_username     = var.ssh_username
  ssh_public_key_path  = var.ssh_public_key_path
  ssh_private_key_path = var.ssh_private_key_path

  providers = {
    aws = aws
  }

  count = var.cloud_provider == "aws" ? 1 : 0
}

# GCP Infrastructure
module "gcp_infra" {
  source           = "./gcp"
  project_id       = var.gcp_project
  regions          = var.gcp_regions
  ssh_username     = var.ssh_username
  ssh_public_key_path  = var.ssh_public_key_path
  ssh_private_key_path = var.ssh_private_key_path

  providers = {
    google = google
  }

  count = var.cloud_provider == "gcp" ? 1 : 0
}

# Azure Infrastructure
module "azure_infra" {
  source               = "./azure"
  location             = var.azure_location
  resource_group_name  = var.azure_resource_group
  vnet_name            = var.azure_vnet_name
  ssh_username         = var.ssh_username
  ssh_public_key_path  = var.ssh_public_key_path
  ssh_private_key_path = var.ssh_private_key_path

  providers = {
    azurerm = azurerm.azure
  }

  count = var.cloud_provider == "azure" ? 1 : 0
}