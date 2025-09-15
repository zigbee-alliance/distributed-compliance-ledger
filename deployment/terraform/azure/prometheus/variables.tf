variable "ssh_public_key_path" {
  description = "SSH public key file path"
  default     = "~/.ssh/id_rsa.pub"
}

variable "ssh_private_key_path" {
  description = "SSH private key file path"
  default     = "~/.ssh/id_rsa"
}

variable "ssh_username" {
  description = "SSH username"
  default     = "ubuntu"
}

variable "instance_type" {
  description = "Prometheus instance type"
}

variable "endpoints" {
  description = "Prometheus endpoints"
}

variable "vpc" {
  description = "Prometheus VPC"
}