variable "peer_resource_group_name" {
  description = "Resource group of the peer VNet"
  type        = string
}

variable "peer_vnet_name" {
  description = "Name of the peer VNet"
  type        = string
}

resource "azurerm_virtual_network" "main" {
  name                = var.vnet_name
  address_space       = var.vnet_address_space
  location            = var.location
  resource_group_name = var.resource_group_name
}

resource "azurerm_subnet" "subnets" {
  for_each             = { for s in var.subnet_configs : s.name => s }
  name                  = each.value.name
  resource_group_name   = var.resource_group_name
  virtual_network_name  = azurerm_virtual_network.main.name
  address_prefixes      = [each.value.address_prefix]
}