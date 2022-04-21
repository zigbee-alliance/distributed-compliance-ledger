output "public_ips" {
    value = var.enable_ipv6 ? [for node in aws_instance.this_nodes : node.ipv6_addresses[0]] : aws_instance.this_nodes.*.public_ip
}

output "seed_public_ips" {
    value = [var.enable_ipv6 ? aws_instance.this_seed_node.ipv6_addresses[0] : aws_instance.this_seed_node.public_ip]
}