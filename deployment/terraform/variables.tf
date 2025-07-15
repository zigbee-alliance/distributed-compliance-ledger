variable "cloud_provider" {
  description = "Cloud provider to deploy resources to (aws or gcp)"
  type        = string
  default     = "aws"
}

variable "region" {
  description = "AWS region"
  type        = string
  default     = "us-west-1"
}

variable "gcp_project" {
  description = "GCP project ID"
  type        = string
  default     = "DCL"
}

variable "gcp_regions" {
  description = "List of GCP regions to deploy resources"
  type        = list(string)
  default     = ["us-east1", "us-west1"]
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