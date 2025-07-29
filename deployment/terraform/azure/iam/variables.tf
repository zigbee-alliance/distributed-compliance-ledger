variable "identity_name" {
  description = "Name of the User Assigned Managed Identity"
  type        = string
}

variable "resource_group_name" {
  description = "Name of the Resource Group where identity is created"
  type        = string
}

variable "location" {
  description = "Azure region for resources"
  type        = string
}

variable "roles" {
  description = "List of role definition names to assign to the identity"
  type        = list(string)
  default     = ["Contributor"]
}

data "azurerm_subscription" "primary" {}