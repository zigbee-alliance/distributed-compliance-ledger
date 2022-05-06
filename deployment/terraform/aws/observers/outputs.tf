output "public_ips" {
  value = aws_instance.this_nodes.*.public_ip
}