variable "tags" {
  description = "A map of tags to add to all resources"
  type        = map(string)
  default     = {}
}

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

variable "nodes_count" {
  description = "Number of Observer nodes"
}

variable "instance_type" {
  description = "Type of AWS instances"
}

variable "region_index" {
  description = "Observer Region Index"
}

variable "enable_tls" {
  description = "Enable TLS on LB listeners"
}

variable "root_domain_name" {
  description = "Root domain name"
}

variable "peer_vpc" {
  description = "Peer VPC"
}

variable "iam_instance_profile" {
  description = "IAM instance profile"
}
