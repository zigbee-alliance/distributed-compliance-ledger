variable "tags" {
  description = "A map of tags to add to all resources"
  type        = map(string)
  default     = {}
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

variable "instance_type" {
  description = "Type of Azure instances"
}

variable "iam_instance_profile" {
  description = "IAM instance profile"
}
