resource "azurerm_resource_group" "rg" {
  name     = var.resource_group_name
  location = var.location
}

module "vnet" {
  source              = "./vpc.tf"
  resource_group_name = azurerm_resource_group.rg.name
  location            = azurerm_resource_group.rg.location
  vnet_name           = var.vnet_name
  vnet_address_space  = var.vnet_address_space
  subnet_configs      = var.subnet_configs
}

module "security" {
  source              = "./security.tf"
  resource_group_name = azurerm_resource_group.rg.name
  location            = azurerm_resource_group.rg.location
  nsg_name            = var.nsg_name
}

module "peering" {
  source                       = "./vpc_peering.tf"
  resource_group_name          = azurerm_resource_group.rg.name
  location                     = azurerm_resource_group.rg.location
  vnet_name                    = var.vnet_name
  peering_name                 = var.peering_name
  peer_resource_group_name     = var.peer_resource_group_name
  peer_vnet_name               = var.peer_vnet_name
}