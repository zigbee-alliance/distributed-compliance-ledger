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
  default     = null
}

variable "enable_encryption_at_host" {
  description = "Enables encryption at host for the node's managed disks"
  type        = bool
  default     = false
}

variable "disable_instance_protection" {
  description = "Disable the protection that prevents the validator instance from being accidentally terminated"
  type        = bool
  default     = false
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

variable "instance_size" {
  description = "Type of Azure instances"
}

# FIXME
#variable "iam_instance_profile" {
#  description = "IAM instance profile"
#}
