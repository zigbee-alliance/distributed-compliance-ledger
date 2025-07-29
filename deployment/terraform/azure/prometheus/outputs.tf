output "private_ips" {
  value = [azurerm_linux_virtual_machine.this_node.private_ip_address]
}

output "public_ips" {
  value = [azurerm_public_ip.this_node.ip_address]
}

output "prometheus_endpoint" {
  value = azurerm_public_ip.this_node.ip_address
}