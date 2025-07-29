output "private_ips" {
  value = azurerm_linux_virtual_machine.this_nodes[*].private_ip_address
}

output "public_ips" {
  value = azurerm_public_ip.this_nodes[*].ip_address
}