output "vpc" {
  value = module.this_vpc
}

output "private_ips" {
  value = [aws_instance.this_node.private_ip]
}

output "public_ips" {
  value = [aws_instance.this_node.public_ip]
}

output "prometheus_endpoint" {
  value = var.enable_prometheus ? aws_prometheus_workspace.this_amp_workspace[0].prometheus_endpoint : null
}