data "azurerm_virtual_network" "peer_vnet" {
  name                = var.peer_vnet_name
  resource_group_name = var.peer_vnet_resource_group_name

  depends_on = [azurerm_virtual_network.this] # force deferring to apply phase
}

resource "azurerm_virtual_network_peering" "this_public_sentries_to_private_sentries_vnet_peering" {
  name                         = "${local.resource_prefix}-to-private-sentries-vnet-peering"
  resource_group_name          = local.resource_group_name
  virtual_network_name         = azurerm_virtual_network.this.name
  remote_virtual_network_id    = data.azurerm_virtual_network.peer_vnet.id
  allow_virtual_network_access = true
  allow_forwarded_traffic      = true
  allow_gateway_transit        = false
}

# Azure Virtual Network peering between Virtual Network B and A
resource "azurerm_virtual_network_peering" "this_private_sentries_to_public_sentries_vnet_peering" {
  name                         = "private-sentries-to-${local.resource_prefix}-vnet-peering"
  resource_group_name          = var.peer_vnet_resource_group_name
  virtual_network_name         = data.azurerm_virtual_network.peer_vnet.name
  remote_virtual_network_id    = azurerm_virtual_network.this.id
  allow_virtual_network_access = true
  allow_forwarded_traffic      = true
  allow_gateway_transit        = false
}
