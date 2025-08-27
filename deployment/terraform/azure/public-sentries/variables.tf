variable "tags" {
  description = "A map of tags to add to all resources"
  type        = map(string)
  default     = {}
}

variable "resource_group_name" {
  description = "Resource group to use for the validator resources"
  type        = string
}

variable "location" {
  description = "Azure location. By default resource group's location is used"
  default = null
}

variable "location_index" {
  description = "Public Sentries Location Index"
  type        = number
}

variable "resource_suffix" {
  description = "Resource suffix to use for all the resources"
  default = null
}

variable "enable_encryption_at_host" {
  description = "Enables encryption at host for the node's managed disks"
  type        = bool
  default     = false
}

variable "enable_ipv6" {
  description = "Enable public IPv6 addresses"
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

variable "peer_vnet_name" {
  description = "Peer Virtual Network name"
  type        = string
}

variable "peer_vnet_resource_group_name" {
  description = "Peer Virtual Network resource group"
  type        = string
}

variable "nodes_count" {
  description = "Number of Public Sentry nodes"
}

variable "instance_size" {
  description = "Type of Azure instances"
}

# FIXME
#variable "iam_instance_profile" {
#  description = "IAM instance profile"
#}
