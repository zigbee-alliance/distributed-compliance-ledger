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

variable "validator_vpc" {
    description = "Validator VPC"
}

variable "nodes_count" {
    description = "Number of Private Sentry nodes"
    default = 2
}