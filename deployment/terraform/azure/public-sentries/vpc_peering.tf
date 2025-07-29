resource "azurerm_virtual_network_peering" "peer" {
  name                      = var.peering_name
  resource_group_name       = var.resource_group_name
  virtual_network_name      = var.vnet_name
  remote_virtual_network_id = data.azurerm_virtual_network.peer.id
  allow_forwarded_traffic   = true
  allow_gateway_transit     = false
  use_remote_gateways       = false
  depends_on                = [azurerm_virtual_network.main]
}

// Data for peer VNet

data "azurerm_virtual_network" "peer" {
  name                = var.peer_vnet_name
  resource_group_name = var.peer_resource_group_name
}