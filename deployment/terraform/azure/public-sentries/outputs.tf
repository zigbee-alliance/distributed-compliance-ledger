output "vnet_id" {
  description = "ID of the Virtual Network"
  value       = azurerm_virtual_network.main.id
}

output "nsg_id" {
  description = "ID of the Network Security Group"
  value       = azurerm_network_security_group.nsg.id
}

output "peering_id" {
  description = "ID of the VNet Peering"
  value       = azurerm_virtual_network_peering.peer.id
}