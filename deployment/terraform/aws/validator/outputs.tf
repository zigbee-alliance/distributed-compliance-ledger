output "vpc" {
  value = module.this_vpc
}

output "private_ips" {
  value = [aws_instance.this_node.private_ip]
}

output "public_ips" {
  value = [aws_instance.this_node.public_ip]
}