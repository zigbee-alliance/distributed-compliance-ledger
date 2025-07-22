output "vpc" {
  value = module.this_vpc
}

output "private_ips" {
  value = azurerm_instance.this_nodes.*.private_ip
}

output "public_ips" {
  value = concat(slice(azurerm_instance.this_nodes, length(aws_eip.this_eips), length(azurerm_instance.this_nodes)).*.public_ip, aws_eip.this_eips.*.public_ip)
}

output "public_eips" {
  value = aws_eip.this_eips.*.public_ip
}