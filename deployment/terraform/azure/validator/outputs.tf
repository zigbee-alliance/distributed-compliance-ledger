output "vpc" {
  value = module.this_vnet
}

output "private_ips" {
  value = [azurerm_linux_virtual_machine.validator_vm.private_ip]
}

output "public_ips" {
  value = [azurerm_public_ip.validator_public_ip.ip_address]
}
