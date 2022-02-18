output "ssh_console" {
  value = format("ssh -o 'StrictHostKeyChecking=no' ubuntu@%s", aws_instance.dcl_node.public_ip)
}
