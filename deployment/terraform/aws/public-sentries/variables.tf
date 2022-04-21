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
    description = "Number of Public Sentry nodes"
}

variable "region_index" {
    description = "Public Sentries Region Index"
}

variable "enable_ipv6" {
    description = "Enable public IPv6 addresses"
}

variable "peer_vpc" {
    description = "Peer VPC"
}