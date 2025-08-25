output "private_ips" {
  value = azurerm_linux_virtual_machine.this_nodes.*.private_ip_address
}

# TODO ipv6 support
output "public_ips" {
  value = azurerm_linux_virtual_machine.this_nodes.*.public_ip_address
}

output "seed_private_ips" {
  value = [azurerm_linux_virtual_machine.seed.private_ip_address]
}

# TODO ipv6 support
output "seed_public_ips" {
  value = [azurerm_linux_virtual_machine.seed.public_ip_address]
}
