output "private_ips" {
  value = aws_instance.this_nodes.*.private_ip
}

output "public_ips" {
  value = var.enable_ipv6 ? [for node in aws_instance.this_nodes : node.ipv6_addresses[0]] : aws_eip.this_nodes_eips.*.public_ip
}

output "seed_private_ips" {
  value = [aws_instance.this_seed_node.private_ip]
}

output "seed_public_ips" {
  value = [var.enable_ipv6 ? aws_instance.this_seed_node.ipv6_addresses[0] : aws_eip.this_seed_eip[0].public_ip]
}
