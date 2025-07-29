variable "tags" {
  description = "A map of tags to add to all resources"
  type        = map(string)
  default     = {}
}

variable "resource_group_name" {
  description = "Name of the Azure Resource Group"
  type        = string
}

variable "location" {
  description = "Azure region for resources"
  type        = string
}

variable "subnet_id" {
  description = "ID of the subnet to deploy VMs into"
  type        = string
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
  description = "Number of Public Sentry nodes"
}

variable "instance_type" {
  description = "Azure VM size (e.g., Standard_B2s)"
}

variable "enable_ipv6" {
  description = "Enable IPv6 support"
  type        = bool
  default     = false
}