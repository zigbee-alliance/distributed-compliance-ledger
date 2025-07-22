output "private_ips" {
  value = azurerm_instance.this_nodes.*.private_ip
}

output "public_ips" {
  value = azurerm_instance.this_nodes.*.public_ip
}
