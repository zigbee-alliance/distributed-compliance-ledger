output "ssh_console" {
  value = var.cloud_provider == "aws" ?
    format("ssh -o 'StrictHostKeyChecking=no' %s@%s", var.ssh_username, module.aws_infra[0].public_ip) :
    format("ssh -o 'StrictHostKeyChecking=no' %s@%s", var.ssh_username, module.gcp_infra[0].public_ip)
  description = "SSH command to connect to the DCL node"
}