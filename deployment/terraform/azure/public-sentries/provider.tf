terraform {
  required_providers {
    azurerm = {
      source  = "hashicorp/azurerm"
      version = ">= 4.1"
    }
  }
}

provider "azurerm" {
  features {}
}

provider "azurerm" {
  alias    = "peer"
  features {}
}