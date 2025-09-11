output "vnet" {
  value = azurerm_virtual_network.this
}

output "private_ips" {
  value = azurerm_linux_virtual_machine.this_nodes.*.private_ip_address
}

output "public_ips" {
  value = azurerm_linux_virtual_machine.this_nodes.*.public_ip_address
}
