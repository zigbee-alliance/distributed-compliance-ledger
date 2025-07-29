provider "azurerm" {
  features {}
}

variable "location" {
  default = "East US"
}

variable "tags" {
  type    = map(string)
  default = {}
}
