variable "location" {
  description = "Azure region for resource deployment"
  type        = string
  default     = "East US"
}

variable "resource_group_name" {
  description = "Name of the Azure Resource Group"
  type        = string
  default     = "my-resource-group"
}

variable "vnet_name" {
  description = "Name of the Azure Virtual Network"
  type        = string
  default     = "my-vnet"
}

variable "vnet_address_space" {
  description = "Address space for the VNet"
  type        = list(string)
  default     = ["10.0.0.0/16"]
}

variable "subnet_configs" {
  description = "List of subnet configurations"
  type = list(object({
    name          = string
    address_prefix = string
  }))
  default = [
    { name = "subnet1"; address_prefix = "10.0.1.0/24" },
    { name = "subnet2"; address_prefix = "10.0.2.0/24" }
  ]
}

variable "nsg_name" {
  description = "Name of the network security group"
  type        = string
  default     = "my-nsg"
}

variable "peering_name" {
  description = "Name for VNet peering"
  type        = string
  default     = "peer-to-other-vnet"
}